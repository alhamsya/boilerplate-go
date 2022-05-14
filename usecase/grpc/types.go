package grpcUC

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
)

type UCInteractor struct {
	Cfg             *config.ServiceConfig
	DBRepo          repository.DBRepo
	CacheRepo       repository.CacheRepo
	CallWrapperRepo repository.CallWrapperRepo
	UtilsRepo       repository.UtilsRepo
	HelpersRepo     repository.HelpersRepo
	OMDBRepo        repository.OMDBRepo
	SpotifyRepo     repository.SpotifyRepo
}
