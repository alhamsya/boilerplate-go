package restUc

import (
	"bou.ke/monkey"
	"errors"
	modelMovie "github.com/alhamsya/boilerplate-go/domain/models/movie"
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/alhamsya/boilerplate-go/lib/utils/datetime"
	"github.com/alhamsya/boilerplate-go/service/exter/omdb"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
	"net/http"
	"reflect"
	"testing"
)

func TestUcInteractor_DoGetListMovie(t *testing.T) {
	initMock()

	someError := errors.New("some error | TestUcInteractor_DoGetListMovie")

	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	type fields struct {
		Cfg         *config.MainConfig
		ServiceRepo repository.ServiceRepo
		OMDBRepo    repository.OMDBRepo
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
		patch func()
	}{
		{
			name:         "When_GetListMovieReturnError_expectError",
			fields:       fields{
				ServiceRepo: mockServiceRepo,
				OMDBRepo:    mockOMDBRepo,
			},
			args:         args{
				ctx:      &fiber.Ctx{},
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockOMDBRepo.On("GetListMovie", "tes", int64(1)).Return(nil, someError).Once()
			},
		},
		{
			name:         "When_GetListMovieReturnNil_expectError",
			fields:       fields{
				ServiceRepo: mockServiceRepo,
				OMDBRepo:    mockOMDBRepo,
			},
			args:         args{
				ctx:      &fiber.Ctx{},
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockOMDBRepo.On("GetListMovie", "tes", int64(1)).Return(nil, nil).Once()
			},
		},
		{
			name:         "When_ConvertResponseBoolReturnError_expectError",
			fields:       fields{
				ServiceRepo: mockServiceRepo,
				OMDBRepo:    mockOMDBRepo,
			},
			args:         args{
				ctx:      &fiber.Ctx{},
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockOMDBRepo.On("GetListMovie", "tes", int64(1)).Return(&omdb.OMDBList{
					Search:       nil,
					TotalResults: "",
					Response:     "#",
				}, nil).Once()
			},
		},
		{
			name:         "When_ConvertResponseBoolReturnFalse_expectError",
			fields:       fields{
				ServiceRepo: mockServiceRepo,
				OMDBRepo:    mockOMDBRepo,
			},
			args:         args{
				ctx:      &fiber.Ctx{},
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusBadRequest,
			wantErr:      true,
			patch: func() {
				mockOMDBRepo.On("GetListMovie", "tes", int64(1)).Return(&omdb.OMDBList{
					Search:       nil,
					TotalResults: "",
					Response:     "False",
				}, nil).Once()
			},
		},
		{
			name:         "When_ConvertTotalResultsReturnError_expectError",
			fields:       fields{
				ServiceRepo: mockServiceRepo,
				OMDBRepo:    mockOMDBRepo,
			},
			args:         args{
				ctx:      &fiber.Ctx{},
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockOMDBRepo.On("GetListMovie", "tes", int64(1)).Return(&omdb.OMDBList{
					Search:       []omdb.Search{
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
			name:         "When_CurrentTimeFReturnError_expectError",
			fields:       fields{
				ServiceRepo: mockServiceRepo,
				OMDBRepo:    mockOMDBRepo,
			},
			args:         args{
				ctx:      &fiber.Ctx{},
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockOMDBRepo.On("GetListMovie", "tes", int64(1)).Return(&omdb.OMDBList{
					Search:       []omdb.Search{
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

				var guard *monkey.PatchGuard
				guard = monkey.Patch(datetime.CurrentTimeF, func(format string) (string, error) {
					defer guard.Unpatch()
					return "", someError
				})
			},
		},
		{
			name:         "When_CreateHistoryLogReturnError_expectError",
			fields:       fields{
				ServiceRepo: mockServiceRepo,
				OMDBRepo:    mockOMDBRepo,
			},
			args:         args{
				ctx:      ctx,
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     nil,
			wantHttpCode: http.StatusInternalServerError,
			wantErr:      true,
			patch: func() {
				mockOMDBRepo.On("GetListMovie", "tes", int64(1)).Return(&omdb.OMDBList{
					Search:       []omdb.Search{
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

				var guard1, guard2 *monkey.PatchGuard
				guard1 = monkey.Patch(datetime.CurrentTimeF, func(format string) (string, error) {
					defer guard1.Unpatch()
					return "2021-01-01 15:04:05", nil
				})

				guard2 = monkey.PatchInstanceMethod(reflect.TypeOf(&fiber.Ctx{}), "IP", func(c *fiber.Ctx) string {
					defer guard2.Unpatch()
					return ""
				})

				mockServiceRepo.On("CreateHistoryLog", ctx.Context(), mock.Anything).Return(int64(1), someError).Once()
			},
		},
		{
			name:         "When_functionReturnSuccess_expectSuccess",
			fields:       fields{
				ServiceRepo: mockServiceRepo,
				OMDBRepo:    mockOMDBRepo,
			},
			args:         args{
				ctx:      ctx,
				reqClient: &modelMovie.ReqListMovie{
					Search: "tes",
					Page:   1,
				},
			},
			wantResp:     &modelMovie.RespListMovie{
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
				mockOMDBRepo.On("GetListMovie", "tes", int64(1)).Return(&omdb.OMDBList{
					Search:       []omdb.Search{
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

				var guard1, guard2 *monkey.PatchGuard
				guard1 = monkey.Patch(datetime.CurrentTimeF, func(format string) (string, error) {
					defer guard1.Unpatch()
					return "2021-01-01 15:04:05", nil
				})

				guard2 = monkey.PatchInstanceMethod(reflect.TypeOf(&fiber.Ctx{}), "IP", func(c *fiber.Ctx) string {
					defer guard2.Unpatch()
					return ""
				})

				mockServiceRepo.On("CreateHistoryLog", ctx.Context(), mock.Anything).Return(int64(1), nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &UcInteractor{
				Cfg:         tt.fields.Cfg,
				ServiceRepo: mockServiceRepo,
				OMDBRepo:    mockOMDBRepo,
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
