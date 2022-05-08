package consumerHandler

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

func New(this *Handler) *Handler {
	client, err := pubsub.NewClient(context.Background(), this.Cfg.ConsumerServer.ProjectID)
	if err != nil {
		log.Fatalf("[INIT] [PUB/SUB] %v", err)
	}

	this.subscription = make(map[string]*pubsub.Subscription)

	for name, data := range this.Cfg.PubSub {
		this.subscription[name] = client.Subscription(data.Topic)
	}

	return &Handler{
		Cfg:          this.Cfg,
		Interactor:   this.Interactor,
		subscription: this.subscription,
	}
}
