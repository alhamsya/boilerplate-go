package grpcUC

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
)

type UCInteractor struct {
	Cfg         *config.ServiceConfig
	DBRepo      repository.DBRepo
	CacheRepo   repository.CacheRepo
	OMDBRepo    repository.OMDBRepo
	CallWrapper repository.WrapperRepo
}
