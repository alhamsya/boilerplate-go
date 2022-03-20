package definition

import (
	"context"

	pb "github.com/alhamsya/boilerplate-go/protos"
)

type GrpcUsecase interface {
	DoGetListMovie(ctx context.Context, req *pb.GetListMovieReq) (resp *pb.GetListMovieResp, err error)
	DoGetDetailMovie(ctx context.Context, req *pb.GetDetailMovieReq) (resp *pb.GetDetailMovieResp, err error)
}
