package auth

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

func AuthHeaderMatcher(key string) (string, bool) {
	switch key {
	case "X-Authentication":
		return "authentication", true
	default:
		return key, false
	}
}

func UnaryAuthentication(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Print(info.FullMethod)
	if md, ok := metadata.FromIncomingContext(ctx); !ok {
		return nil, status.Error(codes.PermissionDenied, "auth header not exist")
	} else {
		arr := md.Get("authentication")
		if len(arr) == 0 {
			return nil, status.Error(codes.PermissionDenied, "auth header empty")
		} else if arr[0] != "auth-test" {
			return nil, status.Error(codes.PermissionDenied, "auth failed")
		}
	}
	return handler(ctx, req)
}

func StreamAuthentication(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Print(info.FullMethod)
	return handler(srv, ss)
}
