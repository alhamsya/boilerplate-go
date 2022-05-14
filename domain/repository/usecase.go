package repository

import (
	"context"

	"github.com/alhamsya/boilerplate-go/domain/models/request"
	"github.com/alhamsya/boilerplate-go/domain/models/response"
	"github.com/alhamsya/boilerplate-go/lib/managers/custom_error"
	"github.com/gofiber/fiber/v2"

	pb "github.com/alhamsya/boilerplate-go/protos"
)

type RestUsecase interface {
	DoGetListMovie(ctx *fiber.Ctx, reqClient *modelReq.ListMovie) (resp *modelResp.ListMovie, httpCode int, err error)
	DoGetDetailMovie(ctx *fiber.Ctx, movieID string) (resp *modelResp.DetailMovie, httpCode int, err error)
	DoSignup(ctx *fiber.Ctx, reqClient *modelReq.User) (resp *modelResp.User, httpCode int, err error)
	DoSigning(ctx *fiber.Ctx, reqClient *modelReq.User) (resp *modelResp.User, httpCode int, err error)
}

type GrpcUsecase interface {
	DoGetListMovie(ctx context.Context, reqClient *pb.GetListMovieReq) (resp *pb.GetListMovieResp, errResp *customError.Error)
	DoGetDetailMovie(ctx context.Context, req *pb.GetDetailMovieReq) (resp *pb.GetDetailMovieResp, err error)
}

type CronUsecase interface {
	DoCreateDummyData(ctx context.Context) error
}

type ConsumerUsecase interface {
	DoPayment(ctx context.Context) error
}
