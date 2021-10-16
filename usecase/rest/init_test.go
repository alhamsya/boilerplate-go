package restUC

import (
	"github.com/alhamsya/boilerplate-go/domain/repository/mocks"
	"reflect"
	"testing"
)

var mockServiceRepo *mocks.ServiceRepo
var mockOMDBRepo *mocks.OMDBRepo

func initMock()  {
	mockServiceRepo = new(mocks.ServiceRepo)
	mockOMDBRepo = new(mocks.OMDBRepo)
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
