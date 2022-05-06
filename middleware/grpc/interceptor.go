package grpcMiddleware

import (
	"context"

	"github.com/alhamsya/boilerplate-go/lib/helpers/client"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
	"google.golang.org/grpc"
)

//GrpcLoggingInterceptor GRPC log for interceptor
func GrpcLoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	clientIP := client.GrpcGetIP(ctx)

	resp, errHandler := handler(ctx, req)
	if errHandler != nil {
		customLog.ErrorF("[GRPC] %s [CLIENT] %s : %v", info.FullMethod, clientIP, errHandler)
	}

	customLog.InfoF("[GRPC] %s [CLIENT] %s : [REQUEST] %s", info.FullMethod, clientIP, req)
	return resp, nil
}
