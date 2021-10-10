package restRouters

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/gofiber/fiber/v2"
)

func (rest *RestServer) GetListMovie(ctx *fiber.Ctx) error {
	paramPage, err := strconv.ParseInt(ctx.Query("page", "1"), 10, 64)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid request page",
			"data": nil,
		})
	}

	paramSearch := ctx.Query("search")
	if strings.TrimSpace(paramSearch) == "" {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please input search movie",
			"data": nil,
		})
	}

	reqClient := &modelMovie.ReqListMovie{
		Search: paramSearch,
		Page:   paramPage,
	}
	data, httpCode, err := rest.RestInteractor.RestInterface.DoGetListMovie(ctx, reqClient)
	if err != nil {
		return ctx.Status(httpCode).JSON(&fiber.Map{
			"message": err.Error(),
			"data": nil,
		})
	}

	return ctx.Status(httpCode).JSON(&fiber.Map{
		"message": "get all movie successfully",
		"data":   data,
	})
}

func (rest *RestServer) GetDetailMovie(ctx *fiber.Ctx) error {
	data, httpCode, err := rest.RestInteractor.RestInterface.DoGetDetailMovie(ctx, ctx.Params("movieID"))
	if err != nil {
		return ctx.Status(httpCode).JSON(&fiber.Map{
			"message": err.Error(),
			"data": nil,
		})
	}

	return ctx.Status(httpCode).JSON(&fiber.Map{
		"message": "get detail movie successfully",
		"data":   data,
	})
}
