package repository

import (
	omdb2 "github.com/alhamsya/boilerplate-go/service/exter/omdb"
)

type OMDBRepo interface {
	GetListMovie(search string, page int64) (*omdb2.OMDBList, error)
	GetDetailMovie(movieID string) (*omdb2.OMDBDetail, error)
}
