package app

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/infrastructure/caches"
	"github.com/alhamsya/boilerplate-go/infrastructure/databases"
	"github.com/alhamsya/boilerplate-go/infrastructure/external/omdb"
	"github.com/alhamsya/boilerplate-go/infrastructure/firestores"
	"github.com/alhamsya/boilerplate-go/infrastructure/wrappers"
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
	"github.com/alhamsya/boilerplate-go/lib/helpers/database"
	"github.com/alhamsya/boilerplate-go/lib/utils"
	"github.com/alhamsya/boilerplate-go/middleware/rest"
	"github.com/alhamsya/boilerplate-go/transport/inter/consumer/routers"
	"github.com/alhamsya/boilerplate-go/transport/inter/cron/routers"
	"github.com/alhamsya/boilerplate-go/transport/inter/grpc/routers"
	"github.com/alhamsya/boilerplate-go/transport/inter/rest/routers"
	"github.com/alhamsya/boilerplate-go/usecase/consumer"
	"github.com/alhamsya/boilerplate-go/usecase/cron"
	"github.com/alhamsya/boilerplate-go/usecase/grpc"
	"github.com/alhamsya/boilerplate-go/usecase/helpers"
	"github.com/alhamsya/boilerplate-go/usecase/rest"
)

type ModuleRepo struct {
	serviceFirestore *firestores.ServiceFirestore
	serviceDB        *databases.ServiceDB
	serviceCache     *caches.ServiceCache
	omdb             *external.OMDB
	wrapper          *wrappers.Wrapper
	utils            repository.UtilsRepo
	helpers          repository.HelpersRepo
}

//GetConfig get config by name
func GetConfig() (cfg config.ServiceConfig) {
	cfg.ReadConfig("main")
	cfg.ReadConfig("toggle")
	cfg.ReadConfig("scheduler")
	cfg.ReadConfig("pubsub")
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

	middleware := restMiddleware.New(&restMiddleware.Middleware{
		Cfg:       cfg,
		DBRepo:    generalInteractor.serviceDB,
		UtilsRepo: generalInteractor.utils,
	})

	return &restRouters.RestInteractor{
		Usecase:    uc,
		Middleware: middleware,
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
			HelpersRepo:     generalInteractor.helpers,
		},
	)

	return &grpcRouters.GrpcInteractor{
		GrpcInterface: uc,
	}
}

//CronGetInteractor job scheduler interactor and related usecase
func CronGetInteractor(cfg *config.ServiceConfig) *cronRouters.CronInteractor {
	generalInteractor := GeneralInteractor(cfg)

	ucScheduler := cronUC.New(
		&cronUC.UCInteractor{
			Cfg:             cfg,
			DBRepo:          generalInteractor.serviceDB,
			Firestore:       generalInteractor.serviceFirestore,
			CacheRepo:       generalInteractor.serviceCache,
			OMDBRepo:        generalInteractor.omdb,
			CallWrapperRepo: generalInteractor.wrapper,
			UtilsRepo:       generalInteractor.utils,
			HelpersRepo:     generalInteractor.helpers,
		},
	)

	return &cronRouters.CronInteractor{
		CronInterface: ucScheduler,
	}
}

//ConsumerGetInteractor consumer interactor and related usecase
func ConsumerGetInteractor(cfg *config.ServiceConfig) *consumerRouters.ConsumerInteractor {
	generalInteractor := GeneralInteractor(cfg)

	ucConsumer := consumerUC.New(
		&consumerUC.UCInteractor{
			Cfg:             cfg,
			DBRepo:          generalInteractor.serviceDB,
			Firestore:       generalInteractor.serviceFirestore,
			CacheRepo:       generalInteractor.serviceCache,
			OMDBRepo:        generalInteractor.omdb,
			CallWrapperRepo: generalInteractor.wrapper,
			UtilsRepo:       generalInteractor.utils,
			HelpersRepo:     generalInteractor.helpers,
		},
	)

	return &consumerRouters.ConsumerInteractor{
		ConsumerInterface: ucConsumer,
	}
}

//GeneralInteractor general interactor for rest and gRPC
func GeneralInteractor(cfg *config.ServiceConfig) *ModuleRepo {
	utilsService := utils.New()

	dbService := databases.New(
		&databases.ServiceDB{
			Cfg:    cfg,
			Name:   "movie",
			Driver: database.DriverMySQL,
		},
	)

	firestoreService := firestores.New(&firestores.ServiceFirestore{
		Cfg:       cfg,
		UtilsRepo: utilsService,
	})

	cacheService := caches.New(
		&caches.ServiceCache{
			Cfg: cfg,
		},
	)

	omdbRepo := external.New(&external.OMDB{
		Cfg: cfg,
	})

	cw := wrappers.New(
		&wrappers.Wrapper{
			Cfg: cfg,
		},
	)

	helpersService := helpersUC.New(&helpersUC.UCInteractor{
		Cfg:             cfg,
		DBRepo:          dbService,
		FirestoreRepo:   firestoreService,
		CacheRepo:       cacheService,
		CallWrapperRepo: cw,
		UtilsRepo:       utilsService,
	})

	return &ModuleRepo{
		serviceFirestore: firestoreService,
		serviceDB:        dbService,
		serviceCache:     cacheService,
		omdb:             omdbRepo,
		wrapper:          cw,
		utils:            utilsService,
		helpers:          helpersService,
	}
}
