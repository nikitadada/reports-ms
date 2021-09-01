package tarantool

import (
	"code.citik.ru/back/report-action/internal"
	"code.citik.ru/back/report-action/internal/mock"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReportRepository_Create(t *testing.T) {
	t.Run("create error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tntClient := mock.NewMockClient(ctrl)

		report := &internal.Report{
			Id:                internal.ReportId(uuid.New()),
			NavActionNumber:   "test",
			Type:              internal.ReportTypeDetailed,
			Status:            internal.ReportStatusSuccess,
			CampaignStartDate: time.Unix(0, 0),
		}

		tntClient.EXPECT().
			Call17(
				context.TODO(), "API.report.create.v1",
				[]interface{}{
					uuid.UUID(report.Id),
					report.NavActionNumber,
					report.CampaignStartDate.Unix(),
					report.Status,
					report.Type,
				}).
			Return(nil, errors.New("test-error"))

		fileRepository := NewReportRepository(tntClient)
		err := fileRepository.Create(context.TODO(), report)

		assert.EqualError(t, err, "can't create report: test-error")
	})

	t.Run("create success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tntClient := mock.NewMockClient(ctrl)

		report := &internal.Report{
			Id:                internal.ReportId(uuid.New()),
			NavActionNumber:   "test",
			Type:              internal.ReportTypeDetailed,
			Status:            internal.ReportStatusSuccess,
			CampaignStartDate: time.Unix(0, 0),
		}

		tntClient.EXPECT().
			Call17(
				context.TODO(), "API.report.create.v1",
				[]interface{}{
					uuid.UUID(report.Id),
					report.NavActionNumber,
					report.CampaignStartDate.Unix(),
					report.Status,
					report.Type,
				}).
			Return(nil, nil)

		reportRepository := NewReportRepository(tntClient)
		err := reportRepository.Create(context.TODO(), report)

		assert.NoError(t, err)
	})
}

func TestReportRepository_Get(t *testing.T) {
	t.Run("get error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tntClient := mock.NewMockClient(ctrl)
		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")

		tntClient.EXPECT().
			Call17Typed(context.TODO(), "API.report.get.v1", []interface{}{id}, gomock.Any()).
			Return(errors.New("test-error"))

		reportRepository := NewReportRepository(tntClient)
		report, err := reportRepository.Get(context.TODO(), internal.ReportId(id))

		assert.EqualError(t, err, "can't get report: test-error")
		assert.Nil(t, report)
	})

	t.Run("get empty result", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tntClient := mock.NewMockClient(ctrl)
		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")

		tntClient.EXPECT().
			Call17Typed(context.TODO(), "API.report.get.v1", []interface{}{id}, gomock.Any()).
			SetArg(3, []*internal.Report{}).
			Return(nil)

		reportRepository := NewReportRepository(tntClient)
		report, err := reportRepository.Get(context.TODO(), internal.ReportId(id))

		assert.NoError(t, err)
		assert.Nil(t, report)
	})

	t.Run("get success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tntClient := mock.NewMockClient(ctrl)

		report := &internal.Report{
			Id:                internal.ReportId(uuid.New()),
			NavActionNumber:   "test",
			Type:              internal.ReportTypeDetailed,
			Status:            internal.ReportStatusSuccess,
			CampaignStartDate: time.Unix(0, 0),
		}
		var expectedReports []*internal.Report
		expectedReports = append(expectedReports, report)

		tntClient.EXPECT().
			Call17Typed(context.TODO(), "API.report.get.v1", []interface{}{uuid.UUID(report.Id)}, gomock.Any()).
			SetArg(3, expectedReports).
			Return(nil)

		fileRepository := NewReportRepository(tntClient)
		got, err := fileRepository.Get(context.TODO(), report.Id)

		assert.NoError(t, err)
		assert.Equal(t, expectedReports[0], got)
	})
}

func TestReportRepository_Update(t *testing.T) {
	t.Run("update error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tntClient := mock.NewMockClient(ctrl)

		report := &internal.Report{
			Id:                internal.ReportId(uuid.New()),
			NavActionNumber:   "test",
			Type:              internal.ReportTypeDetailed,
			Status:            internal.ReportStatusSuccess,
			CampaignStartDate: time.Unix(0, 0),
		}

		tntClient.EXPECT().
			Call17(
				context.TODO(), "API.report.update.v1",
				[]interface{}{
					uuid.UUID(report.Id),
					report.NavActionNumber,
					report.CampaignStartDate.Unix(),
					report.Status,
					report.Type,
				}).
			Return(nil, errors.New("test-error"))

		fileRepository := NewReportRepository(tntClient)
		err := fileRepository.Update(
			context.TODO(),
			report.Id,
			report.NavActionNumber,
			report.CampaignStartDate.Unix(),
			report.Status,
			report.Type,
		)

		assert.EqualError(t, err, "can't update report: test-error")
	})

	t.Run("update success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tntClient := mock.NewMockClient(ctrl)

		report := &internal.Report{
			Id:                internal.ReportId(uuid.New()),
			NavActionNumber:   "test",
			Type:              internal.ReportTypeDetailed,
			Status:            internal.ReportStatusSuccess,
			CampaignStartDate: time.Unix(0, 0),
		}

		tntClient.EXPECT().
			Call17(
				context.TODO(), "API.report.update.v1",
				[]interface{}{
					uuid.UUID(report.Id),
					report.NavActionNumber,
					report.CampaignStartDate.Unix(),
					report.Status,
					report.Type,
				}).
			Return(nil, nil)
		fileRepository := NewReportRepository(tntClient)
		err := fileRepository.Update(
			context.TODO(),
			report.Id,
			report.NavActionNumber,
			report.CampaignStartDate.Unix(),
			report.Status,
			report.Type,
		)

		assert.NoError(t, err)
	})
}

func TestReportRepository_Filter(t *testing.T) {
	t.Run("filter error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tntClient := mock.NewMockClient(ctrl)
		filter := &internal.Filter{
			NavActionNumber:   "test",
			CampaignStartDate: time.Unix(0, 0).Unix(),
			ValidTime:         time.Unix(0, 0),
			Types:             []internal.ReportType{internal.ReportTypeDetailed},
			Statuses:          []internal.ReportStatus{internal.ReportStatusSuccess},
		}

		var res *TntReportLines
		tntClient.EXPECT().
			Call17Typed(
				context.TODO(), "API.report.filter.v1",
				[]interface{}{
					filter.NavActionNumber,
					filter.CampaignStartDate,
					filter.ValidTime.Unix(),
					filter.Types,
					filter.Statuses,
				}, &res).
			Return(errors.New("test-error"))

		reportRepository := NewReportRepository(tntClient)
		reports, err := reportRepository.Filter(context.TODO(), filter)

		assert.EqualError(t, err, "can't filter reports: test-error")
		assert.Nil(t, reports)
	})

	t.Run("filter success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tntClient := mock.NewMockClient(ctrl)

		filter := &internal.Filter{
			NavActionNumber:   "test",
			CampaignStartDate: time.Unix(0, 0).Unix(),
			ValidTime:         time.Unix(0, 0),
			Types:             []internal.ReportType{internal.ReportTypeDetailed},
			Statuses:          []internal.ReportStatus{internal.ReportStatusSuccess},
		}

		report := &internal.Report{
			Id:                internal.ReportId(uuid.New()),
			NavActionNumber:   "test",
			Type:              internal.ReportTypeDetailed,
			Status:            internal.ReportStatusSuccess,
			CampaignStartDate: time.Unix(0, 0),
		}
		res := TntReportLines{
			Reports: []*internal.Report{report},
		}

		tntClient.EXPECT().
			Call17Typed(
				context.TODO(), "API.report.filter.v1",
				[]interface{}{
					filter.NavActionNumber,
					filter.CampaignStartDate,
					filter.ValidTime.Unix(),
					filter.Types,
					filter.Statuses,
				}, gomock.Any()).
			SetArg(3, &res).
			Return(nil)

		reportRepository := NewReportRepository(tntClient)
		got, err := reportRepository.Filter(context.TODO(), filter)

		assert.NoError(t, err)
		assert.Equal(t, res.Reports, got)
	})

	t.Run("filter empty result", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		tntClient := mock.NewMockClient(ctrl)

		filter := &internal.Filter{
			NavActionNumber:   "test",
			CampaignStartDate: time.Unix(0, 0).Unix(),
			ValidTime:         time.Unix(0, 0),
			Types:             []internal.ReportType{internal.ReportTypeDetailed},
			Statuses:          []internal.ReportStatus{internal.ReportStatusSuccess},
		}

		res := TntReportLines{
			Reports: []*internal.Report{},
		}

		tntClient.EXPECT().
			Call17Typed(
				context.TODO(), "API.report.filter.v1",
				[]interface{}{
					filter.NavActionNumber,
					filter.CampaignStartDate,
					filter.ValidTime.Unix(),
					filter.Types,
					filter.Statuses,
				}, gomock.Any()).
			SetArg(3, &res).
			Return(nil)

		reportRepository := NewReportRepository(tntClient)
		got, err := reportRepository.Filter(context.TODO(), filter)

		assert.NoError(t, err)
		assert.Equal(t, res.Reports, got)
	})
}
