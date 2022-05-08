package repository

import "github.com/alhamsya/boilerplate-go/infrastructure/wrappers"

type CallWrapperRepo interface {
	GetWrapper(usecase string) *wrappers.Wrapper
	Call(req func() (interface{}, error)) (interface{}, error)
}
