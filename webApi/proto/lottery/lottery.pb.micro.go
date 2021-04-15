// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/lottery/lottery.proto

package go_micro_service_lottery

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

// Api Endpoints for Lottery service

func NewLotteryEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Lottery service

type LotteryService interface {
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	GetList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResponse, error)
}

type lotteryService struct {
	c    client.Client
	name string
}

func NewLotteryService(name string, c client.Client) LotteryService {
	return &lotteryService{
		c:    c,
		name: name,
	}
}

func (c *lotteryService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "Lottery.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lotteryService) GetList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResponse, error) {
	req := c.c.NewRequest(c.name, "Lottery.GetList", in)
	out := new(GetListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Lottery service

type LotteryHandler interface {
	Add(context.Context, *AddRequest, *AddResponse) error
	GetList(context.Context, *GetListRequest, *GetListResponse) error
}

func RegisterLotteryHandler(s server.Server, hdlr LotteryHandler, opts ...server.HandlerOption) error {
	type lottery interface {
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
		GetList(ctx context.Context, in *GetListRequest, out *GetListResponse) error
	}
	type Lottery struct {
		lottery
	}
	h := &lotteryHandler{hdlr}
	return s.Handle(s.NewHandler(&Lottery{h}, opts...))
}

type lotteryHandler struct {
	LotteryHandler
}

func (h *lotteryHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.LotteryHandler.Add(ctx, in, out)
}

func (h *lotteryHandler) GetList(ctx context.Context, in *GetListRequest, out *GetListResponse) error {
	return h.LotteryHandler.GetList(ctx, in, out)
}