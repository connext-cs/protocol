// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: host.proto

/*
Package host is a generated protocol buffer package.

It is generated from these files:
	host.proto

It has these top-level messages:
	CheckHostReq
	CheckHostRes
*/
package host

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Host service

type HostService interface {
	CheckHost(ctx context.Context, in *CheckHostReq, opts ...client.CallOption) (*CheckHostRes, error)
}

type hostService struct {
	c    client.Client
	name string
}

func NewHostService(name string, c client.Client) HostService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "host"
	}
	return &hostService{
		c:    c,
		name: name,
	}
}

func (c *hostService) CheckHost(ctx context.Context, in *CheckHostReq, opts ...client.CallOption) (*CheckHostRes, error) {
	req := c.c.NewRequest(c.name, "Host.CheckHost", in)
	out := new(CheckHostRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Host service

type HostHandler interface {
	CheckHost(context.Context, *CheckHostReq, *CheckHostRes) error
}

func RegisterHostHandler(s server.Server, hdlr HostHandler, opts ...server.HandlerOption) {
	type host interface {
		CheckHost(ctx context.Context, in *CheckHostReq, out *CheckHostRes) error
	}
	type Host struct {
		host
	}
	h := &hostHandler{hdlr}
	s.Handle(s.NewHandler(&Host{h}, opts...))
}

type hostHandler struct {
	HostHandler
}

func (h *hostHandler) CheckHost(ctx context.Context, in *CheckHostReq, out *CheckHostRes) error {
	return h.HostHandler.CheckHost(ctx, in, out)
}
