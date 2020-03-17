package forwards

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/grpclog"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
)

var errEmptyResponse = errors.New("empty response")

func MixedReplaceForwardResponseStream(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, recv func() (proto.Message, error), opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	f, ok := w.(http.Flusher)
	if !ok {
		grpclog.Infof("Flush not supported in %T", w)
		http.Error(w, "unexpected type of web server", http.StatusInternalServerError)
		return
	}

	m := multipart.NewWriter(w)

	w.Header().Set("X-Accel-Buffering", "no")
	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary="+m.Boundary())
	if err := handleForwardResponseOptions(ctx, w, nil, opts); err != nil {
		runtime.MuxOrGlobalHTTPError(ctx, mux, marshaler, w, req, err)
		return
	}

	var wroteHeader bool
	h := textproto.MIMEHeader{}
	for {
		resp, err := recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			handleForwardResponseStreamError(ctx, wroteHeader, marshaler, w, req, mux, err)
			return
		}
		if err := handleForwardResponseOptions(ctx, w, resp, opts); err != nil {
			handleForwardResponseStreamError(ctx, wroteHeader, marshaler, w, req, mux, err)
			return
		}

		var buf []byte
		contentType := marshaler.ContentType()
		if body, ok := resp.(*httpbody.HttpBody); ok {
			buf, err = marshaler.Marshal(resp)
			if body.ContentType != "" {
				contentType = body.ContentType
			}
		} else {
			buf, err = marshaler.Marshal(streamChunk(ctx, resp, runtime.DefaultHTTPStreamErrorHandler))
		}
		if err != nil {
			grpclog.Infof("Failed to marshal response chunk: %v", err)
			handleForwardResponseStreamError(ctx, wroteHeader, marshaler, w, req, mux, err)
			return
		}
		h.Set("Content-Type", contentType)
		h.Set("Content-Length", fmt.Sprint(len(buf)))
		mw, err := m.CreatePart(h)
		if err != nil {
			grpclog.Infof("Failed to create part: %v", err)
			return
		}
		if _, err = mw.Write(buf); err != nil {
			grpclog.Infof("Failed to write part: %v", err)
			return
		}
		f.Flush()
	}
}

func handleForwardResponseOptions(ctx context.Context, w http.ResponseWriter, resp proto.Message, opts []func(context.Context, http.ResponseWriter, proto.Message) error) error {
	if len(opts) == 0 {
		return nil
	}
	for _, opt := range opts {
		if err := opt(ctx, w, resp); err != nil {
			grpclog.Infof("Error handling ForwardResponseOptions: %v", err)
			return err
		}
	}
	return nil
}

func handleForwardResponseStreamError(ctx context.Context, wroteHeader bool, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, mux *runtime.ServeMux, err error) {
	serr := streamError(ctx, runtime.DefaultHTTPStreamErrorHandler, err)
	if !wroteHeader {
		w.WriteHeader(int(serr.HttpCode))
	}
	buf, merr := marshaler.Marshal(errorChunk(serr))
	if merr != nil {
		grpclog.Infof("Failed to marshal an error: %v", merr)
		return
	}
	if _, werr := w.Write(buf); werr != nil {
		grpclog.Infof("Failed to notify error to client: %v", werr)
		return
	}
}

func streamError(ctx context.Context, errHandler runtime.StreamErrorHandlerFunc, err error) *runtime.StreamError {
	serr := errHandler(ctx, err)
	if serr != nil {
		return serr
	}
	// TODO: log about misbehaving stream error handler?
	return runtime.DefaultHTTPStreamErrorHandler(ctx, err)
}

func errorChunk(err *runtime.StreamError) map[string]proto.Message {
	return map[string]proto.Message{}
}

// streamChunk returns a chunk in a response stream for the given result. The
// given errHandler is used to render an error chunk if result is nil.
func streamChunk(ctx context.Context, result proto.Message, errHandler runtime.StreamErrorHandlerFunc) map[string]proto.Message {
	if result == nil {
		return errorChunk(streamError(ctx, errHandler, errEmptyResponse))
	}
	return map[string]proto.Message{"result": result}
}
