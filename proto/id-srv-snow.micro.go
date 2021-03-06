// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: id-srv-snow.proto

/*
Package platform_id_srv_snow is a generated protocol buffer package.

It is generated from these files:
	id-srv-snow.proto

It has these top-level messages:
	GetIdReq
	GetIdResp
*/
package platform_id_srv_snow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Snow service

type SnowService interface {
	GetId(ctx context.Context, in *GetIdReq, opts ...client.CallOption) (*GetIdResp, error)
}

type snowService struct {
	c    client.Client
	name string
}

func NewSnowService(name string, c client.Client) SnowService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "platform.id.srv.snow"
	}
	return &snowService{
		c:    c,
		name: name,
	}
}

func (c *snowService) GetId(ctx context.Context, in *GetIdReq, opts ...client.CallOption) (*GetIdResp, error) {
	req := c.c.NewRequest(c.name, "Snow.GetId", in)
	out := new(GetIdResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Snow service

type SnowHandler interface {
	GetId(context.Context, *GetIdReq, *GetIdResp) error
}

func RegisterSnowHandler(s server.Server, hdlr SnowHandler, opts ...server.HandlerOption) error {
	type snow interface {
		GetId(ctx context.Context, in *GetIdReq, out *GetIdResp) error
	}
	type Snow struct {
		snow
	}
	h := &snowHandler{hdlr}
	return s.Handle(s.NewHandler(&Snow{h}, opts...))
}

type snowHandler struct {
	SnowHandler
}

func (h *snowHandler) GetId(ctx context.Context, in *GetIdReq, out *GetIdResp) error {
	return h.SnowHandler.GetId(ctx, in, out)
}
