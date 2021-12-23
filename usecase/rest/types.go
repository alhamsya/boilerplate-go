package restUC

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
)

type UCInteractor struct {
	Cfg         *config.MainConfig
	ServiceRepo repository.ServiceRepo
	OMDBRepo    repository.OMDBRepo
	CallWrapper repository.WrapperRepo
}
