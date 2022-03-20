package modelDB

import "github.com/volatiletech/null"

type User struct {
	ID        int64       `db:"id"`
	Email     string      `db:"email"`
	Password  string      `db:"password"`
	Status    int         `db:"status"`
	CreatedAt string      `db:"created_at"`
	CreatedBy string      `db:"created_by"`
	UpdatedAt null.String `db:"updated_at"`
	UpdatedBy null.String `db:"updated_by"`
	DeletedAt null.String `db:"deleted_at"`
	DeletedBy null.String `db:"deleted_by"`
}
