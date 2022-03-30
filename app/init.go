package app

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/infrastructure/cache"
	"github.com/alhamsya/boilerplate-go/infrastructure/databases"
	"github.com/alhamsya/boilerplate-go/infrastructure/firestore"
	"github.com/alhamsya/boilerplate-go/infrastructure/pubsub"
	"github.com/alhamsya/boilerplate-go/infrastructure/wrapper"
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
	"github.com/alhamsya/boilerplate-go/lib/helpers/database"
	"github.com/alhamsya/boilerplate-go/lib/utils"
	"github.com/alhamsya/boilerplate-go/middleware/rest"
	"github.com/alhamsya/boilerplate-go/transport/exter/omdb"
	"github.com/alhamsya/boilerplate-go/transport/inter/consumer/routers"
	"github.com/alhamsya/boilerplate-go/transport/inter/cron/routers"
	"github.com/alhamsya/boilerplate-go/transport/inter/grpc/routers"
	"github.com/alhamsya/boilerplate-go/transport/inter/rest/routers"
	"github.com/alhamsya/boilerplate-go/usecase/consumer"
	"github.com/alhamsya/boilerplate-go/usecase/cron"
	"github.com/alhamsya/boilerplate-go/usecase/grpc"
	"github.com/alhamsya/boilerplate-go/usecase/rest"
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
		},
	)

	return &grpcRouters.GrpcInteractor{
		GrpcInterface: uc,
	}
}

//CronGetInteractor job scheduler interactor and related usecase
func CronGetInteractor(cfg *config.ServiceConfig) *cronRouters.CronInteractor {
	generalInteractor := GeneralInteractor(cfg)

	firestoreRepo := firestore.New(&firestore.ServiceFirestore{
		Cfg:       cfg,
		UtilsRepo: generalInteractor.utils,
	})

	ucScheduler := cronUC.New(
		&cronUC.UCInteractor{
			Cfg:             cfg,
			DBRepo:          generalInteractor.serviceDB,
			OMDBRepo:        generalInteractor.omdb,
			CallWrapperRepo: generalInteractor.wrapper,
			CacheRepo:       generalInteractor.serviceCache,
			UtilsRepo:       generalInteractor.utils,
			Firestore:       firestoreRepo,
		},
	)

	return &cronRouters.CronInteractor{
		CronInterface: ucScheduler,
	}
}

func ConsumerGetInteractor(cfg *config.ServiceConfig) *consumerRouters.ConsumerInteractor {
	generalInteractor := GeneralInteractor(cfg)

	firestoreRepo := firestore.New(&firestore.ServiceFirestore{
		Cfg:       cfg,
		UtilsRepo: generalInteractor.utils,
	})

	pubsub.New(&pubsub.ServicePubSub{
		Cfg:       cfg,
		UtilsRepo: generalInteractor.utils,
	})

	ucConsumer := consumerUC.New(
		&consumerUC.UCInteractor{
			Cfg:             cfg,
			DBRepo:          generalInteractor.serviceDB,
			OMDBRepo:        generalInteractor.omdb,
			CallWrapperRepo: generalInteractor.wrapper,
			CacheRepo:       generalInteractor.serviceCache,
			UtilsRepo:       generalInteractor.utils,
			Firestore:       firestoreRepo,
		},
	)

	return &consumerRouters.ConsumerInteractor{
		ConsumerInterface: ucConsumer,
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
