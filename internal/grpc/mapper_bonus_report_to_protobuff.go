package grpc

import (
	"code.citik.ru/back/report-action/internal"
	bonusv1 "code.citik.ru/back/report-action/internal/grpc/gen/citilink/reportaction/bonus/v1"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
)

// bonusReportToProtobufMapper маппер отчетов по бонусам из protobuf в доменные структуры
type bonusReportToProtobufMapper struct{}

// NewBonusReportToProtobufMapper создает новый мапер в доменные структуры.
func NewBonusReportToProtobufMapper() *bonusReportToProtobufMapper {
	return &bonusReportToProtobufMapper{}
}

func (m *bonusReportToProtobufMapper) mapType(domainType internal.ReportType) (bonusv1.Type, error) {
	pbType, ok := toProtobufEnumMaps.bonusReportTypes[domainType]
	if !ok {
		return bonusv1.Type_TYPE_INVALID, errors.New("type is invalid")
	}

	return pbType, nil
}

func (m *bonusReportToProtobufMapper) mapStatus(domainStatus internal.ReportStatus) (bonusv1.Status, error) {
	pbStatus, ok := toProtobufEnumMaps.bonusReportStatuses[domainStatus]
	if !ok {
		return bonusv1.Status_STATUS_INVALID, errors.New("status is invalid")
	}

	return pbStatus, nil
}

func (m *bonusReportToProtobufMapper) mapBonusReport(report *internal.Report) (*bonusv1.Info, error) {
	pbType, err := m.mapType(report.Type)
	if err != nil {
		return nil, fmt.Errorf("can't map type: %w", err)
	}

	pbStatus, err := m.mapStatus(report.Status)
	if err != nil {
		return nil, fmt.Errorf("can't map status: %w", err)
	}

	return &bonusv1.Info{
		Id:          uuid.UUID(report.Id).String(),
		Type:        pbType,
		Name:        report.FileName,
		Status:      pbStatus,
		CreatedTime: &timestamp.Timestamp{Seconds: report.LastModified.Unix()},
	}, nil
}
