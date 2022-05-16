package databases

import (
	"context"
	"database/sql"

	"github.com/alhamsya/boilerplate-go/domain/models/database"
	"github.com/alhamsya/boilerplate-go/lib/managers/custom_error"
	"github.com/jmoiron/sqlx"
)

//CreateUserTx create user from request client using transaction
func (db *ServiceDB) CreateUserTx(ctx context.Context, tx *sqlx.Tx, reqDB *modelDB.User) (lastInsertID int64, err error) {
	q := `
		INSERT INTO movie_mst_user (email, password, status, created_at, created_by)
		VALUES (:email, :password, :status, :created_at, :created_by)
	`
	result, err := tx.NamedExecContext(ctx, q, &reqDB)
	if err != nil {
		return -1, customError.WrapFlag(err, "sqlx", "NamedExecContext")
	}

	lastInsertID, err = result.LastInsertId()
	if err != nil {
		return -1, customError.WrapFlag(err, "sqlx", "LastInsertId")
	}

	return lastInsertID, nil
}

func (db *ServiceDB) GetUserByEmail(ctx context.Context, email string) (respDB *modelDB.User, err error) {
	q := `
		SELECT *
		FROM movie_mst_user u
		WHERE u.email=:email AND u.deleted_at IS NULL
	`

	q, args, err := sqlx.Named(q, map[string]interface{}{
		"email": email,
	})
	if err != nil {
		return nil, customError.WrapFlag(err, "sqlx", "Named")
	}

	respDB = new(modelDB.User)
	err = db.db.Slave.GetContext(ctx, respDB, q, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, customError.Wrap(err, "GetContext")
	}

	return respDB, nil
}
