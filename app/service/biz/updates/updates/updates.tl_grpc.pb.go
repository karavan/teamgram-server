//
// WARNING! All changes made in this file will be lost!
// Created from 'scheme.tl' by 'mtprotoc'
//
// Copyright (c) 2024-present,  Teamgram Authors.
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: updates.tl.proto

package updates

import (
	context "context"
	mtproto "github.com/teamgram/proto/mtproto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RPCUpdates_UpdatesGetStateV2_FullMethodName             = "/updates.RPCUpdates/updates_getStateV2"
	RPCUpdates_UpdatesGetDifferenceV2_FullMethodName        = "/updates.RPCUpdates/updates_getDifferenceV2"
	RPCUpdates_UpdatesGetChannelDifferenceV2_FullMethodName = "/updates.RPCUpdates/updates_getChannelDifferenceV2"
)

// RPCUpdatesClient is the client API for RPCUpdates service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCUpdatesClient interface {
	UpdatesGetStateV2(ctx context.Context, in *TLUpdatesGetStateV2, opts ...grpc.CallOption) (*mtproto.Updates_State, error)
	UpdatesGetDifferenceV2(ctx context.Context, in *TLUpdatesGetDifferenceV2, opts ...grpc.CallOption) (*Difference, error)
	UpdatesGetChannelDifferenceV2(ctx context.Context, in *TLUpdatesGetChannelDifferenceV2, opts ...grpc.CallOption) (*ChannelDifference, error)
}

type rPCUpdatesClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCUpdatesClient(cc grpc.ClientConnInterface) RPCUpdatesClient {
	return &rPCUpdatesClient{cc}
}

func (c *rPCUpdatesClient) UpdatesGetStateV2(ctx context.Context, in *TLUpdatesGetStateV2, opts ...grpc.CallOption) (*mtproto.Updates_State, error) {
	out := new(mtproto.Updates_State)
	err := c.cc.Invoke(ctx, RPCUpdates_UpdatesGetStateV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCUpdatesClient) UpdatesGetDifferenceV2(ctx context.Context, in *TLUpdatesGetDifferenceV2, opts ...grpc.CallOption) (*Difference, error) {
	out := new(Difference)
	err := c.cc.Invoke(ctx, RPCUpdates_UpdatesGetDifferenceV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCUpdatesClient) UpdatesGetChannelDifferenceV2(ctx context.Context, in *TLUpdatesGetChannelDifferenceV2, opts ...grpc.CallOption) (*ChannelDifference, error) {
	out := new(ChannelDifference)
	err := c.cc.Invoke(ctx, RPCUpdates_UpdatesGetChannelDifferenceV2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCUpdatesServer is the server API for RPCUpdates service.
// All implementations should embed UnimplementedRPCUpdatesServer
// for forward compatibility
type RPCUpdatesServer interface {
	UpdatesGetStateV2(context.Context, *TLUpdatesGetStateV2) (*mtproto.Updates_State, error)
	UpdatesGetDifferenceV2(context.Context, *TLUpdatesGetDifferenceV2) (*Difference, error)
	UpdatesGetChannelDifferenceV2(context.Context, *TLUpdatesGetChannelDifferenceV2) (*ChannelDifference, error)
}

// UnimplementedRPCUpdatesServer should be embedded to have forward compatible implementations.
type UnimplementedRPCUpdatesServer struct {
}

func (UnimplementedRPCUpdatesServer) UpdatesGetStateV2(context.Context, *TLUpdatesGetStateV2) (*mtproto.Updates_State, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatesGetStateV2 not implemented")
}
func (UnimplementedRPCUpdatesServer) UpdatesGetDifferenceV2(context.Context, *TLUpdatesGetDifferenceV2) (*Difference, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatesGetDifferenceV2 not implemented")
}
func (UnimplementedRPCUpdatesServer) UpdatesGetChannelDifferenceV2(context.Context, *TLUpdatesGetChannelDifferenceV2) (*ChannelDifference, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatesGetChannelDifferenceV2 not implemented")
}

// UnsafeRPCUpdatesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCUpdatesServer will
// result in compilation errors.
type UnsafeRPCUpdatesServer interface {
	mustEmbedUnimplementedRPCUpdatesServer()
}

func RegisterRPCUpdatesServer(s grpc.ServiceRegistrar, srv RPCUpdatesServer) {
	s.RegisterService(&RPCUpdates_ServiceDesc, srv)
}

func _RPCUpdates_UpdatesGetStateV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUpdatesGetStateV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUpdatesServer).UpdatesGetStateV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUpdates_UpdatesGetStateV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUpdatesServer).UpdatesGetStateV2(ctx, req.(*TLUpdatesGetStateV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCUpdates_UpdatesGetDifferenceV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUpdatesGetDifferenceV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUpdatesServer).UpdatesGetDifferenceV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUpdates_UpdatesGetDifferenceV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUpdatesServer).UpdatesGetDifferenceV2(ctx, req.(*TLUpdatesGetDifferenceV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCUpdates_UpdatesGetChannelDifferenceV2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TLUpdatesGetChannelDifferenceV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCUpdatesServer).UpdatesGetChannelDifferenceV2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RPCUpdates_UpdatesGetChannelDifferenceV2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCUpdatesServer).UpdatesGetChannelDifferenceV2(ctx, req.(*TLUpdatesGetChannelDifferenceV2))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCUpdates_ServiceDesc is the grpc.ServiceDesc for RPCUpdates service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCUpdates_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "updates.RPCUpdates",
	HandlerType: (*RPCUpdatesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "updates_getStateV2",
			Handler:    _RPCUpdates_UpdatesGetStateV2_Handler,
		},
		{
			MethodName: "updates_getDifferenceV2",
			Handler:    _RPCUpdates_UpdatesGetDifferenceV2_Handler,
		},
		{
			MethodName: "updates_getChannelDifferenceV2",
			Handler:    _RPCUpdates_UpdatesGetChannelDifferenceV2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "updates.tl.proto",
}