package repository

import (
	"context"

	"github.com/alhamsya/boilerplate-go/service/exter/omdb"
)

type CacheRepo interface {
	SetListMovie(ctx context.Context, search string, page int64, req *omdb.OMDBList) (err error)
	GetListMovie(ctx context.Context, search string, page int64) (resp *omdb.OMDBList, err error)

	SetDetailMovie(ctx context.Context, movieID string, req *omdb.OMDBDetail) (err error)
	GetDetailMovie(ctx context.Context, movieID string) (resp *omdb.OMDBDetail, err error)
}
