// Code generated by protoc-gen-go. DO NOT EDIT.
// source: project.proto

package project

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/micro/go-api/proto"
	math "math"
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

func init() { proto.RegisterFile("project.proto", fileDescriptor_8340e6318dfdfac2) }

var fileDescriptor_8340e6318dfdfac2 = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x97, 0x6f, 0x6b, 0x13, 0x4f,
	0x10, 0xc7, 0x7f, 0x0f, 0x7e, 0x58, 0x18, 0x49, 0x53, 0x63, 0x55, 0xac, 0x05, 0x41, 0x10, 0x1f,
	0x99, 0x60, 0xed, 0x1f, 0x6b, 0x45, 0xac, 0xa9, 0x82, 0x50, 0x8a, 0xa6, 0xea, 0xd3, 0x63, 0xef,
	0x76, 0x7a, 0xd9, 0x76, 0x6f, 0xe7, 0xdc, 0xdd, 0x6b, 0xac, 0xaf, 0xc9, 0x37, 0xe5, 0x3b, 0x91,
	0xcb, 0xad, 0x49, 0x4e, 0x10, 0x66, 0xd3, 0x67, 0x21, 0xb9, 0xcf, 0x77, 0xf6, 0x66, 0x67, 0xe6,
	0x3b, 0x81, 0x4e, 0x69, 0xe9, 0x1c, 0x33, 0xdf, 0x2f, 0x2d, 0x79, 0xda, 0x78, 0x92, 0x2b, 0x3f,
	0xae, 0xd2, 0x7e, 0x46, 0xc5, 0xa0, 0x50, 0x99, 0xa5, 0x41, 0x4e, 0x4f, 0x45, 0xa9, 0x06, 0xd3,
	0x9f, 0x07, 0xa2, 0x54, 0xcd, 0x83, 0x5b, 0x3f, 0xef, 0xc2, 0xca, 0xc7, 0x06, 0xed, 0xed, 0x42,
	0xf7, 0x04, 0x27, 0x41, 0x28, 0x1b, 0x63, 0x76, 0xd1, 0xeb, 0xf6, 0x73, 0xea, 0xd7, 0x4f, 0x8f,
	0xf0, 0x5b, 0x85, 0xce, 0x6f, 0xac, 0xcd, 0xbf, 0x70, 0x25, 0x19, 0x87, 0x8f, 0xfe, 0xeb, 0x3d,
	0x03, 0x98, 0x73, 0x3c, 0x64, 0x0b, 0x6e, 0x86, 0xa8, 0x5a, 0x39, 0x26, 0xb3, 0x0f, 0xbd, 0xc0,
	0x64, 0x54, 0x94, 0x4a, 0xa3, 0x14, 0x5e, 0xf0, 0xd0, 0x3d, 0x58, 0x0b, 0xa8, 0x2a, 0x44, 0x5e,
	0x83, 0xc8, 0x03, 0xdf, 0xc0, 0xfd, 0x00, 0x26, 0x13, 0x61, 0x71, 0x4c, 0x95, 0xc3, 0x44, 0x48,
	0x69, 0xd1, 0x39, 0x76, 0x72, 0x8e, 0x50, 0x47, 0x25, 0x67, 0x17, 0xba, 0x23, 0xd2, 0x3a, 0x15,
	0xd9, 0x45, 0x6c, 0x52, 0x4f, 0x3d, 0x95, 0x51, 0xcc, 0x36, 0x74, 0xbe, 0x94, 0x75, 0x3e, 0xa2,
	0xa8, 0x1d, 0x58, 0x0d, 0x69, 0xb1, 0xa8, 0x51, 0x38, 0x66, 0x36, 0x5f, 0xc1, 0x9d, 0x36, 0x36,
	0x56, 0xce, 0x93, 0xbd, 0xe2, 0xd1, 0x6f, 0xe1, 0x41, 0x9b, 0x3e, 0xa7, 0x54, 0x53, 0x5e, 0x6b,
	0x58, 0xba, 0x8a, 0xad, 0xa1, 0xa0, 0xa1, 0xcc, 0x19, 0xf1, 0xd0, 0x03, 0x58, 0x1f, 0x35, 0x4c,
	0x48, 0x95, 0xf3, 0xc2, 0x57, 0x2e, 0x36, 0x61, 0x25, 0x49, 0x7e, 0xc9, 0xcf, 0x6e, 0x27, 0x23,
	0x89, 0x9a, 0x72, 0x76, 0xfd, 0x84, 0x60, 0x8e, 0x8c, 0xb0, 0x6c, 0xee, 0x05, 0xdc, 0x6a, 0x37,
	0x18, 0x9b, 0x3c, 0x80, 0xf5, 0x40, 0xa6, 0x95, 0xd2, 0x72, 0xda, 0x64, 0x6c, 0xf8, 0x25, 0xdc,
	0x6e, 0xdf, 0x49, 0x26, 0x97, 0x78, 0xd5, 0x34, 0xea, 0x32, 0x77, 0xa1, 0x1b, 0xca, 0x9e, 0xa4,
	0xa9, 0x8a, 0x14, 0x2d, 0x7b, 0x90, 0x9c, 0xa2, 0xb0, 0xd9, 0x38, 0xd4, 0x00, 0xfb, 0xa0, 0x3b,
	0xb0, 0xfa, 0x55, 0xcd, 0x86, 0x24, 0x1b, 0x1b, 0xc2, 0x66, 0x73, 0xce, 0x24, 0x90, 0x89, 0x32,
	0x06, 0x6d, 0x62, 0xd0, 0x4f, 0xc8, 0x5e, 0xb0, 0x87, 0xd8, 0x09, 0x4e, 0xae, 0xa9, 0x70, 0x84,
	0xfa, 0x3a, 0x0a, 0x43, 0xd8, 0x5c, 0x3c, 0x03, 0x7e, 0xf7, 0x68, 0x8d, 0xd0, 0x71, 0x22, 0xc7,
	0xf0, 0x78, 0x58, 0xdb, 0xd2, 0x3f, 0x65, 0x12, 0x49, 0x85, 0x50, 0x86, 0xa7, 0xf6, 0x1e, 0x1e,
	0xfe, 0x95, 0xdb, 0xe5, 0x4e, 0x35, 0x84, 0xcd, 0xc5, 0xe4, 0x2c, 0x27, 0x32, 0x77, 0x28, 0x57,
	0x15, 0x85, 0xb0, 0xea, 0x07, 0xb2, 0xc1, 0x43, 0x29, 0xff, 0x98, 0x36, 0x99, 0x33, 0xc5, 0x6f,
	0xbb, 0x23, 0xd4, 0x38, 0x9b, 0xfc, 0x91, 0x6c, 0xcb, 0x35, 0x62, 0xd8, 0x7d, 0xe8, 0x7d, 0xaa,
	0xd0, 0x5e, 0x49, 0x3c, 0x13, 0x95, 0x8e, 0x42, 0x77, 0x60, 0x75, 0x84, 0x8e, 0x2a, 0x9b, 0x61,
	0x0c, 0x76, 0x00, 0xeb, 0xd3, 0x88, 0xb3, 0x2b, 0x32, 0xe4, 0x55, 0x86, 0xb1, 0x30, 0xc9, 0xa4,
	0x9e, 0xdc, 0x09, 0x5e, 0xa2, 0x61, 0xce, 0xef, 0x43, 0xd8, 0x68, 0x47, 0x9e, 0x0e, 0xe4, 0x24,
	0xc6, 0x39, 0x5e, 0xc3, 0xbd, 0xb6, 0x44, 0x33, 0x24, 0x95, 0xe1, 0x5f, 0xd5, 0xb4, 0x67, 0x02,
	0x7e, 0x89, 0xd6, 0x29, 0x32, 0xec, 0xda, 0x5a, 0x64, 0x8d, 0x28, 0x90, 0xbd, 0xf4, 0x1c, 0x53,
	0xfe, 0x19, 0x8b, 0x52, 0x23, 0x7f, 0xa5, 0x98, 0x23, 0x1f, 0xd8, 0x83, 0x7c, 0x1b, 0x3a, 0x27,
	0x38, 0x89, 0x0d, 0xb6, 0x07, 0x6b, 0x4d, 0xed, 0x47, 0x82, 0x5b, 0xbf, 0xfe, 0x87, 0x95, 0x53,
	0xb4, 0x97, 0x75, 0xf9, 0xd4, 0xeb, 0x56, 0xf3, 0x91, 0x6f, 0xe8, 0xcd, 0xaa, 0xec, 0x82, 0x02,
	0xd7, 0x95, 0x43, 0x98, 0xcc, 0xa2, 0xf0, 0xb8, 0xc4, 0xf6, 0x10, 0x15, 0xaf, 0x59, 0x58, 0xa3,
	0x90, 0xa6, 0x2f, 0xbd, 0xb0, 0x3e, 0x0a, 0xdb, 0x86, 0x4e, 0x78, 0x33, 0x89, 0x5e, 0x28, 0x1d,
	0x4b, 0x95, 0xc2, 0x8a, 0xc2, 0x45, 0xf9, 0x6f, 0x38, 0x1f, 0xdb, 0x7f, 0x67, 0x7e, 0x1f, 0x0b,
	0xee, 0x43, 0xef, 0x5d, 0x30, 0x82, 0x42, 0x18, 0x91, 0x63, 0xc1, 0x1d, 0x1a, 0xe9, 0x8d, 0xe9,
	0x3f, 0xb3, 0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x34, 0x7c, 0x01, 0x07, 0xd3, 0x0d, 0x00,
	0x00,
}