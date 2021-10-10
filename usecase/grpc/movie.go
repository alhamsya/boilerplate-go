package grpcUc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alhamsya/boilerplate-go/lib/helpers/client"
	"strconv"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/alhamsya/boilerplate-go/lib/utils/datetime"

	pb "github.com/alhamsya/boilerplate-go/protos"
)

func (uc *UcInteractor) DoGetListMovie(ctx context.Context, reqClient *pb.GetListMovieReq) (resp *pb.GetListMovieResp, err error) {
	respMovie, err := uc.OMDBRepo.GetListMovie(reqClient.Search, reqClient.Page)
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
		Data: &pb.DataListMovie{
			Items: items,
			Total: total,
		},
	}

	now, err := datetime.CurrentTimeF(constCommon.DateTime)
	if err != nil {
		return nil, err
	}

	reqStr, _ := json.Marshal(reqClient)
	respStr, _ := json.Marshal(resp)
	reqDB := &modelMovie.DBHistoryLog{
		Request:    string(reqStr),
		Response:   string(respStr),
		SourceData: "GRPC",
		CreatedAt:  now,
		CreatedBy:  client.GrpcGetIP(ctx),
	}
	_, err = uc.ServiceRepo.CreateHistoryLog(ctx, reqDB)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
