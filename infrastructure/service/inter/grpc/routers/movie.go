package grpcRouters

import (
	"context"

	pb "github.com/alhamsya/boilerplate-go/protos"
)

func (grpc *GrpcServer) GetListMovie(ctx context.Context, req *pb.GetListMovieReq) (*pb.GetListMovieResp, error) {
	panic("implement me")
}

func (grpc *GrpcServer) GetDetailMovie(ctx context.Context, req *pb.GetDetailMovieReq) (*pb.GetDetailMovieResp, error) {
	panic("implement me")
}
