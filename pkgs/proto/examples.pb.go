// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/zxfishhack/grpc-gateway-examples/pkgs/proto/examples.proto

package proto

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

func init() {
	proto.RegisterFile("github.com/zxfishhack/grpc-gateway-examples/pkgs/proto/examples.proto", fileDescriptor_357adbe592a7cfed)
}

var fileDescriptor_357adbe592a7cfed = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0x47, 0x11, 0x41, 0x30, 0xc7, 0xc5, 0x0a, 0xae, 0x9f, 0xa1, 0x89, 0xb4, 0x20, 0x82, 0x27,
	0x0b, 0x15, 0x3d, 0x78, 0x2b, 0x08, 0x5e, 0x4a, 0x36, 0x3b, 0x4d, 0x86, 0x76, 0x3b, 0x43, 0x33,
	0xc1, 0x5d, 0x3f, 0xbd, 0xec, 0x9f, 0x5c, 0xbc, 0xd9, 0x53, 0xc8, 0xef, 0xf1, 0xde, 0x61, 0xd4,
	0xda, 0xa3, 0x84, 0x54, 0x69, 0x47, 0x8d, 0xf9, 0x69, 0x77, 0x18, 0x43, 0xb0, 0x6e, 0x6f, 0xfc,
	0x89, 0xdd, 0xdc, 0x5b, 0x81, 0x6f, 0xdb, 0xcd, 0xa1, 0xb5, 0x0d, 0x1f, 0x20, 0x1a, 0xde, 0xfb,
	0x68, 0xf8, 0x44, 0x42, 0x26, 0x6f, 0x7a, 0xf8, 0x16, 0xb3, 0x5e, 0xd8, 0x4e, 0xc2, 0x36, 0xc3,
	0xf2, 0xde, 0x13, 0xf9, 0x03, 0x8c, 0x4e, 0x95, 0x76, 0x06, 0x1a, 0x96, 0x6e, 0x74, 0xca, 0xbb,
	0x09, 0x5a, 0x46, 0x13, 0x44, 0xb8, 0xa2, 0x7a, 0x42, 0x8b, 0x77, 0x55, 0x7c, 0xa2, 0x84, 0x97,
	0x24, 0x01, 0x8e, 0x82, 0xce, 0x0a, 0xd2, 0xb1, 0x58, 0xaa, 0xcb, 0x57, 0xa2, 0xe2, 0x56, 0x8f,
	0xa2, 0xce, 0x55, 0xbd, 0xee, 0xab, 0xe5, 0x4d, 0xde, 0x2d, 0xa3, 0x7e, 0x13, 0xe1, 0x15, 0xd5,
	0xdd, 0x62, 0xa3, 0x66, 0x7d, 0x8a, 0x92, 0xfc, 0xa9, 0x3d, 0xab, 0xeb, 0x0f, 0x6c, 0xa1, 0xde,
	0x40, 0x94, 0xff, 0x35, 0x1f, 0x2e, 0x56, 0x4f, 0x5f, 0x8f, 0xe7, 0x1d, 0xae, 0xba, 0x1a, 0x9e,
	0xe5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x4f, 0x42, 0xe6, 0x79, 0x01, 0x00, 0x00,
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
	Metadata: "github.com/zxfishhack/grpc-gateway-examples/pkgs/proto/examples.proto",
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
	Metadata: "github.com/zxfishhack/grpc-gateway-examples/pkgs/proto/examples.proto",
}
