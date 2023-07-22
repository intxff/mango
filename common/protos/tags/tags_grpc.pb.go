// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: tags/tags.proto

package tags

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
	Tags_Create_FullMethodName = "/Tags/Create"
	Tags_Get_FullMethodName    = "/Tags/Get"
	Tags_Update_FullMethodName = "/Tags/Update"
	Tags_List_FullMethodName   = "/Tags/List"
)

// TagsClient is the client API for Tags service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TagsClient interface {
	// create a new tag
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error)
	// fuzzy search tag
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
	// update tag
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error)
	// list all tags
	List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error)
}

type tagsClient struct {
	cc grpc.ClientConnInterface
}

func NewTagsClient(cc grpc.ClientConnInterface) TagsClient {
	return &tagsClient{cc}
}

func (c *tagsClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateResp, error) {
	out := new(CreateResp)
	err := c.cc.Invoke(ctx, Tags_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, Tags_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error) {
	out := new(UpdateResp)
	err := c.cc.Invoke(ctx, Tags_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagsClient) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error) {
	out := new(ListResp)
	err := c.cc.Invoke(ctx, Tags_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagsServer is the server API for Tags service.
// All implementations must embed UnimplementedTagsServer
// for forward compatibility
type TagsServer interface {
	// create a new tag
	Create(context.Context, *CreateReq) (*CreateResp, error)
	// fuzzy search tag
	Get(context.Context, *GetReq) (*GetResp, error)
	// update tag
	Update(context.Context, *UpdateReq) (*UpdateResp, error)
	// list all tags
	List(context.Context, *ListReq) (*ListResp, error)
	mustEmbedUnimplementedTagsServer()
}

// UnimplementedTagsServer must be embedded to have forward compatible implementations.
type UnimplementedTagsServer struct {
}

func (UnimplementedTagsServer) Create(context.Context, *CreateReq) (*CreateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTagsServer) Get(context.Context, *GetReq) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTagsServer) Update(context.Context, *UpdateReq) (*UpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTagsServer) List(context.Context, *ListReq) (*ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTagsServer) mustEmbedUnimplementedTagsServer() {}

// UnsafeTagsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TagsServer will
// result in compilation errors.
type UnsafeTagsServer interface {
	mustEmbedUnimplementedTagsServer()
}

func RegisterTagsServer(s grpc.ServiceRegistrar, srv TagsServer) {
	s.RegisterService(&Tags_ServiceDesc, srv)
}

func _Tags_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tags_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tags_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tags_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tags_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tags_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tags_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Tags_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagsServer).List(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Tags_ServiceDesc is the grpc.ServiceDesc for Tags service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tags_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Tags",
	HandlerType: (*TagsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Tags_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Tags_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Tags_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Tags_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tags/tags.proto",
}