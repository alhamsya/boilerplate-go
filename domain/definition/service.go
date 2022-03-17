package definition

import (
	"context"

	"github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/gofiber/fiber/v2"

	pb "github.com/alhamsya/boilerplate-go/protos"
)

type RestUsecase interface {
	DoGetListMovie(ctx *fiber.Ctx, reqClient *modelMovie.ReqListMovie) (resp *modelMovie.RespListMovie, httpCode int, err error)
	DoGetDetailMovie(ctx *fiber.Ctx, movieID string) (resp *modelMovie.RespDetailMovie, httpCode int, err error)
}

type GrpcUsecase interface {
	DoGetListMovie(ctx context.Context, req *pb.GetListMovieReq) (resp *pb.GetListMovieResp, err error)
	DoGetDetailMovie(ctx context.Context, req *pb.GetDetailMovieReq) (resp *pb.GetDetailMovieResp, err error)
}

type SchedulerUsecase interface {
	DoCreateDummyData() error
	DoChunkCountingData() error
}
