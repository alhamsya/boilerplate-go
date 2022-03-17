package restMiddleware

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func (m *Middleware) Authorize(h fiber.Handler) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		search := ctx.Query("search")
		if strings.Contains(search, "bat") {
			return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
				"message": "Testing",
				"data":    nil,
			})
		}
		return h(ctx)
	}
}
