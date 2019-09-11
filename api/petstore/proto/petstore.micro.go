// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/petstore/proto/petstore.proto

package petstore

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

// Client API for PetStore service

type PetStoreService interface {
}

type petStoreService struct {
	c    client.Client
	name string
}

func NewPetStoreService(name string, c client.Client) PetStoreService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "petstore"
	}
	return &petStoreService{
		c:    c,
		name: name,
	}
}

// Server API for PetStore service

type PetStoreHandler interface {
}

func RegisterPetStoreHandler(s server.Server, hdlr PetStoreHandler, opts ...server.HandlerOption) error {
	type petStore interface {
	}
	type PetStore struct {
		petStore
	}
	h := &petStoreHandler{hdlr}
	return s.Handle(s.NewHandler(&PetStore{h}, opts...))
}

type petStoreHandler struct {
	PetStoreHandler
}
