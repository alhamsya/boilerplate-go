package grpcRouters

import (
	"context"
	constCommon "github.com/alhamsya/boilerplate-go/domain/constants"

	pb "github.com/alhamsya/boilerplate-go/protos"
)

func (grpc *GrpcServer) GetListMovie(ctx context.Context, req *pb.GetListMovieReq) (*pb.GetListMovieResp, error) {
	resp, err := grpc.GrpcInteractor.GrpcInterface.DoGetListMovie(ctx, req)
	if err != nil {
		return &pb.GetListMovieResp{
			Status: &pb.RPCStatus{
				Code:    constCommon.GRPCStatusInternal,
				Message: err.Error(),
			},
		}, nil
	}

	return resp, nil
}

func (grpc *GrpcServer) GetDetailMovie(ctx context.Context, req *pb.GetDetailMovieReq) (*pb.GetDetailMovieResp, error) {
	data, err := grpc.GrpcInteractor.GrpcInterface.DoGetDetailMovie(ctx, req)
	if err != nil {
		return &pb.GetDetailMovieResp{
			Status: &pb.RPCStatus{
				Code:    constCommon.GRPCStatusInternal,
				Message: err.Error(),
			},
		}, nil
	}

	return &pb.GetDetailMovieResp{
		Status: &pb.RPCStatus{
			Code:    constCommon.GRPCStatusOk,
			Message: "get detail movie successfully",
		},
		Data:   data,
	}, nil
}
