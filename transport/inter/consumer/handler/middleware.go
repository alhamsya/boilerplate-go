package consumerHandler

import (
	"context"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
)

func (h *Handler) Consume(sub *pubsub.Subscription, fun func(ctx context.Context, msg *pubsub.Message) (interface{}, error)) {
	err := sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		now := time.Now()
		resp, err := fun(ctx, msg)
		if err != nil {
			msg.Nack()
			customLog.ErrorF("[CONSUMER] %s: %v", sub.ID(), err)
			return
		}

		msg.Ack()
		customLog.InfoF("[CONSUMER] %s: success | elapsed time: %f secs | response func: %+v", sub.ID(), time.Since(now).Seconds(), resp)

	})

	if err != nil {
		customLog.ErrorF("[CONSUMER] %s: %v", sub.ID(), err)
	}
}
