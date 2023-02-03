// Code generated by MockGen. DO NOT EDIT.
// Source: ../exchange_rates_grpc.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	exchange_rates "github.com/kappapay/backend/lib/golang/src/kappa/services/exchange_rates"
	grpc "google.golang.org/grpc"
)

// MockExchangeRatesServiceClient is a mock of ExchangeRatesServiceClient interface.
type MockExchangeRatesServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeRatesServiceClientMockRecorder
}

// MockExchangeRatesServiceClientMockRecorder is the mock recorder for MockExchangeRatesServiceClient.
type MockExchangeRatesServiceClientMockRecorder struct {
	mock *MockExchangeRatesServiceClient
}

// NewMockExchangeRatesServiceClient creates a new mock instance.
func NewMockExchangeRatesServiceClient(ctrl *gomock.Controller) *MockExchangeRatesServiceClient {
	mock := &MockExchangeRatesServiceClient{ctrl: ctrl}
	mock.recorder = &MockExchangeRatesServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchangeRatesServiceClient) EXPECT() *MockExchangeRatesServiceClientMockRecorder {
	return m.recorder
}

// GetBlendedExchangeRate mocks base method.
func (m *MockExchangeRatesServiceClient) GetBlendedExchangeRate(ctx context.Context, in *exchange_rates.GetBlendedExchangeRateRequest, opts ...grpc.CallOption) (*exchange_rates.GetBlendedExchangeRateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBlendedExchangeRate", varargs...)
	ret0, _ := ret[0].(*exchange_rates.GetBlendedExchangeRateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlendedExchangeRate indicates an expected call of GetBlendedExchangeRate.
func (mr *MockExchangeRatesServiceClientMockRecorder) GetBlendedExchangeRate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlendedExchangeRate", reflect.TypeOf((*MockExchangeRatesServiceClient)(nil).GetBlendedExchangeRate), varargs...)
}

// GetDataSetBySourceAndName mocks base method.
func (m *MockExchangeRatesServiceClient) GetDataSetBySourceAndName(ctx context.Context, in *exchange_rates.GetDataSetBySourceAndNameRequest, opts ...grpc.CallOption) (*exchange_rates.GetDataSetBySourceAndNameResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDataSetBySourceAndName", varargs...)
	ret0, _ := ret[0].(*exchange_rates.GetDataSetBySourceAndNameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataSetBySourceAndName indicates an expected call of GetDataSetBySourceAndName.
func (mr *MockExchangeRatesServiceClientMockRecorder) GetDataSetBySourceAndName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataSetBySourceAndName", reflect.TypeOf((*MockExchangeRatesServiceClient)(nil).GetDataSetBySourceAndName), varargs...)
}

// GetExchangeRate mocks base method.
func (m *MockExchangeRatesServiceClient) GetExchangeRate(ctx context.Context, in *exchange_rates.GetExchangeRateRequest, opts ...grpc.CallOption) (*exchange_rates.GetExchangeRateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetExchangeRate", varargs...)
	ret0, _ := ret[0].(*exchange_rates.GetExchangeRateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRate indicates an expected call of GetExchangeRate.
func (mr *MockExchangeRatesServiceClientMockRecorder) GetExchangeRate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockExchangeRatesServiceClient)(nil).GetExchangeRate), varargs...)
}

// MockExchangeRatesServiceServer is a mock of ExchangeRatesServiceServer interface.
type MockExchangeRatesServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockExchangeRatesServiceServerMockRecorder
}

// MockExchangeRatesServiceServerMockRecorder is the mock recorder for MockExchangeRatesServiceServer.
type MockExchangeRatesServiceServerMockRecorder struct {
	mock *MockExchangeRatesServiceServer
}

// NewMockExchangeRatesServiceServer creates a new mock instance.
func NewMockExchangeRatesServiceServer(ctrl *gomock.Controller) *MockExchangeRatesServiceServer {
	mock := &MockExchangeRatesServiceServer{ctrl: ctrl}
	mock.recorder = &MockExchangeRatesServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExchangeRatesServiceServer) EXPECT() *MockExchangeRatesServiceServerMockRecorder {
	return m.recorder
}

// GetBlendedExchangeRate mocks base method.
func (m *MockExchangeRatesServiceServer) GetBlendedExchangeRate(arg0 context.Context, arg1 *exchange_rates.GetBlendedExchangeRateRequest) (*exchange_rates.GetBlendedExchangeRateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlendedExchangeRate", arg0, arg1)
	ret0, _ := ret[0].(*exchange_rates.GetBlendedExchangeRateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlendedExchangeRate indicates an expected call of GetBlendedExchangeRate.
func (mr *MockExchangeRatesServiceServerMockRecorder) GetBlendedExchangeRate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlendedExchangeRate", reflect.TypeOf((*MockExchangeRatesServiceServer)(nil).GetBlendedExchangeRate), arg0, arg1)
}

// GetDataSetBySourceAndName mocks base method.
func (m *MockExchangeRatesServiceServer) GetDataSetBySourceAndName(arg0 context.Context, arg1 *exchange_rates.GetDataSetBySourceAndNameRequest) (*exchange_rates.GetDataSetBySourceAndNameResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataSetBySourceAndName", arg0, arg1)
	ret0, _ := ret[0].(*exchange_rates.GetDataSetBySourceAndNameResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataSetBySourceAndName indicates an expected call of GetDataSetBySourceAndName.
func (mr *MockExchangeRatesServiceServerMockRecorder) GetDataSetBySourceAndName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataSetBySourceAndName", reflect.TypeOf((*MockExchangeRatesServiceServer)(nil).GetDataSetBySourceAndName), arg0, arg1)
}

// GetExchangeRate mocks base method.
func (m *MockExchangeRatesServiceServer) GetExchangeRate(arg0 context.Context, arg1 *exchange_rates.GetExchangeRateRequest) (*exchange_rates.GetExchangeRateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchangeRate", arg0, arg1)
	ret0, _ := ret[0].(*exchange_rates.GetExchangeRateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExchangeRate indicates an expected call of GetExchangeRate.
func (mr *MockExchangeRatesServiceServerMockRecorder) GetExchangeRate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeRate", reflect.TypeOf((*MockExchangeRatesServiceServer)(nil).GetExchangeRate), arg0, arg1)
}

// mustEmbedUnimplementedExchangeRatesServiceServer mocks base method.
func (m *MockExchangeRatesServiceServer) mustEmbedUnimplementedExchangeRatesServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedExchangeRatesServiceServer")
}

// mustEmbedUnimplementedExchangeRatesServiceServer indicates an expected call of mustEmbedUnimplementedExchangeRatesServiceServer.
func (mr *MockExchangeRatesServiceServerMockRecorder) mustEmbedUnimplementedExchangeRatesServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedExchangeRatesServiceServer", reflect.TypeOf((*MockExchangeRatesServiceServer)(nil).mustEmbedUnimplementedExchangeRatesServiceServer))
}

// MockUnsafeExchangeRatesServiceServer is a mock of UnsafeExchangeRatesServiceServer interface.
type MockUnsafeExchangeRatesServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeExchangeRatesServiceServerMockRecorder
}

// MockUnsafeExchangeRatesServiceServerMockRecorder is the mock recorder for MockUnsafeExchangeRatesServiceServer.
type MockUnsafeExchangeRatesServiceServerMockRecorder struct {
	mock *MockUnsafeExchangeRatesServiceServer
}

// NewMockUnsafeExchangeRatesServiceServer creates a new mock instance.
func NewMockUnsafeExchangeRatesServiceServer(ctrl *gomock.Controller) *MockUnsafeExchangeRatesServiceServer {
	mock := &MockUnsafeExchangeRatesServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeExchangeRatesServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeExchangeRatesServiceServer) EXPECT() *MockUnsafeExchangeRatesServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedExchangeRatesServiceServer mocks base method.
func (m *MockUnsafeExchangeRatesServiceServer) mustEmbedUnimplementedExchangeRatesServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedExchangeRatesServiceServer")
}

// mustEmbedUnimplementedExchangeRatesServiceServer indicates an expected call of mustEmbedUnimplementedExchangeRatesServiceServer.
func (mr *MockUnsafeExchangeRatesServiceServerMockRecorder) mustEmbedUnimplementedExchangeRatesServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedExchangeRatesServiceServer", reflect.TypeOf((*MockUnsafeExchangeRatesServiceServer)(nil).mustEmbedUnimplementedExchangeRatesServiceServer))
}
