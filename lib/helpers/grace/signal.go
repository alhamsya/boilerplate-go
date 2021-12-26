package grace

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// WaitTermSig wait for termination signal and then execute the given handler
// The handler is usually service shutdown, so we can properly shutdown our server upon SIGTERM.
// It returns channel which will be closed after the signal received and the handler executed.
// We can use the signal to wait for the shutdown to be finished.
func WaitTermSig(handler func(context.Context) error) <-chan struct{} {
	stoppedCh := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)

		// wait for the sigterm
		signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-signals

		// We received an os signal, shut down.
		if err := handler(context.Background()); err != nil {
			log.Printf("[GRACEFUL SHUTDOWN] is failed: %v", err)
		} else {
			log.Println("[GRACEFUL SHUTDOWN] succeed")
		}

		close(stoppedCh)

	}()
	return stoppedCh
}
