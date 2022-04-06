package consumerRouters

import (
	"context"

	"cloud.google.com/go/pubsub"
)

func (consume *ConsumerServer) Payment(ctx context.Context, msg *pubsub.Message) (interface{}, error) {
	err := consume.ConsumerInteractor.ConsumerInterface.DoPayment(ctx)
	return nil, err
}
