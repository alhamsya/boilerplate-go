package databases

import (
	"database/sql"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alhamsya/boilerplate-go/lib/helpers/database"
	"github.com/jmoiron/sqlx"
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

func TestNew(t *testing.T) {
	type args struct {
		this *ServiceDB
	}
	tests := []struct {
		name string
		args args
		want *ServiceDB
	}{
		{
			name: "When_init_expectSuccess",
			args: args{
				this: &ServiceDB{},
			},
			want: &ServiceDB{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.this); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
