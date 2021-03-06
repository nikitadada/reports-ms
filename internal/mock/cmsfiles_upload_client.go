// Code generated by MockGen. DO NOT EDIT.
// Source: code.citik.ru/back/report-action/internal/grpcclient/gen/citilink/cmsfiles/file/v1 (interfaces: FileAPI_UploadClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	filev1 "code.citik.ru/back/report-action/internal/grpcclient/gen/citilink/cmsfiles/file/v1"
	gomock "github.com/golang/mock/gomock"
	metadata "google.golang.org/grpc/metadata"
)

// MockFileAPI_UploadClient is a mock of FileAPI_UploadClient interface.
type MockFileAPI_UploadClient struct {
	ctrl     *gomock.Controller
	recorder *MockFileAPI_UploadClientMockRecorder
}

// MockFileAPI_UploadClientMockRecorder is the mock recorder for MockFileAPI_UploadClient.
type MockFileAPI_UploadClientMockRecorder struct {
	mock *MockFileAPI_UploadClient
}

// NewMockFileAPI_UploadClient creates a new mock instance.
func NewMockFileAPI_UploadClient(ctrl *gomock.Controller) *MockFileAPI_UploadClient {
	mock := &MockFileAPI_UploadClient{ctrl: ctrl}
	mock.recorder = &MockFileAPI_UploadClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileAPI_UploadClient) EXPECT() *MockFileAPI_UploadClientMockRecorder {
	return m.recorder
}

// CloseAndRecv mocks base method.
func (m *MockFileAPI_UploadClient) CloseAndRecv() (*filev1.UploadResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAndRecv")
	ret0, _ := ret[0].(*filev1.UploadResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAndRecv indicates an expected call of CloseAndRecv.
func (mr *MockFileAPI_UploadClientMockRecorder) CloseAndRecv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAndRecv", reflect.TypeOf((*MockFileAPI_UploadClient)(nil).CloseAndRecv))
}

// CloseSend mocks base method.
func (m *MockFileAPI_UploadClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockFileAPI_UploadClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockFileAPI_UploadClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockFileAPI_UploadClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockFileAPI_UploadClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockFileAPI_UploadClient)(nil).Context))
}

// Header mocks base method.
func (m *MockFileAPI_UploadClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockFileAPI_UploadClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockFileAPI_UploadClient)(nil).Header))
}

// RecvMsg mocks base method.
func (m *MockFileAPI_UploadClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockFileAPI_UploadClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockFileAPI_UploadClient)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockFileAPI_UploadClient) Send(arg0 *filev1.UploadRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockFileAPI_UploadClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockFileAPI_UploadClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m *MockFileAPI_UploadClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockFileAPI_UploadClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockFileAPI_UploadClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockFileAPI_UploadClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockFileAPI_UploadClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockFileAPI_UploadClient)(nil).Trailer))
}
