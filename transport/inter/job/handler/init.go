package jobHandler

import (
	"github.com/alhamsya/boilerplate-go/transport/inter/job/handler/scheduler"
	"golang.org/x/sync/errgroup"
)

func New(this *Handler) *Handler {
	return &Handler{
		Cfg:        this.Cfg,
		Interactor: this.Interactor,
	}
}

func (h *Handler) Run() error {
	var eg errgroup.Group

	eg.Go(func() error {
		return initScheduler(&schedulerHandler.Handler{
			Cfg:        h.Cfg,
			Interactor: h.Interactor,
		})
	})

	return eg.Wait()
}
