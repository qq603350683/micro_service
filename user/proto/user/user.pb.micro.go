// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_service_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	GetInfoByUniqueId(ctx context.Context, in *GetInfoByUniqueIdRequest, opts ...client.CallOption) (*GetInfoByUniqueIdResponse, error)
	GetInfoByUserId(ctx context.Context, in *GetInfoByUserIdRequest, opts ...client.CallOption) (*GetInfoByUserIdResponse, error)
	GetListByUserId(ctx context.Context, in *GetListByUserIdRequest, opts ...client.CallOption) (*GetListByUserIdResponse, error)
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) GetInfoByUniqueId(ctx context.Context, in *GetInfoByUniqueIdRequest, opts ...client.CallOption) (*GetInfoByUniqueIdResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetInfoByUniqueId", in)
	out := new(GetInfoByUniqueIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetInfoByUserId(ctx context.Context, in *GetInfoByUserIdRequest, opts ...client.CallOption) (*GetInfoByUserIdResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetInfoByUserId", in)
	out := new(GetInfoByUserIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetListByUserId(ctx context.Context, in *GetListByUserIdRequest, opts ...client.CallOption) (*GetListByUserIdResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetListByUserId", in)
	out := new(GetListByUserIdResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "User.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	GetInfoByUniqueId(context.Context, *GetInfoByUniqueIdRequest, *GetInfoByUniqueIdResponse) error
	GetInfoByUserId(context.Context, *GetInfoByUserIdRequest, *GetInfoByUserIdResponse) error
	GetListByUserId(context.Context, *GetListByUserIdRequest, *GetListByUserIdResponse) error
	Add(context.Context, *AddRequest, *AddResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		GetInfoByUniqueId(ctx context.Context, in *GetInfoByUniqueIdRequest, out *GetInfoByUniqueIdResponse) error
		GetInfoByUserId(ctx context.Context, in *GetInfoByUserIdRequest, out *GetInfoByUserIdResponse) error
		GetListByUserId(ctx context.Context, in *GetListByUserIdRequest, out *GetListByUserIdResponse) error
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
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

func (h *userHandler) GetInfoByUniqueId(ctx context.Context, in *GetInfoByUniqueIdRequest, out *GetInfoByUniqueIdResponse) error {
	return h.UserHandler.GetInfoByUniqueId(ctx, in, out)
}

func (h *userHandler) GetInfoByUserId(ctx context.Context, in *GetInfoByUserIdRequest, out *GetInfoByUserIdResponse) error {
	return h.UserHandler.GetInfoByUserId(ctx, in, out)
}

func (h *userHandler) GetListByUserId(ctx context.Context, in *GetListByUserIdRequest, out *GetListByUserIdResponse) error {
	return h.UserHandler.GetListByUserId(ctx, in, out)
}

func (h *userHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.UserHandler.Add(ctx, in, out)
}
