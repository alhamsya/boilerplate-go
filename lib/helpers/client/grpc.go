package client

import (
	"context"

	"google.golang.org/grpc/peer"
)

func GrpcGetIP(ctx context.Context) (clientIp string) {
	peerRpc, ok := peer.FromContext(ctx)
	if ok {
		clientIp = peerRpc.Addr.String()
	} else {
		clientIp = "-"
	}

	return clientIp
}
