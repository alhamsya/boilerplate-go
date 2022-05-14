package consumerRouters

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
)

type ConsumerServer struct {
	Cfg                *config.ServiceConfig
	ConsumerInteractor *ConsumerInteractor
}

type ConsumerInteractor struct {
	ConsumerInterface repository.ConsumerUsecase
}
