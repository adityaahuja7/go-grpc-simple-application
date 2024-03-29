// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: dscd_a1.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MarketClient is the client API for Market service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketClient interface {
	RegisterSeller(ctx context.Context, in *RegisterSellerRequest, opts ...grpc.CallOption) (*RegisterSellerResponse, error)
	SellItem(ctx context.Context, in *SellItemRequest, opts ...grpc.CallOption) (*SellItemResponse, error)
	UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error)
	DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error)
	DisplayItems(ctx context.Context, in *DisplayItemsRequest, opts ...grpc.CallOption) (*DisplayItemsResponse, error)
	SearchItems(ctx context.Context, in *SearchItemRequest, opts ...grpc.CallOption) (*SearchItemResponse, error)
	BuyItem(ctx context.Context, in *BuyItemRequest, opts ...grpc.CallOption) (*BuyItemResponse, error)
	AddToWishlist(ctx context.Context, in *AddToWishlistRequest, opts ...grpc.CallOption) (*AddToWishlistResponse, error)
	RateItem(ctx context.Context, in *RateItemRequest, opts ...grpc.CallOption) (*RateItemResponse, error)
}

type marketClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketClient(cc grpc.ClientConnInterface) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) RegisterSeller(ctx context.Context, in *RegisterSellerRequest, opts ...grpc.CallOption) (*RegisterSellerResponse, error) {
	out := new(RegisterSellerResponse)
	err := c.cc.Invoke(ctx, "/market/registerSeller", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) SellItem(ctx context.Context, in *SellItemRequest, opts ...grpc.CallOption) (*SellItemResponse, error) {
	out := new(SellItemResponse)
	err := c.cc.Invoke(ctx, "/market/sellItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*UpdateItemResponse, error) {
	out := new(UpdateItemResponse)
	err := c.cc.Invoke(ctx, "/market/updateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error) {
	out := new(DeleteItemResponse)
	err := c.cc.Invoke(ctx, "/market/deleteItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) DisplayItems(ctx context.Context, in *DisplayItemsRequest, opts ...grpc.CallOption) (*DisplayItemsResponse, error) {
	out := new(DisplayItemsResponse)
	err := c.cc.Invoke(ctx, "/market/displayItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) SearchItems(ctx context.Context, in *SearchItemRequest, opts ...grpc.CallOption) (*SearchItemResponse, error) {
	out := new(SearchItemResponse)
	err := c.cc.Invoke(ctx, "/market/searchItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) BuyItem(ctx context.Context, in *BuyItemRequest, opts ...grpc.CallOption) (*BuyItemResponse, error) {
	out := new(BuyItemResponse)
	err := c.cc.Invoke(ctx, "/market/buyItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) AddToWishlist(ctx context.Context, in *AddToWishlistRequest, opts ...grpc.CallOption) (*AddToWishlistResponse, error) {
	out := new(AddToWishlistResponse)
	err := c.cc.Invoke(ctx, "/market/addToWishlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) RateItem(ctx context.Context, in *RateItemRequest, opts ...grpc.CallOption) (*RateItemResponse, error) {
	out := new(RateItemResponse)
	err := c.cc.Invoke(ctx, "/market/rateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketServer is the server API for Market service.
// All implementations must embed UnimplementedMarketServer
// for forward compatibility
type MarketServer interface {
	RegisterSeller(context.Context, *RegisterSellerRequest) (*RegisterSellerResponse, error)
	SellItem(context.Context, *SellItemRequest) (*SellItemResponse, error)
	UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error)
	DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error)
	DisplayItems(context.Context, *DisplayItemsRequest) (*DisplayItemsResponse, error)
	SearchItems(context.Context, *SearchItemRequest) (*SearchItemResponse, error)
	BuyItem(context.Context, *BuyItemRequest) (*BuyItemResponse, error)
	AddToWishlist(context.Context, *AddToWishlistRequest) (*AddToWishlistResponse, error)
	RateItem(context.Context, *RateItemRequest) (*RateItemResponse, error)
	mustEmbedUnimplementedMarketServer()
}

// UnimplementedMarketServer must be embedded to have forward compatible implementations.
type UnimplementedMarketServer struct {
}

func (UnimplementedMarketServer) RegisterSeller(context.Context, *RegisterSellerRequest) (*RegisterSellerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterSeller not implemented")
}
func (UnimplementedMarketServer) SellItem(context.Context, *SellItemRequest) (*SellItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellItem not implemented")
}
func (UnimplementedMarketServer) UpdateItem(context.Context, *UpdateItemRequest) (*UpdateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}
func (UnimplementedMarketServer) DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
func (UnimplementedMarketServer) DisplayItems(context.Context, *DisplayItemsRequest) (*DisplayItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayItems not implemented")
}
func (UnimplementedMarketServer) SearchItems(context.Context, *SearchItemRequest) (*SearchItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchItems not implemented")
}
func (UnimplementedMarketServer) BuyItem(context.Context, *BuyItemRequest) (*BuyItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyItem not implemented")
}
func (UnimplementedMarketServer) AddToWishlist(context.Context, *AddToWishlistRequest) (*AddToWishlistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToWishlist not implemented")
}
func (UnimplementedMarketServer) RateItem(context.Context, *RateItemRequest) (*RateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateItem not implemented")
}
func (UnimplementedMarketServer) mustEmbedUnimplementedMarketServer() {}

// UnsafeMarketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketServer will
// result in compilation errors.
type UnsafeMarketServer interface {
	mustEmbedUnimplementedMarketServer()
}

func RegisterMarketServer(s grpc.ServiceRegistrar, srv MarketServer) {
	s.RegisterService(&Market_ServiceDesc, srv)
}

func _Market_RegisterSeller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterSellerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).RegisterSeller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market/registerSeller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).RegisterSeller(ctx, req.(*RegisterSellerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_SellItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SellItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).SellItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market/sellItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).SellItem(ctx, req.(*SellItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market/updateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).UpdateItem(ctx, req.(*UpdateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market/deleteItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).DeleteItem(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_DisplayItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).DisplayItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market/displayItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).DisplayItems(ctx, req.(*DisplayItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_SearchItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).SearchItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market/searchItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).SearchItems(ctx, req.(*SearchItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_BuyItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).BuyItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market/buyItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).BuyItem(ctx, req.(*BuyItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_AddToWishlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToWishlistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).AddToWishlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market/addToWishlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).AddToWishlist(ctx, req.(*AddToWishlistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_RateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).RateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market/rateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).RateItem(ctx, req.(*RateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Market_ServiceDesc is the grpc.ServiceDesc for Market service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Market_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registerSeller",
			Handler:    _Market_RegisterSeller_Handler,
		},
		{
			MethodName: "sellItem",
			Handler:    _Market_SellItem_Handler,
		},
		{
			MethodName: "updateItem",
			Handler:    _Market_UpdateItem_Handler,
		},
		{
			MethodName: "deleteItem",
			Handler:    _Market_DeleteItem_Handler,
		},
		{
			MethodName: "displayItems",
			Handler:    _Market_DisplayItems_Handler,
		},
		{
			MethodName: "searchItems",
			Handler:    _Market_SearchItems_Handler,
		},
		{
			MethodName: "buyItem",
			Handler:    _Market_BuyItem_Handler,
		},
		{
			MethodName: "addToWishlist",
			Handler:    _Market_AddToWishlist_Handler,
		},
		{
			MethodName: "rateItem",
			Handler:    _Market_RateItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dscd_a1.proto",
}

// MarketSellerClient is the client API for MarketSeller service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketSellerClient interface {
	NotifySeller(ctx context.Context, in *NotifySellerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type marketSellerClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketSellerClient(cc grpc.ClientConnInterface) MarketSellerClient {
	return &marketSellerClient{cc}
}

func (c *marketSellerClient) NotifySeller(ctx context.Context, in *NotifySellerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/market_seller/NotifySeller", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketSellerServer is the server API for MarketSeller service.
// All implementations must embed UnimplementedMarketSellerServer
// for forward compatibility
type MarketSellerServer interface {
	NotifySeller(context.Context, *NotifySellerRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMarketSellerServer()
}

// UnimplementedMarketSellerServer must be embedded to have forward compatible implementations.
type UnimplementedMarketSellerServer struct {
}

func (UnimplementedMarketSellerServer) NotifySeller(context.Context, *NotifySellerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifySeller not implemented")
}
func (UnimplementedMarketSellerServer) mustEmbedUnimplementedMarketSellerServer() {}

// UnsafeMarketSellerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketSellerServer will
// result in compilation errors.
type UnsafeMarketSellerServer interface {
	mustEmbedUnimplementedMarketSellerServer()
}

func RegisterMarketSellerServer(s grpc.ServiceRegistrar, srv MarketSellerServer) {
	s.RegisterService(&MarketSeller_ServiceDesc, srv)
}

func _MarketSeller_NotifySeller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifySellerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketSellerServer).NotifySeller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market_seller/NotifySeller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketSellerServer).NotifySeller(ctx, req.(*NotifySellerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MarketSeller_ServiceDesc is the grpc.ServiceDesc for MarketSeller service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketSeller_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "market_seller",
	HandlerType: (*MarketSellerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifySeller",
			Handler:    _MarketSeller_NotifySeller_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dscd_a1.proto",
}

// MarketBuyerClient is the client API for MarketBuyer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketBuyerClient interface {
	NotifyBuyer(ctx context.Context, in *NotifyBuyerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type marketBuyerClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketBuyerClient(cc grpc.ClientConnInterface) MarketBuyerClient {
	return &marketBuyerClient{cc}
}

func (c *marketBuyerClient) NotifyBuyer(ctx context.Context, in *NotifyBuyerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/market_buyer/NotifyBuyer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketBuyerServer is the server API for MarketBuyer service.
// All implementations must embed UnimplementedMarketBuyerServer
// for forward compatibility
type MarketBuyerServer interface {
	NotifyBuyer(context.Context, *NotifyBuyerRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedMarketBuyerServer()
}

// UnimplementedMarketBuyerServer must be embedded to have forward compatible implementations.
type UnimplementedMarketBuyerServer struct {
}

func (UnimplementedMarketBuyerServer) NotifyBuyer(context.Context, *NotifyBuyerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyBuyer not implemented")
}
func (UnimplementedMarketBuyerServer) mustEmbedUnimplementedMarketBuyerServer() {}

// UnsafeMarketBuyerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketBuyerServer will
// result in compilation errors.
type UnsafeMarketBuyerServer interface {
	mustEmbedUnimplementedMarketBuyerServer()
}

func RegisterMarketBuyerServer(s grpc.ServiceRegistrar, srv MarketBuyerServer) {
	s.RegisterService(&MarketBuyer_ServiceDesc, srv)
}

func _MarketBuyer_NotifyBuyer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyBuyerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketBuyerServer).NotifyBuyer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market_buyer/NotifyBuyer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketBuyerServer).NotifyBuyer(ctx, req.(*NotifyBuyerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MarketBuyer_ServiceDesc is the grpc.ServiceDesc for MarketBuyer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MarketBuyer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "market_buyer",
	HandlerType: (*MarketBuyerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyBuyer",
			Handler:    _MarketBuyer_NotifyBuyer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dscd_a1.proto",
}
