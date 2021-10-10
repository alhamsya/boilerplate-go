package restUc

import (
	"context"
	"net/http"
)

func (uc *UcInteractor) DoGetListMovie(ctx context.Context, page int, search string) (data interface{}, httpCode int, err error) {
	respMovie, err := uc.OMDBRepo.GetListMovie(page, search)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return respMovie, http.StatusOK, nil
}
