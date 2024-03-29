// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// FirestoreRepo is an autogenerated mock type for the FirestoreRepo type
type FirestoreRepo struct {
	mock.Mock
}

type FirestoreRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *FirestoreRepo) EXPECT() *FirestoreRepo_Expecter {
	return &FirestoreRepo_Expecter{mock: &_m.Mock}
}

// GetAccount provides a mock function with given fields: ctx
func (_m *FirestoreRepo) GetAccount(ctx context.Context) (interface{}, error) {
	ret := _m.Called(ctx)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context) interface{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
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

// FirestoreRepo_GetAccount_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccount'
type FirestoreRepo_GetAccount_Call struct {
	*mock.Call
}

// GetAccount is a helper method to define mock.On call
//   - ctx context.Context
func (_e *FirestoreRepo_Expecter) GetAccount(ctx interface{}) *FirestoreRepo_GetAccount_Call {
	return &FirestoreRepo_GetAccount_Call{Call: _e.mock.On("GetAccount", ctx)}
}

func (_c *FirestoreRepo_GetAccount_Call) Run(run func(ctx context.Context)) *FirestoreRepo_GetAccount_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *FirestoreRepo_GetAccount_Call) Return(oa interface{}, err error) *FirestoreRepo_GetAccount_Call {
	_c.Call.Return(oa, err)
	return _c
}

type mockConstructorTestingTNewFirestoreRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewFirestoreRepo creates a new instance of FirestoreRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFirestoreRepo(t mockConstructorTestingTNewFirestoreRepo) *FirestoreRepo {
	mock := &FirestoreRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
