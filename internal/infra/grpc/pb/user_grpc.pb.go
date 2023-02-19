// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: internal/infra/grpc/proto/user.proto

package pb

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

// ManageUserClient is the client API for ManageUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManageUserClient interface {
	RegisterUser(ctx context.Context, in *UserCreationInput, opts ...grpc.CallOption) (*ResponseMessage, error)
}

type manageUserClient struct {
	cc grpc.ClientConnInterface
}

func NewManageUserClient(cc grpc.ClientConnInterface) ManageUserClient {
	return &manageUserClient{cc}
}

func (c *manageUserClient) RegisterUser(ctx context.Context, in *UserCreationInput, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/pb.ManageUser/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManageUserServer is the server API for ManageUser service.
// All implementations must embed UnimplementedManageUserServer
// for forward compatibility
type ManageUserServer interface {
	RegisterUser(context.Context, *UserCreationInput) (*ResponseMessage, error)
	mustEmbedUnimplementedManageUserServer()
}

// UnimplementedManageUserServer must be embedded to have forward compatible implementations.
type UnimplementedManageUserServer struct {
}

func (UnimplementedManageUserServer) RegisterUser(context.Context, *UserCreationInput) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedManageUserServer) mustEmbedUnimplementedManageUserServer() {}

// UnsafeManageUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManageUserServer will
// result in compilation errors.
type UnsafeManageUserServer interface {
	mustEmbedUnimplementedManageUserServer()
}

func RegisterManageUserServer(s grpc.ServiceRegistrar, srv ManageUserServer) {
	s.RegisterService(&ManageUser_ServiceDesc, srv)
}

func _ManageUser_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCreationInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManageUserServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ManageUser/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManageUserServer).RegisterUser(ctx, req.(*UserCreationInput))
	}
	return interceptor(ctx, in, info, handler)
}

// ManageUser_ServiceDesc is the grpc.ServiceDesc for ManageUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManageUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ManageUser",
	HandlerType: (*ManageUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _ManageUser_RegisterUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/infra/grpc/proto/user.proto",
}
