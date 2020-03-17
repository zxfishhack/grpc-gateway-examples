// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkgs/proto/examples.proto

package grpc_gateway_examples

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("pkgs/proto/examples.proto", fileDescriptor_e550c6e64b952f0a) }

var fileDescriptor_e550c6e64b952f0a = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd0, 0x3f, 0x4b, 0x03, 0x41,
	0x10, 0x05, 0x70, 0x44, 0x10, 0xdc, 0xf2, 0x30, 0x42, 0xce, 0x8f, 0x20, 0xb7, 0x2b, 0x49, 0x69,
	0x65, 0x40, 0xd1, 0xc2, 0x2e, 0x20, 0xd8, 0x84, 0xbd, 0xbd, 0xc9, 0xee, 0x90, 0x5c, 0x66, 0xc8,
	0xce, 0xe2, 0x9d, 0x9f, 0x5e, 0xee, 0xcf, 0x36, 0x76, 0x29, 0x67, 0x1e, 0xef, 0x57, 0x3c, 0xb5,
	0xe4, 0x83, 0x8f, 0x86, 0xcf, 0x24, 0x64, 0xa0, 0xb3, 0x2d, 0x1f, 0x21, 0xea, 0xf1, 0x2c, 0x16,
	0xfe, 0xcc, 0x6e, 0xe7, 0xad, 0xc0, 0x8f, 0xed, 0x77, 0x39, 0x2c, 0x1f, 0x3c, 0x91, 0x3f, 0xc2,
	0xd4, 0xa9, 0xd3, 0xde, 0x40, 0xcb, 0xd2, 0x4f, 0x9d, 0x72, 0x39, 0x87, 0x96, 0xd1, 0x04, 0x11,
	0xae, 0xa9, 0x99, 0xa3, 0xd5, 0x87, 0x2a, 0xbe, 0x50, 0xc2, 0x4b, 0x92, 0x00, 0x27, 0x41, 0x67,
	0x05, 0xe9, 0x54, 0xac, 0xd5, 0xf5, 0x1b, 0x51, 0x71, 0xaf, 0xa7, 0xa2, 0xce, 0xaa, 0x7e, 0x1d,
	0xd4, 0xf2, 0x2e, 0xff, 0x2d, 0xa3, 0x7e, 0x17, 0xe1, 0x0d, 0x35, 0xfd, 0x6a, 0xab, 0x16, 0x03,
	0x45, 0x49, 0xfe, 0x69, 0xcf, 0xea, 0xf6, 0x13, 0x3b, 0x68, 0xb6, 0x10, 0xe5, 0x32, 0xf3, 0xe9,
	0x6a, 0x53, 0x7d, 0x3f, 0x7a, 0x94, 0x90, 0x6a, 0xed, 0xa8, 0x35, 0xbf, 0xdd, 0x1e, 0x63, 0x08,
	0xd6, 0x1d, 0xcc, 0xb0, 0x43, 0x35, 0xef, 0x50, 0xe5, 0x1d, 0xea, 0x9b, 0x91, 0x5d, 0xff, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x14, 0x8a, 0x50, 0xdb, 0x42, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WithAuthenticationClient is the client API for WithAuthentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WithAuthenticationClient interface {
	Foo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
}

type withAuthenticationClient struct {
	cc *grpc.ClientConn
}

func NewWithAuthenticationClient(cc *grpc.ClientConn) WithAuthenticationClient {
	return &withAuthenticationClient{cc}
}

func (c *withAuthenticationClient) Foo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/grpc_gateway_examples.WithAuthentication/Foo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WithAuthenticationServer is the server API for WithAuthentication service.
type WithAuthenticationServer interface {
	Foo(context.Context, *empty.Empty) (*httpbody.HttpBody, error)
}

func RegisterWithAuthenticationServer(s *grpc.Server, srv WithAuthenticationServer) {
	s.RegisterService(&_WithAuthentication_serviceDesc, srv)
}

func _WithAuthentication_Foo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithAuthenticationServer).Foo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_gateway_examples.WithAuthentication/Foo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithAuthenticationServer).Foo(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _WithAuthentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_gateway_examples.WithAuthentication",
	HandlerType: (*WithAuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Foo",
			Handler:    _WithAuthentication_Foo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkgs/proto/examples.proto",
}

// WithoutAuthenticationClient is the client API for WithoutAuthentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WithoutAuthenticationClient interface {
	MixedTest(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (WithoutAuthentication_MixedTestClient, error)
}

type withoutAuthenticationClient struct {
	cc *grpc.ClientConn
}

func NewWithoutAuthenticationClient(cc *grpc.ClientConn) WithoutAuthenticationClient {
	return &withoutAuthenticationClient{cc}
}

func (c *withoutAuthenticationClient) MixedTest(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (WithoutAuthentication_MixedTestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WithoutAuthentication_serviceDesc.Streams[0], "/grpc_gateway_examples.WithoutAuthentication/MixedTest", opts...)
	if err != nil {
		return nil, err
	}
	x := &withoutAuthenticationMixedTestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WithoutAuthentication_MixedTestClient interface {
	Recv() (*httpbody.HttpBody, error)
	grpc.ClientStream
}

type withoutAuthenticationMixedTestClient struct {
	grpc.ClientStream
}

func (x *withoutAuthenticationMixedTestClient) Recv() (*httpbody.HttpBody, error) {
	m := new(httpbody.HttpBody)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WithoutAuthenticationServer is the server API for WithoutAuthentication service.
type WithoutAuthenticationServer interface {
	MixedTest(*empty.Empty, WithoutAuthentication_MixedTestServer) error
}

func RegisterWithoutAuthenticationServer(s *grpc.Server, srv WithoutAuthenticationServer) {
	s.RegisterService(&_WithoutAuthentication_serviceDesc, srv)
}

func _WithoutAuthentication_MixedTest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WithoutAuthenticationServer).MixedTest(m, &withoutAuthenticationMixedTestServer{stream})
}

type WithoutAuthentication_MixedTestServer interface {
	Send(*httpbody.HttpBody) error
	grpc.ServerStream
}

type withoutAuthenticationMixedTestServer struct {
	grpc.ServerStream
}

func (x *withoutAuthenticationMixedTestServer) Send(m *httpbody.HttpBody) error {
	return x.ServerStream.SendMsg(m)
}

var _WithoutAuthentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_gateway_examples.WithoutAuthentication",
	HandlerType: (*WithoutAuthenticationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MixedTest",
			Handler:       _WithoutAuthentication_MixedTest_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkgs/proto/examples.proto",
}