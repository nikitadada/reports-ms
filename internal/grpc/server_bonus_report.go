package grpc

import (
	"code.citik.ru/back/report-action/internal"
	"code.citik.ru/back/report-action/internal/distributed_work"
	bonusv1 "code.citik.ru/back/report-action/internal/grpc/gen/citilink/reportaction/bonus/v1"
	filev1 "code.citik.ru/back/report-action/internal/grpcclient/gen/citilink/cmsfiles/file/v1"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/go-version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

// BonusReportServer сервер отчетов по бонусам
type BonusReportServer struct {
	queue            distributed_work.Inserter
	storage          internal.ReportRepository
	appVersion       *version.Version
	cmsFilesClient   filev1.FileAPIClient
	toDomainMapper   *bonusReportToDomainMapper
	toProtobufMapper *bonusReportToProtobufMapper
}

func NewBonusReportServer(
	queue distributed_work.Inserter,
	storage internal.ReportRepository,
	appVersion *version.Version,
	cmsFilesClient filev1.FileAPIClient,
	toDomainMapper *bonusReportToDomainMapper,
	toProtobufMapper *bonusReportToProtobufMapper,
) *BonusReportServer {
	return &BonusReportServer{
		queue:            queue,
		storage:          storage,
		appVersion:       appVersion,
		cmsFilesClient:   cmsFilesClient,
		toDomainMapper:   toDomainMapper,
		toProtobufMapper: toProtobufMapper,
	}
}

// Create инициирует создание нового отчета по бонусам акций
func (s *BonusReportServer) Create(ctx context.Context, r *bonusv1.CreateRequest) (*bonusv1.CreateResponse, error) {
	num := r.GetNavisionActionNumber()
	if num == "" {
		return nil, status.Error(codes.InvalidArgument, "navision action number is required")
	}

	startTime := r.GetActionStartTime()
	if startTime.Seconds < 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid start time")
	}

	domainType, err := s.toDomainMapper.mapType(r.GetType())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't map type: %s", err))
	}

	// пробуем получить отчет переданного типа сгенерированный в течение 2-х последний часов без ошибки
	reports, err := s.storage.Filter(ctx,
		&internal.Filter{
			NavActionNumber:   num,
			CampaignStartDate: startTime.Seconds,
			ValidTime:         time.Now().Add(-2 * time.Hour),
			Types:             []internal.ReportType{domainType},
			Statuses: []internal.ReportStatus{
				internal.ReportStatusCreated,
				internal.ReportStatusInProcess,
				internal.ReportStatusSuccess,
			},
		})
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't find report: %s", err))
	}

	createdAt := time.Now()
	reportStatus := internal.ReportStatusCreated
	filename := fmt.Sprintf("bonus_report_%d-%02d-%02d", createdAt.Year(), createdAt.Month(), createdAt.Day())

	var id uuid.UUID
	// генерируем новый отчет только в случае, если нет отчетов с аналогичными параметрами за 2 последних часа
	if len(reports) == 0 {
		res, err := s.cmsFilesClient.Create(ctx, &filev1.CreateRequest{
			Name:      filename,
			Type:      filev1.Type_TYPE_BONUS_ACTION,
			Extension: internal.ReportFileExtension,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("can't create file in microservice cms-files: %s", err))
		}

		id, err = uuid.Parse(res.GetId())
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("can't parse file id: %s", err))
		}

		err = s.storage.Create(ctx,
			&internal.Report{
				Id:                internal.ReportId(id),
				NavActionNumber:   num,
				CampaignStartDate: startTime.AsTime().Local(),
				Status:            reportStatus,
				Type:              domainType,
			},
		)
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("can't save report: %s", err))
		}

		payload, err := json.Marshal(internal.Payload{
			ReportId:        id.String(),
			ActionNumber:    num,
			ActionStartTime: startTime.Seconds,
			ReportType:      domainType,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("can't marshal task payload: %s", err))
		}

		_, err = s.queue.Insert(internal.BonusTaskType, s.appVersion, payload, distributed_work.TaskOptions{})
		if err != nil {
			return nil, status.Error(codes.Internal, fmt.Sprintf("can't insert task: %s", err))
		}
	} else {
		id = uuid.UUID(reports[0].Id)
		createdAt = reports[0].LastModified
		reportStatus = reports[0].Status
	}

	bonusReport := &internal.Report{
		Id:           internal.ReportId(id),
		Type:         domainType,
		FileName:     filename,
		Status:       reportStatus,
		LastModified: createdAt,
	}

	pbBonusReport, err := s.toProtobufMapper.mapBonusReport(bonusReport)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't map bonus report: %s", err))
	}

	return &bonusv1.CreateResponse{Info: pbBonusReport}, nil
}

// Get возвращает отчет по бонусам акций
func (s *BonusReportServer) Get(ctx context.Context, r *bonusv1.GetRequest) (*bonusv1.GetResponse, error) {
	id, err := uuid.Parse(r.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid argument id: %s", err))
	}

	res, err := s.cmsFilesClient.Get(ctx, &filev1.GetRequest{Id: r.GetId()})
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't get file: %s", err))
	}

	file := res.GetFile()
	report, err := s.storage.Get(ctx, internal.ReportId(id))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't get report: %s", err))
	}
	if report == nil {
		return nil, status.Error(codes.NotFound, "report not found")
	}

	bonusReport := &internal.Report{
		Id:           report.Id,
		Type:         report.Type,
		FileName:     file.GetName(),
		Status:       report.Status,
		LastModified: file.GetModifiedTime().AsTime(),
	}

	pbBonusReport, err := s.toProtobufMapper.mapBonusReport(bonusReport)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't map report: %s", err))
	}

	return &bonusv1.GetResponse{Info: pbBonusReport}, nil
}
