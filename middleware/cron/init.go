package cronMiddleware

import (
	"context"
	"time"

	customLog "github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
)

func Interceptor(ctx context.Context, funcOrigins map[string]FuncOrigin, name string) FuncScheduler {
	// prevent panic if the function does not exist
	if funcOrigins[name] == nil {
		return nil
	}

	return func() {
		now := time.Now()

		ctx = context.WithValue(ctx, "cron_name", name)

		customLog.InfoF("[RUN CRON] %s: start", name)
		resp, err := funcOrigins[name](ctx)
		if err != nil {
			customLog.ErrorF("[ROUTERS] scheduler name %s: %v", name, err)
			return
		}

		customLog.InfoF("[RUN CRON] %s: success | elapsed time: %f secs | response func: %+v", name, time.Since(now).Seconds(), resp)
	}
}
