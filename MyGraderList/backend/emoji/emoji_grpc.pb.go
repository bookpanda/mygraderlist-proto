// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: emoji.proto

package emoji

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

const (
	EmojiService_FindAll_FullMethodName      = "/emoji.EmojiService/FindAll"
	EmojiService_FindByUserId_FullMethodName = "/emoji.EmojiService/FindByUserId"
	EmojiService_Create_FullMethodName       = "/emoji.EmojiService/Create"
	EmojiService_Delete_FullMethodName       = "/emoji.EmojiService/Delete"
)

// EmojiServiceClient is the client API for EmojiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmojiServiceClient interface {
	FindAll(ctx context.Context, in *FindByUserIdEmojiRequest, opts ...grpc.CallOption) (*FindByUserIdEmojiResponse, error)
	FindByUserId(ctx context.Context, in *FindByUserIdEmojiRequest, opts ...grpc.CallOption) (*FindByUserIdEmojiResponse, error)
	Create(ctx context.Context, in *CreateEmojiRequest, opts ...grpc.CallOption) (*CreateEmojiResponse, error)
	Delete(ctx context.Context, in *DeleteEmojiRequest, opts ...grpc.CallOption) (*DeleteEmojiResponse, error)
}

type emojiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmojiServiceClient(cc grpc.ClientConnInterface) EmojiServiceClient {
	return &emojiServiceClient{cc}
}

func (c *emojiServiceClient) FindAll(ctx context.Context, in *FindByUserIdEmojiRequest, opts ...grpc.CallOption) (*FindByUserIdEmojiResponse, error) {
	out := new(FindByUserIdEmojiResponse)
	err := c.cc.Invoke(ctx, EmojiService_FindAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emojiServiceClient) FindByUserId(ctx context.Context, in *FindByUserIdEmojiRequest, opts ...grpc.CallOption) (*FindByUserIdEmojiResponse, error) {
	out := new(FindByUserIdEmojiResponse)
	err := c.cc.Invoke(ctx, EmojiService_FindByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emojiServiceClient) Create(ctx context.Context, in *CreateEmojiRequest, opts ...grpc.CallOption) (*CreateEmojiResponse, error) {
	out := new(CreateEmojiResponse)
	err := c.cc.Invoke(ctx, EmojiService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emojiServiceClient) Delete(ctx context.Context, in *DeleteEmojiRequest, opts ...grpc.CallOption) (*DeleteEmojiResponse, error) {
	out := new(DeleteEmojiResponse)
	err := c.cc.Invoke(ctx, EmojiService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmojiServiceServer is the server API for EmojiService service.
// All implementations should embed UnimplementedEmojiServiceServer
// for forward compatibility
type EmojiServiceServer interface {
	FindAll(context.Context, *FindByUserIdEmojiRequest) (*FindByUserIdEmojiResponse, error)
	FindByUserId(context.Context, *FindByUserIdEmojiRequest) (*FindByUserIdEmojiResponse, error)
	Create(context.Context, *CreateEmojiRequest) (*CreateEmojiResponse, error)
	Delete(context.Context, *DeleteEmojiRequest) (*DeleteEmojiResponse, error)
}

// UnimplementedEmojiServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEmojiServiceServer struct {
}

func (UnimplementedEmojiServiceServer) FindAll(context.Context, *FindByUserIdEmojiRequest) (*FindByUserIdEmojiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedEmojiServiceServer) FindByUserId(context.Context, *FindByUserIdEmojiRequest) (*FindByUserIdEmojiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByUserId not implemented")
}
func (UnimplementedEmojiServiceServer) Create(context.Context, *CreateEmojiRequest) (*CreateEmojiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedEmojiServiceServer) Delete(context.Context, *DeleteEmojiRequest) (*DeleteEmojiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeEmojiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmojiServiceServer will
// result in compilation errors.
type UnsafeEmojiServiceServer interface {
	mustEmbedUnimplementedEmojiServiceServer()
}

func RegisterEmojiServiceServer(s grpc.ServiceRegistrar, srv EmojiServiceServer) {
	s.RegisterService(&EmojiService_ServiceDesc, srv)
}

func _EmojiService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByUserIdEmojiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmojiServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmojiService_FindAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmojiServiceServer).FindAll(ctx, req.(*FindByUserIdEmojiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmojiService_FindByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByUserIdEmojiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmojiServiceServer).FindByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmojiService_FindByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmojiServiceServer).FindByUserId(ctx, req.(*FindByUserIdEmojiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmojiService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmojiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmojiServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmojiService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmojiServiceServer).Create(ctx, req.(*CreateEmojiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmojiService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEmojiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmojiServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmojiService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmojiServiceServer).Delete(ctx, req.(*DeleteEmojiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmojiService_ServiceDesc is the grpc.ServiceDesc for EmojiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmojiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emoji.EmojiService",
	HandlerType: (*EmojiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAll",
			Handler:    _EmojiService_FindAll_Handler,
		},
		{
			MethodName: "FindByUserId",
			Handler:    _EmojiService_FindByUserId_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _EmojiService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EmojiService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "emoji.proto",
}
