package schedulerHandler

func New(this *Handler) *Handler {
	return &Handler{
		Cfg:        this.Cfg,
		Interactor: this.Interactor,
	}
}
