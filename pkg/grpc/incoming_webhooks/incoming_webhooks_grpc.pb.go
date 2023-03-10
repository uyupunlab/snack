// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: api/incoming_webhooks.proto

package incoming_webhooks

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

// IncomingWebhooksServiceClient is the client API for IncomingWebhooksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IncomingWebhooksServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
}

type incomingWebhooksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIncomingWebhooksServiceClient(cc grpc.ClientConnInterface) IncomingWebhooksServiceClient {
	return &incomingWebhooksServiceClient{cc}
}

func (c *incomingWebhooksServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/snack.IncomingWebhooksService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incomingWebhooksServiceClient) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/snack.IncomingWebhooksService/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IncomingWebhooksServiceServer is the server API for IncomingWebhooksService service.
// All implementations must embed UnimplementedIncomingWebhooksServiceServer
// for forward compatibility
type IncomingWebhooksServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Post(context.Context, *PostRequest) (*PostResponse, error)
	mustEmbedUnimplementedIncomingWebhooksServiceServer()
}

// UnimplementedIncomingWebhooksServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIncomingWebhooksServiceServer struct {
}

func (UnimplementedIncomingWebhooksServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedIncomingWebhooksServiceServer) Post(context.Context, *PostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedIncomingWebhooksServiceServer) mustEmbedUnimplementedIncomingWebhooksServiceServer() {
}

// UnsafeIncomingWebhooksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IncomingWebhooksServiceServer will
// result in compilation errors.
type UnsafeIncomingWebhooksServiceServer interface {
	mustEmbedUnimplementedIncomingWebhooksServiceServer()
}

func RegisterIncomingWebhooksServiceServer(s grpc.ServiceRegistrar, srv IncomingWebhooksServiceServer) {
	s.RegisterService(&IncomingWebhooksService_ServiceDesc, srv)
}

func _IncomingWebhooksService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncomingWebhooksServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.IncomingWebhooksService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncomingWebhooksServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncomingWebhooksService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncomingWebhooksServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.IncomingWebhooksService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncomingWebhooksServiceServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IncomingWebhooksService_ServiceDesc is the grpc.ServiceDesc for IncomingWebhooksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IncomingWebhooksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "snack.IncomingWebhooksService",
	HandlerType: (*IncomingWebhooksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _IncomingWebhooksService_Register_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _IncomingWebhooksService_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/incoming_webhooks.proto",
}
