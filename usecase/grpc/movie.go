package grpcUc

import (
	"context"
	pb "github.com/alhamsya/boilerplate-go/protos"
)

func (uc *UcInteractor) DoGetListMovie(ctx context.Context, req *pb.GetListMovieReq) (resp *pb.GetListMovieResp, err error) {
	uc.OMDBRepo.GetListMovie(1, "batman")
	return nil, nil
}
