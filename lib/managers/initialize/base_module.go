package initialize

import (
	"github.com/alhamsya/boilerplate-go/domain/repositories"
	"github.com/alhamsya/boilerplate-go/infrastructure/internals/caches"
	"github.com/alhamsya/boilerplate-go/infrastructure/internals/databases"
	"github.com/alhamsya/boilerplate-go/infrastructure/internals/firestores"
	"github.com/alhamsya/boilerplate-go/infrastructure/wrappers"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
	"github.com/alhamsya/boilerplate-go/lib/managers/database"
	"github.com/alhamsya/boilerplate-go/lib/utils"
	"github.com/alhamsya/boilerplate-go/usecase/helpers"
)

func NewBaseModule() *ModuleRepo {
	return &ModuleRepo{}
}

func (m *ModuleRepo) UtilsInteractor() repositories.UtilsRepo {
	return utils.New()
}

func (m *ModuleRepo) DBInteractor(cfg *config.ServiceConfig) repositories.DBRepo {
	return databases.New(
		&databases.ServiceDB{
			Cfg:    cfg,
			Name:   "movie",
			Driver: database.DriverMySQL,
		},
	)
}

func (m *ModuleRepo) CacheInteractor(cfg *config.ServiceConfig) repositories.CacheRepo {
	return caches.New(
		&caches.ServiceCache{
			Cfg: cfg,
		},
	)
}

func (m *ModuleRepo) CallWrapperInteractor(cfg *config.ServiceConfig) repositories.CallWrapperRepo {
	return wrappers.New(
		&wrappers.Wrapper{
			Cfg: cfg,
		},
	)
}

func (m *ModuleRepo) FirestoreInteractor(cfg *config.ServiceConfig) repositories.FirestoreRepo {
	return firestores.New(&firestores.ServiceFirestore{
		Cfg:       cfg,
		UtilsRepo: m.UtilsInteractor(),
	})
}

func (m *ModuleRepo) HelpersInteractor(cfg *config.ServiceConfig) repositories.HelpersRepo {
	return helpersUC.New(&helpersUC.UCInteractor{
		Cfg:             cfg,
		DBRepo:          m.DBInteractor(cfg),
		CacheRepo:       m.CacheInteractor(cfg),
		CallWrapperRepo: m.CallWrapperInteractor(cfg),
		UtilsRepo:       m.UtilsInteractor(),
	})
}
