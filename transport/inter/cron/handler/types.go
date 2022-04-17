package cronHandler

import (
	cronMiddleware "github.com/alhamsya/boilerplate-go/middleware/cron"

	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
	"github.com/alhamsya/boilerplate-go/transport/inter/cron/routers"
	"github.com/robfig/cron/v3"
)

type Handler struct {
	Cfg         *config.ServiceConfig
	Interactor  *cronRouters.CronInteractor
	cron        *cron.Cron
	funcOrigins map[string]cronMiddleware.FuncOrigin
}
