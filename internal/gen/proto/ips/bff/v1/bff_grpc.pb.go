// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: ips/bff/v1/bff.proto

package bffv1

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

// BFFServiceClient is the client API for BFFService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BFFServiceClient interface {
	// Get
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type bFFServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBFFServiceClient(cc grpc.ClientConnInterface) BFFServiceClient {
	return &bFFServiceClient{cc}
}

func (c *bFFServiceClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/ips.bff.v1.BFFService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BFFServiceServer is the server API for BFFService service.
// All implementations must embed UnimplementedBFFServiceServer
// for forward compatibility
type BFFServiceServer interface {
	// Get
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	mustEmbedUnimplementedBFFServiceServer()
}

// UnimplementedBFFServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBFFServiceServer struct {
}

func (UnimplementedBFFServiceServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedBFFServiceServer) mustEmbedUnimplementedBFFServiceServer() {}

// UnsafeBFFServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BFFServiceServer will
// result in compilation errors.
type UnsafeBFFServiceServer interface {
	mustEmbedUnimplementedBFFServiceServer()
}

func RegisterBFFServiceServer(s grpc.ServiceRegistrar, srv BFFServiceServer) {
	s.RegisterService(&BFFService_ServiceDesc, srv)
}

func _BFFService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFFServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ips.bff.v1.BFFService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFFServiceServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BFFService_ServiceDesc is the grpc.ServiceDesc for BFFService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BFFService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ips.bff.v1.BFFService",
	HandlerType: (*BFFServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _BFFService_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ips/bff/v1/bff.proto",
}