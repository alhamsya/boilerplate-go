package app

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/infrastructure/cache"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/alhamsya/boilerplate-go/infrastructure/databases"
	"github.com/alhamsya/boilerplate-go/infrastructure/firestore"
	"github.com/alhamsya/boilerplate-go/infrastructure/wrapper"
	"github.com/alhamsya/boilerplate-go/lib/helpers/database"
	"github.com/alhamsya/boilerplate-go/lib/utils"
	"github.com/alhamsya/boilerplate-go/transport/exter/omdb"
	"github.com/alhamsya/boilerplate-go/transport/inter/grpc/routers"
	"github.com/alhamsya/boilerplate-go/transport/inter/job/routers"
	"github.com/alhamsya/boilerplate-go/transport/inter/rest/routers"
	"github.com/alhamsya/boilerplate-go/usecase/grpc"
	"github.com/alhamsya/boilerplate-go/usecase/rest"
	"github.com/alhamsya/boilerplate-go/usecase/scheduler"
)

type ModuleRepo struct {
	serviceDB    *databases.ServiceDB
	serviceCache *cache.ServiceCache
	omdb         *omdb.OMDB
	wrapper      *wrapper.Wrapper
	utils        repository.UtilsRepo
}

//GetConfig get config by name
func GetConfig() (cfg config.ServiceConfig) {
	cfg.ReadConfig("main")
	cfg.ReadConfig("toggle")
	cfg.ReadConfig("scheduler")
	return cfg
}

//RestGetInteractor rest get interactor and related usecase
func RestGetInteractor(cfg *config.ServiceConfig) *restRouters.RestInteractor {
	generalInteractor := GeneralInteractor(cfg)

	uc := restUC.New(&restUC.UCInteractor{
		Cfg:             cfg,
		DBRepo:          generalInteractor.serviceDB,
		OMDBRepo:        generalInteractor.omdb,
		CallWrapperRepo: generalInteractor.wrapper,
		CacheRepo:       generalInteractor.serviceCache,
		UtilsRepo:       generalInteractor.utils,
	})

	return &restRouters.RestInteractor{
		RestInterface: uc,
	}
}

//GrpcGetInteractor gRPC get interactor and related usecase
func GrpcGetInteractor(cfg *config.ServiceConfig) *grpcRouters.GrpcInteractor {
	generalInteractor := GeneralInteractor(cfg)

	uc := grpcUC.New(
		&grpcUC.UCInteractor{
			Cfg:             cfg,
			DBRepo:          generalInteractor.serviceDB,
			OMDBRepo:        generalInteractor.omdb,
			CallWrapperRepo: generalInteractor.wrapper,
			CacheRepo:       generalInteractor.serviceCache,
			UtilsRepo:       generalInteractor.utils,
		},
	)

	return &grpcRouters.GrpcInteractor{
		GrpcInterface: uc,
	}
}

func JobGetInteractor(cfg *config.ServiceConfig) *jobRouters.JobInteractor {
	generalInteractor := GeneralInteractor(cfg)

	firestoreRepo := firestore.New(&firestore.ServiceFirestore{
		Cfg:       cfg,
		UtilsRepo: generalInteractor.utils,
	})

	ucScheduler := schedulerUC.New(
		&schedulerUC.UCInteractor{
			Cfg:             cfg,
			DBRepo:          generalInteractor.serviceDB,
			OMDBRepo:        generalInteractor.omdb,
			CallWrapperRepo: generalInteractor.wrapper,
			CacheRepo:       generalInteractor.serviceCache,
			UtilsRepo:       generalInteractor.utils,
			Firestore:       firestoreRepo,
		},
	)

	return &jobRouters.JobInteractor{
		SchedulerInterface: ucScheduler,
	}
}

//GeneralInteractor general interactor for rest and gRPC
func GeneralInteractor(cfg *config.ServiceConfig) *ModuleRepo {
	dbService := databases.New(
		&databases.ServiceDB{
			Cfg:    cfg,
			Name:   "movie",
			Driver: database.DriverMySQL,
		},
	)

	cacheService := cache.New(
		&cache.ServiceCache{
			Cfg: cfg,
		},
	)

	omdbRepo := omdb.New(&omdb.OMDB{
		Cfg: cfg,
	})

	cw := wrapper.New(
		&wrapper.Wrapper{
			Cfg: cfg,
		},
	)

	return &ModuleRepo{
		serviceDB:    dbService,
		serviceCache: cacheService,
		omdb:         omdbRepo,
		wrapper:      cw,
		utils:        utils.New(),
	}
}
