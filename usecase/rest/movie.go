package restUc

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alhamsya/boilerplate-go/domain/models/movie"
)

func (uc *UcInteractor) DoGetListMovie(ctx context.Context, search string, page int64) (resp *modelMovie.RespListMovie, httpCode int, err error) {
	respMovie, err := uc.OMDBRepo.GetListMovie(search, page)
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

	return resp, http.StatusOK, nil
}
