package consumerHandler

import (
	"cloud.google.com/go/pubsub"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
	"github.com/alhamsya/boilerplate-go/transport/consumer/routers"
)

type Handler struct {
	Cfg          *config.ServiceConfig
	Interactor   *consumerRouters.ConsumerInteractor
	subscription map[string]*pubsub.Subscription
}
