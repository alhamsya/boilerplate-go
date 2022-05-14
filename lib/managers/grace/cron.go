package grace

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

//ServeCron start the cron server on the given app and add graceful shutdown handler
func ServeCron(app *cron.Cron, graceTimeout time.Duration) error {
	stoppedCh := WaitTermSig(func(ctx context.Context) error {
		if graceTimeout == 0 {
			graceTimeout = 10 * time.Second
		}
		stopped := make(chan struct{})
		ctx, cancel := context.WithTimeout(ctx, graceTimeout)
		defer cancel()
		go func() {
			app.Stop()
			close(stopped)
		}()

		select {
		case <-ctx.Done():
			return errors.New("[CRON] server shutdown timed out")
		case <-stopped:

		}

		return nil
	})

	log.Printf("[CRON] server running")

	// start serving
	app.Run()

	<-stoppedCh
	log.Println("[CRON] server stopped")

	return nil
}
