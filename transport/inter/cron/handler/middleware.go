package cronHandler

import (
	"context"
	"time"

	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
)

func (h *Handler) addScheduler(name, schedule string) {
	customLog.InfoF("[CRON] %s: will be running at %s", name, schedule)

	_, err := h.cron.AddFunc(
		schedule, h.middleware(name),
	)
	if err != nil {
		customLog.ErrorLn(err)
	}
}

func (h *Handler) middleware(name string) funcScheduler {
	// prevent panic if the function does not exist
	if h.funcOrigins[name] == nil {
		return nil
	}

	return func() {
		now := time.Now()

		ctx := context.WithValue(context.Background(), "cron_name", name)

		customLog.InfoF("[RUN CRON] %s: start", name)
		resp, err := h.funcOrigins[name](ctx)
		h.metricInterceptor(name, err)
		if err != nil {
			customLog.ErrorF("[ROUTERS] scheduler name %s: %v", name, err)
			return
		}

		customLog.InfoF("[RUN CRON] %s: success | elapsed time: %f secs | response func: %+v", name, time.Since(now).Seconds(), resp)
	}
}

func (h *Handler) metricInterceptor(name string, err error) {

}
