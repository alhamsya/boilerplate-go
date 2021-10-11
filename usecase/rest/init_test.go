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
		c *UcInteractor
	}
	tests := []struct {
		name string
		args args
		want *UcInteractor
	}{
		{
			name: "When_init_expectSuccess",
			args: args{
				c: &UcInteractor{},
			},
			want: &UcInteractor{},
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
