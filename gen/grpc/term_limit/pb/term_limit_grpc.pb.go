// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package term_limitpb

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

// TermLimitClient is the client API for TermLimit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TermLimitClient interface {
	// Get implements get.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type termLimitClient struct {
	cc grpc.ClientConnInterface
}

func NewTermLimitClient(cc grpc.ClientConnInterface) TermLimitClient {
	return &termLimitClient{cc}
}

func (c *termLimitClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/term_limit.TermLimit/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TermLimitServer is the server API for TermLimit service.
// All implementations must embed UnimplementedTermLimitServer
// for forward compatibility
type TermLimitServer interface {
	// Get implements get.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedTermLimitServer()
}

// UnimplementedTermLimitServer must be embedded to have forward compatible implementations.
type UnimplementedTermLimitServer struct {
}

func (UnimplementedTermLimitServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTermLimitServer) mustEmbedUnimplementedTermLimitServer() {}

// UnsafeTermLimitServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TermLimitServer will
// result in compilation errors.
type UnsafeTermLimitServer interface {
	mustEmbedUnimplementedTermLimitServer()
}

func RegisterTermLimitServer(s grpc.ServiceRegistrar, srv TermLimitServer) {
	s.RegisterService(&TermLimit_ServiceDesc, srv)
}

func _TermLimit_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermLimitServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/term_limit.TermLimit/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermLimitServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TermLimit_ServiceDesc is the grpc.ServiceDesc for TermLimit service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TermLimit_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "term_limit.TermLimit",
	HandlerType: (*TermLimitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TermLimit_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "term_limit.proto",
}
