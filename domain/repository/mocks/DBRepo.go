// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	modelDB "github.com/alhamsya/boilerplate-go/domain/models/database"
	mock "github.com/stretchr/testify/mock"

	sqlx "github.com/jmoiron/sqlx"
)

// DBRepo is an autogenerated mock type for the DBRepo type
type DBRepo struct {
	mock.Mock
}

// CreateHistoryLog provides a mock function with given fields: ctx, reqDB
func (_m *DBRepo) CreateHistoryLog(ctx context.Context, reqDB *modelDB.HistoryLog) (int64, error) {
	ret := _m.Called(ctx, reqDB)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *modelDB.HistoryLog) int64); ok {
		r0 = rf(ctx, reqDB)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *modelDB.HistoryLog) error); ok {
		r1 = rf(ctx, reqDB)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateHistoryLogTx provides a mock function with given fields: ctx, tx, reqDB
func (_m *DBRepo) CreateHistoryLogTx(ctx context.Context, tx *sqlx.Tx, reqDB *modelDB.HistoryLog) (int64, error) {
	ret := _m.Called(ctx, tx, reqDB)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, *modelDB.HistoryLog) int64); ok {
		r0 = rf(ctx, tx, reqDB)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.Tx, *modelDB.HistoryLog) error); ok {
		r1 = rf(ctx, tx, reqDB)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUserTx provides a mock function with given fields: ctx, tx, reqDB
func (_m *DBRepo) CreateUserTx(ctx context.Context, tx *sqlx.Tx, reqDB *modelDB.User) (int64, error) {
	ret := _m.Called(ctx, tx, reqDB)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *sqlx.Tx, *modelDB.User) int64); ok {
		r0 = rf(ctx, tx, reqDB)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *sqlx.Tx, *modelDB.User) error); ok {
		r1 = rf(ctx, tx, reqDB)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *DBRepo) GetUserByEmail(ctx context.Context, email string) (*modelDB.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *modelDB.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *modelDB.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*modelDB.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxBegin provides a mock function with given fields: ctx
func (_m *DBRepo) TxBegin(ctx context.Context) (*sqlx.Tx, error) {
	ret := _m.Called(ctx)

	var r0 *sqlx.Tx
	if rf, ok := ret.Get(0).(func(context.Context) *sqlx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Tx)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxCommit provides a mock function with given fields: tx
func (_m *DBRepo) TxCommit(tx *sqlx.Tx) error {
	ret := _m.Called(tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sqlx.Tx) error); ok {
		r0 = rf(tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TxError provides a mock function with given fields: tx, err
func (_m *DBRepo) TxError(tx *sqlx.Tx, err error) error {
	ret := _m.Called(tx, err)

	var r0 error
	if rf, ok := ret.Get(0).(func(*sqlx.Tx, error) error); ok {
		r0 = rf(tx, err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
