package consumerHandler

import (
	"context"

	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
	"github.com/alhamsya/boilerplate-go/middleware/consumer"
	"github.com/alhamsya/boilerplate-go/transport/consumer/routers"
)

func (h *Handler) Run(ctx context.Context) error {

	consumerList := consumerRouters.New(&consumerRouters.ConsumerServer{
		Cfg:                h.Cfg,
		ConsumerInteractor: h.Interactor,
	}).Register()

	for name, val := range h.Cfg.PubSub {
		if val.IsActive {
			for _, fun := range consumerList[name] {
				subscription, ok := h.subscription[name]
				if !ok {
					continue
				}
				consumerMiddleware.InterceptorPubSub(ctx, subscription, fun)
			}
		} else {
			customLog.InfoF("[CONSUMER] %s: is inactive", name)
		}
	}

	return nil
}
