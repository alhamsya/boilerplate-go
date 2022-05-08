package restRouters

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		this *RestServer
	}
	tests := []struct {
		name string
		args args
		want *RestServer
	}{
		{
			name: "When_NewRouter_expectSuccess",
			args: args{
				this: &RestServer{},
			},
			want: &RestServer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.this); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
