// Code generated by protoc-gen-go. DO NOT EDIT.
// source: parentProject.proto

package parentProject

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/micro/go-api/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RpcQueryListByParentFuzzyNameReq struct {
	FuzzyName            string   `protobuf:"bytes,1,opt,name=FuzzyName,proto3" json:"FuzzyName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcQueryListByParentFuzzyNameReq) Reset()         { *m = RpcQueryListByParentFuzzyNameReq{} }
func (m *RpcQueryListByParentFuzzyNameReq) String() string { return proto.CompactTextString(m) }
func (*RpcQueryListByParentFuzzyNameReq) ProtoMessage()    {}
func (*RpcQueryListByParentFuzzyNameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_parentProject_165ee8eaf680be38, []int{0}
}
func (m *RpcQueryListByParentFuzzyNameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcQueryListByParentFuzzyNameReq.Unmarshal(m, b)
}
func (m *RpcQueryListByParentFuzzyNameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcQueryListByParentFuzzyNameReq.Marshal(b, m, deterministic)
}
func (dst *RpcQueryListByParentFuzzyNameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcQueryListByParentFuzzyNameReq.Merge(dst, src)
}
func (m *RpcQueryListByParentFuzzyNameReq) XXX_Size() int {
	return xxx_messageInfo_RpcQueryListByParentFuzzyNameReq.Size(m)
}
func (m *RpcQueryListByParentFuzzyNameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcQueryListByParentFuzzyNameReq.DiscardUnknown(m)
}

var xxx_messageInfo_RpcQueryListByParentFuzzyNameReq proto.InternalMessageInfo

func (m *RpcQueryListByParentFuzzyNameReq) GetFuzzyName() string {
	if m != nil {
		return m.FuzzyName
	}
	return ""
}

type Host struct {
	HostID               int32    `protobuf:"varint,1,opt,name=HostID,proto3" json:"HostID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Host) Reset()         { *m = Host{} }
func (m *Host) String() string { return proto.CompactTextString(m) }
func (*Host) ProtoMessage()    {}
func (*Host) Descriptor() ([]byte, []int) {
	return fileDescriptor_parentProject_165ee8eaf680be38, []int{1}
}
func (m *Host) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Host.Unmarshal(m, b)
}
func (m *Host) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Host.Marshal(b, m, deterministic)
}
func (dst *Host) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Host.Merge(dst, src)
}
func (m *Host) XXX_Size() int {
	return xxx_messageInfo_Host.Size(m)
}
func (m *Host) XXX_DiscardUnknown() {
	xxx_messageInfo_Host.DiscardUnknown(m)
}

var xxx_messageInfo_Host proto.InternalMessageInfo

func (m *Host) GetHostID() int32 {
	if m != nil {
		return m.HostID
	}
	return 0
}

type RpcListRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	QueryAll             bool     `protobuf:"varint,2,opt,name=QueryAll,proto3" json:"QueryAll,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcListRequest) Reset()         { *m = RpcListRequest{} }
func (m *RpcListRequest) String() string { return proto.CompactTextString(m) }
func (*RpcListRequest) ProtoMessage()    {}
func (*RpcListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_parentProject_165ee8eaf680be38, []int{2}
}
func (m *RpcListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcListRequest.Unmarshal(m, b)
}
func (m *RpcListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcListRequest.Marshal(b, m, deterministic)
}
func (dst *RpcListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcListRequest.Merge(dst, src)
}
func (m *RpcListRequest) XXX_Size() int {
	return xxx_messageInfo_RpcListRequest.Size(m)
}
func (m *RpcListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RpcListRequest proto.InternalMessageInfo

func (m *RpcListRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RpcListRequest) GetQueryAll() bool {
	if m != nil {
		return m.QueryAll
	}
	return false
}

type RpcListResponse struct {
	ParentProjects       []*ParentProjectInfo `protobuf:"bytes,1,rep,name=ParentProjects,proto3" json:"ParentProjects,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RpcListResponse) Reset()         { *m = RpcListResponse{} }
func (m *RpcListResponse) String() string { return proto.CompactTextString(m) }
func (*RpcListResponse) ProtoMessage()    {}
func (*RpcListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_parentProject_165ee8eaf680be38, []int{3}
}
func (m *RpcListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcListResponse.Unmarshal(m, b)
}
func (m *RpcListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcListResponse.Marshal(b, m, deterministic)
}
func (dst *RpcListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcListResponse.Merge(dst, src)
}
func (m *RpcListResponse) XXX_Size() int {
	return xxx_messageInfo_RpcListResponse.Size(m)
}
func (m *RpcListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RpcListResponse proto.InternalMessageInfo

func (m *RpcListResponse) GetParentProjects() []*ParentProjectInfo {
	if m != nil {
		return m.ParentProjects
	}
	return nil
}

type RpcQueryListRequest struct {
	ParendProjectId      int32    `protobuf:"varint,1,opt,name=ParendProjectId,proto3" json:"ParendProjectId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcQueryListRequest) Reset()         { *m = RpcQueryListRequest{} }
func (m *RpcQueryListRequest) String() string { return proto.CompactTextString(m) }
func (*RpcQueryListRequest) ProtoMessage()    {}
func (*RpcQueryListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_parentProject_165ee8eaf680be38, []int{4}
}
func (m *RpcQueryListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcQueryListRequest.Unmarshal(m, b)
}
func (m *RpcQueryListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcQueryListRequest.Marshal(b, m, deterministic)
}
func (dst *RpcQueryListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcQueryListRequest.Merge(dst, src)
}
func (m *RpcQueryListRequest) XXX_Size() int {
	return xxx_messageInfo_RpcQueryListRequest.Size(m)
}
func (m *RpcQueryListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcQueryListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RpcQueryListRequest proto.InternalMessageInfo

func (m *RpcQueryListRequest) GetParendProjectId() int32 {
	if m != nil {
		return m.ParendProjectId
	}
	return 0
}

type RpcQueryListByParentProjectIdsReq struct {
	ParentProjectIds     []int32  `protobuf:"varint,1,rep,packed,name=ParentProjectIds,proto3" json:"ParentProjectIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcQueryListByParentProjectIdsReq) Reset()         { *m = RpcQueryListByParentProjectIdsReq{} }
func (m *RpcQueryListByParentProjectIdsReq) String() string { return proto.CompactTextString(m) }
func (*RpcQueryListByParentProjectIdsReq) ProtoMessage()    {}
func (*RpcQueryListByParentProjectIdsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_parentProject_165ee8eaf680be38, []int{5}
}
func (m *RpcQueryListByParentProjectIdsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcQueryListByParentProjectIdsReq.Unmarshal(m, b)
}
func (m *RpcQueryListByParentProjectIdsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcQueryListByParentProjectIdsReq.Marshal(b, m, deterministic)
}
func (dst *RpcQueryListByParentProjectIdsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcQueryListByParentProjectIdsReq.Merge(dst, src)
}
func (m *RpcQueryListByParentProjectIdsReq) XXX_Size() int {
	return xxx_messageInfo_RpcQueryListByParentProjectIdsReq.Size(m)
}
func (m *RpcQueryListByParentProjectIdsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcQueryListByParentProjectIdsReq.DiscardUnknown(m)
}

var xxx_messageInfo_RpcQueryListByParentProjectIdsReq proto.InternalMessageInfo

func (m *RpcQueryListByParentProjectIdsReq) GetParentProjectIds() []int32 {
	if m != nil {
		return m.ParentProjectIds
	}
	return nil
}

type RpcQueryListResponse struct {
	Hosts                []int32  `protobuf:"varint,1,rep,packed,name=Hosts,proto3" json:"Hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcQueryListResponse) Reset()         { *m = RpcQueryListResponse{} }
func (m *RpcQueryListResponse) String() string { return proto.CompactTextString(m) }
func (*RpcQueryListResponse) ProtoMessage()    {}
func (*RpcQueryListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_parentProject_165ee8eaf680be38, []int{6}
}
func (m *RpcQueryListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcQueryListResponse.Unmarshal(m, b)
}
func (m *RpcQueryListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcQueryListResponse.Marshal(b, m, deterministic)
}
func (dst *RpcQueryListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcQueryListResponse.Merge(dst, src)
}
func (m *RpcQueryListResponse) XXX_Size() int {
	return xxx_messageInfo_RpcQueryListResponse.Size(m)
}
func (m *RpcQueryListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcQueryListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RpcQueryListResponse proto.InternalMessageInfo

func (m *RpcQueryListResponse) GetHosts() []int32 {
	if m != nil {
		return m.Hosts
	}
	return nil
}

type ParentProjectInfo struct {
	ParentProjectID      int32    `protobuf:"varint,1,opt,name=ParentProjectID,proto3" json:"ParentProjectID,omitempty"`
	ParentProjectName    string   `protobuf:"bytes,2,opt,name=ParentProjectName,proto3" json:"ParentProjectName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParentProjectInfo) Reset()         { *m = ParentProjectInfo{} }
func (m *ParentProjectInfo) String() string { return proto.CompactTextString(m) }
func (*ParentProjectInfo) ProtoMessage()    {}
func (*ParentProjectInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_parentProject_165ee8eaf680be38, []int{7}
}
func (m *ParentProjectInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParentProjectInfo.Unmarshal(m, b)
}
func (m *ParentProjectInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParentProjectInfo.Marshal(b, m, deterministic)
}
func (dst *ParentProjectInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentProjectInfo.Merge(dst, src)
}
func (m *ParentProjectInfo) XXX_Size() int {
	return xxx_messageInfo_ParentProjectInfo.Size(m)
}
func (m *ParentProjectInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentProjectInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ParentProjectInfo proto.InternalMessageInfo

func (m *ParentProjectInfo) GetParentProjectID() int32 {
	if m != nil {
		return m.ParentProjectID
	}
	return 0
}

func (m *ParentProjectInfo) GetParentProjectName() string {
	if m != nil {
		return m.ParentProjectName
	}
	return ""
}

type Group struct {
	HostGroupID          int32    `protobuf:"varint,1,opt,name=HostGroupID,proto3" json:"HostGroupID,omitempty"`
	HostGroupName        string   `protobuf:"bytes,2,opt,name=HostGroupName,proto3" json:"HostGroupName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_parentProject_165ee8eaf680be38, []int{8}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (dst *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(dst, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetHostGroupID() int32 {
	if m != nil {
		return m.HostGroupID
	}
	return 0
}

func (m *Group) GetHostGroupName() string {
	if m != nil {
		return m.HostGroupName
	}
	return ""
}

type RemoveResult struct {
	Error                string   `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveResult) Reset()         { *m = RemoveResult{} }
func (m *RemoveResult) String() string { return proto.CompactTextString(m) }
func (*RemoveResult) ProtoMessage()    {}
func (*RemoveResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_parentProject_165ee8eaf680be38, []int{9}
}
func (m *RemoveResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveResult.Unmarshal(m, b)
}
func (m *RemoveResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveResult.Marshal(b, m, deterministic)
}
func (dst *RemoveResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveResult.Merge(dst, src)
}
func (m *RemoveResult) XXX_Size() int {
	return xxx_messageInfo_RemoveResult.Size(m)
}
func (m *RemoveResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveResult.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveResult proto.InternalMessageInfo

func (m *RemoveResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type HostInfos struct {
	HostID               int32    `protobuf:"varint,1,opt,name=HostID,proto3" json:"HostID,omitempty"`
	HostName             string   `protobuf:"bytes,2,opt,name=HostName,proto3" json:"HostName,omitempty"`
	HostIP               string   `protobuf:"bytes,3,opt,name=HostIP,proto3" json:"HostIP,omitempty"`
	HostState            int32    `protobuf:"varint,4,opt,name=HostState,proto3" json:"HostState,omitempty"`
	HostSystem           int32    `protobuf:"varint,5,opt,name=HostSystem,proto3" json:"HostSystem,omitempty"`
	ResourceType         int32    `protobuf:"varint,6,opt,name=ResourceType,proto3" json:"ResourceType,omitempty"`
	HostType             int32    `protobuf:"varint,7,opt,name=HostType,proto3" json:"HostType,omitempty"`
	GroupID              int32    `protobuf:"varint,8,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	GroupName            string   `protobuf:"bytes,9,opt,name=GroupName,proto3" json:"GroupName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HostInfos) Reset()         { *m = HostInfos{} }
func (m *HostInfos) String() string { return proto.CompactTextString(m) }
func (*HostInfos) ProtoMessage()    {}
func (*HostInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_parentProject_165ee8eaf680be38, []int{10}
}
func (m *HostInfos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostInfos.Unmarshal(m, b)
}
func (m *HostInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostInfos.Marshal(b, m, deterministic)
}
func (dst *HostInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostInfos.Merge(dst, src)
}
func (m *HostInfos) XXX_Size() int {
	return xxx_messageInfo_HostInfos.Size(m)
}
func (m *HostInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_HostInfos.DiscardUnknown(m)
}

var xxx_messageInfo_HostInfos proto.InternalMessageInfo

func (m *HostInfos) GetHostID() int32 {
	if m != nil {
		return m.HostID
	}
	return 0
}

func (m *HostInfos) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *HostInfos) GetHostIP() string {
	if m != nil {
		return m.HostIP
	}
	return ""
}

func (m *HostInfos) GetHostState() int32 {
	if m != nil {
		return m.HostState
	}
	return 0
}

func (m *HostInfos) GetHostSystem() int32 {
	if m != nil {
		return m.HostSystem
	}
	return 0
}

func (m *HostInfos) GetResourceType() int32 {
	if m != nil {
		return m.ResourceType
	}
	return 0
}

func (m *HostInfos) GetHostType() int32 {
	if m != nil {
		return m.HostType
	}
	return 0
}

func (m *HostInfos) GetGroupID() int32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *HostInfos) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

type UpdateReq struct {
	ParentProjectID      int32    `protobuf:"varint,1,opt,name=ParentProjectID,proto3" json:"ParentProjectID,omitempty"`
	HostIDList           []int32  `protobuf:"varint,2,rep,packed,name=HostIDList,proto3" json:"HostIDList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateReq) Reset()         { *m = UpdateReq{} }
func (m *UpdateReq) String() string { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()    {}
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_parentProject_165ee8eaf680be38, []int{11}
}
func (m *UpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReq.Unmarshal(m, b)
}
func (m *UpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReq.Marshal(b, m, deterministic)
}
func (dst *UpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReq.Merge(dst, src)
}
func (m *UpdateReq) XXX_Size() int {
	return xxx_messageInfo_UpdateReq.Size(m)
}
func (m *UpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReq proto.InternalMessageInfo

func (m *UpdateReq) GetParentProjectID() int32 {
	if m != nil {
		return m.ParentProjectID
	}
	return 0
}

func (m *UpdateReq) GetHostIDList() []int32 {
	if m != nil {
		return m.HostIDList
	}
	return nil
}

type UpdateRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRes) Reset()         { *m = UpdateRes{} }
func (m *UpdateRes) String() string { return proto.CompactTextString(m) }
func (*UpdateRes) ProtoMessage()    {}
func (*UpdateRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_parentProject_165ee8eaf680be38, []int{12}
}
func (m *UpdateRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRes.Unmarshal(m, b)
}
func (m *UpdateRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRes.Marshal(b, m, deterministic)
}
func (dst *UpdateRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRes.Merge(dst, src)
}
func (m *UpdateRes) XXX_Size() int {
	return xxx_messageInfo_UpdateRes.Size(m)
}
func (m *UpdateRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRes proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RpcQueryListByParentFuzzyNameReq)(nil), "RpcQueryListByParentFuzzyNameReq")
	proto.RegisterType((*Host)(nil), "Host")
	proto.RegisterType((*RpcListRequest)(nil), "RpcListRequest")
	proto.RegisterType((*RpcListResponse)(nil), "RpcListResponse")
	proto.RegisterType((*RpcQueryListRequest)(nil), "RpcQueryListRequest")
	proto.RegisterType((*RpcQueryListByParentProjectIdsReq)(nil), "RpcQueryListByParentProjectIdsReq")
	proto.RegisterType((*RpcQueryListResponse)(nil), "RpcQueryListResponse")
	proto.RegisterType((*ParentProjectInfo)(nil), "ParentProjectInfo")
	proto.RegisterType((*Group)(nil), "Group")
	proto.RegisterType((*RemoveResult)(nil), "RemoveResult")
	proto.RegisterType((*HostInfos)(nil), "HostInfos")
	proto.RegisterType((*UpdateReq)(nil), "UpdateReq")
	proto.RegisterType((*UpdateRes)(nil), "UpdateRes")
}

func init() { proto.RegisterFile("parentProject.proto", fileDescriptor_parentProject_165ee8eaf680be38) }

var fileDescriptor_parentProject_165ee8eaf680be38 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0xa5, 0x4e, 0x9b, 0xe9, 0x25, 0xe9, 0x36, 0x20, 0x2b, 0x82, 0x2a, 0x5d, 0x15, 0x11,
	0x95, 0x74, 0x23, 0xca, 0x03, 0x08, 0x09, 0x41, 0xab, 0xb4, 0x10, 0x09, 0xda, 0x60, 0xe8, 0x13,
	0xe2, 0x21, 0xb5, 0xb7, 0x25, 0x90, 0x64, 0x5d, 0xaf, 0x8d, 0x48, 0xff, 0x81, 0x7f, 0xe5, 0x13,
	0xd0, 0xce, 0xfa, 0x9a, 0xde, 0xfc, 0x64, 0xcf, 0x99, 0xb3, 0xc7, 0x33, 0xb3, 0x33, 0x63, 0xd8,
	0x70, 0x87, 0x1e, 0x9f, 0xfa, 0x03, 0x4f, 0xfc, 0xe4, 0xb6, 0xcf, 0x5c, 0x4f, 0xf8, 0xa2, 0xf9,
	0xf4, 0x62, 0xe4, 0xff, 0x08, 0xce, 0x98, 0x2d, 0x26, 0xdd, 0xc9, 0xc8, 0xf6, 0x44, 0xf7, 0x42,
	0xec, 0x0e, 0xdd, 0x51, 0x17, 0xdd, 0xdd, 0xa1, 0x3b, 0xd2, 0x44, 0xfa, 0x0e, 0x5a, 0x96, 0x6b,
	0x7f, 0x0e, 0xb8, 0x37, 0xfb, 0x38, 0x92, 0xfe, 0xc1, 0x6c, 0x80, 0x6a, 0x47, 0xc1, 0xd5, 0xd5,
	0xec, 0x78, 0x38, 0xe1, 0x16, 0xbf, 0x24, 0x8f, 0xa0, 0x1a, 0xdb, 0x66, 0xb1, 0x55, 0x6c, 0x57,
	0xad, 0x04, 0xa0, 0x9b, 0xb0, 0xf0, 0x41, 0x48, 0x9f, 0x3c, 0x84, 0x8a, 0x7a, 0xf6, 0x7b, 0x48,
	0x31, 0xac, 0xd0, 0xa2, 0x3d, 0x58, 0xb3, 0x5c, 0x5b, 0x89, 0x5b, 0xfc, 0x32, 0xe0, 0x9a, 0x79,
	0x2a, 0xb9, 0xd7, 0x77, 0x90, 0x59, 0xb6, 0x42, 0x8b, 0x34, 0x61, 0x09, 0x03, 0xd9, 0x1f, 0x8f,
	0xcd, 0x52, 0xab, 0xd8, 0x5e, 0xb2, 0x62, 0x9b, 0x7e, 0x82, 0x5a, 0xac, 0x22, 0x5d, 0x31, 0x95,
	0x9c, 0xbc, 0x86, 0xb5, 0x41, 0x3a, 0x75, 0x69, 0x16, 0x5b, 0xe5, 0xf6, 0xf2, 0x1e, 0x61, 0x19,
	0xb8, 0x3f, 0x3d, 0x17, 0xd6, 0x1c, 0x93, 0xbe, 0x85, 0x8d, 0x74, 0xda, 0x51, 0x64, 0x6d, 0xa8,
	0x21, 0xd1, 0x89, 0xce, 0x3a, 0x61, 0x32, 0xf3, 0x30, 0x3d, 0x81, 0xad, 0x9b, 0xea, 0x16, 0x13,
	0xa4, 0x2a, 0xdc, 0x0e, 0xd4, 0xe7, 0x61, 0x8c, 0xd1, 0xb0, 0xae, 0xe1, 0xb4, 0x03, 0x8d, 0x6c,
	0x44, 0x61, 0x96, 0x0d, 0x30, 0x54, 0x21, 0xa3, 0x83, 0xda, 0xa0, 0xbf, 0x60, 0xfd, 0x5a, 0x92,
	0x71, 0xf4, 0x31, 0xd8, 0xcb, 0x44, 0x9f, 0xc0, 0xa4, 0x33, 0x77, 0x1c, 0x6f, 0xb6, 0x84, 0x37,
	0x7b, 0xdd, 0x41, 0x4f, 0xc0, 0x78, 0xef, 0x89, 0xc0, 0x25, 0x2d, 0x58, 0x56, 0x9f, 0x47, 0x23,
	0x16, 0x4f, 0x43, 0x64, 0x1b, 0x56, 0x63, 0x33, 0x25, 0x9a, 0x05, 0xe9, 0x36, 0xac, 0x58, 0x7c,
	0x22, 0x7e, 0x73, 0x8b, 0xcb, 0x60, 0xec, 0xab, 0x1c, 0x0f, 0x3d, 0x4f, 0x78, 0x61, 0x73, 0x69,
	0x83, 0xfe, 0x2d, 0x41, 0x15, 0x7b, 0x68, 0x7a, 0x2e, 0xe4, 0x6d, 0xed, 0xa5, 0x9a, 0x46, 0xbd,
	0xa5, 0x3e, 0x16, 0xdb, 0xf1, 0x99, 0x81, 0x59, 0x46, 0x4f, 0x68, 0xa9, 0x86, 0x56, 0x6f, 0x5f,
	0xfc, 0xa1, 0xcf, 0xcd, 0x05, 0x94, 0x4b, 0x00, 0xb2, 0x09, 0x80, 0xc6, 0x4c, 0xfa, 0x7c, 0x62,
	0x1a, 0xe8, 0x4e, 0x21, 0x84, 0xaa, 0xe8, 0xa5, 0x08, 0x3c, 0x9b, 0x7f, 0x9d, 0xb9, 0xdc, 0xac,
	0x20, 0x23, 0x83, 0x45, 0x51, 0xa1, 0x7f, 0x11, 0xfd, 0xb1, 0x4d, 0x4c, 0x58, 0x8c, 0x2a, 0xb8,
	0x84, 0xae, 0xc8, 0x54, 0x71, 0x25, 0x95, 0xab, 0xea, 0x41, 0x4b, 0xaa, 0x76, 0x0a, 0xd5, 0x53,
	0xd7, 0x19, 0xfa, 0x38, 0x93, 0xf9, 0xef, 0x3a, 0x4c, 0xa7, 0xdf, 0x53, 0x6d, 0x65, 0x96, 0xb0,
	0x8b, 0x52, 0x08, 0x5d, 0x4e, 0x64, 0xe5, 0xde, 0xbf, 0x0a, 0xac, 0x66, 0x04, 0xc8, 0x0e, 0x94,
	0xf7, 0x1d, 0x87, 0xd4, 0xd8, 0x85, 0x60, 0x6a, 0x6d, 0x84, 0xa3, 0xd2, 0xac, 0x27, 0x80, 0xee,
	0x54, 0x5a, 0x20, 0xcf, 0x60, 0x41, 0x49, 0xe6, 0x23, 0x77, 0xc0, 0xc0, 0x6e, 0xcf, 0xc7, 0xde,
	0x85, 0x8a, 0x8e, 0x32, 0x1f, 0xfd, 0x09, 0x54, 0x2d, 0xd7, 0x0e, 0x4f, 0x00, 0x8b, 0xeb, 0xd6,
	0x4c, 0xde, 0xa5, 0x56, 0xed, 0xf1, 0x31, 0xcf, 0xab, 0xda, 0x01, 0xe3, 0xf0, 0x4f, 0xee, 0x04,
	0x5f, 0xc1, 0xfa, 0xb1, 0x98, 0x72, 0xbc, 0x40, 0x55, 0xef, 0xfc, 0xa5, 0xd9, 0x86, 0x75, 0x2c,
	0x8d, 0xbe, 0x89, 0x83, 0x19, 0xee, 0x57, 0x83, 0xa9, 0x47, 0xb3, 0xc2, 0x50, 0x90, 0x16, 0x08,
	0x83, 0x86, 0x9e, 0x22, 0x85, 0x1f, 0x79, 0x62, 0xa2, 0xe9, 0x11, 0x71, 0x95, 0xa5, 0x67, 0x8c,
	0x16, 0xc8, 0x4b, 0xa8, 0xef, 0x3b, 0x8e, 0x9e, 0x28, 0x5f, 0xe8, 0x89, 0xce, 0x15, 0x0e, 0x83,
	0xc5, 0x70, 0xf7, 0x92, 0x1a, 0xcb, 0xee, 0xf2, 0x66, 0x9d, 0xcd, 0xad, 0x65, 0x5a, 0x20, 0x6f,
	0x60, 0x25, 0xbd, 0xca, 0x48, 0x83, 0xdd, 0xb0, 0x6b, 0x9b, 0x0f, 0xd8, 0x4d, 0xfb, 0x8e, 0x16,
	0xc8, 0x77, 0xd8, 0xbc, 0x7b, 0xb5, 0x12, 0xca, 0xee, 0xdd, 0xbd, 0xb7, 0xcb, 0x7f, 0x83, 0xc7,
	0x77, 0xfe, 0xf1, 0xc8, 0x16, 0xbb, 0xef, 0x8f, 0x78, 0xbb, 0xf8, 0x73, 0x00, 0x84, 0x71, 0x4b,
	0xe7, 0xaa, 0xee, 0x59, 0x05, 0x7f, 0xc4, 0x2f, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x8b,
	0x84, 0x6c, 0xc8, 0x07, 0x00, 0x00,
}