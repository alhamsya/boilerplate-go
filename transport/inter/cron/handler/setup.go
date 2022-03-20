package cronHandler

import (
	"time"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
	"github.com/alhamsya/boilerplate-go/transport/inter/cron/routers"
	"github.com/robfig/cron/v3"
)

func (h *Handler) Register() error {
	h.funcOrigins = make(map[string]funcOrigin)

	cronParser := cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor))

	//- Recover any panics from jobs (activated by default)
	//- Delay a job's execution if the previous run hasn't completed yet
	//- Skip a job's execution if the previous run hasn't completed yet
	//- Log each job's invocations
	cronWrappers := cron.WithChain(
		cron.SkipIfStillRunning(
			VerbosePrintfLogger(),
		),
	)

	location, err := time.LoadLocation(constCommon.TimeLocalJakarta)
	if err != nil {
		return err
	}

	h.cron = cron.New(cron.WithLocation(location), cronParser, cronWrappers)

	schedulerList := cronRouters.New(&cronRouters.CronServer{
		Cfg:            h.Cfg,
		CronInteractor: h.Interactor,
	}).Register()

	//start scheduler
	for name, val := range h.Cfg.Scheduler {
		if val.IsActive {
			for _, fn := range schedulerList[name] {
				h.funcOrigins[name] = fn
				h.addScheduler(name, val.Schedule)
			}
		} else {
			customLog.InfoF("[CRON] %s: is inactive", name)
		}
	}

	h.cron.Run()

	return nil
}

// VerbosePrintfLogger wraps a Printf-based logger (such as the standard library
// "log") into an implementation of the Logger interface which logs everything.
func VerbosePrintfLogger() cron.Logger {
	return printfLogger{
		logInfo: true,
	}
}

func (pl printfLogger) Info(msg string, keysAndValues ...interface{}) {
	if pl.logInfo {
		customLog.WarnF("[CRON] %s", msg)
	}
}

func (pl printfLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	customLog.ErrorF("[CRON] %s: %v", msg, err)
}
