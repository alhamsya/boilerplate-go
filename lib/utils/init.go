package utils

import "github.com/alhamsya/boilerplate-go/domain/repository"

type thing struct {}

func New() repository.UtilsRepo {
	return &thing{}
}