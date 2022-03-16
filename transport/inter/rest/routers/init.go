package restRouters

import "github.com/alhamsya/boilerplate-go/domain/definition/mocks"

var mockRestInterface *mocks.RestInterface

func initMock()  {
	mockRestInterface = new(mocks.RestInterface)
}

func New(this *RestServer) *RestServer {
	return &RestServer{
		Cfg:            this.Cfg,
		App:            this.App,
		RestInteractor: this.RestInteractor,
	}
}
