package restRouters

import (
	modelReq "github.com/alhamsya/boilerplate-go/domain/models/request"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_resp"
	"github.com/gofiber/fiber/v2"
)

func (rest *RestServer) CreateUser(ctx *fiber.Ctx) error {
	req := new(modelReq.User)
	err := ctx.BodyParser(req)
	if err != nil {
		return customResp.New(ctx).SetErr(err).SetHttpCode(fiber.StatusBadRequest).Send()
	}

	resp, httpCode, err := rest.RestInteractor.Usecase.DoCreateUser(ctx, req)
	if err != nil {
		return customResp.New(ctx).SetErr(err).SetHttpCode(httpCode).Send()
	}

	return customResp.New(ctx).SetHttpCode(httpCode).SetData(resp).SetMessage("create user").Send()
}
