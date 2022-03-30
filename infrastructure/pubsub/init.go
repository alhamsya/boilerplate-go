package pubsub

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

func New(this *ServicePubSub) *ServicePubSub {
	client, err := pubsub.NewClient(context.Background(), this.Cfg.ConsumerServer.ProjectID)
	if err != nil {
		log.Fatalf("[INIT] [PUB/SUB] %v", err)
	}

	this.subscription = make(map[string]*pubsub.Subscription)
	for name, data := range this.Cfg.PubSub {
		this.subscription[name] = client.Subscription(data.Topic)
	}

	result := &ServicePubSub{
		Cfg:          this.Cfg,
		UtilsRepo:    this.UtilsRepo,
		subscription: this.subscription,
	}
	log.Printf("[IGNORE] [PUB/SUB] initialize")
	return result
}
