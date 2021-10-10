package grpcUc

import (
	"context"
	"fmt"
	"strconv"

	"github.com/alhamsya/boilerplate-go/domain/constants"

	pb "github.com/alhamsya/boilerplate-go/protos"
)

func (uc *UcInteractor) DoGetListMovie(ctx context.Context, req *pb.GetListMovieReq) (resp *pb.GetListMovieResp, err error) {
	respMovie, err := uc.OMDBRepo.GetListMovie(req.Search, req.Page)
	if err != nil {
		return nil, err
	}

	var items []*pb.ItemsMovie
	for _, movie := range respMovie.Search {
		items = append(items, &pb.ItemsMovie{
			Title:   movie.Title,
			Year:    movie.Year,
			MovieID: movie.ImdbID,
			Types:   movie.Type,
			Poster:  movie.Poster,
		})
	}

	total, err := strconv.ParseInt(respMovie.TotalResults, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("fail convert total result")
	}

	resp = &pb.GetListMovieResp{
		Status: &pb.RPCStatus{
			Code:    constCommon.GRPCStatusOk,
			Message: "get all movie successfully",
		},
		Items:  items,
		Total:  total,
	}
	return nil, nil
}
