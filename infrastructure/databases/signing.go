package databases

import (
	"context"

	"github.com/alhamsya/boilerplate-go/domain/models/database"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
)

//CreateHistorySigning create history signing from request client
func (db *ServiceDB) CreateHistorySigning(ctx context.Context, reqDB *modelDB.Signing) (lastInsertID int64, err error) {
	q := `
		INSERT INTO movie_hst_signing (email, created_at, created_by)
		VALUES (:email, :created_at, :created_by)
	`
	result, err := db.db.Master.NamedExecContext(ctx, q, &reqDB)
	if err != nil {
		return -1, customError.WrapFlag(err, "sqlx", "NamedExecContext")
	}

	lastInsertID, err = result.LastInsertId()
	if err != nil {
		return -1, customError.WrapFlag(err, "sqlx", "LastInsertId")
	}

	return lastInsertID, nil
}
