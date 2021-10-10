package definition

import (
	"context"

	"github.com/alhamsya/boilerplate-go/domain/models/movie"

	pb "github.com/alhamsya/boilerplate-go/protos"
)

type RestInterface interface {
	DoGetListMovie(ctx context.Context, search string, page int64) (resp *modelMovie.RespListMovie, httpCode int, err error)
}

type GrpcInterface interface {
	DoGetListMovie(ctx context.Context, req *pb.GetListMovieReq) (resp *pb.GetListMovieResp, err error)
}