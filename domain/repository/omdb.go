package repository

import (
	"github.com/alhamsya/boilerplate-go/transport/exter/omdb"
)

type OMDBRepo interface {
	GetListMovie(search string, page int64) (*omdb.OMDBList, error)
	GetDetailMovie(movieID string) (*omdb.OMDBDetail, error)
}
