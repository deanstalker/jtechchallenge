// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/pet/proto/pet.proto

package pet

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

// Client API for PetService service

type PetService interface {
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	ByStatus(ctx context.Context, in *ByStatusRequest, opts ...client.CallOption) (*ByStatusResponse, error)
	ByID(ctx context.Context, in *ByIDRequest, opts ...client.CallOption) (*ByIDResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	UploadImage(ctx context.Context, in *UploadImageRequest, opts ...client.CallOption) (*UploadImageResponse, error)
}

type petService struct {
	c    client.Client
	name string
}

func NewPetService(name string, c client.Client) PetService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "pet"
	}
	return &petService{
		c:    c,
		name: name,
	}
}

func (c *petService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "PetService.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "PetService.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petService) ByStatus(ctx context.Context, in *ByStatusRequest, opts ...client.CallOption) (*ByStatusResponse, error) {
	req := c.c.NewRequest(c.name, "PetService.ByStatus", in)
	out := new(ByStatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petService) ByID(ctx context.Context, in *ByIDRequest, opts ...client.CallOption) (*ByIDResponse, error) {
	req := c.c.NewRequest(c.name, "PetService.ByID", in)
	out := new(ByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "PetService.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petService) UploadImage(ctx context.Context, in *UploadImageRequest, opts ...client.CallOption) (*UploadImageResponse, error) {
	req := c.c.NewRequest(c.name, "PetService.UploadImage", in)
	out := new(UploadImageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PetService service

type PetServiceHandler interface {
	Add(context.Context, *AddRequest, *AddResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	ByStatus(context.Context, *ByStatusRequest, *ByStatusResponse) error
	ByID(context.Context, *ByIDRequest, *ByIDResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	UploadImage(context.Context, *UploadImageRequest, *UploadImageResponse) error
}

func RegisterPetServiceHandler(s server.Server, hdlr PetServiceHandler, opts ...server.HandlerOption) error {
	type petService interface {
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		ByStatus(ctx context.Context, in *ByStatusRequest, out *ByStatusResponse) error
		ByID(ctx context.Context, in *ByIDRequest, out *ByIDResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		UploadImage(ctx context.Context, in *UploadImageRequest, out *UploadImageResponse) error
	}
	type PetService struct {
		petService
	}
	h := &petServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PetService{h}, opts...))
}

type petServiceHandler struct {
	PetServiceHandler
}

func (h *petServiceHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.PetServiceHandler.Add(ctx, in, out)
}

func (h *petServiceHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.PetServiceHandler.Update(ctx, in, out)
}

func (h *petServiceHandler) ByStatus(ctx context.Context, in *ByStatusRequest, out *ByStatusResponse) error {
	return h.PetServiceHandler.ByStatus(ctx, in, out)
}

func (h *petServiceHandler) ByID(ctx context.Context, in *ByIDRequest, out *ByIDResponse) error {
	return h.PetServiceHandler.ByID(ctx, in, out)
}

func (h *petServiceHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.PetServiceHandler.Delete(ctx, in, out)
}

func (h *petServiceHandler) UploadImage(ctx context.Context, in *UploadImageRequest, out *UploadImageResponse) error {
	return h.PetServiceHandler.UploadImage(ctx, in, out)
}
