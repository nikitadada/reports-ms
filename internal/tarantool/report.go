package tarantool

import (
	"code.citik.ru/back/report-action/internal"
	cititararantool "code.citik.ru/gobase/tarantool"
	"context"
	"fmt"
	"github.com/google/uuid"
)

// ReportRepository хранилище отчетов
type ReportRepository struct {
	client cititararantool.Client
}

// NewReportRepository создает новое хранилище отчетов
func NewReportRepository(client cititararantool.Client) *ReportRepository {
	return &ReportRepository{client: client}
}

// TntReportLines ответ тарантула
type TntReportLines struct {
	Reports []*internal.Report
}

// Filter возвращает отфильтрованные отчеты из хранилища
func (r *ReportRepository) Filter(ctx context.Context, filter *internal.Filter) ([]*internal.Report, error) {
	var res *TntReportLines
	err := r.client.Call17Typed(
		ctx, "API.report.filter.v1",
		[]interface{}{
			filter.NavActionNumber,
			filter.CampaignStartDate,
			filter.ValidTime.Unix(),
			filter.Types,
			filter.Statuses,
		}, &res,
	)
	if err != nil {
		return nil, fmt.Errorf("can't filter reports: %w", err)
	}

	return res.Reports, nil
}

// Create Создает новый отчет в хранилище
func (r *ReportRepository) Create(ctx context.Context, report *internal.Report) error {
	_, err := r.client.Call17(
		ctx, "API.report.create.v1",
		[]interface{}{
			uuid.UUID(report.Id),
			report.NavActionNumber,
			report.CampaignStartDate.Unix(),
			report.Status,
			report.Type,
		})
	if err != nil {
		return fmt.Errorf("can't create report: %w", err)
	}

	return nil
}

// Get Ищет отчет в хранилище по идентификатору
func (r *ReportRepository) Get(ctx context.Context, id internal.ReportId) (*internal.Report, error) {
	var res []*internal.Report
	err := r.client.Call17Typed(
		ctx, "API.report.get.v1",
		[]interface{}{uuid.UUID(id)}, &res)
	if err != nil {
		return nil, fmt.Errorf("can't get report: %w", err)
	}

	if len(res) == 0 {
		return nil, nil
	}

	return res[0], nil
}

// Update Обновляет отчет в хранилище
func (r *ReportRepository) Update(
	ctx context.Context,
	id internal.ReportId,
	actionNumber string,
	actionStart int64,
	status internal.ReportStatus,
	reportType internal.ReportType,
) error {
	_, err := r.client.Call17(
		ctx, "API.report.update.v1",
		[]interface{}{
			uuid.UUID(id),
			actionNumber,
			actionStart,
			status,
			reportType,
		},
	)
	if err != nil {
		return fmt.Errorf("can't update report: %w", err)
	}

	return nil
}
