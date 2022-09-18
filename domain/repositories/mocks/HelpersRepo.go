// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// HelpersRepo is an autogenerated mock type for the HelpersRepo type
type HelpersRepo struct {
	mock.Mock
}

type HelpersRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *HelpersRepo) EXPECT() *HelpersRepo_Expecter {
	return &HelpersRepo_Expecter{mock: &_m.Mock}
}

// Tes provides a mock function with given fields: ctx
func (_m *HelpersRepo) Tes(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// HelpersRepo_Tes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Tes'
type HelpersRepo_Tes_Call struct {
	*mock.Call
}

// Tes is a helper method to define mock.On call
//   - ctx context.Context
func (_e *HelpersRepo_Expecter) Tes(ctx interface{}) *HelpersRepo_Tes_Call {
	return &HelpersRepo_Tes_Call{Call: _e.mock.On("Tes", ctx)}
}

func (_c *HelpersRepo_Tes_Call) Run(run func(ctx context.Context)) *HelpersRepo_Tes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *HelpersRepo_Tes_Call) Return(_a0 bool) *HelpersRepo_Tes_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewHelpersRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewHelpersRepo creates a new instance of HelpersRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHelpersRepo(t mockConstructorTestingTNewHelpersRepo) *HelpersRepo {
	mock := &HelpersRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
