// Code generated by MockGen. DO NOT EDIT.
// Source: ../reconciliation_grpc.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reconciliation "github.com/kappapay/backend/lib/golang/src/kappa/services/reconciliation"
	grpc "google.golang.org/grpc"
)

// MockReconciliationServiceClient is a mock of ReconciliationServiceClient interface.
type MockReconciliationServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockReconciliationServiceClientMockRecorder
}

// MockReconciliationServiceClientMockRecorder is the mock recorder for MockReconciliationServiceClient.
type MockReconciliationServiceClientMockRecorder struct {
	mock *MockReconciliationServiceClient
}

// NewMockReconciliationServiceClient creates a new mock instance.
func NewMockReconciliationServiceClient(ctrl *gomock.Controller) *MockReconciliationServiceClient {
	mock := &MockReconciliationServiceClient{ctrl: ctrl}
	mock.recorder = &MockReconciliationServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReconciliationServiceClient) EXPECT() *MockReconciliationServiceClientMockRecorder {
	return m.recorder
}

// CreateExternalTransactionNotification mocks base method.
func (m *MockReconciliationServiceClient) CreateExternalTransactionNotification(ctx context.Context, in *reconciliation.CreateExternalTransactionNotificationRequest, opts ...grpc.CallOption) (*reconciliation.CreateExternalTransactionNotificationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateExternalTransactionNotification", varargs...)
	ret0, _ := ret[0].(*reconciliation.CreateExternalTransactionNotificationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExternalTransactionNotification indicates an expected call of CreateExternalTransactionNotification.
func (mr *MockReconciliationServiceClientMockRecorder) CreateExternalTransactionNotification(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExternalTransactionNotification", reflect.TypeOf((*MockReconciliationServiceClient)(nil).CreateExternalTransactionNotification), varargs...)
}

// SetMatchProbabilities mocks base method.
func (m *MockReconciliationServiceClient) SetMatchProbabilities(ctx context.Context, in *reconciliation.SetMatchProbabilitiesRequest, opts ...grpc.CallOption) (*reconciliation.SetMatchProbabilitiesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetMatchProbabilities", varargs...)
	ret0, _ := ret[0].(*reconciliation.SetMatchProbabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetMatchProbabilities indicates an expected call of SetMatchProbabilities.
func (mr *MockReconciliationServiceClientMockRecorder) SetMatchProbabilities(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMatchProbabilities", reflect.TypeOf((*MockReconciliationServiceClient)(nil).SetMatchProbabilities), varargs...)
}

// SetTransactionID mocks base method.
func (m *MockReconciliationServiceClient) SetTransactionID(ctx context.Context, in *reconciliation.SetTransactionIDRequest, opts ...grpc.CallOption) (*reconciliation.SetTransactionIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetTransactionID", varargs...)
	ret0, _ := ret[0].(*reconciliation.SetTransactionIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetTransactionID indicates an expected call of SetTransactionID.
func (mr *MockReconciliationServiceClientMockRecorder) SetTransactionID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransactionID", reflect.TypeOf((*MockReconciliationServiceClient)(nil).SetTransactionID), varargs...)
}

// SetTransactionReconciliationStatus mocks base method.
func (m *MockReconciliationServiceClient) SetTransactionReconciliationStatus(ctx context.Context, in *reconciliation.SetTransactionReconciliationStatusRequest, opts ...grpc.CallOption) (*reconciliation.SetTransactionReconciliationStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetTransactionReconciliationStatus", varargs...)
	ret0, _ := ret[0].(*reconciliation.SetTransactionReconciliationStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetTransactionReconciliationStatus indicates an expected call of SetTransactionReconciliationStatus.
func (mr *MockReconciliationServiceClientMockRecorder) SetTransactionReconciliationStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransactionReconciliationStatus", reflect.TypeOf((*MockReconciliationServiceClient)(nil).SetTransactionReconciliationStatus), varargs...)
}

// MockReconciliationServiceServer is a mock of ReconciliationServiceServer interface.
type MockReconciliationServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockReconciliationServiceServerMockRecorder
}

// MockReconciliationServiceServerMockRecorder is the mock recorder for MockReconciliationServiceServer.
type MockReconciliationServiceServerMockRecorder struct {
	mock *MockReconciliationServiceServer
}

// NewMockReconciliationServiceServer creates a new mock instance.
func NewMockReconciliationServiceServer(ctrl *gomock.Controller) *MockReconciliationServiceServer {
	mock := &MockReconciliationServiceServer{ctrl: ctrl}
	mock.recorder = &MockReconciliationServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReconciliationServiceServer) EXPECT() *MockReconciliationServiceServerMockRecorder {
	return m.recorder
}

// CreateExternalTransactionNotification mocks base method.
func (m *MockReconciliationServiceServer) CreateExternalTransactionNotification(arg0 context.Context, arg1 *reconciliation.CreateExternalTransactionNotificationRequest) (*reconciliation.CreateExternalTransactionNotificationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExternalTransactionNotification", arg0, arg1)
	ret0, _ := ret[0].(*reconciliation.CreateExternalTransactionNotificationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExternalTransactionNotification indicates an expected call of CreateExternalTransactionNotification.
func (mr *MockReconciliationServiceServerMockRecorder) CreateExternalTransactionNotification(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExternalTransactionNotification", reflect.TypeOf((*MockReconciliationServiceServer)(nil).CreateExternalTransactionNotification), arg0, arg1)
}

// SetMatchProbabilities mocks base method.
func (m *MockReconciliationServiceServer) SetMatchProbabilities(arg0 context.Context, arg1 *reconciliation.SetMatchProbabilitiesRequest) (*reconciliation.SetMatchProbabilitiesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMatchProbabilities", arg0, arg1)
	ret0, _ := ret[0].(*reconciliation.SetMatchProbabilitiesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetMatchProbabilities indicates an expected call of SetMatchProbabilities.
func (mr *MockReconciliationServiceServerMockRecorder) SetMatchProbabilities(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMatchProbabilities", reflect.TypeOf((*MockReconciliationServiceServer)(nil).SetMatchProbabilities), arg0, arg1)
}

// SetTransactionID mocks base method.
func (m *MockReconciliationServiceServer) SetTransactionID(arg0 context.Context, arg1 *reconciliation.SetTransactionIDRequest) (*reconciliation.SetTransactionIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTransactionID", arg0, arg1)
	ret0, _ := ret[0].(*reconciliation.SetTransactionIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetTransactionID indicates an expected call of SetTransactionID.
func (mr *MockReconciliationServiceServerMockRecorder) SetTransactionID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransactionID", reflect.TypeOf((*MockReconciliationServiceServer)(nil).SetTransactionID), arg0, arg1)
}

// SetTransactionReconciliationStatus mocks base method.
func (m *MockReconciliationServiceServer) SetTransactionReconciliationStatus(arg0 context.Context, arg1 *reconciliation.SetTransactionReconciliationStatusRequest) (*reconciliation.SetTransactionReconciliationStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTransactionReconciliationStatus", arg0, arg1)
	ret0, _ := ret[0].(*reconciliation.SetTransactionReconciliationStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetTransactionReconciliationStatus indicates an expected call of SetTransactionReconciliationStatus.
func (mr *MockReconciliationServiceServerMockRecorder) SetTransactionReconciliationStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransactionReconciliationStatus", reflect.TypeOf((*MockReconciliationServiceServer)(nil).SetTransactionReconciliationStatus), arg0, arg1)
}

// mustEmbedUnimplementedReconciliationServiceServer mocks base method.
func (m *MockReconciliationServiceServer) mustEmbedUnimplementedReconciliationServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedReconciliationServiceServer")
}

// mustEmbedUnimplementedReconciliationServiceServer indicates an expected call of mustEmbedUnimplementedReconciliationServiceServer.
func (mr *MockReconciliationServiceServerMockRecorder) mustEmbedUnimplementedReconciliationServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedReconciliationServiceServer", reflect.TypeOf((*MockReconciliationServiceServer)(nil).mustEmbedUnimplementedReconciliationServiceServer))
}

// MockUnsafeReconciliationServiceServer is a mock of UnsafeReconciliationServiceServer interface.
type MockUnsafeReconciliationServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeReconciliationServiceServerMockRecorder
}

// MockUnsafeReconciliationServiceServerMockRecorder is the mock recorder for MockUnsafeReconciliationServiceServer.
type MockUnsafeReconciliationServiceServerMockRecorder struct {
	mock *MockUnsafeReconciliationServiceServer
}

// NewMockUnsafeReconciliationServiceServer creates a new mock instance.
func NewMockUnsafeReconciliationServiceServer(ctrl *gomock.Controller) *MockUnsafeReconciliationServiceServer {
	mock := &MockUnsafeReconciliationServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeReconciliationServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeReconciliationServiceServer) EXPECT() *MockUnsafeReconciliationServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedReconciliationServiceServer mocks base method.
func (m *MockUnsafeReconciliationServiceServer) mustEmbedUnimplementedReconciliationServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedReconciliationServiceServer")
}

// mustEmbedUnimplementedReconciliationServiceServer indicates an expected call of mustEmbedUnimplementedReconciliationServiceServer.
func (mr *MockUnsafeReconciliationServiceServerMockRecorder) mustEmbedUnimplementedReconciliationServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedReconciliationServiceServer", reflect.TypeOf((*MockUnsafeReconciliationServiceServer)(nil).mustEmbedUnimplementedReconciliationServiceServer))
}