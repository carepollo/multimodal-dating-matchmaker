// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: protos/notifications-service.proto

package protos

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

// NotificationsServiceClient is the client API for NotificationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationsServiceClient interface {
	NotifyByEmail(ctx context.Context, in *NotifyEmailRequest, opts ...grpc.CallOption) (*NotifyEmailResponse, error)
}

type notificationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationsServiceClient(cc grpc.ClientConnInterface) NotificationsServiceClient {
	return &notificationsServiceClient{cc}
}

func (c *notificationsServiceClient) NotifyByEmail(ctx context.Context, in *NotifyEmailRequest, opts ...grpc.CallOption) (*NotifyEmailResponse, error) {
	out := new(NotifyEmailResponse)
	err := c.cc.Invoke(ctx, "/NotificationsService/NotifyByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServiceServer is the server API for NotificationsService service.
// All implementations must embed UnimplementedNotificationsServiceServer
// for forward compatibility
type NotificationsServiceServer interface {
	NotifyByEmail(context.Context, *NotifyEmailRequest) (*NotifyEmailResponse, error)
	mustEmbedUnimplementedNotificationsServiceServer()
}

// UnimplementedNotificationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationsServiceServer struct {
}

func (UnimplementedNotificationsServiceServer) NotifyByEmail(context.Context, *NotifyEmailRequest) (*NotifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyByEmail not implemented")
}
func (UnimplementedNotificationsServiceServer) mustEmbedUnimplementedNotificationsServiceServer() {}

// UnsafeNotificationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationsServiceServer will
// result in compilation errors.
type UnsafeNotificationsServiceServer interface {
	mustEmbedUnimplementedNotificationsServiceServer()
}

func RegisterNotificationsServiceServer(s grpc.ServiceRegistrar, srv NotificationsServiceServer) {
	s.RegisterService(&NotificationsService_ServiceDesc, srv)
}

func _NotificationsService_NotifyByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServiceServer).NotifyByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NotificationsService/NotifyByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServiceServer).NotifyByEmail(ctx, req.(*NotifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationsService_ServiceDesc is the grpc.ServiceDesc for NotificationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NotificationsService",
	HandlerType: (*NotificationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyByEmail",
			Handler:    _NotificationsService_NotifyByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/notifications-service.proto",
}