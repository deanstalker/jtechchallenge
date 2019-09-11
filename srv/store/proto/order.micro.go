// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/store/proto/order.proto

package store

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Order service

type OrderService interface {
	Inventory(ctx context.Context, in *InventoryRequest, opts ...client.CallOption) (*InventoryResponse, error)
	Place(ctx context.Context, in *PlaceRequest, opts ...client.CallOption) (*PlaceResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type orderService struct {
	c    client.Client
	name string
}

func NewOrderService(name string, c client.Client) OrderService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "store"
	}
	return &orderService{
		c:    c,
		name: name,
	}
}

func (c *orderService) Inventory(ctx context.Context, in *InventoryRequest, opts ...client.CallOption) (*InventoryResponse, error) {
	req := c.c.NewRequest(c.name, "Order.Inventory", in)
	out := new(InventoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Place(ctx context.Context, in *PlaceRequest, opts ...client.CallOption) (*PlaceResponse, error) {
	req := c.c.NewRequest(c.name, "Order.Place", in)
	out := new(PlaceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "Order.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Order.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Order service

type OrderHandler interface {
	Inventory(context.Context, *InventoryRequest, *InventoryResponse) error
	Place(context.Context, *PlaceRequest, *PlaceResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterOrderHandler(s server.Server, hdlr OrderHandler, opts ...server.HandlerOption) error {
	type order interface {
		Inventory(ctx context.Context, in *InventoryRequest, out *InventoryResponse) error
		Place(ctx context.Context, in *PlaceRequest, out *PlaceResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type Order struct {
		order
	}
	h := &orderHandler{hdlr}
	return s.Handle(s.NewHandler(&Order{h}, opts...))
}

type orderHandler struct {
	OrderHandler
}

func (h *orderHandler) Inventory(ctx context.Context, in *InventoryRequest, out *InventoryResponse) error {
	return h.OrderHandler.Inventory(ctx, in, out)
}

func (h *orderHandler) Place(ctx context.Context, in *PlaceRequest, out *PlaceResponse) error {
	return h.OrderHandler.Place(ctx, in, out)
}

func (h *orderHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.OrderHandler.Get(ctx, in, out)
}

func (h *orderHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.OrderHandler.Delete(ctx, in, out)
}