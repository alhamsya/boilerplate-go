package restRouters

import (
	"github.com/alhamsya/boilerplate-go/domain/definition"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/gofiber/fiber/v2"
)

type RestServer struct {
	Cfg            *config.ServiceConfig
	App            *fiber.App
	RestInteractor *RestInteractor
}

type RestInteractor struct {
	RestInterface definition.RestInterface
}
