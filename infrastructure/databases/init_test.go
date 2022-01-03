package databases

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alhamsya/boilerplate-go/lib/helpers/database"
	"github.com/jmoiron/sqlx"
	"log"
)

func setupMockDB() (mockDB *sql.DB, mockStore *database.Store, mockSQL sqlmock.Sqlmock) {
	mockDB, mockSQL, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	mockStore = &database.Store{
		Master: sqlxDB,
		Slave:  sqlxDB,
	}

	return mockDB, mockStore, mockSQL
}
