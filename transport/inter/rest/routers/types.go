package restRouters

import (
	"github.com/alhamsya/boilerplate-go/domain/definition"
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
	"github.com/alhamsya/boilerplate-go/middleware/rest"
	"github.com/gofiber/fiber/v2"
)

type RestServer struct {
	Cfg            *config.ServiceConfig
	App            *fiber.App
	RestInteractor *RestInteractor
}

type RestInteractor struct {
	Usecase    definition.RestUsecase
	Middleware *restMiddleware.Middleware
}
