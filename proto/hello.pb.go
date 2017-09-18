// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

/*
Package grpc_play_proto is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	Error
	GreetRequest
	GreetResponse
*/
package grpc_play_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Error struct {
	Code uint64 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Error) GetCode() uint64 {
	if m != nil {
		return m.Code
	}
	return 0
}

// GreetRequest is a request to generate a key.
type GreetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GreetRequest) Reset()                    { *m = GreetRequest{} }
func (m *GreetRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()               {}
func (*GreetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GreetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// GreetResponse is a response to a GreetRequest.
type GreetResponse struct {
	Error   *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *GreetResponse) Reset()                    { *m = GreetResponse{} }
func (m *GreetResponse) String() string            { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()               {}
func (*GreetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GreetResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *GreetResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "grpc_play_proto.Error")
	proto.RegisterType((*GreetRequest)(nil), "grpc_play_proto.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "grpc_play_proto.GreetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hello service

type HelloClient interface {
	// Generate generates a new key and initial key version.
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := grpc.Invoke(ctx, "/grpc_play_proto.Hello/Greet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	// Generate generates a new key and initial key version.
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_play_proto.Hello/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_play_proto.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Hello_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4f, 0x2f, 0x2a, 0x48, 0x8e, 0x2f, 0xc8,
	0x49, 0xac, 0x8c, 0x07, 0x0b, 0x28, 0x49, 0x73, 0xb1, 0xba, 0x16, 0x15, 0xe5, 0x17, 0x09, 0x09,
	0x71, 0xb1, 0x24, 0xe7, 0xa7, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0x81, 0xd9, 0x4a,
	0x4a, 0x5c, 0x3c, 0xee, 0x45, 0xa9, 0xa9, 0x25, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x20,
	0x35, 0x79, 0x89, 0xb9, 0x10, 0x35, 0x9c, 0x41, 0x60, 0xb6, 0x52, 0x38, 0x17, 0x2f, 0x54, 0x4d,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x0e, 0x17, 0x6b, 0x2a, 0xc8, 0x44, 0xb0, 0x2a, 0x6e,
	0x23, 0x31, 0x3d, 0x34, 0x2b, 0xf5, 0xc0, 0xf6, 0x05, 0x41, 0x14, 0x09, 0x49, 0x70, 0xb1, 0xe7,
	0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x81, 0x4d, 0x85, 0x71, 0x8d, 0xfc, 0xb9, 0x58,
	0x3d, 0x40, 0x2e, 0x17, 0x72, 0xe3, 0x62, 0x05, 0xdb, 0x20, 0x24, 0x8b, 0x61, 0x14, 0xb2, 0xeb,
	0xa4, 0xe4, 0x70, 0x49, 0x43, 0x1c, 0x96, 0xc4, 0x06, 0x16, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xad, 0x73, 0x35, 0x51, 0x11, 0x01, 0x00, 0x00,
}