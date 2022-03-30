package pubsub

import (
	"cloud.google.com/go/pubsub"
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
)

type ServicePubSub struct {
	Cfg          *config.ServiceConfig
	UtilsRepo    repository.UtilsRepo
	subscription map[string]*pubsub.Subscription
}
