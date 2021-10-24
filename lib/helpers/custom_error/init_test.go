package customError

import (
	"errors"
	"testing"
)

func TestWrap(t *testing.T) {
	type args struct {
		err  error
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "When_parameterNotError_expectSuccess",
			args:    args{
				err:  errors.New("some error"),
				args: []string{"args 1", "args 2"},
			},
			wantErr: true,
		},
		{
			name:    "When_parameterNil_expectSuccess",
			args:    args{
				err:  nil,
				args: []string{"args 1", "args 2"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Wrap(tt.args.err, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("Wrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWrapFlag(t *testing.T) {
	type args struct {
		err  error
		flag string
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "When_paramNotNil_expectSuccess",
			args:    args{
				err:  errors.New("some error"),
				flag: "database",
				args: []string{"args 1", "args 2"},
			},
			wantErr: true,
		},
		{
			name:    "When_paramNil_expectSuccess",
			args:    args{
				err:  nil,
				flag: "database",
				args: []string{"args 1", "args 2"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WrapFlag(tt.args.err, tt.args.flag, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("WrapFlag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCause(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "When_errorNotNil_expectSuccess",
			args:    args{
				err: WrapFlag(errors.New("some error"), "database", "args 1"),
			},
			wantErr: true,
		},
		{
			name:    "When_errorNil_expectSuccess",
			args:    args{
				err: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Cause(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("Cause() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}