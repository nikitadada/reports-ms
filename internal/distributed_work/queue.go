package distributed_work

import (
	"github.com/hashicorp/go-version"
	"time"
)

type TaskType string

// Task Задание
type Task interface {
	Type() TaskType
	Payload() []byte
	Version() *version.Version
}

func NewSimpleTask(taskType TaskType, payload []byte, ver *version.Version) *SimpleTask {
	return &SimpleTask{taskType: taskType, payload: payload, ver: ver}
}

type SimpleTask struct {
	taskType TaskType
	payload  []byte
	ver      *version.Version
}

func (st *SimpleTask) Type() TaskType {
	return st.taskType
}

func (st *SimpleTask) Payload() []byte {
	return st.payload
}

func (st *SimpleTask) Version() *version.Version {
	return st.ver
}

type TaskOptions struct {
	Ttl      time.Duration
	Ttr      time.Duration
	Priority int
	Delay    time.Duration
}

func (to TaskOptions) IsEmpty() bool {
	return to.Ttl == 0 && to.Ttr == 0 && to.Priority == 0 && to.Delay == 0
}

type Queue interface {
	Inserter
	Take() (Task, error)
	Ack(task Task) error
	Release(task Task, opts TaskOptions) error
	Delete(task Task) error
}

type Inserter interface {
	Insert(taskType TaskType, version *version.Version, payload []byte, options TaskOptions) (string, error)
}
