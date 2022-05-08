package restRouters

import (
	"github.com/alhamsya/boilerplate-go/domain/models/request"
	"net/url"
	"strconv"
	"strings"

	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_resp"
	"github.com/gofiber/fiber/v2"
)

func (rest *RestServer) GetListMovie(ctx *fiber.Ctx) error {
	paramPage, err := strconv.ParseInt(ctx.Query("page", "1"), 10, 64)
	if err != nil {
		return customResp.New(ctx).SetHttpCode(fiber.StatusBadRequest).SetMessage("invalid request page").Send()
	}

	paramSearch := ctx.Query("search")
	if strings.TrimSpace(paramSearch) == "" {
		return customResp.New(ctx).SetHttpCode(fiber.StatusBadRequest).SetMessage("please input search movie").Send()
	}

	reqClient := &modelReq.ListMovie{
		Search: url.QueryEscape(strings.ToLower(paramSearch)),
		Page:   paramPage,
	}
	resp, httpCode, err := rest.RestInteractor.Usecase.DoGetListMovie(ctx, reqClient)
	if err != nil {
		return customResp.New(ctx).SetHttpCode(httpCode).SetErr(err).Send("DoGetListMovie")
	}

	return customResp.New(ctx).SetHttpCode(httpCode).SetData(resp).Send("get all movie")
}

func (rest *RestServer) GetDetailMovie(ctx *fiber.Ctx) error {
	resp, httpCode, err := rest.RestInteractor.Usecase.DoGetDetailMovie(ctx, url.QueryEscape(strings.ToLower(ctx.Params("movieID"))))
	if err != nil {
		return customResp.New(ctx).SetHttpCode(httpCode).SetErr(err).Send("DoGetDetailMovie")
	}

	return customResp.New(ctx).SetHttpCode(httpCode).SetData(resp).Send("get detail movie")
}
