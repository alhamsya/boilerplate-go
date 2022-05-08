package cronHandler

import (
	"context"
	"fmt"
	"time"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
	"github.com/alhamsya/boilerplate-go/middleware/cron"
	"github.com/alhamsya/boilerplate-go/transport/cron/routers"
	"github.com/robfig/cron/v3"
)

func (h *Handler) Register(ctx context.Context) (*cron.Cron, error) {
	cronParser := cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor))

	location, err := time.LoadLocation(constCommon.TimeLocalJakarta)
	if err != nil {
		return nil, customError.WrapFlag(err, "time", "LoadLocation")
	}

	cronJob := cron.New(cron.WithLocation(location), cronParser)

	schedulerList := cronRouters.New(&cronRouters.CronServer{
		Cfg:            h.Cfg,
		CronInteractor: h.Interactor,
	}).Register()

	//start list scheduler
	for name, val := range h.Cfg.Scheduler {
		if val.IsActive {
			for _, fn := range schedulerList[name] {
				errSch := h.addScheduler(ctx, &scheduler{
					cron:                cronJob,
					name:                name,
					isDelayStillRunning: val.IsDelayStillRunning,
					standardSpec:        val.Schedule,
					function:            fn,
				})
				if errSch != nil {
					customLog.ErrorF("[CRON] scheduler name %s: %v", name, errSch)
				}
			}
		} else {
			customLog.InfoF("[CRON] %s: is inactive", name)
		}
	}

	return cronJob, nil
}

func (h *Handler) addScheduler(ctx context.Context, scheduler *scheduler) error {
	//handle param scheduler is nil
	if scheduler == nil {
		return fmt.Errorf("job scheduler is nil")
	}

	customLog.InfoF("[CRON] %s: will be running at (%s) and delay still running is (%t)", scheduler.name, scheduler.standardSpec, scheduler.isDelayStillRunning)

	//parse string scheduler to the standard cron scheduler
	schedule, err := cron.ParseStandard(scheduler.standardSpec)
	if err != nil {
		return customError.Wrap(err, "ParseStandard")
	}

	//setup wrapper for each job scheduler
	wrapper := []cron.JobWrapper{
		Recover(ctx, scheduler.name),
	}

	//mapping is delay still running
	if scheduler.isDelayStillRunning {
		wrapper = append(wrapper, DelayIfStillRunning(ctx, scheduler.name))
	} else {
		wrapper = append(wrapper, SkipIfStillRunning(ctx, scheduler.name))
	}

	//add chain for job scheduler
	scheduler.cron.Schedule(
		schedule,
		cron.NewChain(wrapper...).Then(
			cronMiddleware.Interceptor(
				ctx,
				scheduler.name,
				scheduler.function,
			),
		),
	)

	return nil
}
