package customError

import (
	"fmt"
	"github.com/friendsofgo/errors"
	"google.golang.org/grpc/codes"
	"strings"
)

func New(err error) *Error {
	return &Error{
		Err: err,
	}
}

func (e *Error) WrapFlag(flag string, args ...string) *Error {
	if flag = strings.TrimSpace(flag); flag != "" {
		flag = fmt.Sprintf("[%s] ", strings.ToUpper(flag))
	}

	msg := strings.Join(args, " | ")
	msg = flag + msg

	if e.Err == nil {
		e.Err = fmt.Errorf("[FOR DEVELOPER] forget to set error in the: %s", flag)
	}

	e.Err = errors.Wrap(e.Err, msg)

	return e
}

func (e *Error) Wrap(args ...string) *Error {
	msg := strings.Join(args, " | ")
	if e.Err == nil {
		e.Err = fmt.Errorf("[FOR DEVELOPER] forget to set error in the: %s", msg)
	}

	e.Err = errors.Wrap(e.Err, msg)
	return e
}

func (e *Error) SetGrpcCode(grpcCode codes.Code) *Error {
	e.GrpcCode = grpcCode
	return e
}

func (e *Error) SetHttpCode(httpCode int) *Error {
	e.HttpCode = httpCode
	return e
}
