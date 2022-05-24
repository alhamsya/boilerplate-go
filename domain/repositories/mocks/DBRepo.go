// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	modelDB "github.com/alhamsya/boilerplate-go/domain/models/database"
	mock "github.com/stretchr/testify/mock"

	sqlx "github.com/jmoiron/sqlx"

	testing "testing"
)

// DBRepo is an autogenerated mock type for the DBRepo type
type DBRepo struct {
	mock.Mock
}

type DBRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *DBRepo) EXPECT() *DBRepo_Expecter {
	return &DBRepo_Expecter{mock: &_m.Mock}
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

// DBRepo_CreateHistoryLog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateHistoryLog'
type DBRepo_CreateHistoryLog_Call struct {
	*mock.Call
}

// CreateHistoryLog is a helper method to define mock.On call
//  - ctx context.Context
//  - reqDB *modelDB.HistoryLog
func (_e *DBRepo_Expecter) CreateHistoryLog(ctx interface{}, reqDB interface{}) *DBRepo_CreateHistoryLog_Call {
	return &DBRepo_CreateHistoryLog_Call{Call: _e.mock.On("CreateHistoryLog", ctx, reqDB)}
}

func (_c *DBRepo_CreateHistoryLog_Call) Run(run func(ctx context.Context, reqDB *modelDB.HistoryLog)) *DBRepo_CreateHistoryLog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*modelDB.HistoryLog))
	})
	return _c
}

func (_c *DBRepo_CreateHistoryLog_Call) Return(lastInsertID int64, err error) *DBRepo_CreateHistoryLog_Call {
	_c.Call.Return(lastInsertID, err)
	return _c
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

// DBRepo_CreateHistoryLogTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateHistoryLogTx'
type DBRepo_CreateHistoryLogTx_Call struct {
	*mock.Call
}

// CreateHistoryLogTx is a helper method to define mock.On call
//  - ctx context.Context
//  - tx *sqlx.Tx
//  - reqDB *modelDB.HistoryLog
func (_e *DBRepo_Expecter) CreateHistoryLogTx(ctx interface{}, tx interface{}, reqDB interface{}) *DBRepo_CreateHistoryLogTx_Call {
	return &DBRepo_CreateHistoryLogTx_Call{Call: _e.mock.On("CreateHistoryLogTx", ctx, tx, reqDB)}
}

func (_c *DBRepo_CreateHistoryLogTx_Call) Run(run func(ctx context.Context, tx *sqlx.Tx, reqDB *modelDB.HistoryLog)) *DBRepo_CreateHistoryLogTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*sqlx.Tx), args[2].(*modelDB.HistoryLog))
	})
	return _c
}

func (_c *DBRepo_CreateHistoryLogTx_Call) Return(lastInsertID int64, err error) *DBRepo_CreateHistoryLogTx_Call {
	_c.Call.Return(lastInsertID, err)
	return _c
}

// CreateHistorySigning provides a mock function with given fields: ctx, reqDB
func (_m *DBRepo) CreateHistorySigning(ctx context.Context, reqDB *modelDB.Signing) (int64, error) {
	ret := _m.Called(ctx, reqDB)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *modelDB.Signing) int64); ok {
		r0 = rf(ctx, reqDB)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *modelDB.Signing) error); ok {
		r1 = rf(ctx, reqDB)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DBRepo_CreateHistorySigning_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateHistorySigning'
type DBRepo_CreateHistorySigning_Call struct {
	*mock.Call
}

// CreateHistorySigning is a helper method to define mock.On call
//  - ctx context.Context
//  - reqDB *modelDB.Signing
func (_e *DBRepo_Expecter) CreateHistorySigning(ctx interface{}, reqDB interface{}) *DBRepo_CreateHistorySigning_Call {
	return &DBRepo_CreateHistorySigning_Call{Call: _e.mock.On("CreateHistorySigning", ctx, reqDB)}
}

func (_c *DBRepo_CreateHistorySigning_Call) Run(run func(ctx context.Context, reqDB *modelDB.Signing)) *DBRepo_CreateHistorySigning_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*modelDB.Signing))
	})
	return _c
}

func (_c *DBRepo_CreateHistorySigning_Call) Return(lastInsertID int64, err error) *DBRepo_CreateHistorySigning_Call {
	_c.Call.Return(lastInsertID, err)
	return _c
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

// DBRepo_CreateUserTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUserTx'
type DBRepo_CreateUserTx_Call struct {
	*mock.Call
}

// CreateUserTx is a helper method to define mock.On call
//  - ctx context.Context
//  - tx *sqlx.Tx
//  - reqDB *modelDB.User
func (_e *DBRepo_Expecter) CreateUserTx(ctx interface{}, tx interface{}, reqDB interface{}) *DBRepo_CreateUserTx_Call {
	return &DBRepo_CreateUserTx_Call{Call: _e.mock.On("CreateUserTx", ctx, tx, reqDB)}
}

func (_c *DBRepo_CreateUserTx_Call) Run(run func(ctx context.Context, tx *sqlx.Tx, reqDB *modelDB.User)) *DBRepo_CreateUserTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*sqlx.Tx), args[2].(*modelDB.User))
	})
	return _c
}

func (_c *DBRepo_CreateUserTx_Call) Return(lastInsertID int64, err error) *DBRepo_CreateUserTx_Call {
	_c.Call.Return(lastInsertID, err)
	return _c
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

// DBRepo_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type DBRepo_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//  - ctx context.Context
//  - email string
func (_e *DBRepo_Expecter) GetUserByEmail(ctx interface{}, email interface{}) *DBRepo_GetUserByEmail_Call {
	return &DBRepo_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", ctx, email)}
}

func (_c *DBRepo_GetUserByEmail_Call) Run(run func(ctx context.Context, email string)) *DBRepo_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *DBRepo_GetUserByEmail_Call) Return(respDB *modelDB.User, err error) *DBRepo_GetUserByEmail_Call {
	_c.Call.Return(respDB, err)
	return _c
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

// DBRepo_TxBegin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TxBegin'
type DBRepo_TxBegin_Call struct {
	*mock.Call
}

// TxBegin is a helper method to define mock.On call
//  - ctx context.Context
func (_e *DBRepo_Expecter) TxBegin(ctx interface{}) *DBRepo_TxBegin_Call {
	return &DBRepo_TxBegin_Call{Call: _e.mock.On("TxBegin", ctx)}
}

func (_c *DBRepo_TxBegin_Call) Run(run func(ctx context.Context)) *DBRepo_TxBegin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *DBRepo_TxBegin_Call) Return(_a0 *sqlx.Tx, _a1 error) *DBRepo_TxBegin_Call {
	_c.Call.Return(_a0, _a1)
	return _c
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

// DBRepo_TxCommit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TxCommit'
type DBRepo_TxCommit_Call struct {
	*mock.Call
}

// TxCommit is a helper method to define mock.On call
//  - tx *sqlx.Tx
func (_e *DBRepo_Expecter) TxCommit(tx interface{}) *DBRepo_TxCommit_Call {
	return &DBRepo_TxCommit_Call{Call: _e.mock.On("TxCommit", tx)}
}

func (_c *DBRepo_TxCommit_Call) Run(run func(tx *sqlx.Tx)) *DBRepo_TxCommit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*sqlx.Tx))
	})
	return _c
}

func (_c *DBRepo_TxCommit_Call) Return(_a0 error) *DBRepo_TxCommit_Call {
	_c.Call.Return(_a0)
	return _c
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

// DBRepo_TxError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TxError'
type DBRepo_TxError_Call struct {
	*mock.Call
}

// TxError is a helper method to define mock.On call
//  - tx *sqlx.Tx
//  - err error
func (_e *DBRepo_Expecter) TxError(tx interface{}, err interface{}) *DBRepo_TxError_Call {
	return &DBRepo_TxError_Call{Call: _e.mock.On("TxError", tx, err)}
}

func (_c *DBRepo_TxError_Call) Run(run func(tx *sqlx.Tx, err error)) *DBRepo_TxError_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*sqlx.Tx), args[1].(error))
	})
	return _c
}

func (_c *DBRepo_TxError_Call) Return(_a0 error) *DBRepo_TxError_Call {
	_c.Call.Return(_a0)
	return _c
}

// NewDBRepo creates a new instance of DBRepo. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewDBRepo(t testing.TB) *DBRepo {
	mock := &DBRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
