// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: github.com/zxfishhack/grpc-gateway-examples/pkgs/proto/examples.proto

/*
Package proto is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package proto

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_WithAuthentication_Foo_0(ctx context.Context, marshaler runtime.Marshaler, client WithAuthenticationClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq empty.Empty
	var metadata runtime.ServerMetadata

	msg, err := client.Foo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_WithoutAuthentication_MixedTest_0(ctx context.Context, marshaler runtime.Marshaler, client WithoutAuthenticationClient, req *http.Request, pathParams map[string]string) (WithoutAuthentication_MixedTestClient, runtime.ServerMetadata, error) {
	var protoReq empty.Empty
	var metadata runtime.ServerMetadata

	stream, err := client.MixedTest(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterWithAuthenticationHandlerFromEndpoint is same as RegisterWithAuthenticationHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterWithAuthenticationHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterWithAuthenticationHandler(ctx, mux, conn)
}

// RegisterWithAuthenticationHandler registers the http handlers for service WithAuthentication to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterWithAuthenticationHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterWithAuthenticationHandlerClient(ctx, mux, NewWithAuthenticationClient(conn))
}

// RegisterWithAuthenticationHandlerClient registers the http handlers for service WithAuthentication
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "WithAuthenticationClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "WithAuthenticationClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "WithAuthenticationClient" to call the correct interceptors.
func RegisterWithAuthenticationHandlerClient(ctx context.Context, mux *runtime.ServeMux, client WithAuthenticationClient) error {

	mux.Handle("GET", pattern_WithAuthentication_Foo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_WithAuthentication_Foo_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WithAuthentication_Foo_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_WithAuthentication_Foo_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"foo"}, ""))
)

var (
	forward_WithAuthentication_Foo_0 = runtime.ForwardResponseMessage
)

// RegisterWithoutAuthenticationHandlerFromEndpoint is same as RegisterWithoutAuthenticationHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterWithoutAuthenticationHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterWithoutAuthenticationHandler(ctx, mux, conn)
}

// RegisterWithoutAuthenticationHandler registers the http handlers for service WithoutAuthentication to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterWithoutAuthenticationHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterWithoutAuthenticationHandlerClient(ctx, mux, NewWithoutAuthenticationClient(conn))
}

// RegisterWithoutAuthenticationHandlerClient registers the http handlers for service WithoutAuthentication
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "WithoutAuthenticationClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "WithoutAuthenticationClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "WithoutAuthenticationClient" to call the correct interceptors.
func RegisterWithoutAuthenticationHandlerClient(ctx context.Context, mux *runtime.ServeMux, client WithoutAuthenticationClient) error {

	mux.Handle("GET", pattern_WithoutAuthentication_MixedTest_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_WithoutAuthentication_MixedTest_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_WithoutAuthentication_MixedTest_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_WithoutAuthentication_MixedTest_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"mixed"}, ""))
)

var (
	forward_WithoutAuthentication_MixedTest_0 = runtime.ForwardResponseStream
)