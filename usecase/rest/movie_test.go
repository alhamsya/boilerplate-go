package restUC

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/domain/models/request"
	"github.com/alhamsya/boilerplate-go/domain/models/response"
	"github.com/alhamsya/boilerplate-go/domain/repositories"
	"github.com/alhamsya/boilerplate-go/infrastructure/externals/omdb"
	"github.com/alhamsya/boilerplate-go/infrastructure/wrappers"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
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
		DBRepo      repositories.DBRepo
		CacheRepo   repositories.CacheRepo
		OMDBRepo    repositories.OMDBRepo
		WrapperRepo repositories.CallWrapperRepo
		UtilsRepo   repositories.UtilsRepo
	}
	type args struct {
		ctx       *fiber.Ctx
		reqClient *modelReq.ListMovie
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResp     *modelResp.ListMovie
		wantHttpCode int
		wantErr      bool
		patch        func()
	}{
		{
			name: "When_GetListMovieReturnError_expectError",
			args: args{
				ctx: ctx,
				reqClient: &modelReq.ListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.EXPECT().GetWrapper("omdb").Return(&wrappers.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.EXPECT().GetListMovie(mock.Anything, "tes", int64(1)).Return(nil, someError).Once()

				mockOMDBRepo.EXPECT().GetListMovie("tes", int64(1)).Return(nil, someError).Once()
			},
		},
		{
			name: "When_GetListMovieReturnNil_expectError",
			args: args{
				ctx: ctx,
				reqClient: &modelReq.ListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.EXPECT().GetWrapper("omdb").Return(&wrappers.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.EXPECT().GetListMovie(mock.Anything, "tes", int64(1)).Return(nil, nil).Once()
			},
		},
		{
			name: "When_GetListMovieResponseCannotConvertReturnError_expectError",
			args: args{
				ctx: ctx,
				reqClient: &modelReq.ListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusConflict,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.EXPECT().GetWrapper("omdb").Return(&wrappers.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.EXPECT().GetListMovie(mock.Anything, "tes", int64(1)).Return(&omdb.OMDBList{
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
				reqClient: &modelReq.ListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusBadRequest,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.EXPECT().GetWrapper("omdb").Return(&wrappers.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.EXPECT().GetListMovie(mock.Anything, "tes", int64(1)).Return(nil, someError).Once()

				mockOMDBRepo.EXPECT().GetListMovie("tes", int64(1)).Return(&omdb.OMDBList{
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
				reqClient: &modelReq.ListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.EXPECT().GetWrapper("omdb").Return(&wrappers.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.EXPECT().GetListMovie(mock.Anything, "tes", int64(1)).Return(&omdb.OMDBList{
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
				reqClient: &modelReq.ListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.EXPECT().GetWrapper("omdb").Return(&wrappers.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.EXPECT().GetListMovie(mock.Anything, "tes", int64(1)).Return(&omdb.OMDBList{
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

				mockUtilsRepo.EXPECT().CurrentTimeF(constCommon.DateTime).Return("", someError).Once()
			},
		},
		{
			name: "When_CreateHistoryLogReturnError_expectError",
			args: args{
				ctx: ctx,
				reqClient: &modelReq.ListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockCallWrapperRepo.EXPECT().GetWrapper("omdb").Return(&wrappers.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.EXPECT().GetListMovie(mock.Anything, "tes", int64(1)).Return(&omdb.OMDBList{
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

				mockUtilsRepo.EXPECT().CurrentTimeF(constCommon.DateTime).Return(constCommon.DateTime, nil).Once()

				mockServiceRepo.EXPECT().CreateHistoryLog(ctx.Context(), mock.Anything).Return(int64(1), someError).Once()
			},
		},
		{
			name: "When_DoGetListMovieReturnSuccess_expectSuccess",
			args: args{
				ctx: ctx,
				reqClient: &modelReq.ListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp: &modelResp.ListMovie{
				Items: []modelResp.Items{
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
				mockCallWrapperRepo.EXPECT().GetWrapper("omdb").Return(&wrappers.Wrapper{
					Err: errors.New("ignore"),
				}).Once()

				mockCacheRepo.EXPECT().GetListMovie(mock.Anything, "tes", int64(1)).Return(&omdb.OMDBList{
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

				mockUtilsRepo.EXPECT().CurrentTimeF(constCommon.DateTime).Return(constCommon.DateTime, nil).Once()

				mockServiceRepo.EXPECT().CreateHistoryLog(ctx.Context(), mock.Anything).Return(int64(1), nil).Once()
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
