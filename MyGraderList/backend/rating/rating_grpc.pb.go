// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: MyGraderList/backend/rating/rating.proto

package rating

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
	RatingService_FindAll_FullMethodName      = "/rating.RatingService/FindAll"
	RatingService_FindByUserId_FullMethodName = "/rating.RatingService/FindByUserId"
	RatingService_Create_FullMethodName       = "/rating.RatingService/Create"
	RatingService_Update_FullMethodName       = "/rating.RatingService/Update"
	RatingService_Delete_FullMethodName       = "/rating.RatingService/Delete"
)

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	FindAll(ctx context.Context, in *FindAllRatingRequest, opts ...grpc.CallOption) (*FindAllRatingResponse, error)
	FindByUserId(ctx context.Context, in *FindByUserIdRatingRequest, opts ...grpc.CallOption) (*FindByUserIdRatingResponse, error)
	Create(ctx context.Context, in *CreateRatingRequest, opts ...grpc.CallOption) (*CreateRatingResponse, error)
	Update(ctx context.Context, in *UpdateRatingRequest, opts ...grpc.CallOption) (*UpdateRatingResponse, error)
	Delete(ctx context.Context, in *DeleteRatingRequest, opts ...grpc.CallOption) (*DeleteRatingResponse, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) FindAll(ctx context.Context, in *FindAllRatingRequest, opts ...grpc.CallOption) (*FindAllRatingResponse, error) {
	out := new(FindAllRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_FindAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) FindByUserId(ctx context.Context, in *FindByUserIdRatingRequest, opts ...grpc.CallOption) (*FindByUserIdRatingResponse, error) {
	out := new(FindByUserIdRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_FindByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) Create(ctx context.Context, in *CreateRatingRequest, opts ...grpc.CallOption) (*CreateRatingResponse, error) {
	out := new(CreateRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) Update(ctx context.Context, in *UpdateRatingRequest, opts ...grpc.CallOption) (*UpdateRatingResponse, error) {
	out := new(UpdateRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) Delete(ctx context.Context, in *DeleteRatingRequest, opts ...grpc.CallOption) (*DeleteRatingResponse, error) {
	out := new(DeleteRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	FindAll(context.Context, *FindAllRatingRequest) (*FindAllRatingResponse, error)
	FindByUserId(context.Context, *FindByUserIdRatingRequest) (*FindByUserIdRatingResponse, error)
	Create(context.Context, *CreateRatingRequest) (*CreateRatingResponse, error)
	Update(context.Context, *UpdateRatingRequest) (*UpdateRatingResponse, error)
	Delete(context.Context, *DeleteRatingRequest) (*DeleteRatingResponse, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) FindAll(context.Context, *FindAllRatingRequest) (*FindAllRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAll not implemented")
}
func (UnimplementedRatingServiceServer) FindByUserId(context.Context, *FindByUserIdRatingRequest) (*FindByUserIdRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByUserId not implemented")
}
func (UnimplementedRatingServiceServer) Create(context.Context, *CreateRatingRequest) (*CreateRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRatingServiceServer) Update(context.Context, *UpdateRatingRequest) (*UpdateRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRatingServiceServer) Delete(context.Context, *DeleteRatingRequest) (*DeleteRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_FindAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).FindAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_FindAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).FindAll(ctx, req.(*FindAllRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_FindByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByUserIdRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).FindByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_FindByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).FindByUserId(ctx, req.(*FindByUserIdRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).Create(ctx, req.(*CreateRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).Update(ctx, req.(*UpdateRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).Delete(ctx, req.(*DeleteRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rating.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindAll",
			Handler:    _RatingService_FindAll_Handler,
		},
		{
			MethodName: "FindByUserId",
			Handler:    _RatingService_FindByUserId_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RatingService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RatingService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RatingService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "MyGraderList/backend/rating/rating.proto",
}
