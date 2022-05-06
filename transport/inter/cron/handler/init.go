package cronHandler

import (
	"context"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
	"github.com/alhamsya/boilerplate-go/lib/helpers/grace"
)

func New(this *Handler) *Handler {
	return &Handler{
		Cfg:        this.Cfg,
		Interactor: this.Interactor,
	}
}

func (h *Handler) Run(ctx context.Context) error {
	app, err := h.Register(ctx)
	if err != nil {
		return customError.WrapFlag(err, "handler", "Register")
	}

	return grace.ServeCron(app, constCommon.DefaultServerCRONGraceTimeout)
}
