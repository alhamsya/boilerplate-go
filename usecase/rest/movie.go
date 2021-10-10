package restUc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/alhamsya/boilerplate-go/lib/utils/datetime"
)

func (uc *UcInteractor) DoGetListMovie(ctx context.Context, reqClient *modelMovie.ReqListMovie) (resp *modelMovie.RespListMovie, httpCode int, err error) {
	respMovie, err := uc.OMDBRepo.GetListMovie(reqClient.Search, reqClient.Page)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	var items []modelMovie.Items
	for _, movie := range respMovie.Search {
		items = append(items, modelMovie.Items{
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

	resp = &modelMovie.RespListMovie{
		Items: items,
		Total: total,
	}

	now, err := datetime.CurrentTimeF(constCommon.DateTime)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	reqStr, _ := json.Marshal(reqClient)
	respStr, _ := json.Marshal(resp)
	reqDB := &modelMovie.DBHistoryLog{
		Request:   string(reqStr),
		Response:  string(respStr),
		SourceData: "REST",
		CreatedAt: now,
		CreatedBy: reqClient.CreatedBy,
	}
	_, err = uc.ServiceRepo.CreateHistoryLog(ctx, reqDB)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return resp, http.StatusOK, nil
}
