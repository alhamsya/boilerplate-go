package definition

import (
	"github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/gofiber/fiber/v2"
)

type RestUsecase interface {
	DoGetListMovie(ctx *fiber.Ctx, reqClient *modelMovie.ReqListMovie) (resp *modelMovie.RespListMovie, httpCode int, err error)
	DoGetDetailMovie(ctx *fiber.Ctx, movieID string) (resp *modelMovie.RespDetailMovie, httpCode int, err error)
}
