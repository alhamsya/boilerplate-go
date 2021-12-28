package repository

import (
	"context"
	"github.com/jmoiron/sqlx"

	"github.com/alhamsya/boilerplate-go/domain/models/movie"
)

type DBRepo interface {
	TxBegin(ctx context.Context) (*sqlx.Tx, error)
	TxError(tx *sqlx.Tx, err error) error
	TxCommit(tx *sqlx.Tx) error

	CreateHistoryLog(ctx context.Context, reqDB *modelMovie.DBHistoryLog) (lastInsertID int64, err error)
}
