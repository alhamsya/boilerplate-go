package grpcUC

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/alhamsya/boilerplate-go/lib/helpers/client"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
	"github.com/alhamsya/boilerplate-go/service/exter/omdb"
	"github.com/volatiletech/null"
	"google.golang.org/grpc"

	pb "github.com/alhamsya/boilerplate-go/protos"
)

//DoGetListMovie get list movie
func (uc *UCInteractor) DoGetListMovie(ctx context.Context, reqClient *pb.GetListMovieReq) (resp *pb.GetListMovieResp, err error) {
	//implement call wrapping and on purpose do not use error wrapping
	respWrapper, err := uc.CallWrapperRepo.GetWrapper("omdb").Call(func() (interface{}, error) {
		//get data from redis
		respMovie, err := uc.CacheRepo.GetListMovie(ctx, reqClient.Search, reqClient.Page)
		if err == nil {
			return respMovie, nil
		}

		//api call to the OMDB
		respMovie, err = uc.OMDBRepo.GetListMovie(reqClient.Search, reqClient.Page)
		if err != nil {
			return nil, err
		}

		//ignore for return error
		uc.CacheRepo.SetListMovie(ctx, reqClient.Search, reqClient.Page, respMovie)
		return respMovie, nil
	})

	//handle error for API call
	if err != nil {
		return nil, customError.WrapFlag(err, "OMDBRepo", "GetListMovie")
	}

	//handle response wrapper is nil
	if respWrapper == nil {
		return nil, fmt.Errorf("data from api call does not exist")
	}

	//force data to struct
	respMovie := respWrapper.(*omdb.OMDBList)

	status, err := strconv.ParseBool(respMovie.Response)
	if err != nil {
		return nil, customError.Wrap(err, "ParseBool")
	}

	if !status {
		return nil, customError.WrapFlag(fmt.Errorf(respMovie.Error), "OMDBRepo", "status third party")
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
		return nil, customError.Wrap(err, "ParseInt")
	}

	now, err := uc.UtilsRepo.CurrentTimeF(constCommon.DateTime)
	if err != nil {
		return nil, customError.Wrap(err, "CurrentTimeF")
	}

	endPoint, _ := grpc.Method(ctx)
	reqStr, _ := json.Marshal(reqClient)
	respStr, _ := json.Marshal(respMovie)
	reqDB := &modelMovie.DBHistoryLog{
		Endpoint:   null.StringFrom(endPoint),
		Request:    string(reqStr),
		Response:   string(respStr),
		SourceData: constCommon.TypeGRPC,
		CreatedAt:  now,
		CreatedBy:  client.GrpcGetIP(ctx),
	}

	_, err = uc.DBRepo.CreateHistoryLog(ctx, reqDB)
	if err != nil {
		return nil, customError.WrapFlag(err, "database", "CreateHistoryLog")
	}

	resp = &pb.GetListMovieResp{
		Data: &pb.DataListMovie{
			Items: items,
			Total: total,
		},
	}

	return resp, nil
}

func (uc *UCInteractor) DoGetDetailMovie(ctx context.Context, reqClient *pb.GetDetailMovieReq) (data *pb.GetDetailMovieResp, err error) {
	//implement call wrapping and on purpose do not use error wrapping
	respWrapper, err := uc.CallWrapperRepo.GetWrapper("omdb").Call(func() (interface{}, error) {
		//get data from redis
		respMovie, err := uc.CacheRepo.GetDetailMovie(ctx, reqClient.MovieID)
		if err == nil {
			return respMovie, nil
		}

		//api call to the OMDB
		respMovie, err = uc.OMDBRepo.GetDetailMovie(reqClient.MovieID)
		if err != nil {
			return nil, err
		}

		//ignore for return error
		uc.CacheRepo.SetDetailMovie(ctx, reqClient.MovieID, respMovie)
		return respMovie, nil
	})

	//handle error for API call
	if err != nil {
		return nil, customError.WrapFlag(err, "OMDBRepo", "GetDetailMovie")
	}

	//handle response wrapper is nil
	if respWrapper == nil {
		return nil, fmt.Errorf("data from api call does not exist")
	}

	//force data to struct
	respMovie := respWrapper.(*omdb.OMDBDetail)

	status, err := strconv.ParseBool(respMovie.Response)
	if err != nil {
		return nil, customError.Wrap(err, "response from api third party there is a problem")
	}

	if !status {
		return nil, customError.WrapFlag(fmt.Errorf(respMovie.Error), "OMDBRepo", "status third party")
	}

	var ratings []*pb.Ratings
	for _, rating := range respMovie.Ratings {
		ratings = append(ratings, &pb.Ratings{
			Source: rating.Source,
			Value:  rating.Value,
		})
	}

	now, err := uc.UtilsRepo.CurrentTimeF(constCommon.DateTime)
	if err != nil {
		return nil, customError.Wrap(err, "CurrentTimeF")
	}

	endPoint, _ := grpc.Method(ctx)
	reqStr, _ := json.Marshal(reqClient)
	respStr, _ := json.Marshal(respMovie)
	reqDB := &modelMovie.DBHistoryLog{
		Endpoint:   null.StringFrom(endPoint),
		Request:    string(reqStr),
		Response:   string(respStr),
		SourceData: constCommon.TypeGRPC,
		CreatedAt:  now,
		CreatedBy:  client.GrpcGetIP(ctx),
	}
	_, err = uc.DBRepo.CreateHistoryLog(ctx, reqDB)
	if err != nil {
		return nil, customError.Wrap(err, "database", "CreateHistoryLog")
	}

	data = &pb.GetDetailMovieResp{
		Data: &pb.DataDetailMovie{
			Title:      respMovie.Title,
			Year:       respMovie.Year,
			Rated:      respMovie.Rated,
			Released:   respMovie.Released,
			Runtime:    respMovie.Runtime,
			Genre:      respMovie.Genre,
			Director:   respMovie.Director,
			Writer:     respMovie.Writer,
			Actors:     respMovie.Actors,
			Plot:       respMovie.Plot,
			Language:   respMovie.Language,
			Country:    respMovie.Country,
			Awards:     respMovie.Awards,
			Poster:     respMovie.Poster,
			Ratings:    ratings,
			MetaScore:  respMovie.Metascore,
			ImdbRating: respMovie.ImdbRating,
			ImdbVotes:  respMovie.ImdbVotes,
			ImdbID:     respMovie.ImdbID,
			Type:       respMovie.Type,
			DVD:        respMovie.DVD,
			BoxOffice:  respMovie.BoxOffice,
			Production: respMovie.Production,
			Website:    respMovie.Website,
		},
	}

	return data, nil
}
