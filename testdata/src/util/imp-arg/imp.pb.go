// Code generated by protoc-gen-go. DO NOT EDIT.
// source: util/imp-arg/imp.gunk

/*
Package imp is a generated protocol buffer package.

It is generated from these files:
	util/imp-arg/imp.gunk

It has these top-level messages:
	Message
*/
package imp

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

// Message is a Echo message.
type Message struct {
	// Msg holds a message.
	Msg string `protobuf:"bytes,0,,name=Msg" json:"Msg,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "util/imp-arg.Message")
}

func init() { proto.RegisterFile("util/imp-arg/imp.gunk", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 86 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2d, 0xc9, 0xcc,
	0xd1, 0xcf, 0xcc, 0x2d, 0xd0, 0x4d, 0x2c, 0x4a, 0x07, 0xd1, 0x7a, 0xe9, 0xa5, 0x79, 0xd9, 0x42,
	0x3c, 0xc8, 0xc2, 0x4a, 0x22, 0x5c, 0xec, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x9c,
	0x5c, 0xcc, 0xbe, 0xc5, 0xe9, 0x12, 0x0c, 0x1a, 0x9c, 0x4e, 0xac, 0x51, 0xcc, 0x99, 0xb9, 0x05,
	0x49, 0x6c, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x2b,
	0x45, 0x67, 0x4a, 0x00, 0x00, 0x00,
}
