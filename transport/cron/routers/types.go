package cronRouters

import (
	"github.com/alhamsya/boilerplate-go/domain/repositorys"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
)

type CronServer struct {
	Cfg            *config.ServiceConfig
	CronInteractor *CronInteractor
}

type CronInteractor struct {
	CronInterface repositorys.CronUsecase
}
