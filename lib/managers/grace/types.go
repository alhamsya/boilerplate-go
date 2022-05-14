package grace

import (
	"context"
	"net"
)

const ProtocolTCP4 = "tcp4"

// HTTPServer represents an HTTP server
type HTTPServer interface {
	Shutdown(ctx context.Context) error
	Serve(l net.Listener) error
}

// GRPCServer represents interface for grpc server
type GRPCServer interface {
	GracefulStop()
	Stop()
	Serve(l net.Listener) error
}
