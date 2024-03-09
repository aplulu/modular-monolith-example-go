// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: example/user/v1/internal.proto

package user

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

const (
	InternalUserService_GetUser_FullMethodName = "/example.user.v1.InternalUserService/GetUser"
)

// InternalUserServiceClient is the client API for InternalUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalUserServiceClient interface {
	// ユーザを取得
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type internalUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalUserServiceClient(cc grpc.ClientConnInterface) InternalUserServiceClient {
	return &internalUserServiceClient{cc}
}

func (c *internalUserServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, InternalUserService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalUserServiceServer is the server API for InternalUserService service.
// All implementations should embed UnimplementedInternalUserServiceServer
// for forward compatibility
type InternalUserServiceServer interface {
	// ユーザを取得
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
}

// UnimplementedInternalUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInternalUserServiceServer struct {
}

func (UnimplementedInternalUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

// UnsafeInternalUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalUserServiceServer will
// result in compilation errors.
type UnsafeInternalUserServiceServer interface {
	mustEmbedUnimplementedInternalUserServiceServer()
}

func RegisterInternalUserServiceServer(s grpc.ServiceRegistrar, srv InternalUserServiceServer) {
	s.RegisterService(&InternalUserService_ServiceDesc, srv)
}

func _InternalUserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalUserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InternalUserService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalUserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InternalUserService_ServiceDesc is the grpc.ServiceDesc for InternalUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.user.v1.InternalUserService",
	HandlerType: (*InternalUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _InternalUserService_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/user/v1/internal.proto",
}
