package modelDB

type Signing struct {
	ID        int64  `db:"id"`
	Email     string `db:"email"`
	CreatedAt string `db:"created_at"`
	CreatedBy string `db:"created_by"`
}
