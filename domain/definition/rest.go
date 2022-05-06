package definition

import (
	"github.com/alhamsya/boilerplate-go/domain/models/request"
	"github.com/alhamsya/boilerplate-go/domain/models/response"
	"github.com/gofiber/fiber/v2"
)

type RestUsecase interface {
	DoGetListMovie(ctx *fiber.Ctx, reqClient *modelReq.ListMovie) (resp *modelResp.ListMovie, httpCode int, err error)
	DoGetDetailMovie(ctx *fiber.Ctx, movieID string) (resp *modelResp.DetailMovie, httpCode int, err error)
	DoSignup(ctx *fiber.Ctx, reqClient *modelReq.User) (resp *modelResp.User, httpCode int, err error)
	DoSigning(ctx *fiber.Ctx, reqClient *modelReq.User) (resp *modelResp.User, httpCode int, err error)
}
