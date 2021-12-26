package grace

import (
	"context"
	"errors"
	"log"
	"net"
	"time"
)

func ServeGRPC(server GRPCServer, address string, graceTimeout time.Duration) error {
	lis, err :=  net.Listen(ProtocolTCP4, address)
	if err != nil {
		return err
	}

	stoppedCh := WaitTermSig(func(ctx context.Context) error {
		if graceTimeout == 0 {
			graceTimeout = 10 * time.Second
		}
		stopped := make(chan struct{})
		go func() {
			server.GracefulStop()
			close(stopped)
		}()

		select {
		case <-time.After(graceTimeout):
			server.Stop()
			return errors.New("[GRPC] server shutdown timed out")
		case <-stopped:
		}
		return nil
	})

	log.Printf("[GRPC] server running on address %v", address)

	if err := server.Serve(lis); err != nil {
		// Error starting or closing listener:
		return err
	}

	<-stoppedCh
	log.Println("[GRPC] server stopped")
	return nil
}