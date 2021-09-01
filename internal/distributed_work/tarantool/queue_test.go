package tarantool

import (
	"code.citik.ru/back/report-action/internal"
	"code.citik.ru/back/report-action/internal/distributed_work"
	internal_mock "code.citik.ru/back/report-action/internal/mock"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Insert(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockTntQueue := internal_mock.NewMockQueue(ctrl)
		inserter := NewQueue(mockTntQueue, QueueOpts{})

		mockTntQueue.EXPECT().PutWithOpts(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))

		_, err := inserter.Insert(
			internal.BonusTaskType,
			version.Must(version.NewVersion("v1.0.0")), []byte(""),
			distributed_work.TaskOptions{},
		)

		assert.EqualError(t, err, "fail insert task in tarantool queue: some error")
	})
}

func TestQueue_Take(t *testing.T) {
	t.Run("returns error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockTntQueue := internal_mock.NewMockQueue(ctrl)
		inserter := NewQueue(mockTntQueue, QueueOpts{})

		mockTntQueue.EXPECT().TakeTypedTimeout(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error"))
		_, err := inserter.Take()

		assert.EqualError(t, err, "getting task from the queue: some error")
	})

	t.Run("returns nil", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockTntQueue := internal_mock.NewMockQueue(ctrl)
		inserter := NewQueue(mockTntQueue, QueueOpts{})

		mockTntQueue.EXPECT().TakeTypedTimeout(gomock.Any(), gomock.Any()).Return(nil, nil)
		got, err := inserter.Take()

		assert.Nil(t, got)
		assert.NoError(t, err)
	})
}

func TestQueue_Ack(t *testing.T) {
	t.Run("invalid task type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockTntQueue := internal_mock.NewMockQueue(ctrl)
		inserter := NewQueue(mockTntQueue, QueueOpts{})
		simpleTask := distributed_work.NewSimpleTask(internal.BonusTaskType, []byte("1"), version.Must(version.NewVersion("v1.0.0")))

		err := inserter.Ack(simpleTask)

		assert.EqualError(t, err, "this queue can process only TntTask")
	})
}

func TestQueue_Release(t *testing.T) {
	t.Run("invalid task type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockTntQueue := internal_mock.NewMockQueue(ctrl)
		inserter := NewQueue(mockTntQueue, QueueOpts{})
		simpleTask := distributed_work.NewSimpleTask(internal.BonusTaskType, []byte("1"), version.Must(version.NewVersion("v1.0.0")))

		err := inserter.Release(simpleTask, distributed_work.TaskOptions{})

		assert.EqualError(t, err, "this queue can process only TntTask")
	})
}

func TestQueue_Delete(t *testing.T) {
	t.Run("invalid task type", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockTntQueue := internal_mock.NewMockQueue(ctrl)
		inserter := NewQueue(mockTntQueue, QueueOpts{})
		simpleTask := distributed_work.NewSimpleTask(internal.BonusTaskType, []byte("1"), version.Must(version.NewVersion("v1.0.0")))

		err := inserter.Delete(simpleTask)

		assert.EqualError(t, err, "this queue can process only TntTask")
	})
}
