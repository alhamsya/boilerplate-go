package restRouters

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
)

func TestRestServer_GetDetailMovie(t *testing.T) {
	initMock()

	someError := errors.New("some error | TestRestServer_GetDetailMovie")

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	type fields struct {
		Cfg            *config.ServiceConfig
		App            *fiber.App
		RestInteractor *RestInteractor
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		patch   func()
	}{
		{
			name: "When_DoGetDetailMovieReturnError_expectError",
			args: args{
				ctx: ctx,
			},
			wantErr: false,
			patch: func() {
				mockRestInterface.On("DoGetDetailMovie", mock.Anything, mock.Anything).Return(nil, http.StatusInternalServerError, someError).Once()

				var guard *monkey.PatchGuard
				guard = monkey.PatchInstanceMethod(reflect.TypeOf(&fiber.Ctx{}), "Params", func(c *fiber.Ctx, key string, defaultValue ...string) string {
					defer guard.Unpatch()
					return "1"
				})
			},
		},
		{
			name: "When_GetDetailMovieReturnSuccess_expectSuccess",
			args: args{
				ctx: ctx,
			},
			wantErr: false,
			patch: func() {
				mockRestInterface.On("DoGetDetailMovie", mock.Anything, mock.Anything).Return(&modelMovie.RespDetailMovie{
					Title: "One Day",
				}, http.StatusOK, nil).Once()

				var guard *monkey.PatchGuard
				guard = monkey.PatchInstanceMethod(reflect.TypeOf(&fiber.Ctx{}), "Params", func(c *fiber.Ctx, key string, defaultValue ...string) string {
					defer guard.Unpatch()
					return "1"
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rest := &RestServer{
				Cfg: tt.fields.Cfg,
				App: tt.fields.App,
				RestInteractor: &RestInteractor{
					RestInterface: mockRestInterface,
				},
			}
			tt.patch()
			if err := rest.GetDetailMovie(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("GetDetailMovie() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRestServer_GetListMovie(t *testing.T) {
	initMock()

	someError := errors.New("some error | TestRestServer_GetListMovie")

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	type fields struct {
		Cfg            *config.ServiceConfig
		App            *fiber.App
		RestInteractor *RestInteractor
	}
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		patch   func()
	}{
		{
			name: "When_convertQueryParamReturnError_expectError",
			args: args{
				ctx: ctx,
			},
			wantErr: false,
			patch: func() {
				ctx.Request().URI().SetQueryString("search=batman&page=1")

				var guard1 *monkey.PatchGuard
				guard1 = monkey.Patch(strconv.ParseInt, func(s string, base int, bitSize int) (i int64, err error) {
					defer guard1.Unpatch()
					return 0, someError
				})
			},
		},
		{
			name: "When_requestQuerySearchEmpty_expectError",
			args: args{
				ctx: ctx,
			},
			wantErr: false,
			patch: func() {
				ctx.Request().URI().SetQueryString("search=&page=1")

				var guard1 *monkey.PatchGuard
				guard1 = monkey.Patch(strconv.ParseInt, func(s string, base int, bitSize int) (i int64, err error) {
					defer guard1.Unpatch()
					return 1, nil
				})
			},
		},
		{
			name: "When_DoGetListMovieReturnError_expectError",
			args: args{
				ctx: ctx,
			},
			wantErr: false,
			patch: func() {
				ctx.Request().URI().SetQueryString("search=batman&page=1")

				var guard1 *monkey.PatchGuard
				guard1 = monkey.Patch(strconv.ParseInt, func(s string, base int, bitSize int) (i int64, err error) {
					defer guard1.Unpatch()
					return 1, nil
				})

				mockRestInterface.On("DoGetListMovie", ctx, mock.Anything).Return(nil, http.StatusInternalServerError, someError).Once()
			},
		},
		{
			name: "When_GetListMovieReturnSuccess_expectSuccess",
			args: args{
				ctx: ctx,
			},
			wantErr: false,
			patch: func() {
				ctx.Request().URI().SetQueryString("search=batman&page=1")

				var guard1 *monkey.PatchGuard
				guard1 = monkey.Patch(strconv.ParseInt, func(s string, base int, bitSize int) (i int64, err error) {
					defer guard1.Unpatch()
					return 1, nil
				})

				mockRestInterface.On("DoGetListMovie", ctx, mock.Anything).Return(&modelMovie.RespListMovie{
					Items: []modelMovie.Items{
						{
							Title:   "One Day",
						},
					},
					Total: 10,
				}, http.StatusOK, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rest := &RestServer{
				Cfg: tt.fields.Cfg,
				App: tt.fields.App,
				RestInteractor: &RestInteractor{
					RestInterface: mockRestInterface,
				},
			}
			tt.patch()
			if err := rest.GetListMovie(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("GetListMovie() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
