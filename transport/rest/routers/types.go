package restRouters

import (
	"github.com/alhamsya/boilerplate-go/domain/repositorys"
	"github.com/alhamsya/boilerplate-go/infrastructure/middlewares/rest"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
	"github.com/gofiber/fiber/v2"
)

type RestServer struct {
	Cfg            *config.ServiceConfig
	App            *fiber.App
	RestInteractor *RestInteractor
}

type RestInteractor struct {
	Usecase    repositorys.RestUsecase
	Middleware *restMiddleware.Middleware
}
