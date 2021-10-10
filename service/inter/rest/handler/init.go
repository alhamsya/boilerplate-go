package restHandler

import (
	"fmt"
	"github.com/alhamsya/boilerplate-go/service/inter/rest/routers"
	"time"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New(this *Handler) *Handler {
	app := fiber.New(fiber.Config{
		Prefork:                   false,
		ServerHeader:              "",
		StrictRouting:             false,
		CaseSensitive:             false,
		Immutable:                 false,
		UnescapePath:              false,
		ETag:                      false,
		BodyLimit:                 0,
		Concurrency:               0,
		Views:                     nil,
		ReadTimeout:               this.Cfg.RestServer.Timeout * time.Second,
		WriteTimeout:              0,
		IdleTimeout:               0,
		ReadBufferSize:            0,
		WriteBufferSize:           0,
		CompressedFileSuffix:      "",
		ProxyHeader:               "",
		GETOnly:                   false,
		ErrorHandler:              nil,
		DisableKeepalive:          false,
		DisableDefaultDate:        false,
		DisableDefaultContentType: false,
		DisableHeaderNormalizing:  false,
		DisableStartupMessage:     false,
		ReduceMemoryUsage:         false,
		JSONEncoder:               nil,
		Network:                   "",
	})

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n",
		TimeFormat: constCommon.DateTime,
		TimeZone:   constCommon.TimeLocalJakarta,
	}))

	routeHandler := &restRouters.RestServer{
		Cfg:            this.Cfg,
		App:            app,
		RestInteractor: this.Interactor,
	}
	restRouters.New(routeHandler).Register()
	return &Handler{
		Cfg:        this.Cfg,
		App:        app,
		Interactor: this.Interactor,
	}
}

func (h *Handler) Run() error {
	return h.App.Listen(fmt.Sprintf(":%s", h.Cfg.RestServer.Port))
}
