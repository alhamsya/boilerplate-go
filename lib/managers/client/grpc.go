package client

import (
	"context"

	"google.golang.org/grpc/peer"
)

//GrpcGetIP get IP for GRPC
func GrpcGetIP(ctx context.Context) (clientIp string) {
	peerRpc, ok := peer.FromContext(ctx)
	if ok {
		clientIp = peerRpc.Addr.String()
	} else {
		clientIp = "-"
	}

	return clientIp
}
