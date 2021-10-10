package restRouters

func (rest *RestServer) Register()  {
	apiGroup := rest.App.Group("/api")
	apiGroup.Get("/movie" , rest.GetAllMovie)
}