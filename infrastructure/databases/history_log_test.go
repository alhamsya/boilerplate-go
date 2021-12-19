package databases

import (
	"context"
	"errors"
	constCommon "github.com/alhamsya/boilerplate-go/domain/constants"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/alhamsya/boilerplate-go/lib/helpers/database"
	"github.com/volatiletech/null"
)

func TestDBService_CreateHistoryLog(t *testing.T) {
	someError := errors.New("some error | TestDBService_CreateHistoryLog")

	mockDB, mockStore, mockSQl := setupMockDB()
	defer mockDB.Close()

	type fields struct {
		DB *database.Store
	}
	type args struct {
		ctx   context.Context
		reqDB *modelMovie.DBHistoryLog
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		patch            func()
		wantLastInsertID int64
		wantErr          bool
	}{
		{
			name:   "When_NamedExecContextReturnError_expectError",
			fields: fields{},
			args: args{
				ctx:   context.TODO(),
				reqDB: &modelMovie.DBHistoryLog{
					Endpoint:   null.StringFrom("/api/movie"),
					Request:    "",
					Response:   "",
					SourceData: constCommon.TypeREST,
					CreatedAt:  "2021-12-19 21:55:48",
					CreatedBy:  "127.0.0.1",
				},
			},
			patch: func() {
				mockSQl.ExpectExec("INSERT").WithArgs(null.StringFrom("/api/movie"), "", "", constCommon.TypeREST, "2021-12-19 21:55:48", "127.0.0.1").WillReturnError(someError)
			},
			wantLastInsertID: -1,
			wantErr:          true,
		},
		{
			name:   "When_CreateHistoryLogReturnSuccess_expectSuccess",
			fields: fields{},
			args: args{
				ctx:   context.TODO(),
				reqDB: &modelMovie.DBHistoryLog{
					Endpoint:   null.StringFrom("/api/movie"),
					Request:    "",
					Response:   "",
					SourceData: constCommon.TypeREST,
					CreatedAt:  "2021-12-19 21:55:48",
					CreatedBy:  "127.0.0.1",
				},
			},
			patch: func() {
				mockSQl.ExpectExec("INSERT").WithArgs(null.StringFrom("/api/movie"), "", "", constCommon.TypeREST, "2021-12-19 21:55:48", "127.0.0.1").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantLastInsertID: 1,
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DBService{
				DB: mockStore,
			}
			tt.patch()
			gotLastInsertID, err := db.CreateHistoryLog(tt.args.ctx, tt.args.reqDB)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateHistoryLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLastInsertID != tt.wantLastInsertID {
				t.Errorf("CreateHistoryLog() gotLastInsertID = %v, want %v", gotLastInsertID, tt.wantLastInsertID)
			}
		})
	}
}
