package repository

import (
	"context"
	"github.com/alhamsya/boilerplate-go/domain/models/movie"
)

type ServiceRepo interface {
	CreateHistoryLog(ctx context.Context, reqDB *modelMovie.DBHistoryLog) (lastInsertID int64, err error)
}
