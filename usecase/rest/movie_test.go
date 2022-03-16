package restUC

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/alhamsya/boilerplate-go/infrastructure/wrapper"
	"github.com/alhamsya/boilerplate-go/transport/exter/omdb"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
)

func TestUcInteractor_DoGetListMovie(t *testing.T) {
	initMock()

	someError := errors.New("some error | TestUcInteractor_DoGetListMovie")

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	type fields struct {
		Cfg         *config.ServiceConfig
		DBRepo      repository.DBRepo
		CacheRepo   repository.CacheRepo
		OMDBRepo    repository.OMDBRepo
		WrapperRepo repository.CallWrapperRepo
		UtilsRepo   repository.UtilsRepo
	}
	type args struct {
		ctx       *fiber.Ctx
		reqClient *modelMovie.ReqListMovie
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResp     *modelMovie.RespListMovie
		wantHttpCode int
		wantErr      bool
		patch        func()
	}{
		{
			name: "When_GetListMovieReturnError_expectError",
			args: args{
				ctx: ctx,
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.On("GetWrapper", "omdb").Return(&wrapper.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.On("GetListMovie", mock.Anything, "tes", int64(1)).Return(nil, someError).Once()

				mockOMDBRepo.On("GetListMovie", "tes", int64(1)).Return(nil, someError).Once()
			},
		},
		{
			name: "When_GetListMovieReturnNil_expectError",
			args: args{
				ctx: ctx,
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.On("GetWrapper", "omdb").Return(&wrapper.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.On("GetListMovie", mock.Anything, "tes", int64(1)).Return(nil, nil).Once()
			},
		},
		{
			name: "When_GetListMovieResponseCannotConvertReturnError_expectError",
			args: args{
				ctx: ctx,
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusConflict,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.On("GetWrapper", "omdb").Return(&wrapper.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.On("GetListMovie", mock.Anything, "tes", int64(1)).Return(&omdb.OMDBList{
					Search: []omdb.Search{
						{
							Title:  "One Day",
							Year:   "2021",
							ImdbID: "alham123",
							Type:   "Film",
							Poster: "www.google.com",
						},
					},
					Response: "#",
				}, nil).Once()
			},
		},
		{
			name: "When_GetListMovieResponseReturnFalse_expectError",
			args: args{
				ctx: ctx,
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusBadRequest,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.On("GetWrapper", "omdb").Return(&wrapper.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.On("GetListMovie", mock.Anything, "tes", int64(1)).Return(nil, someError).Once()

				mockOMDBRepo.On("GetListMovie", "tes", int64(1)).Return(&omdb.OMDBList{
					Search: []omdb.Search{
						{
							Title:  "One Day",
							Year:   "2021",
							ImdbID: "alham123",
							Type:   "Film",
							Poster: "www.google.com",
						},
					},
					Response: "False",
				}, nil).Once()

				mockCacheRepo.On("SetListMovie", mock.Anything, "tes", int64(1), mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "When_ConvertTotalResultsReturnError_expectError",
			args: args{
				ctx: &fiber.Ctx{},
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.On("GetWrapper", "omdb").Return(&wrapper.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.On("GetListMovie", mock.Anything, "tes", int64(1)).Return(&omdb.OMDBList{
					Search: []omdb.Search{
						{
							Title:  "One Day",
							Year:   "2021",
							ImdbID: "alham123",
							Type:   "Film",
							Poster: "www.google.com",
						},
					},
					TotalResults: "#",
					Response:     "True",
				}, nil).Once()
			},
		},
		{
			name: "When_CurrentTimeFReturnError_expectError",
			args: args{
				ctx: ctx,
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.On("GetWrapper", "omdb").Return(&wrapper.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.On("GetListMovie", mock.Anything, "tes", int64(1)).Return(&omdb.OMDBList{
					Search: []omdb.Search{
						{
							Title:  "One Day",
							Year:   "2021",
							ImdbID: "alham123",
							Type:   "Film",
							Poster: "www.google.com",
						},
					},
					TotalResults: "100",
					Response:     "True",
				}, nil).Once()

				mockUtilsRepo.On("CurrentTimeF", constCommon.DateTime).Return("", someError).Once()
			},
		},
		{
			name: "When_CreateHistoryLogReturnError_expectError",
			args: args{
				ctx: ctx,
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.On("GetWrapper", "omdb").Return(&wrapper.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.On("GetListMovie", mock.Anything, "tes", int64(1)).Return(&omdb.OMDBList{
					Search: []omdb.Search{
						{
							Title:  "One Day",
							Year:   "2021",
							ImdbID: "alham123",
							Type:   "Film",
							Poster: "www.google.com",
						},
					},
					TotalResults: "100",
					Response:     "True",
				}, nil).Once()

				mockUtilsRepo.On("CurrentTimeF", constCommon.DateTime).Return(constCommon.DateTime, nil).Once()

				mockServiceRepo.On("CreateHistoryLog", ctx.Context(), mock.Anything).Return(int64(1), someError).Once()
			},
		},
		{
			name: "When_DoGetListMovieReturnSuccess_expectSuccess",
			args: args{
				ctx: ctx,
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp: &modelMovie.RespListMovie{
				Items: []modelMovie.Items{
					{
						Title:   "One Day",
						Year:    "2021",
						MovieID: "alham123",
						Types:   "Film",
						Poster:  "www.google.com",
					},
				},
				Total: 100,
			},
			wantHttpCode: http.StatusOK,
			wantErr:      false,
			patch: func() {
				mockCallWrapperRepo.On("GetWrapper", "omdb").Return(&wrapper.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.On("GetListMovie", mock.Anything, "tes", int64(1)).Return(&omdb.OMDBList{
					Search: []omdb.Search{
						{
							Title:  "One Day",
							Year:   "2021",
							ImdbID: "alham123",
							Type:   "Film",
							Poster: "www.google.com",
						},
					},
					TotalResults: "100",
					Response:     "True",
				}, nil).Once()

				mockUtilsRepo.On("CurrentTimeF", constCommon.DateTime).Return(constCommon.DateTime, nil).Once()

				mockServiceRepo.On("CreateHistoryLog", ctx.Context(), mock.Anything).Return(int64(1), nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &UCInteractor{
				Cfg:             tt.fields.Cfg,
				DBRepo:          mockServiceRepo,
				OMDBRepo:        mockOMDBRepo,
				CallWrapperRepo: mockCallWrapperRepo,
				UtilsRepo:       mockUtilsRepo,
				CacheRepo:       mockCacheRepo,
			}
			tt.patch()
			gotResp, gotHttpCode, err := uc.DoGetListMovie(tt.args.ctx, tt.args.reqClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoGetListMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("DoGetListMovie() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
			if gotHttpCode != tt.wantHttpCode {
				t.Errorf("DoGetListMovie() gotHttpCode = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}
		})
	}
}
