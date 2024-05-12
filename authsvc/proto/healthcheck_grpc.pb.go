// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: proto/healthcheck.proto

package proto

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

const (
	HealthCheckService_Ping_FullMethodName           = "/HealthCheckService/Ping"
	HealthCheckService_GetHealthCheck_FullMethodName = "/HealthCheckService/GetHealthCheck"
	HealthCheckService_GetReadiness_FullMethodName   = "/HealthCheckService/GetReadiness"
)

// HealthCheckServiceClient is the client API for HealthCheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthCheckServiceClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error)
	GetHealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetHealthCheckResponse, error)
	GetReadiness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetReadinessResponse, error)
}

type healthCheckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthCheckServiceClient(cc grpc.ClientConnInterface) HealthCheckServiceClient {
	return &healthCheckServiceClient{cc}
}

func (c *healthCheckServiceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, HealthCheckService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthCheckServiceClient) GetHealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetHealthCheckResponse, error) {
	out := new(GetHealthCheckResponse)
	err := c.cc.Invoke(ctx, HealthCheckService_GetHealthCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthCheckServiceClient) GetReadiness(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetReadinessResponse, error) {
	out := new(GetReadinessResponse)
	err := c.cc.Invoke(ctx, HealthCheckService_GetReadiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthCheckServiceServer is the server API for HealthCheckService service.
// All implementations must embed UnimplementedHealthCheckServiceServer
// for forward compatibility
type HealthCheckServiceServer interface {
	Ping(context.Context, *emptypb.Empty) (*PingResponse, error)
	GetHealthCheck(context.Context, *emptypb.Empty) (*GetHealthCheckResponse, error)
	GetReadiness(context.Context, *emptypb.Empty) (*GetReadinessResponse, error)
	mustEmbedUnimplementedHealthCheckServiceServer()
}

// UnimplementedHealthCheckServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHealthCheckServiceServer struct {
}

func (UnimplementedHealthCheckServiceServer) Ping(context.Context, *emptypb.Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedHealthCheckServiceServer) GetHealthCheck(context.Context, *emptypb.Empty) (*GetHealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealthCheck not implemented")
}
func (UnimplementedHealthCheckServiceServer) GetReadiness(context.Context, *emptypb.Empty) (*GetReadinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReadiness not implemented")
}
func (UnimplementedHealthCheckServiceServer) mustEmbedUnimplementedHealthCheckServiceServer() {}

// UnsafeHealthCheckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthCheckServiceServer will
// result in compilation errors.
type UnsafeHealthCheckServiceServer interface {
	mustEmbedUnimplementedHealthCheckServiceServer()
}

func RegisterHealthCheckServiceServer(s grpc.ServiceRegistrar, srv HealthCheckServiceServer) {
	s.RegisterService(&HealthCheckService_ServiceDesc, srv)
}

func _HealthCheckService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthCheckService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServiceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthCheckService_GetHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServiceServer).GetHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthCheckService_GetHealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServiceServer).GetHealthCheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthCheckService_GetReadiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServiceServer).GetReadiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HealthCheckService_GetReadiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServiceServer).GetReadiness(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthCheckService_ServiceDesc is the grpc.ServiceDesc for HealthCheckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthCheckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HealthCheckService",
	HandlerType: (*HealthCheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _HealthCheckService_Ping_Handler,
		},
		{
			MethodName: "GetHealthCheck",
			Handler:    _HealthCheckService_GetHealthCheck_Handler,
		},
		{
			MethodName: "GetReadiness",
			Handler:    _HealthCheckService_GetReadiness_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/healthcheck.proto",
}
