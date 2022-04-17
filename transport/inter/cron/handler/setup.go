package cronHandler

import (
	"context"
	"sync"
	"time"

	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
	"github.com/alhamsya/boilerplate-go/middleware/cron"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
	"github.com/alhamsya/boilerplate-go/transport/inter/cron/routers"
	"github.com/robfig/cron/v3"
)

func (h *Handler) Register(ctx context.Context) error {
	h.funcOrigins = make(map[string]cronMiddleware.FuncOrigin)

	cronParser := cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor))

	//- Recover any panics from jobs (activated by default)
	//- Delay a job's execution if the previous run hasn't completed yet
	//- Skip a job's execution if the previous run hasn't completed yet
	//- Log each job's invocations
	cronWrappers := cron.WithChain(
		delayIfStillRunning(ctx),
	)

	location, err := time.LoadLocation(constCommon.TimeLocalJakarta)
	if err != nil {
		return customError.WrapFlag(err, "time", "LoadLocation")
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
				h.addScheduler(ctx, name, val.Schedule)
			}
		} else {
			customLog.InfoF("[CRON] %s: is inactive", name)
		}
	}

	h.cron.Run()

	return nil
}

func (h *Handler) addScheduler(ctx context.Context, name, schedule string) {
	customLog.InfoF("[CRON] %s: will be running at %s", name, schedule)

	_, err := h.cron.AddFunc(
		schedule, cronMiddleware.Interceptor(ctx, h.funcOrigins, name),
	)
	if err != nil {
		customLog.ErrorLn(err)
	}
}

// delayIfStillRunning serializes jobs, delaying subsequent runs until the
// previous one is complete. Jobs running after a delay of more than a minute
// have the delay logged at Info.
func delayIfStillRunning(ctx context.Context) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		var mu sync.Mutex
		return cron.FuncJob(func() {
			start := time.Now()
			mu.Lock()
			defer mu.Unlock()
			if dur := time.Since(start); dur > time.Minute {
				customLog.WarnF("[CRON] delaying subsequent runs until the previous one is complete | duration %v", dur)
			}
			j.Run()
		})
	}
}
