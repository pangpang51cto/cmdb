// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mongodb

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	SaveMongoDB(ctx context.Context, in *MongoDB, opts ...grpc.CallOption) (*MongoDB, error)
	QueryMongoDB(ctx context.Context, in *QueryMongoDBRequest, opts ...grpc.CallOption) (*Set, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) SaveMongoDB(ctx context.Context, in *MongoDB, opts ...grpc.CallOption) (*MongoDB, error) {
	out := new(MongoDB)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.mongodb.Service/SaveMongoDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryMongoDB(ctx context.Context, in *QueryMongoDBRequest, opts ...grpc.CallOption) (*Set, error) {
	out := new(Set)
	err := c.cc.Invoke(ctx, "/infraboard.cmdb.mongodb.Service/QueryMongoDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	SaveMongoDB(context.Context, *MongoDB) (*MongoDB, error)
	QueryMongoDB(context.Context, *QueryMongoDBRequest) (*Set, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) SaveMongoDB(context.Context, *MongoDB) (*MongoDB, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveMongoDB not implemented")
}
func (UnimplementedServiceServer) QueryMongoDB(context.Context, *QueryMongoDBRequest) (*Set, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMongoDB not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_SaveMongoDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MongoDB)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).SaveMongoDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.mongodb.Service/SaveMongoDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).SaveMongoDB(ctx, req.(*MongoDB))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryMongoDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMongoDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryMongoDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infraboard.cmdb.mongodb.Service/QueryMongoDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryMongoDB(ctx, req.(*QueryMongoDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "infraboard.cmdb.mongodb.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveMongoDB",
			Handler:    _Service_SaveMongoDB_Handler,
		},
		{
			MethodName: "QueryMongoDB",
			Handler:    _Service_QueryMongoDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/mongodb/pb/mongodb.proto",
}
