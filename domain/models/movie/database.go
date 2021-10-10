package modelMovie

import (
	"github.com/volatiletech/null"
)

type DBHistoryLog struct {
	ID         int64       `db:"id"`
	Endpoint   null.String `db:"endpoint"`
	Request    string      `db:"request"`
	Response   string      `db:"response"`
	SourceData string      `db:"source_data"`
	CreatedAt  string      `db:"created_at"`
	CreatedBy  string      `db:"created_by"`
}
