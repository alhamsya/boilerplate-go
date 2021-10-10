package routers

func (rest *RestServer) Register()  {
	apiGroup := rest.App.Group("/api")
	apiGroup.Get("/tes" , rest.GetAllMovie)
}