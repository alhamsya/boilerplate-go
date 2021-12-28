package restHandler

import (
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/alhamsya/boilerplate-go/service/inter/rest/routers"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Cfg        *config.ServiceConfig
	App        *fiber.App
	Interactor *restRouters.RestInteractor
}
