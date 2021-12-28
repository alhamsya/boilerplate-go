package grpcRouters

import (
	"context"
	constCommon "github.com/alhamsya/boilerplate-go/domain/constants"
	"net/url"
	"strings"

	pb "github.com/alhamsya/boilerplate-go/protos"
)

func (grpc *GrpcServer) GetListMovie(ctx context.Context, req *pb.GetListMovieReq) (*pb.GetListMovieResp, error) {
	if strings.TrimSpace(req.Search) == "" {
		return &pb.GetListMovieResp{
			Status: &pb.RPCStatus{
				Code:    constCommon.GRPCStatusInvalidArgument,
				Message: "request search not empty",
			},
		}, nil
	}

	if req.Page < 1 {
		req.Page = 1
	}

	req = &pb.GetListMovieReq{
		Search: url.QueryEscape(strings.ToLower(req.Search)),
		Page:   req.Page,
	}
	resp, err := grpc.GrpcInteractor.GrpcInterface.DoGetListMovie(ctx, req)
	if err != nil {
		return &pb.GetListMovieResp{
			Status: &pb.RPCStatus{
				Code:    constCommon.GRPCStatusInternal,
				Message: err.Error(),
			},
		}, nil
	}

	if resp.Data == nil {
		return &pb.GetListMovieResp{
			Status: resp.Status,
		}, nil
	}

	return &pb.GetListMovieResp{
		Status: &pb.RPCStatus{
			Code:    constCommon.GRPCStatusOk,
			Message: "get all movie successfully",
		},
		Data: resp.Data,
	}, nil
}

func (grpc *GrpcServer) GetDetailMovie(ctx context.Context, req *pb.GetDetailMovieReq) (*pb.GetDetailMovieResp, error) {
	resp, err := grpc.GrpcInteractor.GrpcInterface.DoGetDetailMovie(ctx, req)
	if err != nil {
		return &pb.GetDetailMovieResp{
			Status: &pb.RPCStatus{
				Code:    constCommon.GRPCStatusInternal,
				Message: err.Error(),
			},
		}, nil
	}

	if resp.Data == nil {
		return &pb.GetDetailMovieResp{
			Status: resp.Status,
		}, nil
	}

	return &pb.GetDetailMovieResp{
		Status: &pb.RPCStatus{
			Code:    constCommon.GRPCStatusOk,
			Message: "get detail movie successfully",
		},
		Data: resp.Data,
	}, nil
}
