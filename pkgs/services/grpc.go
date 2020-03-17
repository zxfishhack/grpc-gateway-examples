package services

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/zxfishhack/grpc-gateway-examples/pkgs/assets"
	"github.com/zxfishhack/grpc-gateway-examples/pkgs/proto"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/metadata"
	"time"
)

type GrpcServices struct {
	datas [][]byte
}

func NewGrpcServices() (svc *GrpcServices) {
	svc = &GrpcServices{}
	svc.datas = append(svc.datas, assets.MustAsset("web/streams/r.jpg"))
	svc.datas = append(svc.datas, assets.MustAsset("web/streams/g.jpg"))
	svc.datas = append(svc.datas, assets.MustAsset("web/streams/b.jpg"))
	return
}

func (svc *GrpcServices) Foo(context.Context, *empty.Empty) (*httpbody.HttpBody, error) {
	return &httpbody.HttpBody{
		ContentType: "text/plain",
		Data:        []byte("you got it\r\n"),
	}, nil
}

func (svc *GrpcServices) MixedTest(_ *empty.Empty, resp proto.WithoutAuthentication_MixedTestServer) error {
	ctx := resp.Context()
	driver := time.NewTimer(0)
	d := 500 * time.Millisecond

	streamMd := metadata.MD{}
	streamMd.Set("Content-type", "text/plain")
	i := 0
	for {
		select {
		case <-ctx.Done():
			break
		case <-driver.C:
			resp.Send(&httpbody.HttpBody{
				ContentType: "image/jpeg",
				Data:        svc.datas[i%len(svc.datas)],
			})
			i = i + 1
			driver.Reset(d)
		}
	}
}
