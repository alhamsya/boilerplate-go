package restHandler

import (
	"fmt"
	"strings"
	"time"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/alhamsya/boilerplate-go/service/inter/rest/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New(this *Handler) *Handler {
	app := fiber.New(fiber.Config{
		ReadTimeout: this.Cfg.RestServer.Timeout * time.Second,
	})

	//set middleware fiber framework
	app.Use(getLoggerConfig(this), getCORSConfig(this.Cfg), getLimiterConfig(this.Cfg))

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

func getLoggerConfig(this *Handler) logger.Config {
	loggerConfig := logger.Config{
		Format:     "[${latency}] ${status} - ${method} ${path}\n",
		TimeFormat: constCommon.DateTime,
		TimeZone:   constCommon.TimeLocalJakarta,
	}
	return loggerConfig
}

func getCORSConfig(cfg *config.MainConfig) cors.Config {
	corsConfig := cors.Config{
		AllowOrigins: strings.Join(cfg.CORS.AllowOrigins, ", "),
		AllowHeaders: "Origin, Content-Type, Accept",
	}
	return corsConfig
}

func getLimiterConfig(cfg *config.MainConfig) limiter.Config {
	limiterConfig := limiter.Config{
		Max:        20,               // max count of connections
		Expiration: 30 * time.Second, // expiration time of the limit
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
	}

	return limiterConfig
}

func (h *Handler) Run() error {
	return h.App.Listen(fmt.Sprintf(":%s", h.Cfg.RestServer.Port))
}
