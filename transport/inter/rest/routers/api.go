package restRouters

func (rest *RestServer) Register() {
	apiGroup := rest.App.Group("/api")
	apiGroup.Post("/user", rest.CreateUser)
	apiGroup.Get("/movie", rest.RestInteractor.Middleware.Authorize(rest.GetListMovie))
	apiGroup.Get("/movie/:movieID", rest.RestInteractor.Middleware.Authorize(rest.GetDetailMovie))
}
