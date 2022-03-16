package restHandler

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/alhamsya/boilerplate-go/lib/helpers/grace"
	"github.com/alhamsya/boilerplate-go/transport/inter/rest/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New(this *Handler) *Handler {
	app := fiber.New(fiber.Config{
		ReadTimeout: this.Cfg.RestServer.Timeout * time.Second,
		IdleTimeout: 5 * time.Second,
	})

	//set middleware fiber framework
	app.Use(
		getLoggerConfig(this.Cfg),
		getCORSConfig(this.Cfg),
		getLimiterConfig(this.Cfg),
		getRecoverConfig(this.Cfg),
	)

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

func getLoggerConfig(cfg *config.ServiceConfig) fiber.Handler {
	loggerConfig := logger.New(logger.Config{
		Format:     "[${latency}] ${status} - ${method} ${path}\n",
		TimeFormat: constCommon.DateTime,
		TimeZone:   constCommon.TimeLocalJakarta,
	})
	return loggerConfig
}

func getCORSConfig(cfg *config.ServiceConfig) fiber.Handler {
	corsConfig := cors.New(cors.Config{
		AllowOrigins: strings.Join(cfg.CORS.AllowOrigins, ", "),
		AllowHeaders: "Origin, Content-Type, Accept",
	})
	return corsConfig
}

func getLimiterConfig(cfg *config.ServiceConfig) fiber.Handler {
	limiterConfig := limiter.New(limiter.Config{
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
	})

	return limiterConfig
}

func getRecoverConfig(cfg *config.ServiceConfig) fiber.Handler {
	recoverConfig := recover.New(recover.Config{
		Next: func(c *fiber.Ctx) bool {
			c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
				"message": "please try again",
				"data":    nil,
			})
			return false
		},
		EnableStackTrace:  true,
		StackTraceHandler: nil,
	})

	return recoverConfig
}

func (h *Handler) Run() error {
	return grace.ServeRESTWithFiber(h.App, fmt.Sprintf(":%s", h.Cfg.RestServer.Port), 10*time.Second)
}
