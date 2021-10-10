package grace

import "net"

const ProtocolTCP4 = "tcp4"

type GRPCServer interface {
	GracefulStop()
	Stop()
	Serve(l net.Listener) error
}
