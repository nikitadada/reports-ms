package grpc

//go:generate mockgen -destination=./mock/report.go -package=mock code.citik.ru/back/report-action/internal ReportRepository

import (
	"code.citik.ru/back/report-action/internal"
	"code.citik.ru/back/report-action/internal/distributed_work"
	bonusv1 "code.citik.ru/back/report-action/internal/grpc/gen/citilink/reportaction/bonus/v1"
	"code.citik.ru/back/report-action/internal/grpc/mock"
	filev1 "code.citik.ru/back/report-action/internal/grpcclient/gen/citilink/cmsfiles/file/v1"
	internal_mock "code.citik.ru/back/report-action/internal/mock"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func TestBonusReportServer_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reportRepository := mock.NewMockReportRepository(ctrl)
	mockQueue := internal_mock.NewMockInserter(ctrl)
	ver := version.Must(version.NewVersion("v1.0.0"))
	cmsFilesClient := internal_mock.NewMockFileAPIClient(ctrl)
	id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")

	s := NewBonusReportServer(
		mockQueue,
		reportRepository,
		ver, cmsFilesClient,
		NewBonusReportToDomainMapper(),
		NewBonusReportToProtobufMapper(),
	)

	type args struct {
		ctx context.Context
		r   *bonusv1.GetRequest
	}
	tests := []struct {
		name string
		args args
		want func() (*bonusv1.GetResponse, error)
	}{
		{
			name: "invalid uuid",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.GetRequest{
					Id: "not uuid",
				},
			},
			want: func() (*bonusv1.GetResponse, error) {
				err := errors.New("invalid UUID length: 8")
				return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("invalid argument id: %s", err))
			},
		},
		{
			name: "cms file client error",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.GetRequest{
					Id: id.String(),
				},
			},
			want: func() (*bonusv1.GetResponse, error) {
				cmsFilesClient.EXPECT().
					Get(context.Background(), &filev1.GetRequest{Id: id.String()}).
					Return(nil, errors.New("some error"))

				return nil, status.Error(codes.Internal, "can't get file: some error")
			},
		},
		{
			name: "storage error",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.GetRequest{
					Id: id.String(),
				},
			},
			want: func() (*bonusv1.GetResponse, error) {
				cmsFilesClient.EXPECT().
					Get(context.Background(), &filev1.GetRequest{Id: id.String()}).
					Return(&filev1.GetResponse{File: &filev1.File{
						Id:           id.String(),
						Name:         "test",
						Type:         filev1.Type_TYPE_BONUS_ACTION,
						Status:       filev1.Status_STATUS_SUCCESS,
						ModifiedTime: &timestamp.Timestamp{Seconds: 0, Nanos: 0},
					}}, nil)

				reportRepository.EXPECT().
					Get(
						context.Background(),
						internal.ReportId(id),
					).Return(nil, errors.New("some error")).Times(1)

				return nil, status.Error(codes.Internal, "can't get report: some error")
			},
		},
		{
			name: "not found",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.GetRequest{
					Id: id.String(),
				},
			},
			want: func() (*bonusv1.GetResponse, error) {
				cmsFilesClient.EXPECT().
					Get(context.Background(), &filev1.GetRequest{Id: id.String()}).
					Return(&filev1.GetResponse{File: &filev1.File{
						Id:           id.String(),
						Name:         "test",
						Type:         filev1.Type_TYPE_BONUS_ACTION,
						Status:       filev1.Status_STATUS_SUCCESS,
						ModifiedTime: &timestamp.Timestamp{Seconds: 0, Nanos: 0},
					}}, nil)

				reportRepository.EXPECT().
					Get(
						context.Background(),
						internal.ReportId(id),
					).Return(nil, nil).Times(1)

				return nil, status.Error(codes.NotFound, "report not found")
			},
		},
		{
			name: "mapping error",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.GetRequest{
					Id: id.String(),
				},
			},
			want: func() (*bonusv1.GetResponse, error) {
				cmsFilesClient.EXPECT().
					Get(context.Background(), &filev1.GetRequest{Id: id.String()}).
					Return(&filev1.GetResponse{File: &filev1.File{
						Id:           id.String(),
						Name:         "test",
						Type:         filev1.Type_TYPE_BONUS_ACTION,
						Status:       filev1.Status_STATUS_SUCCESS,
						ModifiedTime: &timestamp.Timestamp{Seconds: 0, Nanos: 0},
					}}, nil)

				reportRepository.EXPECT().
					Get(
						context.Background(),
						internal.ReportId(id),
					).
					Return(&internal.Report{
						Id: internal.ReportId(id),
					}, nil).Times(1)

				return nil, status.Error(codes.Internal, "can't map report: can't map type: type is invalid")
			},
		},
		{
			name: "correct",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.GetRequest{
					Id: id.String(),
				},
			},
			want: func() (*bonusv1.GetResponse, error) {
				cmsFilesClient.EXPECT().
					Get(context.Background(), &filev1.GetRequest{Id: id.String()}).
					Return(&filev1.GetResponse{File: &filev1.File{
						Id:           id.String(),
						Name:         "test",
						Type:         filev1.Type_TYPE_BONUS_ACTION,
						Status:       filev1.Status_STATUS_SUCCESS,
						ModifiedTime: &timestamp.Timestamp{Seconds: 0, Nanos: 0},
					}}, nil)

				reportRepository.EXPECT().
					Get(
						context.Background(),
						internal.ReportId(id),
					).
					Return(&internal.Report{
						Id:                internal.ReportId(id),
						NavActionNumber:   "num",
						Type:              internal.ReportTypeDetailed,
						Status:            internal.ReportStatusSuccess,
						FileName:          "test",
						CampaignStartDate: time.Unix(0, 0),
						LastModified:      time.Unix(0, 0),
					}, nil).Times(1)

				return &bonusv1.GetResponse{
					Info: &bonusv1.Info{
						Id:          id.String(),
						Type:        bonusv1.Type_TYPE_DETAILED,
						Name:        "test",
						Status:      bonusv1.Status_STATUS_SUCCESS,
						CreatedTime: &timestamp.Timestamp{Seconds: 0, Nanos: 0},
					},
				}, nil
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			want, wantErr := tt.want()
			got, err := s.Get(tt.args.ctx, tt.args.r)
			if wantErr != nil {
				assert.Error(t, err)
				assert.EqualErrorf(t, err, wantErr.Error(), err.Error())
				assert.Nil(t, got)
			} else {
				assert.IsType(t, want, got)
				assert.Nil(t, err)
			}
		})
	}
}

func TestBonusReportServer_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reportRepository := mock.NewMockReportRepository(ctrl)
	mockQueue := internal_mock.NewMockInserter(ctrl)
	ver := version.Must(version.NewVersion("v1.0.0"))
	cmsFilesClient := internal_mock.NewMockFileAPIClient(ctrl)

	s := NewBonusReportServer(
		mockQueue,
		reportRepository,
		ver, cmsFilesClient,
		NewBonusReportToDomainMapper(),
		NewBonusReportToProtobufMapper(),
	)

	type args struct {
		ctx context.Context
		r   *bonusv1.CreateRequest
	}
	tests := []struct {
		name string
		args args
		want func() (*bonusv1.CreateResponse, error)
	}{
		{
			name: "empty action num",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.CreateRequest{
					NavisionActionNumber: "",
					ActionStartTime:      &timestamp.Timestamp{Seconds: 0, Nanos: 0},
					Type:                 bonusv1.Type_TYPE_DETAILED,
				},
			},
			want: func() (*bonusv1.CreateResponse, error) {
				return nil, status.Error(codes.InvalidArgument, "navision action number is required")
			},
		},
		{
			name: "invalid start time",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.CreateRequest{
					NavisionActionNumber: "test num",
					ActionStartTime:      &timestamp.Timestamp{Seconds: -1, Nanos: 0},
					Type:                 bonusv1.Type_TYPE_DETAILED,
				},
			},
			want: func() (*bonusv1.CreateResponse, error) {
				return nil, status.Error(codes.InvalidArgument, "invalid start time")
			},
		},
		{
			name: "cms file client error",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.CreateRequest{
					NavisionActionNumber: "test num",
					ActionStartTime:      &timestamp.Timestamp{Seconds: 946674001, Nanos: 0},
					Type:                 bonusv1.Type_TYPE_DETAILED,
				},
			},
			want: func() (*bonusv1.CreateResponse, error) {
				reportRepository.EXPECT().
					Filter(context.Background(), gomock.Any()).Return(nil, errors.New("some error"))

				return nil, status.Error(codes.Internal, "can't find report: some error")
			},
		},
		{
			name: "map type error",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.CreateRequest{
					NavisionActionNumber: "test num",
					ActionStartTime:      &timestamp.Timestamp{Seconds: 946674001, Nanos: 0},
					Type:                 bonusv1.Type(99999999),
				},
			},
			want: func() (*bonusv1.CreateResponse, error) {
				return nil, status.Error(codes.Internal, "can't map type: type is invalid")
			},
		},
		{
			name: "cms files client error",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.CreateRequest{
					NavisionActionNumber: "test num",
					ActionStartTime:      &timestamp.Timestamp{Seconds: 946674001, Nanos: 0},
					Type:                 bonusv1.Type_TYPE_DETAILED,
				},
			},
			want: func() (*bonusv1.CreateResponse, error) {
				reportRepository.EXPECT().Filter(context.Background(), gomock.Any()).Return(nil, nil)
				cmsFilesClient.EXPECT().Create(context.Background(), gomock.Any()).Return(nil, errors.New("some error"))

				return nil, status.Error(codes.Internal, "can't create file in microservice cms-files: some error")
			},
		},
		{
			name: "invalid uuid",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.CreateRequest{
					NavisionActionNumber: "test num",
					ActionStartTime:      &timestamp.Timestamp{Seconds: 946674001, Nanos: 0},
					Type:                 bonusv1.Type_TYPE_DETAILED,
				},
			},
			want: func() (*bonusv1.CreateResponse, error) {
				reportRepository.EXPECT().Filter(context.Background(), gomock.Any()).Return(nil, nil)
				cmsFilesClient.EXPECT().Create(context.Background(), gomock.Any()).Return(
					&filev1.CreateResponse{
						Id: "not uuid",
					}, nil)

				return nil, status.Error(codes.Internal, "can't parse file id: invalid UUID length: 8")
			},
		},
		{
			name: "storage error",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.CreateRequest{
					NavisionActionNumber: "test num",
					ActionStartTime:      &timestamp.Timestamp{Seconds: 1624320000, Nanos: 0},
					Type:                 bonusv1.Type_TYPE_DETAILED,
				},
			},
			want: func() (*bonusv1.CreateResponse, error) {
				id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")

				reportRepository.EXPECT().Filter(context.Background(), gomock.Any()).Return(nil, nil)
				cmsFilesClient.EXPECT().Create(context.Background(), gomock.Any()).Return(
					&filev1.CreateResponse{
						Id: id.String(),
					}, nil)
				campaignStart := timestamp.Timestamp{Seconds: 1624320000, Nanos: 0}
				reportRepository.EXPECT().Create(
					context.Background(),
					&internal.Report{
						Id:                internal.ReportId(id),
						NavActionNumber:   "test num",
						CampaignStartDate: campaignStart.AsTime().Local(),
						Status:            internal.ReportStatusCreated,
						Type:              internal.ReportTypeDetailed,
					},
				).Return(errors.New("some error"))

				return nil, status.Error(codes.Internal, "can't save report: some error")
			},
		},
		{
			name: "insert to queue error",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.CreateRequest{
					NavisionActionNumber: "test num",
					ActionStartTime:      &timestamp.Timestamp{Seconds: 1624320000, Nanos: 0},
					Type:                 bonusv1.Type_TYPE_DETAILED,
				},
			},
			want: func() (*bonusv1.CreateResponse, error) {
				id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")

				reportRepository.EXPECT().Filter(context.Background(), gomock.Any()).Return(nil, nil)
				cmsFilesClient.EXPECT().Create(context.Background(), gomock.Any()).Return(
					&filev1.CreateResponse{
						Id: id.String(),
					}, nil)
				reportRepository.EXPECT().Create(
					context.Background(),
					&internal.Report{
						Id:                internal.ReportId(id),
						NavActionNumber:   "test num",
						CampaignStartDate: time.Unix(1624320000, 0),
						Status:            internal.ReportStatusCreated,
						Type:              internal.ReportTypeDetailed,
					},
				).Return(nil)

				payload, _ := json.Marshal(internal.Payload{
					ReportId:        id.String(),
					ActionNumber:    "test num",
					ActionStartTime: int64(1624320000),
					ReportType:      internal.ReportTypeDetailed,
				})

				mockQueue.EXPECT().
					Insert(internal.BonusTaskType, s.appVersion, payload, distributed_work.TaskOptions{}).
					Return("", errors.New("some error"))

				return nil, status.Error(codes.Internal, "can't insert task: some error")
			},
		},
		{
			name: "report already exists",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.CreateRequest{
					NavisionActionNumber: "test num",
					ActionStartTime:      &timestamp.Timestamp{Seconds: 946674001, Nanos: 0},
					Type:                 bonusv1.Type(2),
				},
			},
			want: func() (*bonusv1.CreateResponse, error) {
				id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
				reportRepository.EXPECT().Filter(context.Background(), gomock.Any()).Return([]*internal.Report{
					{
						Id:              internal.ReportId(id),
						NavActionNumber: "test num",
						Status:          internal.ReportStatusSuccess,
						Type:            internal.ReportTypeDetailed,
						LastModified:    time.Unix(0, 0),
					},
				}, nil)

				return &bonusv1.CreateResponse{Info: &bonusv1.Info{
					Id:          id.String(),
					Type:        bonusv1.Type_TYPE_DETAILED,
					Name:        "",
					Status:      bonusv1.Status_STATUS_SUCCESS,
					CreatedTime: &timestamp.Timestamp{Seconds: 0, Nanos: 0},
				}}, nil
			},
		},
		{
			name: "report already exists error map",
			args: args{
				ctx: context.Background(),
				r: &bonusv1.CreateRequest{
					NavisionActionNumber: "test num",
					ActionStartTime:      &timestamp.Timestamp{Seconds: 946674001, Nanos: 0},
					Type:                 bonusv1.Type(2),
				},
			},
			want: func() (*bonusv1.CreateResponse, error) {
				id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
				reportRepository.EXPECT().Filter(context.Background(), gomock.Any()).Return([]*internal.Report{
					{
						Id:              internal.ReportId(id),
						NavActionNumber: "test num",
						Status:          internal.ReportStatus(-999),
						Type:            internal.ReportTypeDetailed,
						LastModified:    time.Unix(0, 0),
					},
				}, nil)

				return nil, status.Error(codes.Internal, "can't map bonus report: can't map status: status is invalid")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			want, wantErr := tt.want()
			got, err := s.Create(tt.args.ctx, tt.args.r)
			if wantErr != nil {
				assert.Error(t, err)
				assert.EqualErrorf(t, err, wantErr.Error(), err.Error())
				assert.Nil(t, got)
			} else {
				assert.IsType(t, want, got)
				assert.Nil(t, err)
			}
		})
	}
}
