// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: manageJobList.proto

/*
Package manageJobList is a generated protocol buffer package.

It is generated from these files:
	manageJobList.proto

It has these top-level messages:
	JobProcessReportRequest
	JobProcessReportResponse
	ClusterNodeJobProcessReportRequest
	ClusterNodeJobProcessReportResponse
*/
package manageJobList

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

// Client API for ManageJobList service

type ManageJobListService interface {
	JobProcessReport(ctx context.Context, in *JobProcessReportRequest, opts ...client.CallOption) (*JobProcessReportResponse, error)
	ClusterNodeJobProcessReport(ctx context.Context, in *ClusterNodeJobProcessReportRequest, opts ...client.CallOption) (*ClusterNodeJobProcessReportResponse, error)
}

type manageJobListService struct {
	c    client.Client
	name string
}

func NewManageJobListService(name string, c client.Client) ManageJobListService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "managejoblist"
	}
	return &manageJobListService{
		c:    c,
		name: name,
	}
}

func (c *manageJobListService) JobProcessReport(ctx context.Context, in *JobProcessReportRequest, opts ...client.CallOption) (*JobProcessReportResponse, error) {
	req := c.c.NewRequest(c.name, "ManageJobList.JobProcessReport", in)
	out := new(JobProcessReportResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *manageJobListService) ClusterNodeJobProcessReport(ctx context.Context, in *ClusterNodeJobProcessReportRequest, opts ...client.CallOption) (*ClusterNodeJobProcessReportResponse, error) {
	req := c.c.NewRequest(c.name, "ManageJobList.ClusterNodeJobProcessReport", in)
	out := new(ClusterNodeJobProcessReportResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ManageJobList service

type ManageJobListHandler interface {
	JobProcessReport(context.Context, *JobProcessReportRequest, *JobProcessReportResponse) error
	ClusterNodeJobProcessReport(context.Context, *ClusterNodeJobProcessReportRequest, *ClusterNodeJobProcessReportResponse) error
}

func RegisterManageJobListHandler(s server.Server, hdlr ManageJobListHandler, opts ...server.HandlerOption) {
	type manageJobList interface {
		JobProcessReport(ctx context.Context, in *JobProcessReportRequest, out *JobProcessReportResponse) error
		ClusterNodeJobProcessReport(ctx context.Context, in *ClusterNodeJobProcessReportRequest, out *ClusterNodeJobProcessReportResponse) error
	}
	type ManageJobList struct {
		manageJobList
	}
	h := &manageJobListHandler{hdlr}
	s.Handle(s.NewHandler(&ManageJobList{h}, opts...))
}

type manageJobListHandler struct {
	ManageJobListHandler
}

func (h *manageJobListHandler) JobProcessReport(ctx context.Context, in *JobProcessReportRequest, out *JobProcessReportResponse) error {
	return h.ManageJobListHandler.JobProcessReport(ctx, in, out)
}

func (h *manageJobListHandler) ClusterNodeJobProcessReport(ctx context.Context, in *ClusterNodeJobProcessReportRequest, out *ClusterNodeJobProcessReportResponse) error {
	return h.ManageJobListHandler.ClusterNodeJobProcessReport(ctx, in, out)
}
