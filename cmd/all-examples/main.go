package main

import (
	"context"
	"flag"
	"github.com/zxfishhack/grpc-gateway-examples/pkgs/services"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	httpLis := ":8080"
	grpcLis := "127.0.0.1:8081"

	flag.StringVar(&httpLis, "http", ":8080", "http listen endpoint")
	flag.StringVar(&grpcLis, "grpc", "127.0.0.1:8081", "grpc listen endpoint")
	flag.Parsed()

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		log.Print("got exit signal.")
		cancel()
		return nil
	})

	cs := services.NewGrpcServices()
	g.Go(func() error {
		return services.RunHttpService(ctx, httpLis, grpcLis, cs, cs)
	})

	log.Print("misc go started.")
	err := g.Wait()
	log.Print("misc go exit.", err)
}
