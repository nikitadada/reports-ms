package distributed_work

//go:generate mockgen -destination=./queue_mock_test.go -package=distributed_work code.citik.ru/back/report-action/internal/distributed_work Queue,Consumer,Task

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-version"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestConsumerDistributor_Start(t *testing.T) {
	t.Run("should wait after received error on task take", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx, cancel := context.WithTimeout(context.Background(), 70*time.Millisecond)
		defer cancel()

		mockQueue := NewMockQueue(ctrl)
		mockQueue.EXPECT().Take().Return(nil, errors.New("some error"))

		NewConsumerDistributor(mockQueue,
			[]Consumer{},
			zap.NewNop(), ConsumerDistributorOpts{SameTimeExecutionLimit: 1},
		).Start(ctx)
		<-ctx.Done()
	})

	t.Run("should delete task if consumer version verdict is delete", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		defer cancel()

		mockConsumer := NewMockConsumer(ctrl)
		mockQueue := NewMockQueue(ctrl)

		payload := []byte("{\"reportId\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\"," +
			"\"reportType\":1," +
			" \"actionNumber\":\"num\"," +
			"\"actionStartTime\":199999999}")

		task := NewSimpleTask("bonus", payload, version.Must(version.NewVersion("v1.0.0")))

		mockQueue.EXPECT().Take().Return(task, nil).AnyTimes()
		mockConsumer.EXPECT().CanConsume(task).Return(true).AnyTimes()
		mockConsumer.EXPECT().ResolveVersion(version.Must(version.NewVersion("v1.0.0"))).Return(VersionResolverVerdictDelete).AnyTimes()
		mockQueue.EXPECT().Delete(task).Return(errors.New("some error")).AnyTimes()

		NewConsumerDistributor(mockQueue,
			[]Consumer{mockConsumer},
			zap.NewNop(), ConsumerDistributorOpts{SameTimeExecutionLimit: 1},
		).Start(ctx)
		<-ctx.Done()
	})

	t.Run("should release task if consumer version verdict is release delay", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		defer cancel()

		mockConsumer := NewMockConsumer(ctrl)
		mockQueue := NewMockQueue(ctrl)

		payload := []byte("{\"reportId\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\"," +
			"\"reportType\":1," +
			" \"actionNumber\":\"num\"," +
			"\"actionStartTime\":199999999}")

		task := NewSimpleTask("bonus", payload, version.Must(version.NewVersion("v1.0.0")))

		mockQueue.EXPECT().Take().Return(task, nil).AnyTimes()
		mockConsumer.EXPECT().CanConsume(task).Return(true).AnyTimes()
		mockConsumer.EXPECT().ResolveVersion(version.Must(version.NewVersion("v1.0.0"))).Return(VersionResolverVerdictReleaseWithDelay).AnyTimes()
		mockQueue.EXPECT().Release(task, TaskOptions{Delay: 5 * time.Second}).Return(errors.New("some error")).AnyTimes()

		NewConsumerDistributor(mockQueue,
			[]Consumer{mockConsumer},
			zap.NewNop(), ConsumerDistributorOpts{SameTimeExecutionLimit: 1},
		).Start(ctx)
		<-ctx.Done()
	})

	t.Run("consume error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		defer cancel()

		mockConsumer := NewMockConsumer(ctrl)
		mockQueue := NewMockQueue(ctrl)

		payload := []byte("{\"reportId\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\"," +
			"\"reportType\":1," +
			" \"actionNumber\":\"num\"," +
			"\"actionStartTime\":199999999}")

		task := NewSimpleTask("bonus", payload, version.Must(version.NewVersion("v1.0.0")))

		mockQueue.EXPECT().Take().Return(task, nil).AnyTimes()
		mockConsumer.EXPECT().CanConsume(task).Return(true).AnyTimes()
		mockConsumer.EXPECT().ResolveVersion(version.Must(version.NewVersion("v1.0.0"))).Return(VersionResolverVerdictOK).AnyTimes()
		mockConsumer.EXPECT().Consume(ctx, task).Return(errors.New("some error")).AnyTimes()
		mockQueue.EXPECT().Release(task, TaskOptions{}).Return(errors.New("some error")).AnyTimes()

		NewConsumerDistributor(mockQueue,
			[]Consumer{mockConsumer},
			zap.NewNop(), ConsumerDistributorOpts{SameTimeExecutionLimit: 1},
		).Start(ctx)
		<-ctx.Done()
	})

	t.Run("consume success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		defer cancel()

		mockConsumer := NewMockConsumer(ctrl)
		mockQueue := NewMockQueue(ctrl)

		payload := []byte("{\"reportId\":\"40d206a2-06d1-46cb-9c05-e3540c1eb0cf\"," +
			"\"reportType\":1," +
			" \"actionNumber\":\"num\"," +
			"\"actionStartTime\":199999999}")

		task := NewSimpleTask("bonus", payload, version.Must(version.NewVersion("v1.0.0")))

		mockQueue.EXPECT().Take().Return(task, nil).AnyTimes()
		mockConsumer.EXPECT().CanConsume(task).Return(true).AnyTimes()
		mockConsumer.EXPECT().ResolveVersion(version.Must(version.NewVersion("v1.0.0"))).Return(VersionResolverVerdictOK).AnyTimes()
		mockConsumer.EXPECT().Consume(ctx, task).Return(nil).AnyTimes()
		mockQueue.EXPECT().Ack(task).Return(errors.New("some error")).AnyTimes()

		NewConsumerDistributor(mockQueue,
			[]Consumer{mockConsumer},
			zap.NewNop(), ConsumerDistributorOpts{SameTimeExecutionLimit: 1},
		).Start(ctx)
		<-ctx.Done()
	})
}
