package internal

//go:generate mockgen -source=bonus.go -destination=./bonus_mock_test.go -package=internal
//go:generate mockgen -source=report.go -destination=./report_mock_test.go -package=internal
//go:generate mockgen -source=uploader.go -destination=./uploader_mock_test.go -package=internal

import (
	"bytes"
	"code.citik.ru/back/report-action/internal/distributed_work"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestConsumer_CanConsume(t *testing.T) {
	tests := []struct {
		name           string
		task           *distributed_work.SimpleTask
		expectedResult bool
	}{
		{
			name: "can consume task with valid type",
			task: distributed_work.NewSimpleTask(
				BonusTaskType,
				[]byte("{\"report_id\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\"}"),
				version.Must(version.NewVersion("v1.0.0"))),
			expectedResult: true,
		},
		{
			name: "cannot consume task with invalid type",
			task: distributed_work.NewSimpleTask("invalid", []byte("{\"report_id\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\"}"),
				version.Must(version.NewVersion("v1.0.0"))),
			expectedResult: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockDbStorage := NewMockBonusStorage(ctrl)
			mockTntStorage := NewMockReportRepository(ctrl)
			mockFileGenerator := NewMockFileGenerator(ctrl)
			mockFileUploader := NewMockFileUploader(ctrl)
			ver := version.Must(version.NewVersion("v1.0.0"))

			consumer := NewBonusConsumer(mockDbStorage, mockTntStorage, zap.NewNop(), ver, mockFileGenerator, mockFileUploader)

			consumer.CanConsume(tt.task)
		})
	}
}

func TestConsumer_ResolveVersion(t *testing.T) {
	tests := []struct {
		name        string
		appVersion  *version.Version
		taskVersion *version.Version
		verdict     distributed_work.VersionResolverVerdict
	}{
		{
			name:        "should return delete-verdict if version of task less then app",
			appVersion:  version.Must(version.NewVersion("v2.0.0")),
			taskVersion: version.Must(version.NewVersion("v1.0.0")),
			verdict:     distributed_work.VersionResolverVerdictDelete,
		},
		{
			name:        "should return OK-verdict if version of task equal app",
			appVersion:  version.Must(version.NewVersion("v2.0.0")),
			taskVersion: version.Must(version.NewVersion("v2.0.0")),
			verdict:     distributed_work.VersionResolverVerdictOK,
		},
		{
			name:        "should return release-with-delay-verdict if version of task greater then app",
			appVersion:  version.Must(version.NewVersion("v2.0.0")),
			taskVersion: version.Must(version.NewVersion("v3.0.0")),
			verdict:     distributed_work.VersionResolverVerdictReleaseWithDelay,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockDbStorage := NewMockBonusStorage(ctrl)
			mockTntStorage := NewMockReportRepository(ctrl)
			mockFileGenerator := NewMockFileGenerator(ctrl)
			mockFileUploader := NewMockFileUploader(ctrl)

			consumer := NewBonusConsumer(mockDbStorage, mockTntStorage, zap.NewNop(), tt.appVersion, mockFileGenerator, mockFileUploader)
			assert.Equal(t, tt.verdict, consumer.ResolveVersion(tt.taskVersion))
		})
	}
}

func TestConsumer_Consume(t *testing.T) {
	t.Run("should return err if can't unmarshal task payload", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbStorage := NewMockBonusStorage(ctrl)
		mockTntStorage := NewMockReportRepository(ctrl)
		mockFileGenerator := NewMockFileGenerator(ctrl)
		mockFileUploader := NewMockFileUploader(ctrl)

		ver := version.Must(version.NewVersion("v1.0.0"))
		task := distributed_work.NewSimpleTask(BonusTaskType, []byte("{\"report_id\":\"invalid,,}"), ver)
		consumer := NewBonusConsumer(mockDbStorage, mockTntStorage, zap.NewNop(), ver, mockFileGenerator, mockFileUploader)

		err := consumer.Consume(context.Background(), task)
		assert.EqualError(t, err, "can't unmarshal payload of bonus task: unexpected end of JSON input")
	})

	t.Run("should return err if invalid uuid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbStorage := NewMockBonusStorage(ctrl)
		mockTntStorage := NewMockReportRepository(ctrl)
		mockFileGenerator := NewMockFileGenerator(ctrl)
		mockFileUploader := NewMockFileUploader(ctrl)

		ver := version.Must(version.NewVersion("v1.0.0"))
		task := distributed_work.NewSimpleTask(BonusTaskType, []byte("{\"reportId\":\"not uuid\"}"), ver)
		consumer := NewBonusConsumer(mockDbStorage, mockTntStorage, zap.NewNop(), ver, mockFileGenerator, mockFileUploader)

		err := consumer.Consume(context.Background(), task)
		assert.EqualError(t, err, "can't parse reportId to uuid: invalid UUID length: 8")
	})
	t.Run("should return err if error storage for general report", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbStorage := NewMockBonusStorage(ctrl)
		mockTntStorage := NewMockReportRepository(ctrl)
		mockFileGenerator := NewMockFileGenerator(ctrl)
		mockFileUploader := NewMockFileUploader(ctrl)

		ver := version.Must(version.NewVersion("v1.0.0"))
		task := distributed_work.NewSimpleTask(BonusTaskType, []byte("{\"reportId\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\",\"reportType\":1}"), ver)
		consumer := NewBonusConsumer(mockDbStorage, mockTntStorage, zap.NewNop(), ver, mockFileGenerator, mockFileUploader)

		var payload Payload
		_ = json.Unmarshal(task.Payload(), &payload)

		mockDbStorage.EXPECT().FindGeneral(
			context.Background(),
			payload.ActionNumber,
			time.Unix(payload.ActionStartTime, 0),
		).Return(nil, errors.New("some error"))

		err := consumer.Consume(context.Background(), task)
		assert.EqualError(t, err, "failed getting general bonuses: some error")
	})
	t.Run("should no err if error file generation for general report", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbStorage := NewMockBonusStorage(ctrl)
		mockTntStorage := NewMockReportRepository(ctrl)
		mockFileGenerator := NewMockFileGenerator(ctrl)
		mockFileUploader := NewMockFileUploader(ctrl)

		ver := version.Must(version.NewVersion("v1.0.0"))
		task := distributed_work.NewSimpleTask(
			BonusTaskType,
			[]byte("{\"reportId\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\","+
				"\"reportType\":1,"+
				" \"actionNumber\":\"num\","+
				"\"actionStartTime\":199999999}"),
			ver,
		)
		consumer := NewBonusConsumer(mockDbStorage, mockTntStorage, zap.NewNop(), ver, mockFileGenerator, mockFileUploader)

		var payload Payload
		_ = json.Unmarshal(task.Payload(), &payload)

		mockDbStorage.EXPECT().FindGeneral(
			context.Background(),
			payload.ActionNumber,
			time.Unix(payload.ActionStartTime, 0),
		).Return([]*BonusGeneral{}, nil)

		mockFileGenerator.EXPECT().GenerateGeneral([]*BonusGeneral{}).Return(nil, errors.New("some error"))

		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
		mockTntStorage.EXPECT().Update(
			context.Background(),
			ReportId(id),
			"num",
			int64(199999999),
			ReportStatusError,
			ReportTypeGeneral,
		)

		err := consumer.Consume(context.Background(), task)
		assert.NoError(t, err)
	})
	t.Run("should return err if cannot update report in storage", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbStorage := NewMockBonusStorage(ctrl)
		mockTntStorage := NewMockReportRepository(ctrl)
		mockFileGenerator := NewMockFileGenerator(ctrl)
		mockFileUploader := NewMockFileUploader(ctrl)
		ctx := context.Background()

		ver := version.Must(version.NewVersion("v1.0.0"))
		task := distributed_work.NewSimpleTask(
			BonusTaskType,
			[]byte("{\"reportId\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\","+
				"\"reportType\":1,"+
				" \"actionNumber\":\"num\","+
				"\"actionStartTime\":199999999}"),
			ver,
		)
		consumer := NewBonusConsumer(mockDbStorage, mockTntStorage, zap.NewNop(), ver, mockFileGenerator, mockFileUploader)

		var payload Payload
		_ = json.Unmarshal(task.Payload(), &payload)

		bonusesData := []*BonusGeneral{
			{
				NavActionNumber:               "num",
				Bonus:                         100,
				CountClients:                  1,
				CampaignStartDate:             time.Unix(199999999, 0),
				CampaignFinishDate:            time.Unix(299999999, 0),
				CountClientsWithActiveCard:    1,
				CountClientsSendActivation:    1,
				CountClientsSuccessActivation: 1,
				ActivationPercent:             "100",
			},
		}

		mockDbStorage.EXPECT().FindGeneral(
			ctx,
			payload.ActionNumber,
			time.Unix(payload.ActionStartTime, 0),
		).Return(bonusesData, nil)

		mockFileGenerator.EXPECT().GenerateGeneral(bonusesData).Return([]byte("1"), nil)

		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
		mockTntStorage.EXPECT().Update(
			ctx,
			ReportId(id),
			payload.ActionNumber,
			payload.ActionStartTime,
			ReportStatusInProcess,
			payload.ReportType,
		).Return(errors.New("some error"))

		err := consumer.Consume(ctx, task)
		assert.EqualError(t, err, "can't update report in process status: some error")
	})
	t.Run("should return err if cannot upload file to s3", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbStorage := NewMockBonusStorage(ctrl)
		mockTntStorage := NewMockReportRepository(ctrl)
		mockFileGenerator := NewMockFileGenerator(ctrl)
		mockFileUploader := NewMockFileUploader(ctrl)
		ctx := context.Background()

		ver := version.Must(version.NewVersion("v1.0.0"))
		task := distributed_work.NewSimpleTask(
			BonusTaskType,
			[]byte("{\"reportId\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\","+
				"\"reportType\":1,"+
				" \"actionNumber\":\"num\","+
				"\"actionStartTime\":199999999}"),
			ver,
		)
		consumer := NewBonusConsumer(mockDbStorage, mockTntStorage, zap.NewNop(), ver, mockFileGenerator, mockFileUploader)

		var payload Payload
		_ = json.Unmarshal(task.Payload(), &payload)

		bonusesData := []*BonusGeneral{
			{
				NavActionNumber:               "num",
				Bonus:                         100,
				CountClients:                  1,
				CampaignStartDate:             time.Unix(199999999, 0),
				CampaignFinishDate:            time.Unix(299999999, 0),
				CountClientsWithActiveCard:    1,
				CountClientsSendActivation:    1,
				CountClientsSuccessActivation: 1,
				ActivationPercent:             "100",
			},
		}

		mockDbStorage.EXPECT().FindGeneral(
			ctx,
			payload.ActionNumber,
			time.Unix(payload.ActionStartTime, 0),
		).Return(bonusesData, nil)

		mockFileGenerator.EXPECT().GenerateGeneral(bonusesData).Return([]byte("1"), nil)

		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
		mockTntStorage.EXPECT().Update(
			ctx,
			ReportId(id),
			payload.ActionNumber,
			payload.ActionStartTime,
			ReportStatusInProcess,
			payload.ReportType,
		).Return(nil)

		mockFileUploader.EXPECT().Upload(ctx, bytes.NewReader([]byte("1")), ReportId(id)).Return(errors.New("some error"))

		mockTntStorage.EXPECT().Update(
			ctx,
			ReportId(id),
			payload.ActionNumber,
			payload.ActionStartTime,
			ReportStatusError,
			payload.ReportType,
		).Return(errors.New("update status error"))

		err := consumer.Consume(ctx, task)
		assert.EqualError(t, err, "some error: can't set error status: can't update report: update status error")
	})
	t.Run("should return err if cannot update success status", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbStorage := NewMockBonusStorage(ctrl)
		mockTntStorage := NewMockReportRepository(ctrl)
		mockFileGenerator := NewMockFileGenerator(ctrl)
		mockFileUploader := NewMockFileUploader(ctrl)
		ctx := context.Background()

		ver := version.Must(version.NewVersion("v1.0.0"))
		task := distributed_work.NewSimpleTask(
			BonusTaskType,
			[]byte("{\"reportId\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\","+
				"\"reportType\":1,"+
				" \"actionNumber\":\"num\","+
				"\"actionStartTime\":199999999}"),
			ver,
		)
		consumer := NewBonusConsumer(mockDbStorage, mockTntStorage, zap.NewNop(), ver, mockFileGenerator, mockFileUploader)

		var payload Payload
		_ = json.Unmarshal(task.Payload(), &payload)

		bonusesData := []*BonusGeneral{
			{
				NavActionNumber:               "num",
				Bonus:                         100,
				CountClients:                  1,
				CampaignStartDate:             time.Unix(199999999, 0),
				CampaignFinishDate:            time.Unix(299999999, 0),
				CountClientsWithActiveCard:    1,
				CountClientsSendActivation:    1,
				CountClientsSuccessActivation: 1,
				ActivationPercent:             "100",
			},
		}

		mockDbStorage.EXPECT().
			FindGeneral(ctx, payload.ActionNumber, time.Unix(payload.ActionStartTime, 0)).
			Return(bonusesData, nil)

		mockFileGenerator.EXPECT().GenerateGeneral(bonusesData).Return([]byte("1"), nil)

		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
		mockTntStorage.EXPECT().Update(
			ctx,
			ReportId(id),
			payload.ActionNumber,
			payload.ActionStartTime,
			ReportStatusInProcess,
			payload.ReportType,
		).Return(nil)

		mockFileUploader.EXPECT().Upload(ctx, bytes.NewReader([]byte("1")), ReportId(id)).Return(nil)

		mockTntStorage.EXPECT().Update(
			ctx,
			ReportId(id),
			payload.ActionNumber,
			payload.ActionStartTime,
			ReportStatusSuccess,
			payload.ReportType,
		).Return(errors.New("some error"))

		err := consumer.Consume(ctx, task)
		assert.NoError(t, err)
	})
	t.Run("general success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbStorage := NewMockBonusStorage(ctrl)
		mockTntStorage := NewMockReportRepository(ctrl)
		mockFileGenerator := NewMockFileGenerator(ctrl)
		mockFileUploader := NewMockFileUploader(ctrl)

		ver := version.Must(version.NewVersion("v1.0.0"))
		task := distributed_work.NewSimpleTask(
			BonusTaskType,
			[]byte("{\"reportId\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\","+
				"\"reportType\":1,"+
				" \"actionNumber\":\"num\","+
				"\"actionStartTime\":199999999}"),
			ver,
		)
		consumer := NewBonusConsumer(mockDbStorage, mockTntStorage, zap.NewNop(), ver, mockFileGenerator, mockFileUploader)

		var payload Payload
		_ = json.Unmarshal(task.Payload(), &payload)

		bonusesData := []*BonusGeneral{
			{
				NavActionNumber:               "num",
				Bonus:                         100,
				CountClients:                  1,
				CampaignStartDate:             time.Unix(199999999, 0),
				CampaignFinishDate:            time.Unix(299999999, 0),
				CountClientsWithActiveCard:    1,
				CountClientsSendActivation:    1,
				CountClientsSuccessActivation: 1,
				ActivationPercent:             "100",
			},
		}

		mockDbStorage.EXPECT().FindGeneral(
			context.Background(),
			payload.ActionNumber,
			time.Unix(payload.ActionStartTime, 0),
		).Return(bonusesData, nil)

		mockFileGenerator.EXPECT().GenerateGeneral(bonusesData).Return([]byte("1"), nil)

		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
		mockTntStorage.EXPECT().Update(
			context.Background(),
			ReportId(id),
			payload.ActionNumber,
			payload.ActionStartTime,
			ReportStatusInProcess,
			payload.ReportType,
		).Return(nil)

		mockFileUploader.EXPECT().
			Upload(context.Background(), bytes.NewReader([]byte("1")), ReportId(id)).Return(nil)

		mockTntStorage.EXPECT().Update(
			context.Background(),
			ReportId(id),
			payload.ActionNumber,
			payload.ActionStartTime,
			ReportStatusSuccess,
			payload.ReportType,
		).Return(nil)

		err := consumer.Consume(context.Background(), task)
		assert.NoError(t, err)
	})
	t.Run("should return err if error storage for detailed report", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbStorage := NewMockBonusStorage(ctrl)
		mockTntStorage := NewMockReportRepository(ctrl)
		mockFileGenerator := NewMockFileGenerator(ctrl)
		mockFileUploader := NewMockFileUploader(ctrl)

		ver := version.Must(version.NewVersion("v1.0.0"))
		task := distributed_work.NewSimpleTask(BonusTaskType, []byte("{\"reportId\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\",\"reportType\":2}"), ver)
		consumer := NewBonusConsumer(mockDbStorage, mockTntStorage, zap.NewNop(), ver, mockFileGenerator, mockFileUploader)

		var payload Payload
		_ = json.Unmarshal(task.Payload(), &payload)

		mockDbStorage.EXPECT().FindDetailed(
			context.Background(),
			payload.ActionNumber,
			time.Unix(payload.ActionStartTime, 0),
		).Return(nil, errors.New("some error"))

		err := consumer.Consume(context.Background(), task)
		assert.EqualError(t, err, "failed getting detailed bonuses: some error")
	})
	t.Run("should no err if error file generation for detailed report", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbStorage := NewMockBonusStorage(ctrl)
		mockTntStorage := NewMockReportRepository(ctrl)
		mockFileGenerator := NewMockFileGenerator(ctrl)
		mockFileUploader := NewMockFileUploader(ctrl)

		ver := version.Must(version.NewVersion("v1.0.0"))
		task := distributed_work.NewSimpleTask(
			BonusTaskType,
			[]byte("{\"reportId\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\","+
				"\"reportType\":2,"+
				" \"actionNumber\":\"num\","+
				"\"actionStartTime\":199999999}"),
			ver,
		)
		consumer := NewBonusConsumer(mockDbStorage, mockTntStorage, zap.NewNop(), ver, mockFileGenerator, mockFileUploader)

		var payload Payload
		_ = json.Unmarshal(task.Payload(), &payload)

		ch := make(chan *BonusDetailed, 2)

		mockDbStorage.EXPECT().FindDetailed(
			context.Background(),
			payload.ActionNumber,
			time.Unix(payload.ActionStartTime, 0),
		).Return(ch, nil)

		mockFileGenerator.EXPECT().GenerateDetailed(ch).Return(nil, errors.New("some error"))

		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
		mockTntStorage.EXPECT().Update(
			context.Background(),
			ReportId(id),
			"num",
			int64(199999999),
			ReportStatusError,
			ReportTypeDetailed,
		)

		err := consumer.Consume(context.Background(), task)
		assert.NoError(t, err)
	})
	t.Run("should return err if invalid report type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockDbStorage := NewMockBonusStorage(ctrl)
		mockTntStorage := NewMockReportRepository(ctrl)
		mockFileGenerator := NewMockFileGenerator(ctrl)
		mockFileUploader := NewMockFileUploader(ctrl)

		ver := version.Must(version.NewVersion("v1.0.0"))
		task := distributed_work.NewSimpleTask(BonusTaskType, []byte("{\"reportId\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\",\"reportType\":-398}"), ver)
		consumer := NewBonusConsumer(mockDbStorage, mockTntStorage, zap.NewNop(), ver, mockFileGenerator, mockFileUploader)

		err := consumer.Consume(context.Background(), task)
		assert.NoError(t, err)
	})
}
