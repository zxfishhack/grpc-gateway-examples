package services

import (
	"context"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"github.com/zxfishhack/grpc-gateway-examples/pkgs/assets"
	"github.com/zxfishhack/grpc-gateway-examples/pkgs/auth"
	"github.com/zxfishhack/grpc-gateway-examples/pkgs/proto"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func RunHttpService(ctx context.Context, httpAddr, grpcAddr string, cs1 proto.WithAuthenticationServer,
	cs2 proto.WithoutAuthenticationServer) (err error) {
	var addr *net.TCPAddr
	addr, err = net.ResolveTCPAddr("tcp", grpcAddr)
	if err != nil {
		return
	}

	var grpcLis net.Listener
	grpcLis, err = net.Listen("tcp", grpcAddr)
	if err != nil {
		return
	}
	defer grpcLis.Close()

	var httpLis net.Listener
	httpLis, err = net.Listen("tcp", httpAddr)
	if err != nil {
		return
	}
	defer httpLis.Close()
	grpcSvc := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(auth.UnaryAuthentication),
		grpc_middleware.WithStreamServerChain(auth.StreamAuthentication),
	)
	defer grpcSvc.Stop()

	proto.RegisterWithAuthenticationServer(grpcSvc, cs1)
	proto.RegisterWithoutAuthenticationServer(grpcSvc, cs2)
	reflection.Register(grpcSvc)

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "PUT"},
		AllowedHeaders:   []string{"*"},
	})

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{OrigName: true, EmitDefaults: true, EnumsAsInts: true},
		),
		runtime.WithIncomingHeaderMatcher(auth.AuthHeaderMatcher), // auth header matcher
	)

	runtime.SetHTTPBodyMarshaler(mux)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	if addr.IP.IsUnspecified() {
		addr.IP = net.IPv4(127, 0, 0, 1)
	}
	err = proto.RegisterWithAuthenticationHandlerFromEndpoint(ctx, mux, addr.String(), opts)
	if err != nil {
		return
	}

	err = proto.RegisterWithoutAuthenticationHandlerFromEndpoint(ctx, mux, addr.String(), opts)
	if err != nil {
		return
	}

	mux2 := http.NewServeMux()

	mux2.Handle("/api/", http.StripPrefix("/api", c.Handler(mux)))

	if strings.Contains(os.Args[0], "go-build") {
		log.Print("running with go run, use raw web content.")
		mux2.Handle("/", http.FileServer(http.Dir("pkg/assets/web")))
	} else {
		mux2.Handle("/", http.FileServer(&assetfs.AssetFS{
			Asset:     assets.Asset,
			AssetDir:  assets.AssetDir,
			AssetInfo: assets.AssetInfo,
			Prefix:    "/web",
		}))
	}

	httpSvc := &http.Server{
		Addr:    httpAddr,
		Handler: mux2,
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return grpcSvc.Serve(grpcLis)
	})

	g.Go(func() error {
		err := httpSvc.Serve(httpLis)
		if err == http.ErrServerClosed {
			err = nil
		}
		return err
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
		}
		_ = httpSvc.Close()
		grpcSvc.Stop()
		return nil
	})

	return g.Wait()
}
