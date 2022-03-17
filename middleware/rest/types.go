package restMiddleware

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
)

type Middleware struct {
	Cfg       *config.ServiceConfig
	DBRepo    repository.DBRepo
	UtilsRepo repository.UtilsRepo
}
