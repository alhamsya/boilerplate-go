package restUC

import (
	"github.com/alhamsya/boilerplate-go/domain/repositories"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
)

type UCInteractor struct {
	Cfg             *config.ServiceConfig
	DBRepo          repositories.DBRepo
	CacheRepo       repositories.CacheRepo
	CallWrapperRepo repositories.CallWrapperRepo
	UtilsRepo       repositories.UtilsRepo
	HelpersRepo     repositories.HelpersRepo
	OMDBRepo        repositories.OMDBRepo
	SpotifyRepo     repositories.SpotifyRepo
}
