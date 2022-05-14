package restRouters

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
	"github.com/alhamsya/boilerplate-go/middleware/rest"
	"github.com/gofiber/fiber/v2"
)

type RestServer struct {
	Cfg            *config.ServiceConfig
	App            *fiber.App
	RestInteractor *RestInteractor
}

type RestInteractor struct {
	Usecase    repository.RestUsecase
	Middleware *restMiddleware.Middleware
}
