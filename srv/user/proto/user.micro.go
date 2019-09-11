// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/user/proto/user.proto

package user

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

// Client API for User service

type UserService interface {
	Create(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Create(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*UserResponse, error) {
	req := c.c.NewRequest(c.name, "User.Create", in)
	out := new(UserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Logout(ctx context.Context, in *LogoutRequest, opts ...client.CallOption) (*LogoutResponse, error) {
	req := c.c.NewRequest(c.name, "User.Logout", in)
	out := new(LogoutResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "User.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "User.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "User.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Create(context.Context, *UserRequest, *UserResponse) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
	Logout(context.Context, *LogoutRequest, *LogoutResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Create(ctx context.Context, in *UserRequest, out *UserResponse) error
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
		Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Create(ctx context.Context, in *UserRequest, out *UserResponse) error {
	return h.UserHandler.Create(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) Logout(ctx context.Context, in *LogoutRequest, out *LogoutResponse) error {
	return h.UserHandler.Logout(ctx, in, out)
}

func (h *userHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.UserHandler.Get(ctx, in, out)
}

func (h *userHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.UserHandler.Update(ctx, in, out)
}

func (h *userHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.UserHandler.Delete(ctx, in, out)
}
