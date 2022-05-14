package helpersUC

import (
	"github.com/alhamsya/boilerplate-go/domain/repositorys"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
)

type UCInteractor struct {
	Cfg             *config.ServiceConfig
	DBRepo          repositorys.DBRepo
	CacheRepo       repositorys.CacheRepo
	CallWrapperRepo repositorys.CallWrapperRepo
	UtilsRepo       repositorys.UtilsRepo
}
