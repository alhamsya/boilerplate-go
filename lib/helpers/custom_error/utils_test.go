package customError

import (
	"bou.ke/monkey"
	"runtime"
	"testing"
)

func Test_trace(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		want1 int
		want2 string
		patch func()
	}{
		{
			name:  "When_callerReturnError_expectError",
			want:  "?",
			want1: 0,
			want2: "?",
			patch: func() {
				var guard1 *monkey.PatchGuard
				guard1 = monkey.Patch(runtime.Caller, func(skip int) (pc uintptr, file string, line int, ok bool) {
					defer guard1.Unpatch()
					return 0, "", 0, false
				})
			},
		},
		{
			name:  "When_funcForPCReturnError_expectError",
			want:  "",
			want1: 0,
			want2: "?",
			patch: func() {
				var guard1, guard2 *monkey.PatchGuard
				guard1 = monkey.Patch(runtime.Caller, func(skip int) (pc uintptr, file string, line int, ok bool) {
					defer guard1.Unpatch()
					return 0, "", 0, true
				})

				guard2 = monkey.Patch(runtime.FuncForPC, func(pc uintptr) *runtime.Func {
					defer guard2.Unpatch()
					return nil
				})
			},
		},
		{
			name:  "When_trace_expectSuccess",
			want:  "",
			want1: 10,
			want2: "",
			patch: func() {
				var guard1, guard2 *monkey.PatchGuard
				guard1 = monkey.Patch(runtime.Caller, func(skip int) (pc uintptr, file string, line int, ok bool) {
					defer guard1.Unpatch()
					return 1, "", 10, true
				})

				guard2 = monkey.Patch(runtime.FuncForPC, func(pc uintptr) *runtime.Func {
					defer guard2.Unpatch()
					return &runtime.Func{}
				})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.patch()
			got, got1, got2 := trace()
			if got != tt.want {
				t.Errorf("trace() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("trace() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("trace() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
