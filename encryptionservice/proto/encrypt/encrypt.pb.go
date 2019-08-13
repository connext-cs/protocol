// Code generated by protoc-gen-go. DO NOT EDIT.
// source: encrypt.proto

/*
Package encrypt is a generated protocol buffer package.

It is generated from these files:
	encrypt.proto

It has these top-level messages:
	QueryEncryptSaltReq
	QueryEncryptSaltRes
*/
package encrypt

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type QueryEncryptSaltReq struct {
	SaltKey string `protobuf:"bytes,1,opt,name=SaltKey" json:"SaltKey,omitempty"`
}

func (m *QueryEncryptSaltReq) Reset()                    { *m = QueryEncryptSaltReq{} }
func (m *QueryEncryptSaltReq) String() string            { return proto.CompactTextString(m) }
func (*QueryEncryptSaltReq) ProtoMessage()               {}
func (*QueryEncryptSaltReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *QueryEncryptSaltReq) GetSaltKey() string {
	if m != nil {
		return m.SaltKey
	}
	return ""
}

type QueryEncryptSaltRes struct {
	Salt string `protobuf:"bytes,1,opt,name=Salt" json:"Salt,omitempty"`
}

func (m *QueryEncryptSaltRes) Reset()                    { *m = QueryEncryptSaltRes{} }
func (m *QueryEncryptSaltRes) String() string            { return proto.CompactTextString(m) }
func (*QueryEncryptSaltRes) ProtoMessage()               {}
func (*QueryEncryptSaltRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *QueryEncryptSaltRes) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryEncryptSaltReq)(nil), "QueryEncryptSaltReq")
	proto.RegisterType((*QueryEncryptSaltRes)(nil), "QueryEncryptSaltRes")
}

func init() { proto.RegisterFile("encrypt.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xcd, 0x4b, 0x2e,
	0xaa, 0x2c, 0x28, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xd2, 0xe7, 0x12, 0x0e, 0x2c, 0x4d,
	0x2d, 0xaa, 0x74, 0x85, 0x88, 0x06, 0x27, 0xe6, 0x94, 0x04, 0xa5, 0x16, 0x0a, 0x49, 0x70, 0xb1,
	0x83, 0x98, 0xde, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x92, 0x26,
	0x36, 0x0d, 0xc5, 0x42, 0x42, 0x5c, 0x2c, 0x20, 0x26, 0x54, 0x35, 0x98, 0x6d, 0xe4, 0xcd, 0xc5,
	0x0e, 0x55, 0x25, 0xe4, 0xc0, 0x25, 0x80, 0xae, 0x4b, 0x48, 0x44, 0x0f, 0x8b, 0xcd, 0x52, 0xd8,
	0x44, 0x8b, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0xee, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xea,
	0x2c, 0x06, 0xde, 0xc0, 0x00, 0x00, 0x00,
}
