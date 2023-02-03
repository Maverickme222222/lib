// Code generated by MockGen. DO NOT EDIT.
// Source: ../invites_grpc.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	invites "github.com/kappapay/backend/lib/golang/src/kappa/services/invites"
	grpc "google.golang.org/grpc"
)

// MockInvitesServiceClient is a mock of InvitesServiceClient interface.
type MockInvitesServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockInvitesServiceClientMockRecorder
}

// MockInvitesServiceClientMockRecorder is the mock recorder for MockInvitesServiceClient.
type MockInvitesServiceClientMockRecorder struct {
	mock *MockInvitesServiceClient
}

// NewMockInvitesServiceClient creates a new mock instance.
func NewMockInvitesServiceClient(ctrl *gomock.Controller) *MockInvitesServiceClient {
	mock := &MockInvitesServiceClient{ctrl: ctrl}
	mock.recorder = &MockInvitesServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvitesServiceClient) EXPECT() *MockInvitesServiceClientMockRecorder {
	return m.recorder
}

// CreateInvite mocks base method.
func (m *MockInvitesServiceClient) CreateInvite(ctx context.Context, in *invites.CreateInviteRequest, opts ...grpc.CallOption) (*invites.CreateInviteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateInvite", varargs...)
	ret0, _ := ret[0].(*invites.CreateInviteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInvite indicates an expected call of CreateInvite.
func (mr *MockInvitesServiceClientMockRecorder) CreateInvite(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvite", reflect.TypeOf((*MockInvitesServiceClient)(nil).CreateInvite), varargs...)
}

// MockInvitesServiceServer is a mock of InvitesServiceServer interface.
type MockInvitesServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockInvitesServiceServerMockRecorder
}

// MockInvitesServiceServerMockRecorder is the mock recorder for MockInvitesServiceServer.
type MockInvitesServiceServerMockRecorder struct {
	mock *MockInvitesServiceServer
}

// NewMockInvitesServiceServer creates a new mock instance.
func NewMockInvitesServiceServer(ctrl *gomock.Controller) *MockInvitesServiceServer {
	mock := &MockInvitesServiceServer{ctrl: ctrl}
	mock.recorder = &MockInvitesServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInvitesServiceServer) EXPECT() *MockInvitesServiceServerMockRecorder {
	return m.recorder
}

// CreateInvite mocks base method.
func (m *MockInvitesServiceServer) CreateInvite(arg0 context.Context, arg1 *invites.CreateInviteRequest) (*invites.CreateInviteResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvite", arg0, arg1)
	ret0, _ := ret[0].(*invites.CreateInviteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInvite indicates an expected call of CreateInvite.
func (mr *MockInvitesServiceServerMockRecorder) CreateInvite(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvite", reflect.TypeOf((*MockInvitesServiceServer)(nil).CreateInvite), arg0, arg1)
}

// mustEmbedUnimplementedInvitesServiceServer mocks base method.
func (m *MockInvitesServiceServer) mustEmbedUnimplementedInvitesServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedInvitesServiceServer")
}

// mustEmbedUnimplementedInvitesServiceServer indicates an expected call of mustEmbedUnimplementedInvitesServiceServer.
func (mr *MockInvitesServiceServerMockRecorder) mustEmbedUnimplementedInvitesServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedInvitesServiceServer", reflect.TypeOf((*MockInvitesServiceServer)(nil).mustEmbedUnimplementedInvitesServiceServer))
}

// MockUnsafeInvitesServiceServer is a mock of UnsafeInvitesServiceServer interface.
type MockUnsafeInvitesServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeInvitesServiceServerMockRecorder
}

// MockUnsafeInvitesServiceServerMockRecorder is the mock recorder for MockUnsafeInvitesServiceServer.
type MockUnsafeInvitesServiceServerMockRecorder struct {
	mock *MockUnsafeInvitesServiceServer
}

// NewMockUnsafeInvitesServiceServer creates a new mock instance.
func NewMockUnsafeInvitesServiceServer(ctrl *gomock.Controller) *MockUnsafeInvitesServiceServer {
	mock := &MockUnsafeInvitesServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeInvitesServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeInvitesServiceServer) EXPECT() *MockUnsafeInvitesServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedInvitesServiceServer mocks base method.
func (m *MockUnsafeInvitesServiceServer) mustEmbedUnimplementedInvitesServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedInvitesServiceServer")
}

// mustEmbedUnimplementedInvitesServiceServer indicates an expected call of mustEmbedUnimplementedInvitesServiceServer.
func (mr *MockUnsafeInvitesServiceServerMockRecorder) mustEmbedUnimplementedInvitesServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedInvitesServiceServer", reflect.TypeOf((*MockUnsafeInvitesServiceServer)(nil).mustEmbedUnimplementedInvitesServiceServer))
}