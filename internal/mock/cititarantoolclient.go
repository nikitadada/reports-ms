// Code generated by MockGen. DO NOT EDIT.
// Source: code.citik.ru/gobase/tarantool (interfaces: Client)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tarantool "github.com/tarantool/go-tarantool"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Call17 mocks base method.
func (m *MockClient) Call17(arg0 context.Context, arg1 string, arg2 interface{}) (*tarantool.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call17", arg0, arg1, arg2)
	ret0, _ := ret[0].(*tarantool.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call17 indicates an expected call of Call17.
func (mr *MockClientMockRecorder) Call17(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call17", reflect.TypeOf((*MockClient)(nil).Call17), arg0, arg1, arg2)
}

// Call17Async mocks base method.
func (m *MockClient) Call17Async(arg0 context.Context, arg1 string, arg2 interface{}) *tarantool.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call17Async", arg0, arg1, arg2)
	ret0, _ := ret[0].(*tarantool.Future)
	return ret0
}

// Call17Async indicates an expected call of Call17Async.
func (mr *MockClientMockRecorder) Call17Async(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call17Async", reflect.TypeOf((*MockClient)(nil).Call17Async), arg0, arg1, arg2)
}

// Call17Typed mocks base method.
func (m *MockClient) Call17Typed(arg0 context.Context, arg1 string, arg2, arg3 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call17Typed", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Call17Typed indicates an expected call of Call17Typed.
func (mr *MockClientMockRecorder) Call17Typed(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call17Typed", reflect.TypeOf((*MockClient)(nil).Call17Typed), arg0, arg1, arg2, arg3)
}

// Close mocks base method.
func (m *MockClient) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close), arg0)
}

// Ping mocks base method.
func (m *MockClient) Ping(arg0 context.Context) (*tarantool.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(*tarantool.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Ping indicates an expected call of Ping.
func (mr *MockClientMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockClient)(nil).Ping), arg0)
}