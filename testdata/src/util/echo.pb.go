// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testdata/src/util/echo.gunk

/*
Package util is a generated protocol buffer package.

package util contains a simple Echo service.

It is generated from these files:
	testdata/src/util/echo.gunk
	testdata/src/util/types.gunk

It has these top-level messages:
	Message
	CheckStatusResponse
*/
package util

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Empty from public import google/protobuf/empty.proto
type Empty google_protobuf.Empty

func (m *Empty) Reset()         { (*google_protobuf.Empty)(m).Reset() }
func (m *Empty) String() string { return (*google_protobuf.Empty)(m).String() }
func (*Empty) ProtoMessage()    {}

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

// CheckStatusResponse is the response for a check status.
type CheckStatusResponse struct {
	Status Status `protobuf:"varint,1,,name=Status,enum=util.Status" json:"Status,omitempty"`
}

func (m *CheckStatusResponse) Reset()                    { *m = CheckStatusResponse{} }
func (m *CheckStatusResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckStatusResponse) ProtoMessage()               {}
func (*CheckStatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CheckStatusResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_Unknown
}

func init() {
	proto.RegisterType((*Message)(nil), "util.Message")
	proto.RegisterType((*CheckStatusResponse)(nil), "util.CheckStatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Util service

type UtilClient interface {
	// Echo echoes a message.
	Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	// CheckStatus sends the server health status.
	CheckStatus(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*CheckStatusResponse, error)
}

type utilClient struct {
	cc *grpc.ClientConn
}

func NewUtilClient(cc *grpc.ClientConn) UtilClient {
	return &utilClient{cc}
}

func (c *utilClient) Echo(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/util.Util/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *utilClient) CheckStatus(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*CheckStatusResponse, error) {
	out := new(CheckStatusResponse)
	err := grpc.Invoke(ctx, "/util.Util/CheckStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Util service

type UtilServer interface {
	// Echo echoes a message.
	Echo(context.Context, *Message) (*Message, error)
	// CheckStatus sends the server health status.
	CheckStatus(context.Context, *google_protobuf.Empty) (*CheckStatusResponse, error)
}

func RegisterUtilServer(s *grpc.Server, srv UtilServer) {
	s.RegisterService(&_Util_serviceDesc, srv)
}

func _Util_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/util.Util/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilServer).Echo(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Util_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UtilServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/util.Util/CheckStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UtilServer).CheckStatus(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Util_serviceDesc = grpc.ServiceDesc{
	ServiceName: "util.Util",
	HandlerType: (*UtilServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Util_Echo_Handler,
		},
		{
			MethodName: "CheckStatus",
			Handler:    _Util_CheckStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testdata/src/util/echo.gunk",
}

func init() { proto.RegisterFile("testdata/src/util/echo.gunk", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0x57, 0x2c, 0x95, 0x1d, 0xff, 0x1c, 0xa2, 0xc8, 0x9a, 0xf5, 0xb4, 0x78, 0xd8, 0x53,
	0x82, 0xeb, 0x17, 0x10, 0x65, 0x8f, 0x0b, 0xa2, 0x78, 0xf1, 0x96, 0xc6, 0x31, 0x2d, 0xad, 0x4d,
	0xe8, 0x4c, 0x0e, 0xfd, 0xf6, 0xd2, 0x44, 0x41, 0x71, 0x8f, 0xef, 0x85, 0xf7, 0xcb, 0xfc, 0x60,
	0xc9, 0x48, 0xfc, 0x6e, 0xd8, 0x68, 0x1a, 0xac, 0x8e, 0xdc, 0x74, 0x1a, 0x6d, 0xed, 0x95, 0x8b,
	0x7d, 0x2b, 0x8a, 0x29, 0xcb, 0xa5, 0xf3, 0xde, 0x75, 0xa8, 0xc3, 0xe0, 0xd9, 0x57, 0xf1, 0x43,
	0xe3, 0x67, 0xe0, 0x51, 0xa5, 0x28, 0xaf, 0xff, 0xef, 0x79, 0x0c, 0x48, 0x09, 0xb0, 0xba, 0x80,
	0xa3, 0x1d, 0x12, 0x19, 0x87, 0x62, 0x0e, 0x87, 0x3b, 0x72, 0x8b, 0xd9, 0x7a, 0xbe, 0xba, 0x85,
	0xf3, 0xc7, 0x1a, 0x6d, 0xfb, 0xc2, 0x86, 0x23, 0x3d, 0x23, 0x05, 0xdf, 0x13, 0x0a, 0x09, 0x65,
	0x6e, 0x16, 0x07, 0xeb, 0xb3, 0xcd, 0x89, 0x9a, 0x60, 0x2a, 0x37, 0x9b, 0x1e, 0x8a, 0x57, 0x6e,
	0x3a, 0x71, 0x03, 0xc5, 0xd6, 0xd6, 0x5e, 0x9c, 0xe6, 0xd7, 0x6f, 0xb8, 0xfc, 0x1b, 0xc5, 0x3d,
	0x1c, 0xff, 0xfa, 0x40, 0x5c, 0xaa, 0x6c, 0xa0, 0x7e, 0x0c, 0xd4, 0x76, 0x32, 0x90, 0x57, 0x79,
	0xb5, 0xe7, 0x96, 0x87, 0xf2, 0x2d, 0xb9, 0x3f, 0xcd, 0xaa, 0x32, 0x8d, 0xee, 0xbe, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x03, 0x1d, 0x10, 0xc9, 0x29, 0x01, 0x00, 0x00,
}