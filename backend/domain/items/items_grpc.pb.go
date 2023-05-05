// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: items.proto

package items

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
	ItemsService_Create_FullMethodName = "/items.ItemsService/Create"
	ItemsService_Get_FullMethodName    = "/items.ItemsService/Get"
	ItemsService_Update_FullMethodName = "/items.ItemsService/Update"
	ItemsService_Delete_FullMethodName = "/items.ItemsService/Delete"
	ItemsService_GetAll_FullMethodName = "/items.ItemsService/GetAll"
)

// ItemsServiceClient is the client API for ItemsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemsServiceClient interface {
	Create(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error)
	Get(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error)
	Update(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error)
	Delete(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error)
	GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ItemsList, error)
}

type itemsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewItemsServiceClient(cc grpc.ClientConnInterface) ItemsServiceClient {
	return &itemsServiceClient{cc}
}

func (c *itemsServiceClient) Create(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, ItemsService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsServiceClient) Get(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, ItemsService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsServiceClient) Update(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, ItemsService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsServiceClient) Delete(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Item, error) {
	out := new(Item)
	err := c.cc.Invoke(ctx, ItemsService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemsServiceClient) GetAll(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ItemsList, error) {
	out := new(ItemsList)
	err := c.cc.Invoke(ctx, ItemsService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemsServiceServer is the server API for ItemsService service.
// All implementations must embed UnimplementedItemsServiceServer
// for forward compatibility
type ItemsServiceServer interface {
	Create(context.Context, *Item) (*Item, error)
	Get(context.Context, *Item) (*Item, error)
	Update(context.Context, *Item) (*Item, error)
	Delete(context.Context, *Item) (*Item, error)
	GetAll(context.Context, *Empty) (*ItemsList, error)
	mustEmbedUnimplementedItemsServiceServer()
}

// UnimplementedItemsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedItemsServiceServer struct {
}

func (UnimplementedItemsServiceServer) Create(context.Context, *Item) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedItemsServiceServer) Get(context.Context, *Item) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedItemsServiceServer) Update(context.Context, *Item) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedItemsServiceServer) Delete(context.Context, *Item) (*Item, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedItemsServiceServer) GetAll(context.Context, *Empty) (*ItemsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedItemsServiceServer) mustEmbedUnimplementedItemsServiceServer() {}

// UnsafeItemsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemsServiceServer will
// result in compilation errors.
type UnsafeItemsServiceServer interface {
	mustEmbedUnimplementedItemsServiceServer()
}

func RegisterItemsServiceServer(s grpc.ServiceRegistrar, srv ItemsServiceServer) {
	s.RegisterService(&ItemsService_ServiceDesc, srv)
}

func _ItemsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemsService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServiceServer).Create(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemsService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServiceServer).Get(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemsService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServiceServer).Update(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemsService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServiceServer).Delete(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemsService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemsServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ItemsService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemsServiceServer).GetAll(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ItemsService_ServiceDesc is the grpc.ServiceDesc for ItemsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "items.ItemsService",
	HandlerType: (*ItemsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ItemsService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ItemsService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ItemsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ItemsService_Delete_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ItemsService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "items.proto",
}
