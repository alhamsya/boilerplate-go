package consumerRouters

import (
	"context"

	"cloud.google.com/go/pubsub"
)

func New(this *ConsumerServer) *ConsumerServer {
	return &ConsumerServer{
		Cfg:                this.Cfg,
		ConsumerInteractor: this.ConsumerInteractor,
	}
}

func (consume *ConsumerServer) Register() map[string][]func(context.Context, *pubsub.Message) (interface{}, error) {
	return map[string][]func(context.Context, *pubsub.Message) (interface{}, error){
		"tes": {
			consume.Payment,
		},
	}
}
