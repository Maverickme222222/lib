// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: kappa/services/authorization/authorization.proto

package authorization

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

// AuthorizationServiceClient is the client API for AuthorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	GetTransactionAuthRequirements(ctx context.Context, in *GetTransactionAuthRequirementsRequest, opts ...grpc.CallOption) (*GetTransactionAuthRequirementsResponse, error)
	GetDestinationAuthRequirements(ctx context.Context, in *GetDestinationAuthRequirementsRequest, opts ...grpc.CallOption) (*GetDestinationAuthRequirementsResponse, error)
	GetLinkedAccountAuthRequirements(ctx context.Context, in *GetLinkedAccountAuthRequirementsRequest, opts ...grpc.CallOption) (*GetLinkedAccountAuthRequirementsResponse, error)
	CreateTransactionAuthRequirements(ctx context.Context, in *CreateTransactionAuthRequirementsRequest, opts ...grpc.CallOption) (*CreateTransactionAuthRequirementsResponse, error)
	CreateDestinationAuthRequirements(ctx context.Context, in *CreateDestinationAuthRequirementsRequest, opts ...grpc.CallOption) (*CreateDestinationAuthRequirementsResponse, error)
	CreateLinkedAccountAuthRequirements(ctx context.Context, in *CreateLinkedAccountAuthRequirementsRequest, opts ...grpc.CallOption) (*CreateLinkedAccountAuthRequirementsResponse, error)
	UpdateAuthRequirement(ctx context.Context, in *UpdateAuthRequirementRequest, opts ...grpc.CallOption) (*UpdateAuthRequirementResponse, error)
	GetAuthRequirementByID(ctx context.Context, in *GetAuthRequirementByIDRequest, opts ...grpc.CallOption) (*GetAuthRequirementByIDResponse, error)
}

type authorizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationServiceClient(cc grpc.ClientConnInterface) AuthorizationServiceClient {
	return &authorizationServiceClient{cc}
}

func (c *authorizationServiceClient) GetTransactionAuthRequirements(ctx context.Context, in *GetTransactionAuthRequirementsRequest, opts ...grpc.CallOption) (*GetTransactionAuthRequirementsResponse, error) {
	out := new(GetTransactionAuthRequirementsResponse)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/GetTransactionAuthRequirements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetDestinationAuthRequirements(ctx context.Context, in *GetDestinationAuthRequirementsRequest, opts ...grpc.CallOption) (*GetDestinationAuthRequirementsResponse, error) {
	out := new(GetDestinationAuthRequirementsResponse)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/GetDestinationAuthRequirements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetLinkedAccountAuthRequirements(ctx context.Context, in *GetLinkedAccountAuthRequirementsRequest, opts ...grpc.CallOption) (*GetLinkedAccountAuthRequirementsResponse, error) {
	out := new(GetLinkedAccountAuthRequirementsResponse)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/GetLinkedAccountAuthRequirements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) CreateTransactionAuthRequirements(ctx context.Context, in *CreateTransactionAuthRequirementsRequest, opts ...grpc.CallOption) (*CreateTransactionAuthRequirementsResponse, error) {
	out := new(CreateTransactionAuthRequirementsResponse)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/CreateTransactionAuthRequirements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) CreateDestinationAuthRequirements(ctx context.Context, in *CreateDestinationAuthRequirementsRequest, opts ...grpc.CallOption) (*CreateDestinationAuthRequirementsResponse, error) {
	out := new(CreateDestinationAuthRequirementsResponse)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/CreateDestinationAuthRequirements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) CreateLinkedAccountAuthRequirements(ctx context.Context, in *CreateLinkedAccountAuthRequirementsRequest, opts ...grpc.CallOption) (*CreateLinkedAccountAuthRequirementsResponse, error) {
	out := new(CreateLinkedAccountAuthRequirementsResponse)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/CreateLinkedAccountAuthRequirements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) UpdateAuthRequirement(ctx context.Context, in *UpdateAuthRequirementRequest, opts ...grpc.CallOption) (*UpdateAuthRequirementResponse, error) {
	out := new(UpdateAuthRequirementResponse)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/UpdateAuthRequirement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetAuthRequirementByID(ctx context.Context, in *GetAuthRequirementByIDRequest, opts ...grpc.CallOption) (*GetAuthRequirementByIDResponse, error) {
	out := new(GetAuthRequirementByIDResponse)
	err := c.cc.Invoke(ctx, "/authorization.AuthorizationService/GetAuthRequirementByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServiceServer is the server API for AuthorizationService service.
// All implementations must embed UnimplementedAuthorizationServiceServer
// for forward compatibility
type AuthorizationServiceServer interface {
	GetTransactionAuthRequirements(context.Context, *GetTransactionAuthRequirementsRequest) (*GetTransactionAuthRequirementsResponse, error)
	GetDestinationAuthRequirements(context.Context, *GetDestinationAuthRequirementsRequest) (*GetDestinationAuthRequirementsResponse, error)
	GetLinkedAccountAuthRequirements(context.Context, *GetLinkedAccountAuthRequirementsRequest) (*GetLinkedAccountAuthRequirementsResponse, error)
	CreateTransactionAuthRequirements(context.Context, *CreateTransactionAuthRequirementsRequest) (*CreateTransactionAuthRequirementsResponse, error)
	CreateDestinationAuthRequirements(context.Context, *CreateDestinationAuthRequirementsRequest) (*CreateDestinationAuthRequirementsResponse, error)
	CreateLinkedAccountAuthRequirements(context.Context, *CreateLinkedAccountAuthRequirementsRequest) (*CreateLinkedAccountAuthRequirementsResponse, error)
	UpdateAuthRequirement(context.Context, *UpdateAuthRequirementRequest) (*UpdateAuthRequirementResponse, error)
	GetAuthRequirementByID(context.Context, *GetAuthRequirementByIDRequest) (*GetAuthRequirementByIDResponse, error)
	mustEmbedUnimplementedAuthorizationServiceServer()
}

// UnimplementedAuthorizationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServiceServer struct {
}

func (UnimplementedAuthorizationServiceServer) GetTransactionAuthRequirements(context.Context, *GetTransactionAuthRequirementsRequest) (*GetTransactionAuthRequirementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionAuthRequirements not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetDestinationAuthRequirements(context.Context, *GetDestinationAuthRequirementsRequest) (*GetDestinationAuthRequirementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinationAuthRequirements not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetLinkedAccountAuthRequirements(context.Context, *GetLinkedAccountAuthRequirementsRequest) (*GetLinkedAccountAuthRequirementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinkedAccountAuthRequirements not implemented")
}
func (UnimplementedAuthorizationServiceServer) CreateTransactionAuthRequirements(context.Context, *CreateTransactionAuthRequirementsRequest) (*CreateTransactionAuthRequirementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransactionAuthRequirements not implemented")
}
func (UnimplementedAuthorizationServiceServer) CreateDestinationAuthRequirements(context.Context, *CreateDestinationAuthRequirementsRequest) (*CreateDestinationAuthRequirementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDestinationAuthRequirements not implemented")
}
func (UnimplementedAuthorizationServiceServer) CreateLinkedAccountAuthRequirements(context.Context, *CreateLinkedAccountAuthRequirementsRequest) (*CreateLinkedAccountAuthRequirementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLinkedAccountAuthRequirements not implemented")
}
func (UnimplementedAuthorizationServiceServer) UpdateAuthRequirement(context.Context, *UpdateAuthRequirementRequest) (*UpdateAuthRequirementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthRequirement not implemented")
}
func (UnimplementedAuthorizationServiceServer) GetAuthRequirementByID(context.Context, *GetAuthRequirementByIDRequest) (*GetAuthRequirementByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthRequirementByID not implemented")
}
func (UnimplementedAuthorizationServiceServer) mustEmbedUnimplementedAuthorizationServiceServer() {}

// UnsafeAuthorizationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizationServiceServer will
// result in compilation errors.
type UnsafeAuthorizationServiceServer interface {
	mustEmbedUnimplementedAuthorizationServiceServer()
}

func RegisterAuthorizationServiceServer(s grpc.ServiceRegistrar, srv AuthorizationServiceServer) {
	s.RegisterService(&AuthorizationService_ServiceDesc, srv)
}

func _AuthorizationService_GetTransactionAuthRequirements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionAuthRequirementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetTransactionAuthRequirements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/GetTransactionAuthRequirements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetTransactionAuthRequirements(ctx, req.(*GetTransactionAuthRequirementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetDestinationAuthRequirements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDestinationAuthRequirementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetDestinationAuthRequirements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/GetDestinationAuthRequirements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetDestinationAuthRequirements(ctx, req.(*GetDestinationAuthRequirementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetLinkedAccountAuthRequirements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinkedAccountAuthRequirementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetLinkedAccountAuthRequirements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/GetLinkedAccountAuthRequirements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetLinkedAccountAuthRequirements(ctx, req.(*GetLinkedAccountAuthRequirementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_CreateTransactionAuthRequirements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransactionAuthRequirementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).CreateTransactionAuthRequirements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/CreateTransactionAuthRequirements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).CreateTransactionAuthRequirements(ctx, req.(*CreateTransactionAuthRequirementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_CreateDestinationAuthRequirements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDestinationAuthRequirementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).CreateDestinationAuthRequirements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/CreateDestinationAuthRequirements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).CreateDestinationAuthRequirements(ctx, req.(*CreateDestinationAuthRequirementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_CreateLinkedAccountAuthRequirements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLinkedAccountAuthRequirementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).CreateLinkedAccountAuthRequirements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/CreateLinkedAccountAuthRequirements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).CreateLinkedAccountAuthRequirements(ctx, req.(*CreateLinkedAccountAuthRequirementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_UpdateAuthRequirement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthRequirementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).UpdateAuthRequirement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/UpdateAuthRequirement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).UpdateAuthRequirement(ctx, req.(*UpdateAuthRequirementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetAuthRequirementByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthRequirementByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).GetAuthRequirementByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.AuthorizationService/GetAuthRequirementByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).GetAuthRequirementByID(ctx, req.(*GetAuthRequirementByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorizationService_ServiceDesc is the grpc.ServiceDesc for AuthorizationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorizationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authorization.AuthorizationService",
	HandlerType: (*AuthorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransactionAuthRequirements",
			Handler:    _AuthorizationService_GetTransactionAuthRequirements_Handler,
		},
		{
			MethodName: "GetDestinationAuthRequirements",
			Handler:    _AuthorizationService_GetDestinationAuthRequirements_Handler,
		},
		{
			MethodName: "GetLinkedAccountAuthRequirements",
			Handler:    _AuthorizationService_GetLinkedAccountAuthRequirements_Handler,
		},
		{
			MethodName: "CreateTransactionAuthRequirements",
			Handler:    _AuthorizationService_CreateTransactionAuthRequirements_Handler,
		},
		{
			MethodName: "CreateDestinationAuthRequirements",
			Handler:    _AuthorizationService_CreateDestinationAuthRequirements_Handler,
		},
		{
			MethodName: "CreateLinkedAccountAuthRequirements",
			Handler:    _AuthorizationService_CreateLinkedAccountAuthRequirements_Handler,
		},
		{
			MethodName: "UpdateAuthRequirement",
			Handler:    _AuthorizationService_UpdateAuthRequirement_Handler,
		},
		{
			MethodName: "GetAuthRequirementByID",
			Handler:    _AuthorizationService_GetAuthRequirementByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kappa/services/authorization/authorization.proto",
}