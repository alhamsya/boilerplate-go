// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	fiber "github.com/gofiber/fiber/v2"
	mock "github.com/stretchr/testify/mock"

	modelReq "github.com/alhamsya/boilerplate-go/domain/models/request"

	modelResp "github.com/alhamsya/boilerplate-go/domain/models/response"
)

// RestUsecase is an autogenerated mock type for the RestUsecase type
type RestUsecase struct {
	mock.Mock
}

// DoCreateUser provides a mock function with given fields: ctx, reqClient
func (_m *RestUsecase) DoCreateUser(ctx *fiber.Ctx, reqClient *modelReq.User) (*modelResp.User, int, error) {
	ret := _m.Called(ctx, reqClient)

	var r0 *modelResp.User
	if rf, ok := ret.Get(0).(func(*fiber.Ctx, *modelReq.User) *modelResp.User); ok {
		r0 = rf(ctx, reqClient)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*modelResp.User)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*fiber.Ctx, *modelReq.User) int); ok {
		r1 = rf(ctx, reqClient)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*fiber.Ctx, *modelReq.User) error); ok {
		r2 = rf(ctx, reqClient)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DoGetDetailMovie provides a mock function with given fields: ctx, movieID
func (_m *RestUsecase) DoGetDetailMovie(ctx *fiber.Ctx, movieID string) (*modelResp.DetailMovie, int, error) {
	ret := _m.Called(ctx, movieID)

	var r0 *modelResp.DetailMovie
	if rf, ok := ret.Get(0).(func(*fiber.Ctx, string) *modelResp.DetailMovie); ok {
		r0 = rf(ctx, movieID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*modelResp.DetailMovie)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*fiber.Ctx, string) int); ok {
		r1 = rf(ctx, movieID)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*fiber.Ctx, string) error); ok {
		r2 = rf(ctx, movieID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// DoGetListMovie provides a mock function with given fields: ctx, reqClient
func (_m *RestUsecase) DoGetListMovie(ctx *fiber.Ctx, reqClient *modelReq.ListMovie) (*modelResp.ListMovie, int, error) {
	ret := _m.Called(ctx, reqClient)

	var r0 *modelResp.ListMovie
	if rf, ok := ret.Get(0).(func(*fiber.Ctx, *modelReq.ListMovie) *modelResp.ListMovie); ok {
		r0 = rf(ctx, reqClient)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*modelResp.ListMovie)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*fiber.Ctx, *modelReq.ListMovie) int); ok {
		r1 = rf(ctx, reqClient)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*fiber.Ctx, *modelReq.ListMovie) error); ok {
		r2 = rf(ctx, reqClient)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
