// Code generated by MockGen. DO NOT EDIT.
// Source: ../phones_grpc.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	phones "github.com/kappapay/backend/lib/golang/src/kappa/services/phones"
	grpc "google.golang.org/grpc"
)

// MockPhoneServiceClient is a mock of PhoneServiceClient interface.
type MockPhoneServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockPhoneServiceClientMockRecorder
}

// MockPhoneServiceClientMockRecorder is the mock recorder for MockPhoneServiceClient.
type MockPhoneServiceClientMockRecorder struct {
	mock *MockPhoneServiceClient
}

// NewMockPhoneServiceClient creates a new mock instance.
func NewMockPhoneServiceClient(ctrl *gomock.Controller) *MockPhoneServiceClient {
	mock := &MockPhoneServiceClient{ctrl: ctrl}
	mock.recorder = &MockPhoneServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhoneServiceClient) EXPECT() *MockPhoneServiceClientMockRecorder {
	return m.recorder
}

// CreatePhoneNumber mocks base method.
func (m *MockPhoneServiceClient) CreatePhoneNumber(ctx context.Context, in *phones.CreatePhoneNumberRequest, opts ...grpc.CallOption) (*phones.CreatePhoneNumberResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreatePhoneNumber", varargs...)
	ret0, _ := ret[0].(*phones.CreatePhoneNumberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePhoneNumber indicates an expected call of CreatePhoneNumber.
func (mr *MockPhoneServiceClientMockRecorder) CreatePhoneNumber(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePhoneNumber", reflect.TypeOf((*MockPhoneServiceClient)(nil).CreatePhoneNumber), varargs...)
}

// DeletePhoneNumber mocks base method.
func (m *MockPhoneServiceClient) DeletePhoneNumber(ctx context.Context, in *phones.DeletePhoneNumberRequest, opts ...grpc.CallOption) (*phones.DeletePhoneNumberResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePhoneNumber", varargs...)
	ret0, _ := ret[0].(*phones.DeletePhoneNumberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePhoneNumber indicates an expected call of DeletePhoneNumber.
func (mr *MockPhoneServiceClientMockRecorder) DeletePhoneNumber(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhoneNumber", reflect.TypeOf((*MockPhoneServiceClient)(nil).DeletePhoneNumber), varargs...)
}

// DeletePhoneNumberByID mocks base method.
func (m *MockPhoneServiceClient) DeletePhoneNumberByID(ctx context.Context, in *phones.DeletePhoneNumberByIDRequest, opts ...grpc.CallOption) (*phones.DeletePhoneNumberByIDResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeletePhoneNumberByID", varargs...)
	ret0, _ := ret[0].(*phones.DeletePhoneNumberByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePhoneNumberByID indicates an expected call of DeletePhoneNumberByID.
func (mr *MockPhoneServiceClientMockRecorder) DeletePhoneNumberByID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhoneNumberByID", reflect.TypeOf((*MockPhoneServiceClient)(nil).DeletePhoneNumberByID), varargs...)
}

// GetOrCreatePhoneNumber mocks base method.
func (m *MockPhoneServiceClient) GetOrCreatePhoneNumber(ctx context.Context, in *phones.GetPhoneNumberRequest, opts ...grpc.CallOption) (*phones.GetPhoneNumberResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrCreatePhoneNumber", varargs...)
	ret0, _ := ret[0].(*phones.GetPhoneNumberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrCreatePhoneNumber indicates an expected call of GetOrCreatePhoneNumber.
func (mr *MockPhoneServiceClientMockRecorder) GetOrCreatePhoneNumber(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreatePhoneNumber", reflect.TypeOf((*MockPhoneServiceClient)(nil).GetOrCreatePhoneNumber), varargs...)
}

// GetPhoneNumber mocks base method.
func (m *MockPhoneServiceClient) GetPhoneNumber(ctx context.Context, in *phones.GetPhoneNumberRequest, opts ...grpc.CallOption) (*phones.GetPhoneNumberResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPhoneNumber", varargs...)
	ret0, _ := ret[0].(*phones.GetPhoneNumberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhoneNumber indicates an expected call of GetPhoneNumber.
func (mr *MockPhoneServiceClientMockRecorder) GetPhoneNumber(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhoneNumber", reflect.TypeOf((*MockPhoneServiceClient)(nil).GetPhoneNumber), varargs...)
}

// GetPhoneNumberByID mocks base method.
func (m *MockPhoneServiceClient) GetPhoneNumberByID(ctx context.Context, in *phones.GetPhoneNumberByIdRequest, opts ...grpc.CallOption) (*phones.GetPhoneNumberByIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPhoneNumberByID", varargs...)
	ret0, _ := ret[0].(*phones.GetPhoneNumberByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhoneNumberByID indicates an expected call of GetPhoneNumberByID.
func (mr *MockPhoneServiceClientMockRecorder) GetPhoneNumberByID(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhoneNumberByID", reflect.TypeOf((*MockPhoneServiceClient)(nil).GetPhoneNumberByID), varargs...)
}

// MarkPhoneVerified mocks base method.
func (m *MockPhoneServiceClient) MarkPhoneVerified(ctx context.Context, in *phones.MarkPhoneVerifiedRequest, opts ...grpc.CallOption) (*phones.MarkPhoneVerifiedResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MarkPhoneVerified", varargs...)
	ret0, _ := ret[0].(*phones.MarkPhoneVerifiedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkPhoneVerified indicates an expected call of MarkPhoneVerified.
func (mr *MockPhoneServiceClientMockRecorder) MarkPhoneVerified(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkPhoneVerified", reflect.TypeOf((*MockPhoneServiceClient)(nil).MarkPhoneVerified), varargs...)
}

// UpdatePhoneNumber mocks base method.
func (m *MockPhoneServiceClient) UpdatePhoneNumber(ctx context.Context, in *phones.UpdatePhoneNumberRequest, opts ...grpc.CallOption) (*phones.UpdatePhoneNumberResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdatePhoneNumber", varargs...)
	ret0, _ := ret[0].(*phones.UpdatePhoneNumberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePhoneNumber indicates an expected call of UpdatePhoneNumber.
func (mr *MockPhoneServiceClientMockRecorder) UpdatePhoneNumber(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoneNumber", reflect.TypeOf((*MockPhoneServiceClient)(nil).UpdatePhoneNumber), varargs...)
}

// MockPhoneServiceServer is a mock of PhoneServiceServer interface.
type MockPhoneServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockPhoneServiceServerMockRecorder
}

// MockPhoneServiceServerMockRecorder is the mock recorder for MockPhoneServiceServer.
type MockPhoneServiceServerMockRecorder struct {
	mock *MockPhoneServiceServer
}

// NewMockPhoneServiceServer creates a new mock instance.
func NewMockPhoneServiceServer(ctrl *gomock.Controller) *MockPhoneServiceServer {
	mock := &MockPhoneServiceServer{ctrl: ctrl}
	mock.recorder = &MockPhoneServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhoneServiceServer) EXPECT() *MockPhoneServiceServerMockRecorder {
	return m.recorder
}

// CreatePhoneNumber mocks base method.
func (m *MockPhoneServiceServer) CreatePhoneNumber(arg0 context.Context, arg1 *phones.CreatePhoneNumberRequest) (*phones.CreatePhoneNumberResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePhoneNumber", arg0, arg1)
	ret0, _ := ret[0].(*phones.CreatePhoneNumberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePhoneNumber indicates an expected call of CreatePhoneNumber.
func (mr *MockPhoneServiceServerMockRecorder) CreatePhoneNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePhoneNumber", reflect.TypeOf((*MockPhoneServiceServer)(nil).CreatePhoneNumber), arg0, arg1)
}

// DeletePhoneNumber mocks base method.
func (m *MockPhoneServiceServer) DeletePhoneNumber(arg0 context.Context, arg1 *phones.DeletePhoneNumberRequest) (*phones.DeletePhoneNumberResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePhoneNumber", arg0, arg1)
	ret0, _ := ret[0].(*phones.DeletePhoneNumberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePhoneNumber indicates an expected call of DeletePhoneNumber.
func (mr *MockPhoneServiceServerMockRecorder) DeletePhoneNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhoneNumber", reflect.TypeOf((*MockPhoneServiceServer)(nil).DeletePhoneNumber), arg0, arg1)
}

// DeletePhoneNumberByID mocks base method.
func (m *MockPhoneServiceServer) DeletePhoneNumberByID(arg0 context.Context, arg1 *phones.DeletePhoneNumberByIDRequest) (*phones.DeletePhoneNumberByIDResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePhoneNumberByID", arg0, arg1)
	ret0, _ := ret[0].(*phones.DeletePhoneNumberByIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePhoneNumberByID indicates an expected call of DeletePhoneNumberByID.
func (mr *MockPhoneServiceServerMockRecorder) DeletePhoneNumberByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhoneNumberByID", reflect.TypeOf((*MockPhoneServiceServer)(nil).DeletePhoneNumberByID), arg0, arg1)
}

// GetOrCreatePhoneNumber mocks base method.
func (m *MockPhoneServiceServer) GetOrCreatePhoneNumber(arg0 context.Context, arg1 *phones.GetPhoneNumberRequest) (*phones.GetPhoneNumberResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrCreatePhoneNumber", arg0, arg1)
	ret0, _ := ret[0].(*phones.GetPhoneNumberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrCreatePhoneNumber indicates an expected call of GetOrCreatePhoneNumber.
func (mr *MockPhoneServiceServerMockRecorder) GetOrCreatePhoneNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreatePhoneNumber", reflect.TypeOf((*MockPhoneServiceServer)(nil).GetOrCreatePhoneNumber), arg0, arg1)
}

// GetPhoneNumber mocks base method.
func (m *MockPhoneServiceServer) GetPhoneNumber(arg0 context.Context, arg1 *phones.GetPhoneNumberRequest) (*phones.GetPhoneNumberResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhoneNumber", arg0, arg1)
	ret0, _ := ret[0].(*phones.GetPhoneNumberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhoneNumber indicates an expected call of GetPhoneNumber.
func (mr *MockPhoneServiceServerMockRecorder) GetPhoneNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhoneNumber", reflect.TypeOf((*MockPhoneServiceServer)(nil).GetPhoneNumber), arg0, arg1)
}

// GetPhoneNumberByID mocks base method.
func (m *MockPhoneServiceServer) GetPhoneNumberByID(arg0 context.Context, arg1 *phones.GetPhoneNumberByIdRequest) (*phones.GetPhoneNumberByIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhoneNumberByID", arg0, arg1)
	ret0, _ := ret[0].(*phones.GetPhoneNumberByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhoneNumberByID indicates an expected call of GetPhoneNumberByID.
func (mr *MockPhoneServiceServerMockRecorder) GetPhoneNumberByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhoneNumberByID", reflect.TypeOf((*MockPhoneServiceServer)(nil).GetPhoneNumberByID), arg0, arg1)
}

// MarkPhoneVerified mocks base method.
func (m *MockPhoneServiceServer) MarkPhoneVerified(arg0 context.Context, arg1 *phones.MarkPhoneVerifiedRequest) (*phones.MarkPhoneVerifiedResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkPhoneVerified", arg0, arg1)
	ret0, _ := ret[0].(*phones.MarkPhoneVerifiedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarkPhoneVerified indicates an expected call of MarkPhoneVerified.
func (mr *MockPhoneServiceServerMockRecorder) MarkPhoneVerified(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkPhoneVerified", reflect.TypeOf((*MockPhoneServiceServer)(nil).MarkPhoneVerified), arg0, arg1)
}

// UpdatePhoneNumber mocks base method.
func (m *MockPhoneServiceServer) UpdatePhoneNumber(arg0 context.Context, arg1 *phones.UpdatePhoneNumberRequest) (*phones.UpdatePhoneNumberResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhoneNumber", arg0, arg1)
	ret0, _ := ret[0].(*phones.UpdatePhoneNumberResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePhoneNumber indicates an expected call of UpdatePhoneNumber.
func (mr *MockPhoneServiceServerMockRecorder) UpdatePhoneNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoneNumber", reflect.TypeOf((*MockPhoneServiceServer)(nil).UpdatePhoneNumber), arg0, arg1)
}

// mustEmbedUnimplementedPhoneServiceServer mocks base method.
func (m *MockPhoneServiceServer) mustEmbedUnimplementedPhoneServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPhoneServiceServer")
}

// mustEmbedUnimplementedPhoneServiceServer indicates an expected call of mustEmbedUnimplementedPhoneServiceServer.
func (mr *MockPhoneServiceServerMockRecorder) mustEmbedUnimplementedPhoneServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPhoneServiceServer", reflect.TypeOf((*MockPhoneServiceServer)(nil).mustEmbedUnimplementedPhoneServiceServer))
}

// MockUnsafePhoneServiceServer is a mock of UnsafePhoneServiceServer interface.
type MockUnsafePhoneServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafePhoneServiceServerMockRecorder
}

// MockUnsafePhoneServiceServerMockRecorder is the mock recorder for MockUnsafePhoneServiceServer.
type MockUnsafePhoneServiceServerMockRecorder struct {
	mock *MockUnsafePhoneServiceServer
}

// NewMockUnsafePhoneServiceServer creates a new mock instance.
func NewMockUnsafePhoneServiceServer(ctrl *gomock.Controller) *MockUnsafePhoneServiceServer {
	mock := &MockUnsafePhoneServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafePhoneServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafePhoneServiceServer) EXPECT() *MockUnsafePhoneServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedPhoneServiceServer mocks base method.
func (m *MockUnsafePhoneServiceServer) mustEmbedUnimplementedPhoneServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedPhoneServiceServer")
}

// mustEmbedUnimplementedPhoneServiceServer indicates an expected call of mustEmbedUnimplementedPhoneServiceServer.
func (mr *MockUnsafePhoneServiceServerMockRecorder) mustEmbedUnimplementedPhoneServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedPhoneServiceServer", reflect.TypeOf((*MockUnsafePhoneServiceServer)(nil).mustEmbedUnimplementedPhoneServiceServer))
}