package restUC

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/volatiletech/null"
	"net/http"
	"strconv"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/alhamsya/boilerplate-go/lib/utils/datetime"
)

func (uc *UcInteractor) DoGetListMovie(ctx *fiber.Ctx, reqClient *modelMovie.ReqListMovie) (resp *modelMovie.RespListMovie, httpCode int, err error) {
	respMovie, err := uc.OMDBRepo.GetListMovie(reqClient.Search, reqClient.Page)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if respMovie == nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("data from api call does not exist")
	}

	resp = new(modelMovie.RespListMovie)
	for _, movie := range respMovie.Search {
		resp.Items = append(resp.Items, modelMovie.Items{
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

	now, err := datetime.CurrentTimeF(constCommon.DateTime)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	reqStr, _ := json.Marshal(reqClient)
	respStr, _ := json.Marshal(respMovie)
	reqDB := &modelMovie.DBHistoryLog{
		Endpoint:   null.StringFrom(ctx.Path()),
		Request:    string(reqStr),
		Response:   string(respStr),
		SourceData: "REST",
		CreatedAt:  now,
		CreatedBy:  ctx.IP(),
	}
	_, err = uc.ServiceRepo.CreateHistoryLog(ctx.Context(), reqDB)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	resp = &modelMovie.RespListMovie{
		Items: resp.Items,
		Total: total,
	}

	return resp, http.StatusOK, nil
}

func (uc *UcInteractor) DoGetDetailMovie(ctx *fiber.Ctx, movieID string) (resp *modelMovie.RespDetailMovie, httpCode int, err error) {
	respMovie, err := uc.OMDBRepo.GetDetailMovie(movieID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if respMovie == nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("data from api call does not exist")
	}

	resp = new(modelMovie.RespDetailMovie)
	for _, rating := range respMovie.Ratings {
		resp.Ratings = append(resp.Ratings, modelMovie.Ratings{
			Source: rating.Source,
			Value:  rating.Value,
		})
	}

	now, err := datetime.CurrentTimeF(constCommon.DateTime)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	reqStr, _ := json.Marshal(movieID)
	respStr, _ := json.Marshal(respMovie)
	reqDB := &modelMovie.DBHistoryLog{
		Endpoint:   null.StringFrom(ctx.Path()),
		Request:    string(reqStr),
		Response:   string(respStr),
		SourceData: "REST",
		CreatedAt:  now,
		CreatedBy:  ctx.IP(),
	}
	_, err = uc.ServiceRepo.CreateHistoryLog(ctx.Context(), reqDB)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	resp = &modelMovie.RespDetailMovie{
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
