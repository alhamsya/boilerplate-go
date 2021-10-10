package grace

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
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
			return errors.New("server shutdown timed out")
		case <-stopped:
		}
		return nil
	})

	log.Printf("GRPC server running on adress %v", address)

	if err := server.Serve(lis); err != nil {
		// Error starting or closing listener:
		return err
	}

	<-stoppedCh
	log.Println("GRPC server stopped")
	return nil
}

func WaitTermSig(handler func(context.Context) error) <-chan struct{} {
	stoppedCh := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)

		// wait for the sigterm
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		// We received an os signal, shut down.
		if err := handler(context.Background()); err != nil {
			log.Printf("graceful shutdown  failed: %v", err)
		} else {
			log.Println("gracefull shutdown succeed")
		}

		close(stoppedCh)

	}()
	return stoppedCh
}