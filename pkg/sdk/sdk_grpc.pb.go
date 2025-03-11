// Copyright 2024 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was autogenerated. Do not edit directly.
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.0
// source: sdk.proto

package sdk

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

// SDKClient is the client API for SDK service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SDKClient interface {
	// Call when the GameServer is ready
	Ready(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Call to self Allocation the GameServer
	Allocate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Call when the GameServer is shutting down
	Shutdown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Send a Empty every d Duration to declare that this GameSever is healthy
	Health(ctx context.Context, opts ...grpc.CallOption) (SDK_HealthClient, error)
	// Retrieve the current GameServer data
	GetGameServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GameServer, error)
	// Send GameServer details whenever the GameServer is updated
	WatchGameServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SDK_WatchGameServerClient, error)
	// Apply a Label to the backing GameServer metadata
	SetLabel(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Empty, error)
	// Apply a Annotation to the backing GameServer metadata
	SetAnnotation(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Empty, error)
	// Marks the GameServer as the Reserved state for Duration
	Reserve(ctx context.Context, in *Duration, opts ...grpc.CallOption) (*Empty, error)
}

type sDKClient struct {
	cc grpc.ClientConnInterface
}

func NewSDKClient(cc grpc.ClientConnInterface) SDKClient {
	return &sDKClient{cc}
}

func (c *sDKClient) Ready(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/Ready", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) Allocate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/Allocate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) Shutdown(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/Shutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) Health(ctx context.Context, opts ...grpc.CallOption) (SDK_HealthClient, error) {
	stream, err := c.cc.NewStream(ctx, &SDK_ServiceDesc.Streams[0], "/agones.dev.sdk.SDK/Health", opts...)
	if err != nil {
		return nil, err
	}
	x := &sDKHealthClient{stream}
	return x, nil
}

type SDK_HealthClient interface {
	Send(*Empty) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type sDKHealthClient struct {
	grpc.ClientStream
}

func (x *sDKHealthClient) Send(m *Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sDKHealthClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sDKClient) GetGameServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GameServer, error) {
	out := new(GameServer)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/GetGameServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) WatchGameServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SDK_WatchGameServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &SDK_ServiceDesc.Streams[1], "/agones.dev.sdk.SDK/WatchGameServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &sDKWatchGameServerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SDK_WatchGameServerClient interface {
	Recv() (*GameServer, error)
	grpc.ClientStream
}

type sDKWatchGameServerClient struct {
	grpc.ClientStream
}

func (x *sDKWatchGameServerClient) Recv() (*GameServer, error) {
	m := new(GameServer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sDKClient) SetLabel(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/SetLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) SetAnnotation(ctx context.Context, in *KeyValue, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/SetAnnotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sDKClient) Reserve(ctx context.Context, in *Duration, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/agones.dev.sdk.SDK/Reserve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SDKServer is the server API for SDK service.
// All implementations should embed UnimplementedSDKServer
// for forward compatibility
type SDKServer interface {
	// Call when the GameServer is ready
	Ready(context.Context, *Empty) (*Empty, error)
	// Call to self Allocation the GameServer
	Allocate(context.Context, *Empty) (*Empty, error)
	// Call when the GameServer is shutting down
	Shutdown(context.Context, *Empty) (*Empty, error)
	// Send a Empty every d Duration to declare that this GameSever is healthy
	Health(SDK_HealthServer) error
	// Retrieve the current GameServer data
	GetGameServer(context.Context, *Empty) (*GameServer, error)
	// Send GameServer details whenever the GameServer is updated
	WatchGameServer(*Empty, SDK_WatchGameServerServer) error
	// Apply a Label to the backing GameServer metadata
	SetLabel(context.Context, *KeyValue) (*Empty, error)
	// Apply a Annotation to the backing GameServer metadata
	SetAnnotation(context.Context, *KeyValue) (*Empty, error)
	// Marks the GameServer as the Reserved state for Duration
	Reserve(context.Context, *Duration) (*Empty, error)
}

// UnimplementedSDKServer should be embedded to have forward compatible implementations.
type UnimplementedSDKServer struct {
}

func (UnimplementedSDKServer) Ready(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ready not implemented")
}
func (UnimplementedSDKServer) Allocate(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allocate not implemented")
}
func (UnimplementedSDKServer) Shutdown(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedSDKServer) Health(SDK_HealthServer) error {
	return status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedSDKServer) GetGameServer(context.Context, *Empty) (*GameServer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameServer not implemented")
}
func (UnimplementedSDKServer) WatchGameServer(*Empty, SDK_WatchGameServerServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchGameServer not implemented")
}
func (UnimplementedSDKServer) SetLabel(context.Context, *KeyValue) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLabel not implemented")
}
func (UnimplementedSDKServer) SetAnnotation(context.Context, *KeyValue) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAnnotation not implemented")
}
func (UnimplementedSDKServer) Reserve(context.Context, *Duration) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reserve not implemented")
}

// UnsafeSDKServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SDKServer will
// result in compilation errors.
type UnsafeSDKServer interface {
	mustEmbedUnimplementedSDKServer()
}

func RegisterSDKServer(s grpc.ServiceRegistrar, srv SDKServer) {
	s.RegisterService(&SDK_ServiceDesc, srv)
}

func _SDK_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/Ready",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).Ready(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Allocate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).Allocate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/Allocate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).Allocate(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/Shutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).Shutdown(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Health_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SDKServer).Health(&sDKHealthServer{stream})
}

type SDK_HealthServer interface {
	SendAndClose(*Empty) error
	Recv() (*Empty, error)
	grpc.ServerStream
}

type sDKHealthServer struct {
	grpc.ServerStream
}

func (x *sDKHealthServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sDKHealthServer) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SDK_GetGameServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).GetGameServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/GetGameServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).GetGameServer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_WatchGameServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SDKServer).WatchGameServer(m, &sDKWatchGameServerServer{stream})
}

type SDK_WatchGameServerServer interface {
	Send(*GameServer) error
	grpc.ServerStream
}

type sDKWatchGameServerServer struct {
	grpc.ServerStream
}

func (x *sDKWatchGameServerServer) Send(m *GameServer) error {
	return x.ServerStream.SendMsg(m)
}

func _SDK_SetLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).SetLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/SetLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).SetLabel(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_SetAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).SetAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/SetAnnotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).SetAnnotation(ctx, req.(*KeyValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _SDK_Reserve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Duration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SDKServer).Reserve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agones.dev.sdk.SDK/Reserve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SDKServer).Reserve(ctx, req.(*Duration))
	}
	return interceptor(ctx, in, info, handler)
}

// SDK_ServiceDesc is the grpc.ServiceDesc for SDK service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SDK_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "agones.dev.sdk.SDK",
	HandlerType: (*SDKServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ready",
			Handler:    _SDK_Ready_Handler,
		},
		{
			MethodName: "Allocate",
			Handler:    _SDK_Allocate_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _SDK_Shutdown_Handler,
		},
		{
			MethodName: "GetGameServer",
			Handler:    _SDK_GetGameServer_Handler,
		},
		{
			MethodName: "SetLabel",
			Handler:    _SDK_SetLabel_Handler,
		},
		{
			MethodName: "SetAnnotation",
			Handler:    _SDK_SetAnnotation_Handler,
		},
		{
			MethodName: "Reserve",
			Handler:    _SDK_Reserve_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Health",
			Handler:       _SDK_Health_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "WatchGameServer",
			Handler:       _SDK_WatchGameServer_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sdk.proto",
}
