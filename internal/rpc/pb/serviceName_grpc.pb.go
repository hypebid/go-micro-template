// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ServiceNameClient is the client API for ServiceName service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceNameClient interface {
	// Used to check on the status of the service and all it's dependencies
	HealthCheck(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthStatus, error)
}

type serviceNameClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceNameClient(cc grpc.ClientConnInterface) ServiceNameClient {
	return &serviceNameClient{cc}
}

func (c *serviceNameClient) HealthCheck(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthStatus, error) {
	out := new(HealthStatus)
	err := c.cc.Invoke(ctx, "/hypebid.ServiceName/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceNameServer is the server API for ServiceName service.
// All implementations must embed UnimplementedServiceNameServer
// for forward compatibility
type ServiceNameServer interface {
	// Used to check on the status of the service and all it's dependencies
	HealthCheck(context.Context, *HealthRequest) (*HealthStatus, error)
	mustEmbedUnimplementedServiceNameServer()
}

// UnimplementedServiceNameServer must be embedded to have forward compatible implementations.
type UnimplementedServiceNameServer struct {
}

func (UnimplementedServiceNameServer) HealthCheck(context.Context, *HealthRequest) (*HealthStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedServiceNameServer) mustEmbedUnimplementedServiceNameServer() {}

// UnsafeServiceNameServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceNameServer will
// result in compilation errors.
type UnsafeServiceNameServer interface {
	mustEmbedUnimplementedServiceNameServer()
}

func RegisterServiceNameServer(s grpc.ServiceRegistrar, srv ServiceNameServer) {
	s.RegisterService(&ServiceName_ServiceDesc, srv)
}

func _ServiceName_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceNameServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hypebid.ServiceName/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceNameServer).HealthCheck(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceName_ServiceDesc is the grpc.ServiceDesc for ServiceName service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceName_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hypebid.ServiceName",
	HandlerType: (*ServiceNameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ServiceName_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serviceName.proto",
}
