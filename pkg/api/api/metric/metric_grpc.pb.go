// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: pkg/api/api/metric/metric.proto

package metric

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

// MetricCollectorClient is the client API for MetricCollector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricCollectorClient interface {
	GetMultiMetric(ctx context.Context, in *Request, opts ...grpc.CallOption) (*MultiMetric, error)
}

type metricCollectorClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricCollectorClient(cc grpc.ClientConnInterface) MetricCollectorClient {
	return &metricCollectorClient{cc}
}

func (c *metricCollectorClient) GetMultiMetric(ctx context.Context, in *Request, opts ...grpc.CallOption) (*MultiMetric, error) {
	out := new(MultiMetric)
	err := c.cc.Invoke(ctx, "/metric.MetricCollector/GetMultiMetric", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricCollectorServer is the server API for MetricCollector service.
// All implementations must embed UnimplementedMetricCollectorServer
// for forward compatibility
type MetricCollectorServer interface {
	GetMultiMetric(context.Context, *Request) (*MultiMetric, error)
	mustEmbedUnimplementedMetricCollectorServer()
}

// UnimplementedMetricCollectorServer must be embedded to have forward compatible implementations.
type UnimplementedMetricCollectorServer struct {
}

func (UnimplementedMetricCollectorServer) GetMultiMetric(context.Context, *Request) (*MultiMetric, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMultiMetric not implemented")
}
func (UnimplementedMetricCollectorServer) mustEmbedUnimplementedMetricCollectorServer() {}

// UnsafeMetricCollectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricCollectorServer will
// result in compilation errors.
type UnsafeMetricCollectorServer interface {
	mustEmbedUnimplementedMetricCollectorServer()
}

func RegisterMetricCollectorServer(s grpc.ServiceRegistrar, srv MetricCollectorServer) {
	s.RegisterService(&MetricCollector_ServiceDesc, srv)
}

func _MetricCollector_GetMultiMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricCollectorServer).GetMultiMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metric.MetricCollector/GetMultiMetric",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricCollectorServer).GetMultiMetric(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricCollector_ServiceDesc is the grpc.ServiceDesc for MetricCollector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricCollector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metric.MetricCollector",
	HandlerType: (*MetricCollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMultiMetric",
			Handler:    _MetricCollector_GetMultiMetric_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/api/metric/metric.proto",
}
