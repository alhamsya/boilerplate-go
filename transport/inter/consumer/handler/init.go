package consumerHandler

func New(this *Handler) *Handler {
	return &Handler{
		Cfg:        this.Cfg,
		Interactor: this.Interactor,
	}
}

func (h *Handler) Consume() {
}
