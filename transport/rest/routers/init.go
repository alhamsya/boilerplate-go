package restRouters

import "github.com/alhamsya/boilerplate-go/domain/repositories/mocks"

var mockRestInterface *mocks.RestUsecase

func initMock() {
	mockRestInterface = new(mocks.RestUsecase)
}

func New(this *RestServer) *RestServer {
	return &RestServer{
		Cfg:            this.Cfg,
		App:            this.App,
		RestInteractor: this.RestInteractor,
	}
}
