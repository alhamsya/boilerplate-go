// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	service "github.com/alhamsya/boilerplate-go/protos"
)

// GrpcUsecase is an autogenerated mock type for the GrpcUsecase type
type GrpcUsecase struct {
	mock.Mock
}

// DoGetDetailMovie provides a mock function with given fields: ctx, req
func (_m *GrpcUsecase) DoGetDetailMovie(ctx context.Context, req *service.GetDetailMovieReq) (*service.GetDetailMovieResp, error) {
	ret := _m.Called(ctx, req)

	var r0 *service.GetDetailMovieResp
	if rf, ok := ret.Get(0).(func(context.Context, *service.GetDetailMovieReq) *service.GetDetailMovieResp); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.GetDetailMovieResp)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.GetDetailMovieReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DoGetListMovie provides a mock function with given fields: ctx, req
func (_m *GrpcUsecase) DoGetListMovie(ctx context.Context, req *service.GetListMovieReq) (*service.GetListMovieResp, error) {
	ret := _m.Called(ctx, req)

	var r0 *service.GetListMovieResp
	if rf, ok := ret.Get(0).(func(context.Context, *service.GetListMovieReq) *service.GetListMovieResp); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*service.GetListMovieResp)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *service.GetListMovieReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
