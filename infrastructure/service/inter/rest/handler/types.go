package restHandler

import (
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/alhamsya/boilerplate-go/infrastructure/service/inter/rest/routers"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Cfg        *config.MainConfig
	App        *fiber.App
	Interactor *routers.RestInteractor
}
