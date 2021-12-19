package databases

import (
	"context"

	"github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
)

//CreateHistoryLog create history log from request client
func (db *DBService) CreateHistoryLog(ctx context.Context, reqDB *modelMovie.DBHistoryLog) (lastInsertID int64, err error) {
	q := `
		INSERT INTO hst_log (endpoint, request, response, source_data, created_at, created_by)
		VALUES (:endpoint, :request, :response, :source_data, :created_at, :created_by)
	`
	result, err := db.DB.Master.NamedExecContext(ctx, q, &reqDB)
	if err != nil {
		return -1, customError.WrapFlag(err, "sqlx", "NamedExecContext")
	}

	lastInsertID, err = result.LastInsertId()
	if err != nil {
		return -1, customError.WrapFlag(err, "sqlx", "LastInsertId")
	}

	return lastInsertID, nil
}
