package routers

func New(this *RestServer) *RestServer {
	return &RestServer{
		Cfg:            this.Cfg,
		App:            this.App,
		RestInteractor: this.RestInteractor,
	}
}
