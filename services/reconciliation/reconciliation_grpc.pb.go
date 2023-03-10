// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: kappa/services/reconciliation/reconciliation.proto

package reconciliation

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ReconciliationServiceClient is the client API for ReconciliationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReconciliationServiceClient interface {
	CreateExternalTransactionNotification(ctx context.Context, in *CreateExternalTransactionNotificationRequest, opts ...grpc.CallOption) (*CreateExternalTransactionNotificationResponse, error)
	SetMatchProbabilities(ctx context.Context, in *SetMatchProbabilitiesRequest, opts ...grpc.CallOption) (*SetMatchProbabilitiesResponse, error)
	SetTransactionID(ctx context.Context, in *SetTransactionIDRequest, opts ...grpc.CallOption) (*SetTransactionIDResponse, error)
	SetTransactionReconciliationStatus(ctx context.Context, in *SetTransactionReconciliationStatusRequest, opts ...grpc.CallOption) (*SetTransactionReconciliationStatusResponse, error)
}

type reconciliationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReconciliationServiceClient(cc grpc.ClientConnInterface) ReconciliationServiceClient {
	return &reconciliationServiceClient{cc}
}

func (c *reconciliationServiceClient) CreateExternalTransactionNotification(ctx context.Context, in *CreateExternalTransactionNotificationRequest, opts ...grpc.CallOption) (*CreateExternalTransactionNotificationResponse, error) {
	out := new(CreateExternalTransactionNotificationResponse)
	err := c.cc.Invoke(ctx, "/payout.ReconciliationService/CreateExternalTransactionNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reconciliationServiceClient) SetMatchProbabilities(ctx context.Context, in *SetMatchProbabilitiesRequest, opts ...grpc.CallOption) (*SetMatchProbabilitiesResponse, error) {
	out := new(SetMatchProbabilitiesResponse)
	err := c.cc.Invoke(ctx, "/payout.ReconciliationService/SetMatchProbabilities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reconciliationServiceClient) SetTransactionID(ctx context.Context, in *SetTransactionIDRequest, opts ...grpc.CallOption) (*SetTransactionIDResponse, error) {
	out := new(SetTransactionIDResponse)
	err := c.cc.Invoke(ctx, "/payout.ReconciliationService/SetTransactionID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reconciliationServiceClient) SetTransactionReconciliationStatus(ctx context.Context, in *SetTransactionReconciliationStatusRequest, opts ...grpc.CallOption) (*SetTransactionReconciliationStatusResponse, error) {
	out := new(SetTransactionReconciliationStatusResponse)
	err := c.cc.Invoke(ctx, "/payout.ReconciliationService/SetTransactionReconciliationStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReconciliationServiceServer is the server API for ReconciliationService service.
// All implementations must embed UnimplementedReconciliationServiceServer
// for forward compatibility
type ReconciliationServiceServer interface {
	CreateExternalTransactionNotification(context.Context, *CreateExternalTransactionNotificationRequest) (*CreateExternalTransactionNotificationResponse, error)
	SetMatchProbabilities(context.Context, *SetMatchProbabilitiesRequest) (*SetMatchProbabilitiesResponse, error)
	SetTransactionID(context.Context, *SetTransactionIDRequest) (*SetTransactionIDResponse, error)
	SetTransactionReconciliationStatus(context.Context, *SetTransactionReconciliationStatusRequest) (*SetTransactionReconciliationStatusResponse, error)
	mustEmbedUnimplementedReconciliationServiceServer()
}

// UnimplementedReconciliationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReconciliationServiceServer struct {
}

func (UnimplementedReconciliationServiceServer) CreateExternalTransactionNotification(context.Context, *CreateExternalTransactionNotificationRequest) (*CreateExternalTransactionNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExternalTransactionNotification not implemented")
}
func (UnimplementedReconciliationServiceServer) SetMatchProbabilities(context.Context, *SetMatchProbabilitiesRequest) (*SetMatchProbabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMatchProbabilities not implemented")
}
func (UnimplementedReconciliationServiceServer) SetTransactionID(context.Context, *SetTransactionIDRequest) (*SetTransactionIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTransactionID not implemented")
}
func (UnimplementedReconciliationServiceServer) SetTransactionReconciliationStatus(context.Context, *SetTransactionReconciliationStatusRequest) (*SetTransactionReconciliationStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTransactionReconciliationStatus not implemented")
}
func (UnimplementedReconciliationServiceServer) mustEmbedUnimplementedReconciliationServiceServer() {}

// UnsafeReconciliationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReconciliationServiceServer will
// result in compilation errors.
type UnsafeReconciliationServiceServer interface {
	mustEmbedUnimplementedReconciliationServiceServer()
}

func RegisterReconciliationServiceServer(s grpc.ServiceRegistrar, srv ReconciliationServiceServer) {
	s.RegisterService(&ReconciliationService_ServiceDesc, srv)
}

func _ReconciliationService_CreateExternalTransactionNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExternalTransactionNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReconciliationServiceServer).CreateExternalTransactionNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payout.ReconciliationService/CreateExternalTransactionNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReconciliationServiceServer).CreateExternalTransactionNotification(ctx, req.(*CreateExternalTransactionNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReconciliationService_SetMatchProbabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMatchProbabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReconciliationServiceServer).SetMatchProbabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payout.ReconciliationService/SetMatchProbabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReconciliationServiceServer).SetMatchProbabilities(ctx, req.(*SetMatchProbabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReconciliationService_SetTransactionID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTransactionIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReconciliationServiceServer).SetTransactionID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payout.ReconciliationService/SetTransactionID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReconciliationServiceServer).SetTransactionID(ctx, req.(*SetTransactionIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReconciliationService_SetTransactionReconciliationStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTransactionReconciliationStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReconciliationServiceServer).SetTransactionReconciliationStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payout.ReconciliationService/SetTransactionReconciliationStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReconciliationServiceServer).SetTransactionReconciliationStatus(ctx, req.(*SetTransactionReconciliationStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReconciliationService_ServiceDesc is the grpc.ServiceDesc for ReconciliationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReconciliationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payout.ReconciliationService",
	HandlerType: (*ReconciliationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExternalTransactionNotification",
			Handler:    _ReconciliationService_CreateExternalTransactionNotification_Handler,
		},
		{
			MethodName: "SetMatchProbabilities",
			Handler:    _ReconciliationService_SetMatchProbabilities_Handler,
		},
		{
			MethodName: "SetTransactionID",
			Handler:    _ReconciliationService_SetTransactionID_Handler,
		},
		{
			MethodName: "SetTransactionReconciliationStatus",
			Handler:    _ReconciliationService_SetTransactionReconciliationStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kappa/services/reconciliation/reconciliation.proto",
}
