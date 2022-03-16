package restRouters

func (rest *RestServer) Register()  {
	apiGroup := rest.App.Group("/api")
	apiGroup.Get("/movie" , rest.GetListMovie)
	apiGroup.Get("/movie/:movieID" , rest.GetDetailMovie)
}