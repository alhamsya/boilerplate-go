package utils

import "github.com/alhamsya/boilerplate-go/domain/repositorys"

type thing struct{}

func New() repositorys.UtilsRepo {
	return &thing{}
}
