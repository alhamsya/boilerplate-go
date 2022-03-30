package consumerHandler

import (
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
	"github.com/alhamsya/boilerplate-go/transport/inter/consumer/routers"
)

func (h *Handler) Run() error {

	consumerList := consumerRouters.New(&consumerRouters.ConsumerServer{
		Cfg:                h.Cfg,
		ConsumerInteractor: h.Interactor,
	}).Register()

	for name, val := range h.Cfg.PubSub {
		if val.IsActive {
			for _, topic := range consumerList {
				h.Consume()
			}
		} else {
			customLog.InfoF("[CRON] %s: is inactive", name)
		}
	}

	return nil
}
