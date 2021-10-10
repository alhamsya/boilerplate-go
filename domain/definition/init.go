package definition

import (
	"context"

	pb "github.com/alhamsya/boilerplate-go/protos"
)

type RestInterface interface {
	DoGetListMovie(ctx context.Context, page int, search string) (data interface{}, httpCode int, err error)
}

type GrpcInterface interface {
	DoGetListMovie(ctx context.Context, req *pb.GetListMovieReq) (resp *pb.GetListMovieResp, err error)
}