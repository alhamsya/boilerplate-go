package utils

import (
	"github.com/alhamsya/boilerplate-go/domain/repositories"
)

type thing struct{}

func New() repositories.UtilsRepo {
	return &thing{}
}
