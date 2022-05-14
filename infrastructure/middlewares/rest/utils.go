package restMiddleware

import "github.com/gofiber/fiber/v2"

func (m *Middleware) GetDataUser(ctx *fiber.Ctx) (result MyClaims) {
	data := ctx.UserContext().Value("user")
	return data.(MyClaims)
}
