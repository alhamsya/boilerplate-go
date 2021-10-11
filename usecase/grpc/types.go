package grpcUC

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
)

type UcInteractor struct {
	Cfg         *config.MainConfig
	ServiceRepo repository.ServiceRepo
	OMDBRepo    repository.OMDBRepo
}
