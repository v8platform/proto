// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apiv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClusterAdminServiceClient is the client API for ClusterAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterAdminServiceClient interface {
	Admins(ctx context.Context, in *AdminsRequest, opts ...grpc.CallOption) (*AdminsResponse, error)
	AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAdmin(ctx context.Context, in *DeleteAdminRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type clusterAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterAdminServiceClient(cc grpc.ClientConnInterface) ClusterAdminServiceClient {
	return &clusterAdminServiceClient{cc}
}

func (c *clusterAdminServiceClient) Admins(ctx context.Context, in *AdminsRequest, opts ...grpc.CallOption) (*AdminsResponse, error) {
	out := new(AdminsResponse)
	err := c.cc.Invoke(ctx, "/ras.api.v1.ClusterAdminService/Admins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterAdminServiceClient) AddAdmin(ctx context.Context, in *AddAdminRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ras.api.v1.ClusterAdminService/AddAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterAdminServiceClient) DeleteAdmin(ctx context.Context, in *DeleteAdminRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ras.api.v1.ClusterAdminService/DeleteAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterAdminServiceServer is the server API for ClusterAdminService service.
// All implementations should embed UnimplementedClusterAdminServiceServer
// for forward compatibility
type ClusterAdminServiceServer interface {
	Admins(context.Context, *AdminsRequest) (*AdminsResponse, error)
	AddAdmin(context.Context, *AddAdminRequest) (*emptypb.Empty, error)
	DeleteAdmin(context.Context, *DeleteAdminRequest) (*emptypb.Empty, error)
}

// UnimplementedClusterAdminServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClusterAdminServiceServer struct {
}

func (UnimplementedClusterAdminServiceServer) Admins(context.Context, *AdminsRequest) (*AdminsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Admins not implemented")
}
func (UnimplementedClusterAdminServiceServer) AddAdmin(context.Context, *AddAdminRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAdmin not implemented")
}
func (UnimplementedClusterAdminServiceServer) DeleteAdmin(context.Context, *DeleteAdminRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAdmin not implemented")
}

// UnsafeClusterAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterAdminServiceServer will
// result in compilation errors.
type UnsafeClusterAdminServiceServer interface {
	mustEmbedUnimplementedClusterAdminServiceServer()
}

func RegisterClusterAdminServiceServer(s grpc.ServiceRegistrar, srv ClusterAdminServiceServer) {
	s.RegisterService(&ClusterAdminService_ServiceDesc, srv)
}

func _ClusterAdminService_Admins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterAdminServiceServer).Admins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.api.v1.ClusterAdminService/Admins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterAdminServiceServer).Admins(ctx, req.(*AdminsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterAdminService_AddAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterAdminServiceServer).AddAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.api.v1.ClusterAdminService/AddAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterAdminServiceServer).AddAdmin(ctx, req.(*AddAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterAdminService_DeleteAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterAdminServiceServer).DeleteAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ras.api.v1.ClusterAdminService/DeleteAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterAdminServiceServer).DeleteAdmin(ctx, req.(*DeleteAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterAdminService_ServiceDesc is the grpc.ServiceDesc for ClusterAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ras.api.v1.ClusterAdminService",
	HandlerType: (*ClusterAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Admins",
			Handler:    _ClusterAdminService_Admins_Handler,
		},
		{
			MethodName: "AddAdmin",
			Handler:    _ClusterAdminService_AddAdmin_Handler,
		},
		{
			MethodName: "DeleteAdmin",
			Handler:    _ClusterAdminService_DeleteAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ras/api/v1/cluster_admin.proto",
}
