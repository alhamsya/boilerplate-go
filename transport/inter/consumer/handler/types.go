package consumerHandler

import (
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
	"github.com/alhamsya/boilerplate-go/transport/inter/consumer/routers"
)

type Handler struct {
	Cfg        *config.ServiceConfig
	Interactor *consumerRouters.ConsumerInteractor
}
