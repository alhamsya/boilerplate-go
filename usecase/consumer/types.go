package consumerUC

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
)

type UCInteractor struct {
	Cfg             *config.ServiceConfig
	DBRepo          repository.DBRepo
	CacheRepo       repository.CacheRepo
	OMDBRepo        repository.OMDBRepo
	CallWrapperRepo repository.CallWrapperRepo
	UtilsRepo       repository.UtilsRepo
	Firestore       repository.FirestoreRepo
	HelpersRepo     repository.HelpersRepo
}
