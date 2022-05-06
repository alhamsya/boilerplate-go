package cronRouters

import (
	"github.com/alhamsya/boilerplate-go/domain/definition"
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
)

type CronServer struct {
	Cfg            *config.ServiceConfig
	CronInteractor *CronInteractor
}

type CronInteractor struct {
	CronInterface definition.CronUsecase
}
