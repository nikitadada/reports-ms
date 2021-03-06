// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tarantool/go-tarantool/queue (interfaces: Queue)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	queue "github.com/tarantool/go-tarantool/queue"
)

// MockQueue is a mock of Queue interface.
type MockQueue struct {
	ctrl     *gomock.Controller
	recorder *MockQueueMockRecorder
}

// MockQueueMockRecorder is the mock recorder for MockQueue.
type MockQueueMockRecorder struct {
	mock *MockQueue
}

// NewMockQueue creates a new mock instance.
func NewMockQueue(ctrl *gomock.Controller) *MockQueue {
	mock := &MockQueue{ctrl: ctrl}
	mock.recorder = &MockQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueue) EXPECT() *MockQueueMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockQueue) Create(arg0 queue.Cfg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockQueueMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockQueue)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockQueue) Delete(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockQueueMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockQueue)(nil).Delete), arg0)
}

// Drop mocks base method.
func (m *MockQueue) Drop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Drop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Drop indicates an expected call of Drop.
func (mr *MockQueueMockRecorder) Drop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drop", reflect.TypeOf((*MockQueue)(nil).Drop))
}

// Exists mocks base method.
func (m *MockQueue) Exists() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockQueueMockRecorder) Exists() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockQueue)(nil).Exists))
}

// Kick mocks base method.
func (m *MockQueue) Kick(arg0 uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kick", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Kick indicates an expected call of Kick.
func (mr *MockQueueMockRecorder) Kick(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kick", reflect.TypeOf((*MockQueue)(nil).Kick), arg0)
}

// Peek mocks base method.
func (m *MockQueue) Peek(arg0 uint64) (*queue.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peek", arg0)
	ret0, _ := ret[0].(*queue.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Peek indicates an expected call of Peek.
func (mr *MockQueueMockRecorder) Peek(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peek", reflect.TypeOf((*MockQueue)(nil).Peek), arg0)
}

// Put mocks base method.
func (m *MockQueue) Put(arg0 interface{}) (*queue.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0)
	ret0, _ := ret[0].(*queue.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put.
func (mr *MockQueueMockRecorder) Put(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockQueue)(nil).Put), arg0)
}

// PutWithOpts mocks base method.
func (m *MockQueue) PutWithOpts(arg0 interface{}, arg1 queue.Opts) (*queue.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutWithOpts", arg0, arg1)
	ret0, _ := ret[0].(*queue.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutWithOpts indicates an expected call of PutWithOpts.
func (mr *MockQueueMockRecorder) PutWithOpts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutWithOpts", reflect.TypeOf((*MockQueue)(nil).PutWithOpts), arg0, arg1)
}

// Statistic mocks base method.
func (m *MockQueue) Statistic() (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Statistic")
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Statistic indicates an expected call of Statistic.
func (mr *MockQueueMockRecorder) Statistic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Statistic", reflect.TypeOf((*MockQueue)(nil).Statistic))
}

// Take mocks base method.
func (m *MockQueue) Take() (*queue.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Take")
	ret0, _ := ret[0].(*queue.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Take indicates an expected call of Take.
func (mr *MockQueueMockRecorder) Take() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Take", reflect.TypeOf((*MockQueue)(nil).Take))
}

// TakeTimeout mocks base method.
func (m *MockQueue) TakeTimeout(arg0 time.Duration) (*queue.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TakeTimeout", arg0)
	ret0, _ := ret[0].(*queue.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TakeTimeout indicates an expected call of TakeTimeout.
func (mr *MockQueueMockRecorder) TakeTimeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeTimeout", reflect.TypeOf((*MockQueue)(nil).TakeTimeout), arg0)
}

// TakeTyped mocks base method.
func (m *MockQueue) TakeTyped(arg0 interface{}) (*queue.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TakeTyped", arg0)
	ret0, _ := ret[0].(*queue.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TakeTyped indicates an expected call of TakeTyped.
func (mr *MockQueueMockRecorder) TakeTyped(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeTyped", reflect.TypeOf((*MockQueue)(nil).TakeTyped), arg0)
}

// TakeTypedTimeout mocks base method.
func (m *MockQueue) TakeTypedTimeout(arg0 time.Duration, arg1 interface{}) (*queue.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TakeTypedTimeout", arg0, arg1)
	ret0, _ := ret[0].(*queue.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TakeTypedTimeout indicates an expected call of TakeTypedTimeout.
func (mr *MockQueueMockRecorder) TakeTypedTimeout(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeTypedTimeout", reflect.TypeOf((*MockQueue)(nil).TakeTypedTimeout), arg0, arg1)
}
