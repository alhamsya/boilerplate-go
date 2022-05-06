package restHandler

import (
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
	"github.com/alhamsya/boilerplate-go/transport/inter/rest/routers"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Cfg        *config.ServiceConfig
	App        *fiber.App
	Interactor *restRouters.RestInteractor
}
