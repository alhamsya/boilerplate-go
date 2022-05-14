package initialize

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/infrastructure/caches"
	"github.com/alhamsya/boilerplate-go/infrastructure/databases"
	"github.com/alhamsya/boilerplate-go/infrastructure/firestores"
	"github.com/alhamsya/boilerplate-go/infrastructure/wrappers"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
	"github.com/alhamsya/boilerplate-go/lib/managers/database"
	"github.com/alhamsya/boilerplate-go/lib/utils"
	"github.com/alhamsya/boilerplate-go/usecase/helpers"
)

func NewBaseModule() *ModuleRepo {
	return &ModuleRepo{}
}

func (m *ModuleRepo) UtilsInteractor() repository.UtilsRepo {
	return utils.New()
}

func (m *ModuleRepo) DBInteractor(cfg *config.ServiceConfig) repository.DBRepo {
	return databases.New(
		&databases.ServiceDB{
			Cfg:    cfg,
			Name:   "movie",
			Driver: database.DriverMySQL,
		},
	)
}

func (m *ModuleRepo) CacheInteractor(cfg *config.ServiceConfig) repository.CacheRepo {
	return caches.New(
		&caches.ServiceCache{
			Cfg: cfg,
		},
	)
}

func (m *ModuleRepo) CallWrapperInteractor(cfg *config.ServiceConfig) repository.CallWrapperRepo {
	return wrappers.New(
		&wrappers.Wrapper{
			Cfg: cfg,
		},
	)
}

func (m *ModuleRepo) FirestoreInteractor(cfg *config.ServiceConfig) repository.FirestoreRepo {
	return firestores.New(&firestores.ServiceFirestore{
		Cfg:       cfg,
		UtilsRepo: m.UtilsInteractor(),
	})
}

func (m *ModuleRepo) HelpersInteractor(cfg *config.ServiceConfig) repository.HelpersRepo {
	return helpersUC.New(&helpersUC.UCInteractor{
		Cfg:             cfg,
		DBRepo:          m.DBInteractor(cfg),
		CacheRepo:       m.CacheInteractor(cfg),
		CallWrapperRepo: m.CallWrapperInteractor(cfg),
		UtilsRepo:       m.UtilsInteractor(),
	})
}
