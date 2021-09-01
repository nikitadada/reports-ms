package distributed_work

import (
	"context"
	"github.com/hashicorp/go-version"
	"go.uber.org/zap"
	"strings"
	"time"
)

const (
	VersionResolverVerdictOK VersionResolverVerdict = iota
	VersionResolverVerdictDelete
	VersionResolverVerdictReleaseWithDelay
)

type VersionResolverVerdict int

type VersionResolver interface {
	Resolve(version *version.Version) VersionResolverVerdict
}

// Consumer Исполнитель задания
type Consumer interface {
	CanConsume(task Task) bool
	Consume(ctx context.Context, task Task) error
	ResolveVersion(v *version.Version) VersionResolverVerdict
}

type ConsumeDistributor struct {
	queue     Queue
	consumers []Consumer
	logger    *zap.Logger
	opts      ConsumerDistributorOpts
}

type ConsumerDistributorOpts struct {
	SameTimeExecutionLimit int
}

func NewConsumerDistributor(
	queue Queue,
	consumers []Consumer,
	logger *zap.Logger,
	opts ConsumerDistributorOpts,
) *ConsumeDistributor {
	if opts.SameTimeExecutionLimit == 0 {
		opts.SameTimeExecutionLimit = 1
	}
	c := &ConsumeDistributor{queue: queue, consumers: consumers, logger: logger, opts: opts}
	return c
}

func (c *ConsumeDistributor) Start(ctx context.Context) {
	worksChan := make(chan bool, c.opts.SameTimeExecutionLimit)
	for {
		select {
		case <-ctx.Done():
			return
		case worksChan <- true:
		}
		go func() {
			defer func() { <-worksChan }()
			task, err := c.queue.Take()
			if err != nil {
				c.logger.Error("error taking task", zap.Error(err))
				select {
				case <-ctx.Done():
				case <-time.After(2 * time.Second):
				}
				return
			}

			if task == nil {
				return
			}

			logger := c.logger.With(zap.String("task_type", string(task.Type())))
			logger.Debug("got task from queue", zap.String("task_version", task.Version().String()),
				zap.ByteString("task_payload", task.Payload()))

			var applicableConsumer Consumer
			for _, e := range c.consumers {
				if !e.CanConsume(task) {
					continue
				}

				applicableConsumer = e
				break
			}

			if applicableConsumer == nil {
				logger.Error("consumer not found, task will not be completed")
				err := c.queue.Release(task, TaskOptions{Delay: 5 * time.Second})
				if err != nil {
					logger.Error("fail release task", zap.Error(err))
				}

				return
			}

			versionResolveVerdict := applicableConsumer.ResolveVersion(task.Version())
			switch versionResolveVerdict {
			case VersionResolverVerdictDelete:
				err := c.queue.Delete(task)
				if err != nil {
					logger.Error("can't delete task after version verdict", zap.Error(err))
				}
				return
			case VersionResolverVerdictReleaseWithDelay:
				err := c.queue.Release(task, TaskOptions{Delay: 5 * time.Second})
				if err != nil {
					logger.Error("can't release task after version verdict", zap.Error(err))
				}
				return
			}

			err = applicableConsumer.Consume(ctx, task)
			if err != nil {
				logger.Error("error while executing task, releasing task", zap.Error(err))
				err := c.queue.Release(task, TaskOptions{})
				if err != nil {
					logger.Error("can't release task", zap.Error(err))
				}
			} else {
				err = c.queue.Ack(task)
				if err != nil {
					// Глушим ошибку связанную с ошибкой в очередях в tarantool модуле queue
					// Issue с описанием здесь https://github.com/tarantool/queue/issues/146
					if !strings.Contains(err.Error(), "Task was not taken in the session") {
						logger.Error("fail ack task", zap.Error(err))
					}
				}
			}
		}()
	}
}
