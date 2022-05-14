package cronHandler

import (
	"github.com/alhamsya/boilerplate-go/infrastructure/middlewares/cron"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
	"github.com/alhamsya/boilerplate-go/transport/cron/routers"
	"github.com/robfig/cron/v3"
)

type scheduler struct {
	cron                *cron.Cron
	name                string
	standardSpec        string
	isDelayStillRunning bool
	function            cronMiddleware.FuncOrigin
}

type Handler struct {
	Cfg        *config.ServiceConfig
	Interactor *cronRouters.CronInteractor
}
