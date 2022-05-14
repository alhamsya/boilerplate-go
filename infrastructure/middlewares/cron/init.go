package cronMiddleware

import (
	"context"
	"time"

	"github.com/alhamsya/boilerplate-go/lib/managers/custom_log"
	"github.com/robfig/cron/v3"

	constCommon "github.com/alhamsya/boilerplate-go/domain/constants"
)

func Interceptor(ctx context.Context, name string, fn FuncOrigin) cron.Job {
	ctx = context.WithValue(ctx, constCommon.ContextKeyCronName, name)
	return cron.FuncJob(func() {
		now := time.Now()
		resp, errFn := fn(ctx)
		if errFn != nil {
			customLog.ErrorF("[ROUTERS] scheduler name %s: %v", name, errFn)
			return
		}

		customLog.InfoF("[RUN CRON] %s: success | elapsed time: %.2f secs | response func: %+v", name, time.Since(now).Seconds(), resp)
	})
}
