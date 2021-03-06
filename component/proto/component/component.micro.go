// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: component.proto

/*
Package component is a generated protocol buffer package.

It is generated from these files:
	component.proto

It has these top-level messages:
*/
package component

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_api "github.com/micro/go-api/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = go_api.Response{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Component service

type ComponentService interface {
	List(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	Add(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	Delete(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	Update(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	Urlcheck(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	Namelist(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	NewGitInfo(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	QueryGitInfo(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	GetProject(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	QueryProjectReadme(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	SyncAllProjectInfo(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	ClearProjectInfo(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
}

type componentService struct {
	c    client.Client
	name string
}

func NewComponentService(name string, c client.Client) ComponentService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "component"
	}
	return &componentService{
		c:    c,
		name: name,
	}
}

func (c *componentService) List(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Component.List", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentService) Add(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Component.Add", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentService) Delete(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Component.Delete", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentService) Update(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Component.Update", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentService) Urlcheck(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Component.Urlcheck", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentService) Namelist(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Component.Namelist", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentService) NewGitInfo(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Component.NewGitInfo", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentService) QueryGitInfo(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Component.QueryGitInfo", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentService) GetProject(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Component.GetProject", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentService) QueryProjectReadme(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Component.QueryProjectReadme", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentService) SyncAllProjectInfo(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Component.SyncAllProjectInfo", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentService) ClearProjectInfo(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.name, "Component.ClearProjectInfo", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Component service

type ComponentHandler interface {
	List(context.Context, *go_api.Request, *go_api.Response) error
	Add(context.Context, *go_api.Request, *go_api.Response) error
	Delete(context.Context, *go_api.Request, *go_api.Response) error
	Update(context.Context, *go_api.Request, *go_api.Response) error
	Urlcheck(context.Context, *go_api.Request, *go_api.Response) error
	Namelist(context.Context, *go_api.Request, *go_api.Response) error
	NewGitInfo(context.Context, *go_api.Request, *go_api.Response) error
	QueryGitInfo(context.Context, *go_api.Request, *go_api.Response) error
	GetProject(context.Context, *go_api.Request, *go_api.Response) error
	QueryProjectReadme(context.Context, *go_api.Request, *go_api.Response) error
	SyncAllProjectInfo(context.Context, *go_api.Request, *go_api.Response) error
	ClearProjectInfo(context.Context, *go_api.Request, *go_api.Response) error
}

func RegisterComponentHandler(s server.Server, hdlr ComponentHandler, opts ...server.HandlerOption) {
	type component interface {
		List(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		Add(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		Delete(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		Update(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		Urlcheck(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		Namelist(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		NewGitInfo(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		QueryGitInfo(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		GetProject(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		QueryProjectReadme(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		SyncAllProjectInfo(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		ClearProjectInfo(ctx context.Context, in *go_api.Request, out *go_api.Response) error
	}
	type Component struct {
		component
	}
	h := &componentHandler{hdlr}
	s.Handle(s.NewHandler(&Component{h}, opts...))
}

type componentHandler struct {
	ComponentHandler
}

func (h *componentHandler) List(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.ComponentHandler.List(ctx, in, out)
}

func (h *componentHandler) Add(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.ComponentHandler.Add(ctx, in, out)
}

func (h *componentHandler) Delete(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.ComponentHandler.Delete(ctx, in, out)
}

func (h *componentHandler) Update(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.ComponentHandler.Update(ctx, in, out)
}

func (h *componentHandler) Urlcheck(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.ComponentHandler.Urlcheck(ctx, in, out)
}

func (h *componentHandler) Namelist(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.ComponentHandler.Namelist(ctx, in, out)
}

func (h *componentHandler) NewGitInfo(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.ComponentHandler.NewGitInfo(ctx, in, out)
}

func (h *componentHandler) QueryGitInfo(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.ComponentHandler.QueryGitInfo(ctx, in, out)
}

func (h *componentHandler) GetProject(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.ComponentHandler.GetProject(ctx, in, out)
}

func (h *componentHandler) QueryProjectReadme(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.ComponentHandler.QueryProjectReadme(ctx, in, out)
}

func (h *componentHandler) SyncAllProjectInfo(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.ComponentHandler.SyncAllProjectInfo(ctx, in, out)
}

func (h *componentHandler) ClearProjectInfo(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.ComponentHandler.ClearProjectInfo(ctx, in, out)
}
