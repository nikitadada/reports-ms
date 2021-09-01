package tarantool

import (
	"code.citik.ru/back/report-action/internal/distributed_work"
	"fmt"
	"github.com/hashicorp/go-version"
	"github.com/tarantool/go-tarantool/queue"
	"strconv"
	"time"
)

type TntTask struct {
	taskType distributed_work.TaskType
	version  *version.Version
	payload  []byte
	tntTask  *queue.Task
}

func (t *TntTask) Type() distributed_work.TaskType {
	return t.taskType
}

func (t *TntTask) Payload() []byte {
	return t.payload
}

func (t *TntTask) Version() *version.Version {
	return t.version
}

type task struct {
	TaskType string
	Payload  []byte
	Version  string
}

type Queue struct {
	tarantoolQueue queue.Queue
	opts           QueueOpts
}

type QueueOpts struct {
	DefaultTaskOptions distributed_work.TaskOptions
	TakeTaskTimeout    time.Duration
}

func NewQueue(tarantoolQueue queue.Queue, opts QueueOpts) *Queue {
	if opts.TakeTaskTimeout == 0 {
		opts.TakeTaskTimeout = 1 * time.Minute
	}

	return &Queue{tarantoolQueue: tarantoolQueue, opts: opts}
}

func (q *Queue) Take() (distributed_work.Task, error) {
	taskBody := task{}
	tntTask, err := q.tarantoolQueue.TakeTypedTimeout(q.opts.TakeTaskTimeout, &taskBody)
	if err != nil {
		return nil, fmt.Errorf("getting task from the queue: %w", err)
	}
	if tntTask == nil {
		return nil, nil
	}

	ver, err := version.NewVersion(taskBody.Version)
	if err != nil {
		ver = version.Must(version.NewVersion("v1.0.0"))
	}

	return &TntTask{
		taskType: distributed_work.TaskType(taskBody.TaskType),
		payload:  taskBody.Payload,
		tntTask:  tntTask,
		version:  ver,
	}, nil
}

func (q *Queue) Insert(
	taskType distributed_work.TaskType,
	version *version.Version,
	payload []byte,
	options distributed_work.TaskOptions,
) (string, error) {
	taskBody := task{
		TaskType: string(taskType),
		Payload:  payload,
		Version:  version.String(),
	}

	opts := mergeQueueOpts(convertOpts(q.opts.DefaultTaskOptions), convertOpts(options))
	tntTask, err := q.tarantoolQueue.PutWithOpts(taskBody, opts)
	if err != nil {
		return "", fmt.Errorf("fail insert task in tarantool queue: %w", err)
	}

	return strconv.Itoa(int(tntTask.Id())), nil
}

func (q *Queue) Ack(task distributed_work.Task) error {
	var suitableTask *TntTask
	var ok bool
	if suitableTask, ok = task.(*TntTask); !ok {
		return fmt.Errorf("this queue can process only TntTask")
	}

	err := suitableTask.tntTask.Ack()
	if err != nil {
		return fmt.Errorf("fail ack task in tarantool: %w", err)
	}

	return nil
}

func (q *Queue) Release(task distributed_work.Task, opts distributed_work.TaskOptions) error {
	var suitableTask *TntTask
	var ok bool
	if suitableTask, ok = task.(*TntTask); !ok {
		return fmt.Errorf("this queue can process only TntTask")
	}

	err := suitableTask.tntTask.ReleaseCfg(convertOpts(opts))
	if err != nil {
		return fmt.Errorf("fail release task in tarantool: %w", err)
	}

	return nil
}

func (q *Queue) Delete(task distributed_work.Task) error {
	var suitableTask *TntTask
	var ok bool
	if suitableTask, ok = task.(*TntTask); !ok {
		return fmt.Errorf("this queue can process only TntTask")
	}

	err := suitableTask.tntTask.Delete()
	if err != nil {
		return fmt.Errorf("fail delete task in tarantool: %w", err)
	}

	return nil
}

func convertOpts(opts distributed_work.TaskOptions) queue.Opts {
	return queue.Opts{
		Pri:   opts.Priority,
		Ttl:   opts.Ttl,
		Ttr:   opts.Ttr,
		Delay: opts.Delay,
		Utube: "",
	}
}

func mergeQueueOpts(opts1 queue.Opts, opts2 queue.Opts) queue.Opts {
	if opts2.Pri != 0 {
		opts1.Pri = opts2.Pri
	}
	if opts2.Ttl != 0 {
		opts1.Ttl = opts2.Ttl
	}
	if opts2.Ttr != 0 {
		opts1.Ttr = opts2.Ttr
	}
	if opts2.Delay != 0 {
		opts1.Delay = opts2.Delay
	}

	return opts1
}
