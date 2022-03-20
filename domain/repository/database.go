package repository

import (
	"context"
	"github.com/alhamsya/boilerplate-go/domain/models/database"
	"github.com/jmoiron/sqlx"
)

type DBRepo interface {
	TxBegin(ctx context.Context) (*sqlx.Tx, error)
	TxError(tx *sqlx.Tx, err error) error
	TxCommit(tx *sqlx.Tx) error

	CreateUserTx(ctx context.Context, tx *sqlx.Tx, reqDB *modelDB.User) (lastInsertID int64, err error)
	GetUserByEmail(ctx context.Context, email string) (respDB *modelDB.User, err error)

	CreateHistoryLog(ctx context.Context, reqDB *modelDB.HistoryLog) (lastInsertID int64, err error)
	CreateHistoryLogTx(ctx context.Context, tx *sqlx.Tx, reqDB *modelDB.HistoryLog) (lastInsertID int64, err error)
}
