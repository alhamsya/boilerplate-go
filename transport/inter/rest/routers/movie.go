package restRouters

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_resp"
	"github.com/gofiber/fiber/v2"
)

func (rest *RestServer) GetListMovie(ctx *fiber.Ctx) error {
	paramPage, err := strconv.ParseInt(ctx.Query("page", "1"), 10, 64)
	if err != nil {
		return customResp.Err(ctx, fiber.StatusBadRequest, err, "invalid request page")
	}

	paramSearch := ctx.Query("search")
	if strings.TrimSpace(paramSearch) == "" {
		return customResp.Err(ctx, fiber.StatusBadRequest, err, "please input search movie")
	}

	reqClient := &modelMovie.ReqListMovie{
		Search: url.QueryEscape(strings.ToLower(paramSearch)),
		Page:   paramPage,
	}
	data, httpCode, err := rest.RestInteractor.Usecase.DoGetListMovie(ctx, reqClient)
	if err != nil {
		return customResp.Err(ctx, httpCode, err, "DoGetListMovie")
	}

	return customResp.Success(ctx, httpCode, data, "get all movie")
}

func (rest *RestServer) GetDetailMovie(ctx *fiber.Ctx) error {
	data, httpCode, err := rest.RestInteractor.Usecase.DoGetDetailMovie(ctx, url.QueryEscape(strings.ToLower(ctx.Params("movieID"))))
	if err != nil {
		return customResp.Err(ctx, httpCode, err, "DoGetDetailMovie")
	}

	return customResp.Success(ctx, httpCode, data, "get detail movie")
}
