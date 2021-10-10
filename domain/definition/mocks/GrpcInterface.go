// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	proto "github.com/alhamsya/boilerplate-go/protos"
)

// GrpcInterface is an autogenerated mock type for the GrpcInterface type
type GrpcInterface struct {
	mock.Mock
}

// DoGetDetailMovie provides a mock function with given fields: ctx, req
func (_m *GrpcInterface) DoGetDetailMovie(ctx context.Context, req *proto.GetDetailMovieReq) (*proto.DataDetailMovie, error) {
	ret := _m.Called(ctx, req)

	var r0 *proto.DataDetailMovie
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetDetailMovieReq) *proto.DataDetailMovie); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.DataDetailMovie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.GetDetailMovieReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DoGetListMovie provides a mock function with given fields: ctx, req
func (_m *GrpcInterface) DoGetListMovie(ctx context.Context, req *proto.GetListMovieReq) (*proto.GetListMovieResp, error) {
	ret := _m.Called(ctx, req)

	var r0 *proto.GetListMovieResp
	if rf, ok := ret.Get(0).(func(context.Context, *proto.GetListMovieReq) *proto.GetListMovieResp); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.GetListMovieResp)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.GetListMovieReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}