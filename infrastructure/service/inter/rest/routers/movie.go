package restRouters

import (
	"github.com/gofiber/fiber/v2"
)

func (rest *RestServer) GetAllMovie(ctx *fiber.Ctx) error {
	data, httpCode, err := rest.RestInteractor.RestInterface.DoGetListMovie(ctx.Context(), 1, "batman")
	if err != nil {
		return err
	}

	return ctx.Status(httpCode).JSON(&fiber.Map{
		"status": "tes",
		"data":   data,
	})
}
