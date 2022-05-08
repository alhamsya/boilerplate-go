package repository

import (
	"context"

	"github.com/alhamsya/boilerplate-go/infrastructure/external/omdb"
)

type CacheRepo interface {
	SetListMovie(ctx context.Context, search string, page int64, req *external.OMDBList) (err error)
	GetListMovie(ctx context.Context, search string, page int64) (resp *external.OMDBList, err error)

	SetDetailMovie(ctx context.Context, movieID string, req *external.OMDBDetail) (err error)
	GetDetailMovie(ctx context.Context, movieID string) (resp *external.OMDBDetail, err error)

	IncrKYCByStatus(ctx context.Context, statusKYC int64) (err error)
	DecrKYCByStatus(ctx context.Context, statusKYC int64) (err error)
	SetKYCByStatus(ctx context.Context, statusKYC int64, value int) (err error)
	GetKYCByStatus(ctx context.Context, statusKYC int64) (total int, err error)
}
