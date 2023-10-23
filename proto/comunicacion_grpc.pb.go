// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: proto/comunicacion.proto

package proto

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

// ComServClient is the client API for ComServ service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComServClient interface {
	ContOMS(ctx context.Context, in *EstadoPersona, opts ...grpc.CallOption) (*EstadoPersona, error)
	OMSDataNode(ctx context.Context, in *IdPersona, opts ...grpc.CallOption) (*IdPersona, error)
	ConsultaDN(ctx context.Context, in *IdPersona, opts ...grpc.CallOption) (*IdPersona, error)
	ConsultaOMS(ctx context.Context, in *EstadoPersona, opts ...grpc.CallOption) (*EstadoPersona, error)
}

type comServClient struct {
	cc grpc.ClientConnInterface
}

func NewComServClient(cc grpc.ClientConnInterface) ComServClient {
	return &comServClient{cc}
}

func (c *comServClient) ContOMS(ctx context.Context, in *EstadoPersona, opts ...grpc.CallOption) (*EstadoPersona, error) {
	out := new(EstadoPersona)
	err := c.cc.Invoke(ctx, "/grpc.ComServ/ContOMS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comServClient) OMSDataNode(ctx context.Context, in *IdPersona, opts ...grpc.CallOption) (*IdPersona, error) {
	out := new(IdPersona)
	err := c.cc.Invoke(ctx, "/grpc.ComServ/OMSDataNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comServClient) ConsultaDN(ctx context.Context, in *IdPersona, opts ...grpc.CallOption) (*IdPersona, error) {
	out := new(IdPersona)
	err := c.cc.Invoke(ctx, "/grpc.ComServ/ConsultaDN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *comServClient) ConsultaOMS(ctx context.Context, in *EstadoPersona, opts ...grpc.CallOption) (*EstadoPersona, error) {
	out := new(EstadoPersona)
	err := c.cc.Invoke(ctx, "/grpc.ComServ/ConsultaOMS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComServServer is the server API for ComServ service.
// All implementations must embed UnimplementedComServServer
// for forward compatibility
type ComServServer interface {
	ContOMS(context.Context, *EstadoPersona) (*EstadoPersona, error)
	OMSDataNode(context.Context, *IdPersona) (*IdPersona, error)
	ConsultaDN(context.Context, *IdPersona) (*IdPersona, error)
	ConsultaOMS(context.Context, *EstadoPersona) (*EstadoPersona, error)
	mustEmbedUnimplementedComServServer()
}

// UnimplementedComServServer must be embedded to have forward compatible implementations.
type UnimplementedComServServer struct {
}

func (UnimplementedComServServer) ContOMS(context.Context, *EstadoPersona) (*EstadoPersona, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContOMS not implemented")
}
func (UnimplementedComServServer) OMSDataNode(context.Context, *IdPersona) (*IdPersona, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OMSDataNode not implemented")
}
func (UnimplementedComServServer) ConsultaDN(context.Context, *IdPersona) (*IdPersona, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsultaDN not implemented")
}
func (UnimplementedComServServer) ConsultaOMS(context.Context, *EstadoPersona) (*EstadoPersona, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConsultaOMS not implemented")
}
func (UnimplementedComServServer) mustEmbedUnimplementedComServServer() {}

// UnsafeComServServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComServServer will
// result in compilation errors.
type UnsafeComServServer interface {
	mustEmbedUnimplementedComServServer()
}

func RegisterComServServer(s grpc.ServiceRegistrar, srv ComServServer) {
	s.RegisterService(&ComServ_ServiceDesc, srv)
}

func _ComServ_ContOMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstadoPersona)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComServServer).ContOMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ComServ/ContOMS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComServServer).ContOMS(ctx, req.(*EstadoPersona))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComServ_OMSDataNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdPersona)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComServServer).OMSDataNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ComServ/OMSDataNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComServServer).OMSDataNode(ctx, req.(*IdPersona))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComServ_ConsultaDN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdPersona)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComServServer).ConsultaDN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ComServ/ConsultaDN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComServServer).ConsultaDN(ctx, req.(*IdPersona))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComServ_ConsultaOMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstadoPersona)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComServServer).ConsultaOMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ComServ/ConsultaOMS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComServServer).ConsultaOMS(ctx, req.(*EstadoPersona))
	}
	return interceptor(ctx, in, info, handler)
}

// ComServ_ServiceDesc is the grpc.ServiceDesc for ComServ service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComServ_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ComServ",
	HandlerType: (*ComServServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ContOMS",
			Handler:    _ComServ_ContOMS_Handler,
		},
		{
			MethodName: "OMSDataNode",
			Handler:    _ComServ_OMSDataNode_Handler,
		},
		{
			MethodName: "ConsultaDN",
			Handler:    _ComServ_ConsultaDN_Handler,
		},
		{
			MethodName: "ConsultaOMS",
			Handler:    _ComServ_ConsultaOMS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/comunicacion.proto",
}
