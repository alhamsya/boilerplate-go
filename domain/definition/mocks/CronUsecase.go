// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CronUsecase is an autogenerated mock type for the CronUsecase type
type CronUsecase struct {
	mock.Mock
}

// DoCreateDummyData provides a mock function with given fields: ctx
func (_m *CronUsecase) DoCreateDummyData(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
