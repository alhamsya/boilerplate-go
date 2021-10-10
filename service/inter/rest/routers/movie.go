package restRouters

import (
	"github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/gofiber/fiber/v2"
)

func (rest *RestServer) GetListMovie(ctx *fiber.Ctx) error {
	reqClient := &modelMovie.ReqListMovie{
		Search: "batman",
		Page:   1,
	}
	data, httpCode, err := rest.RestInteractor.RestInterface.DoGetListMovie(ctx, reqClient)
	if err != nil {
		return err
	}

	return ctx.Status(httpCode).JSON(&fiber.Map{
		"message": "get all movie successfully",
		"data":   data,
	})
}

func (rest *RestServer) GetDetailMovie(ctx *fiber.Ctx) error {
	data, httpCode, err := rest.RestInteractor.RestInterface.DoGetDetailMovie(ctx, ctx.Params("movieID"))
	if err != nil {
		return err
	}

	return ctx.Status(httpCode).JSON(&fiber.Map{
		"message": "get detail movie successfully",
		"data":   data,
	})
}
