package repository

import "github.com/alhamsya/boilerplate-go/infrastructure/wrapper"

type WrapperRepo interface {
	GetWrapper(usecase string) *wrapper.Wrapper
}
