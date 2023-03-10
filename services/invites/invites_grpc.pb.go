// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: kappa/services/invites/invites.proto

package invites

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

// InvitesServiceClient is the client API for InvitesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InvitesServiceClient interface {
	CreateInvite(ctx context.Context, in *CreateInviteRequest, opts ...grpc.CallOption) (*CreateInviteResponse, error)
}

type invitesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvitesServiceClient(cc grpc.ClientConnInterface) InvitesServiceClient {
	return &invitesServiceClient{cc}
}

func (c *invitesServiceClient) CreateInvite(ctx context.Context, in *CreateInviteRequest, opts ...grpc.CallOption) (*CreateInviteResponse, error) {
	out := new(CreateInviteResponse)
	err := c.cc.Invoke(ctx, "/addresses.InvitesService/CreateInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvitesServiceServer is the server API for InvitesService service.
// All implementations must embed UnimplementedInvitesServiceServer
// for forward compatibility
type InvitesServiceServer interface {
	CreateInvite(context.Context, *CreateInviteRequest) (*CreateInviteResponse, error)
	mustEmbedUnimplementedInvitesServiceServer()
}

// UnimplementedInvitesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInvitesServiceServer struct {
}

func (UnimplementedInvitesServiceServer) CreateInvite(context.Context, *CreateInviteRequest) (*CreateInviteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInvite not implemented")
}
func (UnimplementedInvitesServiceServer) mustEmbedUnimplementedInvitesServiceServer() {}

// UnsafeInvitesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InvitesServiceServer will
// result in compilation errors.
type UnsafeInvitesServiceServer interface {
	mustEmbedUnimplementedInvitesServiceServer()
}

func RegisterInvitesServiceServer(s grpc.ServiceRegistrar, srv InvitesServiceServer) {
	s.RegisterService(&InvitesService_ServiceDesc, srv)
}

func _InvitesService_CreateInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvitesServiceServer).CreateInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/addresses.InvitesService/CreateInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvitesServiceServer).CreateInvite(ctx, req.(*CreateInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InvitesService_ServiceDesc is the grpc.ServiceDesc for InvitesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InvitesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "addresses.InvitesService",
	HandlerType: (*InvitesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInvite",
			Handler:    _InvitesService_CreateInvite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kappa/services/invites/invites.proto",
}
