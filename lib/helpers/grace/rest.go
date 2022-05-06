package grace

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
	"github.com/gofiber/fiber/v2"
)

//ServeREST start the http server on the given address and add graceful shutdown handler
func ServeREST(srv HTTPServer, address string, graceTimeout time.Duration) error {
	lis, err := net.Listen(ProtocolTCP4, address)
	if err != nil {
		return customError.Wrap(err, "[REST] net listen")
	}

	stoppedCh := WaitTermSig(func(ctx context.Context) error {
		if graceTimeout == 0 {
			graceTimeout = 10 * time.Second
		}

		stopped := make(chan struct{})
		ctx, cancel := context.WithTimeout(ctx, graceTimeout)
		defer cancel()
		go func() {
			srv.Shutdown(ctx)
			close(stopped)
		}()

		select {
		case <-ctx.Done():
			return errors.New("[REST] server shutdown timed out")
		case <-stopped:

		}

		return nil
	})

	log.Printf("[REST] server running on address: %v", address)

	// start serving
	if err := srv.Serve(lis); err != http.ErrServerClosed {
		return err
	}

	<-stoppedCh
	log.Println("[REST] server stopped")
	return nil
}

//ServeRESTWithFiber start the http server using fiber on the given address and add graceful shutdown handler
func ServeRESTWithFiber(app *fiber.App, address string, graceTimeout time.Duration) error {
	stoppedCh := WaitTermSig(func(ctx context.Context) error {
		if graceTimeout == 0 {
			graceTimeout = 10 * time.Second
		}
		stopped := make(chan struct{})
		ctx, cancel := context.WithTimeout(ctx, graceTimeout)
		defer cancel()
		go func() {
			app.Shutdown()
			close(stopped)
		}()

		select {
		case <-ctx.Done():
			return errors.New("[REST] server shutdown timed out")
		case <-stopped:

		}

		return nil
	})

	log.Printf("[REST] server running on address: %v", address)

	// start serving
	if err := app.Listen(address); err != nil {
		return customError.Wrap(err, "net listen")
	}

	<-stoppedCh
	log.Println("[REST] server stopped")

	return nil
}
