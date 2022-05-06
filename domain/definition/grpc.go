package definition

import (
	"context"
	customError "github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"

	pb "github.com/alhamsya/boilerplate-go/protos"
)

type GrpcUsecase interface {
	DoGetListMovie(ctx context.Context, reqClient *pb.GetListMovieReq) (resp *pb.GetListMovieResp, errResp *customError.Error)
	DoGetDetailMovie(ctx context.Context, req *pb.GetDetailMovieReq) (resp *pb.GetDetailMovieResp, err error)
}
