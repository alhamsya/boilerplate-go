package helpersUC

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
)

type UCInteractor struct {
	Cfg             *config.ServiceConfig
	DBRepo          repository.DBRepo
	FirestoreRepo   repository.FirestoreRepo
	CacheRepo       repository.CacheRepo
	CallWrapperRepo repository.CallWrapperRepo
	UtilsRepo       repository.UtilsRepo
}
