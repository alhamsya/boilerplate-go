package schedulerHandler

import (
	"time"

	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
)

func (h *Handler) addScheduler(name, schedule string) {
	customLog.InfoF("[SCHEDULER] %s: will be running at %s", name, schedule)

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
		customLog.InfoF("[RUN SCHEDULER] %s: start", name)
		err := h.funcOrigins[name]()
		h.metricInterceptor(name, err)
		if err != nil {
			customLog.ErrorF("[USECASE] scheduler name %s: %v", name, err)
			return
		}

		customLog.InfoF("[RUN SCHEDULER] %s: success | elapsed time: %f secs", name, time.Since(now).Seconds())
	}
}

func (h *Handler) metricInterceptor(name string, err error) {

}
