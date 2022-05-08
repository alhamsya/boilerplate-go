package repository

import (
	"github.com/alhamsya/boilerplate-go/infrastructure/external/omdb"
)

type OMDBRepo interface {
	GetListMovie(search string, page int64) (*external.OMDBList, error)
	GetDetailMovie(movieID string) (*external.OMDBDetail, error)
}
