// Code generated by MockGen. DO NOT EDIT.
// Source: ../mobile_money_accounts_grpc.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mobile_money_accounts "github.com/kappapay/backend/lib/golang/src/kappa/services/mobile_money_accounts"
	grpc "google.golang.org/grpc"
)

// MockMobileMoneyAccountsServiceClient is a mock of MobileMoneyAccountsServiceClient interface.
type MockMobileMoneyAccountsServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockMobileMoneyAccountsServiceClientMockRecorder
}

// MockMobileMoneyAccountsServiceClientMockRecorder is the mock recorder for MockMobileMoneyAccountsServiceClient.
type MockMobileMoneyAccountsServiceClientMockRecorder struct {
	mock *MockMobileMoneyAccountsServiceClient
}

// NewMockMobileMoneyAccountsServiceClient creates a new mock instance.
func NewMockMobileMoneyAccountsServiceClient(ctrl *gomock.Controller) *MockMobileMoneyAccountsServiceClient {
	mock := &MockMobileMoneyAccountsServiceClient{ctrl: ctrl}
	mock.recorder = &MockMobileMoneyAccountsServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMobileMoneyAccountsServiceClient) EXPECT() *MockMobileMoneyAccountsServiceClientMockRecorder {
	return m.recorder
}

// CreateMobileMoneyAccount mocks base method.
func (m *MockMobileMoneyAccountsServiceClient) CreateMobileMoneyAccount(ctx context.Context, in *mobile_money_accounts.CreateMobileMoneyAccountRequest, opts ...grpc.CallOption) (*mobile_money_accounts.CreateMobileMoneyAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMobileMoneyAccount", varargs...)
	ret0, _ := ret[0].(*mobile_money_accounts.CreateMobileMoneyAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMobileMoneyAccount indicates an expected call of CreateMobileMoneyAccount.
func (mr *MockMobileMoneyAccountsServiceClientMockRecorder) CreateMobileMoneyAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMobileMoneyAccount", reflect.TypeOf((*MockMobileMoneyAccountsServiceClient)(nil).CreateMobileMoneyAccount), varargs...)
}

// DeleteMobileMoneyAccount mocks base method.
func (m *MockMobileMoneyAccountsServiceClient) DeleteMobileMoneyAccount(ctx context.Context, in *mobile_money_accounts.DeleteMobileMoneyAccountRequest, opts ...grpc.CallOption) (*mobile_money_accounts.DeleteMobileMoneyAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteMobileMoneyAccount", varargs...)
	ret0, _ := ret[0].(*mobile_money_accounts.DeleteMobileMoneyAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMobileMoneyAccount indicates an expected call of DeleteMobileMoneyAccount.
func (mr *MockMobileMoneyAccountsServiceClientMockRecorder) DeleteMobileMoneyAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMobileMoneyAccount", reflect.TypeOf((*MockMobileMoneyAccountsServiceClient)(nil).DeleteMobileMoneyAccount), varargs...)
}

// GetMobileMoneyAccountById mocks base method.
func (m *MockMobileMoneyAccountsServiceClient) GetMobileMoneyAccountById(ctx context.Context, in *mobile_money_accounts.GetMobileMoneyAccountByIdRequest, opts ...grpc.CallOption) (*mobile_money_accounts.GetMobileMoneyAccountByIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMobileMoneyAccountById", varargs...)
	ret0, _ := ret[0].(*mobile_money_accounts.GetMobileMoneyAccountByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMobileMoneyAccountById indicates an expected call of GetMobileMoneyAccountById.
func (mr *MockMobileMoneyAccountsServiceClientMockRecorder) GetMobileMoneyAccountById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMobileMoneyAccountById", reflect.TypeOf((*MockMobileMoneyAccountsServiceClient)(nil).GetMobileMoneyAccountById), varargs...)
}

// UpdateMobileMoneyAccount mocks base method.
func (m *MockMobileMoneyAccountsServiceClient) UpdateMobileMoneyAccount(ctx context.Context, in *mobile_money_accounts.UpdateMobileMoneyAccountRequest, opts ...grpc.CallOption) (*mobile_money_accounts.UpdateMobileMoneyAccountResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateMobileMoneyAccount", varargs...)
	ret0, _ := ret[0].(*mobile_money_accounts.UpdateMobileMoneyAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMobileMoneyAccount indicates an expected call of UpdateMobileMoneyAccount.
func (mr *MockMobileMoneyAccountsServiceClientMockRecorder) UpdateMobileMoneyAccount(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMobileMoneyAccount", reflect.TypeOf((*MockMobileMoneyAccountsServiceClient)(nil).UpdateMobileMoneyAccount), varargs...)
}

// MockMobileMoneyAccountsServiceServer is a mock of MobileMoneyAccountsServiceServer interface.
type MockMobileMoneyAccountsServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockMobileMoneyAccountsServiceServerMockRecorder
}

// MockMobileMoneyAccountsServiceServerMockRecorder is the mock recorder for MockMobileMoneyAccountsServiceServer.
type MockMobileMoneyAccountsServiceServerMockRecorder struct {
	mock *MockMobileMoneyAccountsServiceServer
}

// NewMockMobileMoneyAccountsServiceServer creates a new mock instance.
func NewMockMobileMoneyAccountsServiceServer(ctrl *gomock.Controller) *MockMobileMoneyAccountsServiceServer {
	mock := &MockMobileMoneyAccountsServiceServer{ctrl: ctrl}
	mock.recorder = &MockMobileMoneyAccountsServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMobileMoneyAccountsServiceServer) EXPECT() *MockMobileMoneyAccountsServiceServerMockRecorder {
	return m.recorder
}

// CreateMobileMoneyAccount mocks base method.
func (m *MockMobileMoneyAccountsServiceServer) CreateMobileMoneyAccount(arg0 context.Context, arg1 *mobile_money_accounts.CreateMobileMoneyAccountRequest) (*mobile_money_accounts.CreateMobileMoneyAccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMobileMoneyAccount", arg0, arg1)
	ret0, _ := ret[0].(*mobile_money_accounts.CreateMobileMoneyAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMobileMoneyAccount indicates an expected call of CreateMobileMoneyAccount.
func (mr *MockMobileMoneyAccountsServiceServerMockRecorder) CreateMobileMoneyAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMobileMoneyAccount", reflect.TypeOf((*MockMobileMoneyAccountsServiceServer)(nil).CreateMobileMoneyAccount), arg0, arg1)
}

// DeleteMobileMoneyAccount mocks base method.
func (m *MockMobileMoneyAccountsServiceServer) DeleteMobileMoneyAccount(arg0 context.Context, arg1 *mobile_money_accounts.DeleteMobileMoneyAccountRequest) (*mobile_money_accounts.DeleteMobileMoneyAccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMobileMoneyAccount", arg0, arg1)
	ret0, _ := ret[0].(*mobile_money_accounts.DeleteMobileMoneyAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteMobileMoneyAccount indicates an expected call of DeleteMobileMoneyAccount.
func (mr *MockMobileMoneyAccountsServiceServerMockRecorder) DeleteMobileMoneyAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMobileMoneyAccount", reflect.TypeOf((*MockMobileMoneyAccountsServiceServer)(nil).DeleteMobileMoneyAccount), arg0, arg1)
}

// GetMobileMoneyAccountById mocks base method.
func (m *MockMobileMoneyAccountsServiceServer) GetMobileMoneyAccountById(arg0 context.Context, arg1 *mobile_money_accounts.GetMobileMoneyAccountByIdRequest) (*mobile_money_accounts.GetMobileMoneyAccountByIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMobileMoneyAccountById", arg0, arg1)
	ret0, _ := ret[0].(*mobile_money_accounts.GetMobileMoneyAccountByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMobileMoneyAccountById indicates an expected call of GetMobileMoneyAccountById.
func (mr *MockMobileMoneyAccountsServiceServerMockRecorder) GetMobileMoneyAccountById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMobileMoneyAccountById", reflect.TypeOf((*MockMobileMoneyAccountsServiceServer)(nil).GetMobileMoneyAccountById), arg0, arg1)
}

// UpdateMobileMoneyAccount mocks base method.
func (m *MockMobileMoneyAccountsServiceServer) UpdateMobileMoneyAccount(arg0 context.Context, arg1 *mobile_money_accounts.UpdateMobileMoneyAccountRequest) (*mobile_money_accounts.UpdateMobileMoneyAccountResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMobileMoneyAccount", arg0, arg1)
	ret0, _ := ret[0].(*mobile_money_accounts.UpdateMobileMoneyAccountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMobileMoneyAccount indicates an expected call of UpdateMobileMoneyAccount.
func (mr *MockMobileMoneyAccountsServiceServerMockRecorder) UpdateMobileMoneyAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMobileMoneyAccount", reflect.TypeOf((*MockMobileMoneyAccountsServiceServer)(nil).UpdateMobileMoneyAccount), arg0, arg1)
}

// mustEmbedUnimplementedMobileMoneyAccountsServiceServer mocks base method.
func (m *MockMobileMoneyAccountsServiceServer) mustEmbedUnimplementedMobileMoneyAccountsServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedMobileMoneyAccountsServiceServer")
}

// mustEmbedUnimplementedMobileMoneyAccountsServiceServer indicates an expected call of mustEmbedUnimplementedMobileMoneyAccountsServiceServer.
func (mr *MockMobileMoneyAccountsServiceServerMockRecorder) mustEmbedUnimplementedMobileMoneyAccountsServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedMobileMoneyAccountsServiceServer", reflect.TypeOf((*MockMobileMoneyAccountsServiceServer)(nil).mustEmbedUnimplementedMobileMoneyAccountsServiceServer))
}

// MockUnsafeMobileMoneyAccountsServiceServer is a mock of UnsafeMobileMoneyAccountsServiceServer interface.
type MockUnsafeMobileMoneyAccountsServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeMobileMoneyAccountsServiceServerMockRecorder
}

// MockUnsafeMobileMoneyAccountsServiceServerMockRecorder is the mock recorder for MockUnsafeMobileMoneyAccountsServiceServer.
type MockUnsafeMobileMoneyAccountsServiceServerMockRecorder struct {
	mock *MockUnsafeMobileMoneyAccountsServiceServer
}

// NewMockUnsafeMobileMoneyAccountsServiceServer creates a new mock instance.
func NewMockUnsafeMobileMoneyAccountsServiceServer(ctrl *gomock.Controller) *MockUnsafeMobileMoneyAccountsServiceServer {
	mock := &MockUnsafeMobileMoneyAccountsServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeMobileMoneyAccountsServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeMobileMoneyAccountsServiceServer) EXPECT() *MockUnsafeMobileMoneyAccountsServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedMobileMoneyAccountsServiceServer mocks base method.
func (m *MockUnsafeMobileMoneyAccountsServiceServer) mustEmbedUnimplementedMobileMoneyAccountsServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedMobileMoneyAccountsServiceServer")
}

// mustEmbedUnimplementedMobileMoneyAccountsServiceServer indicates an expected call of mustEmbedUnimplementedMobileMoneyAccountsServiceServer.
func (mr *MockUnsafeMobileMoneyAccountsServiceServerMockRecorder) mustEmbedUnimplementedMobileMoneyAccountsServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedMobileMoneyAccountsServiceServer", reflect.TypeOf((*MockUnsafeMobileMoneyAccountsServiceServer)(nil).mustEmbedUnimplementedMobileMoneyAccountsServiceServer))
}