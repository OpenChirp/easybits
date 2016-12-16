// Code generated by protoc-gen-go.
// source: mio.proto
// DO NOT EDIT!

/*
Package mio is a generated protocol buffer package.

It is generated from these files:
	mio.proto

It has these top-level messages:
	MIOValue
*/
package mio

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

type MIOValue struct {
	Field uint32 `protobuf:"varint,1,opt,name=field" json:"field,omitempty"`
	Value int32  `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *MIOValue) Reset()                    { *m = MIOValue{} }
func (m *MIOValue) String() string            { return proto.CompactTextString(m) }
func (*MIOValue) ProtoMessage()               {}
func (*MIOValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*MIOValue)(nil), "MIOValue")
}

func init() { proto.RegisterFile("mio.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 83 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0xcd, 0xcc, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe3, 0xe2, 0xf0, 0xf5, 0xf4, 0x0f, 0x4b, 0xcc, 0x29,
	0x4d, 0x15, 0x12, 0xe1, 0x62, 0x4d, 0xcb, 0x4c, 0xcd, 0x49, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0d, 0x82, 0x70, 0x40, 0xa2, 0x65, 0x20, 0x69, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x08,
	0x27, 0x89, 0x0d, 0xac, 0xdd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x48, 0xca, 0x86, 0x4b,
	0x00, 0x00, 0x00,
}
