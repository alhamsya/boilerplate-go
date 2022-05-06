package repository

import "github.com/alhamsya/boilerplate-go/infrastructure/wrapper"

type CallWrapperRepo interface {
	GetWrapper(usecase string) *wrapper.Wrapper
	Call(req func() (interface{}, error)) (interface{}, error)
}
