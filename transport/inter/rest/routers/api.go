package restRouters

func (rest *RestServer) Register() {
	apiGroup := rest.App.Group("/api")
	apiGroup.Post("/signing", rest.CreateSigning)
	apiGroup.Post("/signup", rest.CreateSignup)
	apiGroup.Post("/refresh", rest.CreateRefreshToken)
	apiGroup.Get("/movie", rest.RestInteractor.Middleware.Authorize(rest.GetListMovie))
	apiGroup.Get("/movie/:movieID", rest.RestInteractor.Middleware.Authorize(rest.GetDetailMovie))
}
