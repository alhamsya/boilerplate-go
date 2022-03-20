package schedulerHandler

import (
	"context"
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
	"github.com/alhamsya/boilerplate-go/transport/inter/job/routers"
	"github.com/robfig/cron/v3"
)

type Handler struct {
	Cfg         *config.ServiceConfig
	Interactor  *jobRouters.JobInteractor
	cron        *cron.Cron
	funcOrigins map[string]funcOrigin
}

type funcScheduler func()
type funcOrigin func(context.Context) (interface{}, error)

type Logger interface {
	// Info logs routine messages about cron's operation.
	Info(msg string, keysAndValues ...interface{})
	// Error logs an error condition.
	Error(err error, msg string, keysAndValues ...interface{})
}
type printfLogger struct {
	logger  interface{ Printf(string, ...interface{}) }
	logInfo bool
}
