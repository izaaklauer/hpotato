// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: hpotato/v1/hpotato.proto

package hpotatov1

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

// HpotatoServiceClient is the client API for HpotatoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HpotatoServiceClient interface {
	// HelloWorld is the default rpc. Feel free to delete.
	HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error)
}

type hpotatoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHpotatoServiceClient(cc grpc.ClientConnInterface) HpotatoServiceClient {
	return &hpotatoServiceClient{cc}
}

func (c *hpotatoServiceClient) HelloWorld(ctx context.Context, in *HelloWorldRequest, opts ...grpc.CallOption) (*HelloWorldResponse, error) {
	out := new(HelloWorldResponse)
	err := c.cc.Invoke(ctx, "/hpotato.v1.HpotatoService/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HpotatoServiceServer is the server API for HpotatoService service.
// All implementations must embed UnimplementedHpotatoServiceServer
// for forward compatibility
type HpotatoServiceServer interface {
	// HelloWorld is the default rpc. Feel free to delete.
	HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error)
	mustEmbedUnimplementedHpotatoServiceServer()
}

// UnimplementedHpotatoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHpotatoServiceServer struct {
}

func (UnimplementedHpotatoServiceServer) HelloWorld(context.Context, *HelloWorldRequest) (*HelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (UnimplementedHpotatoServiceServer) mustEmbedUnimplementedHpotatoServiceServer() {
}

// UnsafeHpotatoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HpotatoServiceServer will
// result in compilation errors.
type UnsafeHpotatoServiceServer interface {
	mustEmbedUnimplementedHpotatoServiceServer()
}

func RegisterHpotatoServiceServer(s grpc.ServiceRegistrar, srv HpotatoServiceServer) {
	s.RegisterService(&HpotatoService_ServiceDesc, srv)
}

func _HpotatoService_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HpotatoServiceServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hpotato.v1.HpotatoService/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HpotatoServiceServer).HelloWorld(ctx, req.(*HelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HpotatoService_ServiceDesc is the grpc.ServiceDesc for HpotatoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HpotatoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hpotato.v1.HpotatoService",
	HandlerType: (*HpotatoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _HpotatoService_HelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hpotato/v1/hpotato.proto",
}
