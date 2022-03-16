package schedulerHandler

import (
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
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
type funcOrigin func() error

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
