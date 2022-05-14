package restMiddleware

func New(this *Middleware) *Middleware {
	return &Middleware{
		Cfg:       this.Cfg,
		DBRepo:    this.DBRepo,
		UtilsRepo: this.UtilsRepo,
	}
}
