package databases

import (
	"context"

	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
	"github.com/jmoiron/sqlx"
)

func (db *ServiceDB) TxBegin(ctx context.Context) (*sqlx.Tx, error) {
	//initialize transaction
	tx, err := db.DB.Master.BeginTxx(ctx, nil)
	if err != nil {
		return nil, customError.Wrap(err, "Beginx")
	}

	return tx, nil
}

func (db *ServiceDB) TxError(tx *sqlx.Tx, err error) error {
	// skip if not error
	if err == nil {
		return nil
	}

	if errTx := tx.Rollback(); errTx != nil {
		// return error rollback
		return customError.Wrap(err, "Rollback")
	}

	// return actual error
	return err
}

func (db *ServiceDB) TxCommit(tx *sqlx.Tx) error {
	//commit transaction
	if err := tx.Commit(); err != nil {
		return customError.Wrap(err, "Commit")
	}
	return nil
}

