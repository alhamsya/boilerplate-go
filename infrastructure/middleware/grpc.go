package middleware

import (
	"context"
	"log"

	"github.com/alhamsya/boilerplate-go/lib/helpers/client"
	"google.golang.org/grpc"
)

func GrpcLoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	clientIP := client.GrpcGetIP(ctx)

	resp, errHandler := handler(ctx, req)
	if errHandler != nil {
		log.Printf("[GRPC] %s [CLIENT] %s : %v", info.FullMethod, clientIP, errHandler)
	}
	log.Printf("[GRPC] %s [CLIENT] %s : [REQUEST] %s", info.FullMethod, clientIP, req)
	return resp, nil
}
