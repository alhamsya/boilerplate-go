package consumerHandler

import (
	"cloud.google.com/go/pubsub"
	"context"
)

type HandlerPubSub interface {
	HandleMessage(ctx context.Context, msg *pubsub.Message) (interface{}, error)
}

type HandlerFunc func(ctx context.Context, msg *pubsub.Message) (interface{}, error)

func (h HandlerFunc) HandleMessage(ctx context.Context, msg *pubsub.Message) (interface{}, error) {
	return h(ctx, msg)
}
