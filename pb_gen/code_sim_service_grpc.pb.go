// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb_gen

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

// CodeSimClient is the client API for CodeSim service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodeSimClient interface {
	HelloWorld(ctx context.Context, in *CodeSimHelloWorldRequest, opts ...grpc.CallOption) (*CodeSimHelloWorldResponse, error)
	Search(ctx context.Context, in *CodeSimSearchRequest, opts ...grpc.CallOption) (*CodeSimSearchResponse, error)
}

type codeSimClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeSimClient(cc grpc.ClientConnInterface) CodeSimClient {
	return &codeSimClient{cc}
}

func (c *codeSimClient) HelloWorld(ctx context.Context, in *CodeSimHelloWorldRequest, opts ...grpc.CallOption) (*CodeSimHelloWorldResponse, error) {
	out := new(CodeSimHelloWorldResponse)
	err := c.cc.Invoke(ctx, "/pb.CodeSim/HelloWorld", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeSimClient) Search(ctx context.Context, in *CodeSimSearchRequest, opts ...grpc.CallOption) (*CodeSimSearchResponse, error) {
	out := new(CodeSimSearchResponse)
	err := c.cc.Invoke(ctx, "/pb.CodeSim/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeSimServer is the server API for CodeSim service.
// All implementations must embed UnimplementedCodeSimServer
// for forward compatibility
type CodeSimServer interface {
	HelloWorld(context.Context, *CodeSimHelloWorldRequest) (*CodeSimHelloWorldResponse, error)
	Search(context.Context, *CodeSimSearchRequest) (*CodeSimSearchResponse, error)
	mustEmbedUnimplementedCodeSimServer()
}

// UnimplementedCodeSimServer must be embedded to have forward compatible implementations.
type UnimplementedCodeSimServer struct {
}

func (UnimplementedCodeSimServer) HelloWorld(context.Context, *CodeSimHelloWorldRequest) (*CodeSimHelloWorldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloWorld not implemented")
}
func (UnimplementedCodeSimServer) Search(context.Context, *CodeSimSearchRequest) (*CodeSimSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedCodeSimServer) mustEmbedUnimplementedCodeSimServer() {}

// UnsafeCodeSimServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodeSimServer will
// result in compilation errors.
type UnsafeCodeSimServer interface {
	mustEmbedUnimplementedCodeSimServer()
}

func RegisterCodeSimServer(s grpc.ServiceRegistrar, srv CodeSimServer) {
	s.RegisterService(&CodeSim_ServiceDesc, srv)
}

func _CodeSim_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeSimHelloWorldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeSimServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CodeSim/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeSimServer).HelloWorld(ctx, req.(*CodeSimHelloWorldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeSim_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeSimSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeSimServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CodeSim/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeSimServer).Search(ctx, req.(*CodeSimSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CodeSim_ServiceDesc is the grpc.ServiceDesc for CodeSim service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeSim_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CodeSim",
	HandlerType: (*CodeSimServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloWorld",
			Handler:    _CodeSim_HelloWorld_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _CodeSim_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "code_sim/code_sim_service.proto",
}
