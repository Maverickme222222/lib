// Code generated by MockGen. DO NOT EDIT.
// Source: ../tasks_grpc.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tasks "github.com/kappapay/backend/lib/golang/src/kappa/services/tasks"
	grpc "google.golang.org/grpc"
)

// MockTasksServiceClient is a mock of TasksServiceClient interface.
type MockTasksServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTasksServiceClientMockRecorder
}

// MockTasksServiceClientMockRecorder is the mock recorder for MockTasksServiceClient.
type MockTasksServiceClientMockRecorder struct {
	mock *MockTasksServiceClient
}

// NewMockTasksServiceClient creates a new mock instance.
func NewMockTasksServiceClient(ctrl *gomock.Controller) *MockTasksServiceClient {
	mock := &MockTasksServiceClient{ctrl: ctrl}
	mock.recorder = &MockTasksServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTasksServiceClient) EXPECT() *MockTasksServiceClientMockRecorder {
	return m.recorder
}

// GetTasksByEntityID mocks base method.
func (m *MockTasksServiceClient) GetTasksByEntityID(ctx context.Context, in *tasks.GetTasksByEntityIDRequest, opts ...grpc.CallOption) (*tasks.GetTasksByEntityIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTasksByEntityID", varargs...)
	ret0, _ := ret[0].(*tasks.GetTasksByEntityIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksByEntityID indicates an expected call of GetTasksByEntityID.
func (mr *MockTasksServiceClientMockRecorder) GetTasksByEntityID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksByEntityID", reflect.TypeOf((*MockTasksServiceClient)(nil).GetTasksByEntityID), varargs...)
}

// GetTasksByUserID mocks base method.
func (m *MockTasksServiceClient) GetTasksByUserID(ctx context.Context, in *tasks.GetTasksByUserIDRequest, opts ...grpc.CallOption) (*tasks.GetTasksByUserIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTasksByUserID", varargs...)
	ret0, _ := ret[0].(*tasks.GetTasksByUserIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksByUserID indicates an expected call of GetTasksByUserID.
func (mr *MockTasksServiceClientMockRecorder) GetTasksByUserID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksByUserID", reflect.TypeOf((*MockTasksServiceClient)(nil).GetTasksByUserID), varargs...)
}

// MockTasksServiceServer is a mock of TasksServiceServer interface.
type MockTasksServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockTasksServiceServerMockRecorder
}

// MockTasksServiceServerMockRecorder is the mock recorder for MockTasksServiceServer.
type MockTasksServiceServerMockRecorder struct {
	mock *MockTasksServiceServer
}

// NewMockTasksServiceServer creates a new mock instance.
func NewMockTasksServiceServer(ctrl *gomock.Controller) *MockTasksServiceServer {
	mock := &MockTasksServiceServer{ctrl: ctrl}
	mock.recorder = &MockTasksServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTasksServiceServer) EXPECT() *MockTasksServiceServerMockRecorder {
	return m.recorder
}

// GetTasksByEntityID mocks base method.
func (m *MockTasksServiceServer) GetTasksByEntityID(arg0 context.Context, arg1 *tasks.GetTasksByEntityIDRequest) (*tasks.GetTasksByEntityIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasksByEntityID", arg0, arg1)
	ret0, _ := ret[0].(*tasks.GetTasksByEntityIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksByEntityID indicates an expected call of GetTasksByEntityID.
func (mr *MockTasksServiceServerMockRecorder) GetTasksByEntityID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksByEntityID", reflect.TypeOf((*MockTasksServiceServer)(nil).GetTasksByEntityID), arg0, arg1)
}

// GetTasksByUserID mocks base method.
func (m *MockTasksServiceServer) GetTasksByUserID(arg0 context.Context, arg1 *tasks.GetTasksByUserIDRequest) (*tasks.GetTasksByUserIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasksByUserID", arg0, arg1)
	ret0, _ := ret[0].(*tasks.GetTasksByUserIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksByUserID indicates an expected call of GetTasksByUserID.
func (mr *MockTasksServiceServerMockRecorder) GetTasksByUserID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksByUserID", reflect.TypeOf((*MockTasksServiceServer)(nil).GetTasksByUserID), arg0, arg1)
}

// mustEmbedUnimplementedTasksServiceServer mocks base method.
func (m *MockTasksServiceServer) mustEmbedUnimplementedTasksServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTasksServiceServer")
}

// mustEmbedUnimplementedTasksServiceServer indicates an expected call of mustEmbedUnimplementedTasksServiceServer.
func (mr *MockTasksServiceServerMockRecorder) mustEmbedUnimplementedTasksServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTasksServiceServer", reflect.TypeOf((*MockTasksServiceServer)(nil).mustEmbedUnimplementedTasksServiceServer))
}

// MockUnsafeTasksServiceServer is a mock of UnsafeTasksServiceServer interface.
type MockUnsafeTasksServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeTasksServiceServerMockRecorder
}

// MockUnsafeTasksServiceServerMockRecorder is the mock recorder for MockUnsafeTasksServiceServer.
type MockUnsafeTasksServiceServerMockRecorder struct {
	mock *MockUnsafeTasksServiceServer
}

// NewMockUnsafeTasksServiceServer creates a new mock instance.
func NewMockUnsafeTasksServiceServer(ctrl *gomock.Controller) *MockUnsafeTasksServiceServer {
	mock := &MockUnsafeTasksServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeTasksServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeTasksServiceServer) EXPECT() *MockUnsafeTasksServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedTasksServiceServer mocks base method.
func (m *MockUnsafeTasksServiceServer) mustEmbedUnimplementedTasksServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedTasksServiceServer")
}

// mustEmbedUnimplementedTasksServiceServer indicates an expected call of mustEmbedUnimplementedTasksServiceServer.
func (mr *MockUnsafeTasksServiceServerMockRecorder) mustEmbedUnimplementedTasksServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedTasksServiceServer", reflect.TypeOf((*MockUnsafeTasksServiceServer)(nil).mustEmbedUnimplementedTasksServiceServer))
}