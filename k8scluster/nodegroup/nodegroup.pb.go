// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nodegroup.proto

package nodegroup

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

type Group struct {
	GroupID                 int32    `protobuf:"varint,1,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	ClusterNodeGroupName    string   `protobuf:"bytes,2,opt,name=ClusterNodeGroupName,proto3" json:"ClusterNodeGroupName,omitempty"`
	ClusterNodeGroupK8SName string   `protobuf:"bytes,3,opt,name=ClusterNodeGroupK8sName,proto3" json:"ClusterNodeGroupK8sName,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodegroup_07d96ea4ce83560c, []int{0}
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

func (m *Group) GetGroupID() int32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *Group) GetClusterNodeGroupName() string {
	if m != nil {
		return m.ClusterNodeGroupName
	}
	return ""
}

func (m *Group) GetClusterNodeGroupK8SName() string {
	if m != nil {
		return m.ClusterNodeGroupK8SName
	}
	return ""
}

type GroupNodeCount struct {
	NodeCount            int32    `protobuf:"varint,1,opt,name=NodeCount,proto3" json:"NodeCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GroupNodeCount) Reset()         { *m = GroupNodeCount{} }
func (m *GroupNodeCount) String() string { return proto.CompactTextString(m) }
func (*GroupNodeCount) ProtoMessage()    {}
func (*GroupNodeCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_nodegroup_07d96ea4ce83560c, []int{1}
}
func (m *GroupNodeCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GroupNodeCount.Unmarshal(m, b)
}
func (m *GroupNodeCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GroupNodeCount.Marshal(b, m, deterministic)
}
func (dst *GroupNodeCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupNodeCount.Merge(dst, src)
}
func (m *GroupNodeCount) XXX_Size() int {
	return xxx_messageInfo_GroupNodeCount.Size(m)
}
func (m *GroupNodeCount) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupNodeCount.DiscardUnknown(m)
}

var xxx_messageInfo_GroupNodeCount proto.InternalMessageInfo

func (m *GroupNodeCount) GetNodeCount() int32 {
	if m != nil {
		return m.NodeCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Group)(nil), "Group")
	proto.RegisterType((*GroupNodeCount)(nil), "GroupNodeCount")
}

func init() { proto.RegisterFile("nodegroup.proto", fileDescriptor_nodegroup_07d96ea4ce83560c) }

var fileDescriptor_nodegroup_07d96ea4ce83560c = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x1b, 0x4b, 0xab, 0x1d, 0xa5, 0x95, 0x55, 0x30, 0x44, 0x0f, 0x25, 0x17, 0x7b, 0x71,
	0x53, 0x53, 0x84, 0x5e, 0xb5, 0x15, 0x11, 0xa1, 0x60, 0xde, 0x20, 0xed, 0x0e, 0x31, 0x90, 0x64,
	0xe2, 0x66, 0xf7, 0xe0, 0x53, 0xf8, 0xa2, 0x3e, 0x84, 0x64, 0x53, 0x1b, 0x2c, 0x0a, 0xc9, 0x2d,
	0xf3, 0x7d, 0xf9, 0xcd, 0xdf, 0x85, 0x51, 0x46, 0x02, 0x23, 0x49, 0x3a, 0xe7, 0xb9, 0x24, 0x45,
	0xce, 0x75, 0x14, 0xab, 0x37, 0xbd, 0xe6, 0x1b, 0x4a, 0xbd, 0x34, 0xde, 0x48, 0xf2, 0x22, 0xba,
	0x09, 0xf3, 0xd8, 0x33, 0xb6, 0x17, 0xe6, 0x71, 0xf5, 0xa3, 0xfb, 0x69, 0x41, 0xef, 0xa9, 0x04,
	0x99, 0x0d, 0x87, 0xe6, 0xe3, 0x79, 0x69, 0x5b, 0x63, 0x6b, 0xd2, 0x0b, 0x7e, 0x42, 0xe6, 0xc3,
	0xf9, 0x22, 0xd1, 0x85, 0x42, 0xb9, 0x22, 0x81, 0x46, 0x5d, 0x85, 0x29, 0xda, 0x07, 0x63, 0x6b,
	0x32, 0x08, 0xfe, 0xf4, 0xd8, 0x1c, 0x2e, 0xf6, 0xf5, 0x97, 0x79, 0x61, 0xb0, 0xae, 0xc1, 0xfe,
	0xb3, 0x5d, 0x0e, 0xc3, 0x2a, 0x0d, 0x09, 0x5c, 0x90, 0xce, 0x14, 0xbb, 0x82, 0xc1, 0x2e, 0xd8,
	0xf6, 0x56, 0x0b, 0xfe, 0x57, 0xb7, 0xb2, 0xcd, 0xf8, 0xec, 0x16, 0xe0, 0x55, 0xa3, 0xfc, 0xa8,
	0xa2, 0x11, 0x8f, 0x88, 0x97, 0xc3, 0x06, 0xf8, 0xae, 0xb1, 0x50, 0xce, 0x69, 0x2d, 0x14, 0x39,
	0x65, 0x05, 0xba, 0x1d, 0xe6, 0xc1, 0xd1, 0xbd, 0x10, 0x2d, 0x80, 0x29, 0x0c, 0x1e, 0x45, 0xac,
	0x5a, 0x10, 0x3e, 0x1c, 0x2f, 0x31, 0x41, 0x85, 0x2d, 0x98, 0x19, 0x9c, 0x6c, 0xdb, 0x12, 0xe5,
	0x75, 0x1b, 0x43, 0x66, 0xfc, 0x30, 0x49, 0x9a, 0x43, 0x77, 0x30, 0xdc, 0xed, 0xac, 0x45, 0xad,
	0x07, 0xb8, 0xac, 0x6b, 0x29, 0x89, 0x49, 0xa8, 0xb0, 0x5a, 0x63, 0xf3, 0x1c, 0x53, 0x38, 0x33,
	0x39, 0xf6, 0x2e, 0xde, 0xe7, 0x46, 0x70, 0x46, 0xfc, 0xb7, 0xe1, 0x76, 0xd6, 0x7d, 0xf3, 0x6e,
	0x67, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xef, 0x58, 0x84, 0xf3, 0x02, 0x00, 0x00,
}
