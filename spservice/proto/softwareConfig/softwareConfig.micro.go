// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: softwareConfig.proto

/*
Package softwareConfig is a generated protocol buffer package.

It is generated from these files:
	softwareConfig.proto

It has these top-level messages:
	SoftWareConfigListRequest
	SoftWareConfigNode
	SoftWareConfigListResponse
*/
package softwareConfig

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

// Client API for SoftWareConfig service

type SoftWareConfigService interface {
	List(ctx context.Context, in *SoftWareConfigListRequest, opts ...client.CallOption) (*SoftWareConfigListResponse, error)
}

type softWareConfigService struct {
	c    client.Client
	name string
}

func NewSoftWareConfigService(name string, c client.Client) SoftWareConfigService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "softwareconfig"
	}
	return &softWareConfigService{
		c:    c,
		name: name,
	}
}

func (c *softWareConfigService) List(ctx context.Context, in *SoftWareConfigListRequest, opts ...client.CallOption) (*SoftWareConfigListResponse, error) {
	req := c.c.NewRequest(c.name, "SoftWareConfig.List", in)
	out := new(SoftWareConfigListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SoftWareConfig service

type SoftWareConfigHandler interface {
	List(context.Context, *SoftWareConfigListRequest, *SoftWareConfigListResponse) error
}

func RegisterSoftWareConfigHandler(s server.Server, hdlr SoftWareConfigHandler, opts ...server.HandlerOption) {
	type softWareConfig interface {
		List(ctx context.Context, in *SoftWareConfigListRequest, out *SoftWareConfigListResponse) error
	}
	type SoftWareConfig struct {
		softWareConfig
	}
	h := &softWareConfigHandler{hdlr}
	s.Handle(s.NewHandler(&SoftWareConfig{h}, opts...))
}

type softWareConfigHandler struct {
	SoftWareConfigHandler
}

func (h *softWareConfigHandler) List(ctx context.Context, in *SoftWareConfigListRequest, out *SoftWareConfigListResponse) error {
	return h.SoftWareConfigHandler.List(ctx, in, out)
}
