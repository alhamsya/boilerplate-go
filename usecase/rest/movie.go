package restUC

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/domain/models/database"
	"github.com/alhamsya/boilerplate-go/domain/models/request"
	"github.com/alhamsya/boilerplate-go/domain/models/response"
	"github.com/alhamsya/boilerplate-go/infrastructure/external/omdb"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/null"
)

//DoGetListMovie get list movie based on request client
func (uc *UCInteractor) DoGetListMovie(ctx *fiber.Ctx, reqClient *modelReq.ListMovie) (resp *modelResp.ListMovie, httpCode int, err error) {
	//implement call wrapping and on purpose do not use error wrapping
	respWrapper, err := uc.CallWrapperRepo.GetWrapper("omdb").Call(func() (interface{}, error) {
		//get data from redis
		respMovie, err := uc.CacheRepo.GetListMovie(ctx.Context(), reqClient.Search, reqClient.Page)
		if err == nil {
			return respMovie, nil
		}

		//api call to the OMDB
		respMovie, err = uc.OMDBRepo.GetListMovie(reqClient.Search, reqClient.Page)
		if err != nil {
			return nil, err
		}

		//ignore for return error
		uc.CacheRepo.SetListMovie(ctx.Context(), reqClient.Search, reqClient.Page, respMovie)
		return respMovie, nil
	})

	//handle error for API call
	if err != nil {
		return nil, http.StatusInternalServerError, customError.WrapFlag(err, "OMDBRepo", "GetListMovie")
	}

	//force data to struct
	respMovie := respWrapper.(*omdb.OMDBList)

	//handle response wrapper is nil
	if respMovie == nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("data from api call does not exist")
	}

	status, err := strconv.ParseBool(respMovie.Response)
	if err != nil {
		return nil, http.StatusConflict, fmt.Errorf("response from api third party there is a problem")
	}

	if !status {
		return nil, http.StatusBadRequest, fmt.Errorf(respMovie.Error)
	}

	resp = new(modelResp.ListMovie)
	for _, movie := range respMovie.Search {
		resp.Items = append(resp.Items, modelResp.Items{
			Title:   movie.Title,
			Year:    movie.Year,
			MovieID: movie.ImdbID,
			Types:   movie.Type,
			Poster:  movie.Poster,
		})
	}

	total, err := strconv.ParseInt(respMovie.TotalResults, 10, 64)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("fail convert total result")
	}

	now, err := uc.UtilsRepo.CurrentTimeF(constCommon.DateTime)
	if err != nil {
		return nil, http.StatusInternalServerError, customError.WrapFlag(err, "datetime", "CurrentTimeF")
	}

	reqStr, _ := json.Marshal(reqClient)
	respStr, _ := json.Marshal(respMovie)
	reqDB := &modelDB.HistoryLog{
		Endpoint:   null.StringFrom(ctx.Path()),
		Request:    string(reqStr),
		Response:   string(respStr),
		SourceData: constCommon.TypeREST,
		CreatedAt:  now,
		CreatedBy:  ctx.IP(),
	}
	_, err = uc.DBRepo.CreateHistoryLog(ctx.Context(), reqDB)
	if err != nil {
		return nil, http.StatusInternalServerError, customError.WrapFlag(err, "database", "CreateHistoryLog")
	}

	resp = &modelResp.ListMovie{
		Items: resp.Items,
		Total: total,
	}

	return resp, http.StatusOK, nil
}

func (uc *UCInteractor) DoGetDetailMovie(ctx *fiber.Ctx, movieID string) (resp *modelResp.DetailMovie, httpCode int, err error) {
	//implement call wrapping and on purpose do not use error wrapping
	respWrapper, err := uc.CallWrapperRepo.GetWrapper("omdb").Call(func() (interface{}, error) {
		//get data from redis
		respMovie, err := uc.CacheRepo.GetDetailMovie(ctx.Context(), movieID)
		if err == nil {
			return respMovie, nil
		}

		//api call to the OMDB
		respMovie, err = uc.OMDBRepo.GetDetailMovie(movieID)
		if err != nil {
			return nil, err
		}

		//ignore for return error
		uc.CacheRepo.SetDetailMovie(ctx.Context(), movieID, respMovie)
		return respMovie, nil
	})

	//handle error for API call
	if err != nil {
		return nil, http.StatusInternalServerError, customError.WrapFlag(err, "OMDBRepo", "GetDetailMovie")
	}

	//handle response wrapper is nil
	if respWrapper == nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("data from api call does not exist")
	}

	//force data to struct
	respMovie := respWrapper.(*omdb.OMDBDetail)

	status, err := strconv.ParseBool(respMovie.Response)
	if err != nil {
		return nil, http.StatusConflict, customError.Wrap(err, "response from api third party there is a problem")
	}

	if !status {
		return nil, http.StatusBadRequest, customError.WrapFlag(fmt.Errorf(respMovie.Error), "OMDBRepo", "status third party")
	}

	resp = new(modelResp.DetailMovie)
	for _, rating := range respMovie.Ratings {
		resp.Ratings = append(resp.Ratings, modelResp.Ratings{
			Source: rating.Source,
			Value:  rating.Value,
		})
	}

	now, err := uc.UtilsRepo.CurrentTimeF(constCommon.DateTime)
	if err != nil {
		return nil, http.StatusInternalServerError, customError.Wrap(err, "CurrentTimeF")
	}

	reqStr, _ := json.Marshal(movieID)
	respStr, _ := json.Marshal(respMovie)
	reqDB := &modelDB.HistoryLog{
		Endpoint:   null.StringFrom(ctx.Path()),
		Request:    string(reqStr),
		Response:   string(respStr),
		SourceData: constCommon.TypeREST,
		CreatedAt:  now,
		CreatedBy:  ctx.IP(),
	}
	_, err = uc.DBRepo.CreateHistoryLog(ctx.Context(), reqDB)
	if err != nil {
		return nil, http.StatusInternalServerError, customError.WrapFlag(err, "database", "CreateHistoryLog")
	}

	resp = &modelResp.DetailMovie{
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
		Ratings:    resp.Ratings,
		MetaScore:  respMovie.Metascore,
		ImdbRating: respMovie.ImdbRating,
		ImdbVotes:  respMovie.ImdbVotes,
		ImdbID:     respMovie.ImdbID,
		Type:       respMovie.Type,
		DVD:        respMovie.DVD,
		BoxOffice:  respMovie.BoxOffice,
		Production: respMovie.Production,
		Website:    respMovie.Website,
	}
	return resp, http.StatusOK, err
}
