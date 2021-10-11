package datetime

import (
	"bou.ke/monkey"
	"errors"
	constCommon "github.com/alhamsya/boilerplate-go/domain/constants"
	"reflect"
	"testing"
	"time"
)

func TestCurrentTimeF(t *testing.T) {
	someError := errors.New("some error | TestCurrentTimeF")

	location, _ := time.LoadLocation(constCommon.TimeLocalJakarta)

	type args struct {
		format string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		patch func()
	}{
		{
			name:    "When_loadLocationReturnError_ExpectError",
			args:    args{
				format: constCommon.DateTime,
			},
			want:    "",
			wantErr: true,
			patch: func() {
				var guard1 *monkey.PatchGuard
				guard1 = monkey.Patch(time.LoadLocation, func(name string) (*time.Location, error) {
					guard1.Unpatch()
					return nil, someError
				})
			},
		},
		{
			name:    "When_CurrentTimeFReturnSuccess_expectSuccess",
			args:    args{
				format: constCommon.DateTime,
			},
			want:    "2012-12-05 15:00:00",
			wantErr: false,
			patch: func() {
				var guard1, guard2, guard3 *monkey.PatchGuard
				guard1 = monkey.Patch(time.LoadLocation, func(name string) (*time.Location, error) {
					guard1.Unpatch()
					location, _ := time.LoadLocation(constCommon.TimeLocalJakarta)
					return location, nil
				})

				guard2 = monkey.Patch(time.Now, func() time.Time {
					defer guard2.Unpatch()
					return time.Date(2012, time.December, 5, 15, 0, 0, 0, location)
				})

				guard3 = monkey.PatchInstanceMethod(reflect.TypeOf(time.Time{}), "In", func(t time.Time, loc *time.Location) time.Time {
					defer guard3.Unpatch()
					return time.Date(2012, time.December, 5, 15, 0, 0, 0, location)
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.patch()
			got, err := CurrentTimeF(tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("CurrentTimeF() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CurrentTimeF() got = %v, want %v", got, tt.want)
			}
		})
	}
}
