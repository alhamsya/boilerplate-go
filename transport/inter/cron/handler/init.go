package cronHandler

import "context"

func New(this *Handler) *Handler {
	return &Handler{
		Cfg:        this.Cfg,
		Interactor: this.Interactor,
	}
}

func (h *Handler) Run(ctx context.Context) error {
	return h.Register(ctx)
}
