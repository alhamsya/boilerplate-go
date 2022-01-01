package restUC

import (
	"github.com/alhamsya/boilerplate-go/domain/repository/mocks"
	"reflect"
	"testing"
)

var mockServiceRepo *mocks.DBRepo
var mockOMDBRepo *mocks.OMDBRepo
var mockUtilsRepo *mocks.UtilsRepo
var mockWrapperRepo *mocks.WrapperRepo
var mockCacheRepo *mocks.CacheRepo

func initMock()  {
	mockServiceRepo = new(mocks.DBRepo)
	mockOMDBRepo = new(mocks.OMDBRepo)
	mockUtilsRepo = new(mocks.UtilsRepo)
	mockWrapperRepo = new(mocks.WrapperRepo)
	mockCacheRepo = new(mocks.CacheRepo)
}

func TestNew(t *testing.T) {
	type args struct {
		c *UCInteractor
	}
	tests := []struct {
		name string
		args args
		want *UCInteractor
	}{
		{
			name: "When_init_expectSuccess",
			args: args{
				c: &UCInteractor{},
			},
			want: &UCInteractor{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
