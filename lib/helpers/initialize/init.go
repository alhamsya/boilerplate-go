package initialize

import (
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
	"github.com/alhamsya/boilerplate-go/middleware/rest"
	"github.com/alhamsya/boilerplate-go/transport/consumer/routers"
	"github.com/alhamsya/boilerplate-go/transport/cron/routers"
	"github.com/alhamsya/boilerplate-go/transport/grpc/routers"
	"github.com/alhamsya/boilerplate-go/transport/rest/routers"
	"github.com/alhamsya/boilerplate-go/usecase/consumer"
	"github.com/alhamsya/boilerplate-go/usecase/cron"
	"github.com/alhamsya/boilerplate-go/usecase/grpc"
	"github.com/alhamsya/boilerplate-go/usecase/rest"
)

//RestGetInteractor rest get interactor and related usecase
func RestGetInteractor(cfg *config.ServiceConfig) *restRouters.RestInteractor {
	uc := restUC.New(&restUC.UCInteractor{
		Cfg:             cfg,
		HelpersRepo:     NewBaseModule().HelpersInteractor(cfg),
		DBRepo:          NewBaseModule().DBInteractor(cfg),
		CallWrapperRepo: NewBaseModule().CallWrapperInteractor(cfg),
		CacheRepo:       NewBaseModule().CacheInteractor(cfg),
		UtilsRepo:       NewBaseModule().UtilsInteractor(),
		OMDBRepo:        NewExternalModule().OMDBInteractor(cfg),
		SpotifyRepo:     NewExternalModule().SpotifyInteractor(cfg),
	})

	middleware := restMiddleware.New(&restMiddleware.Middleware{
		Cfg:       cfg,
		DBRepo:    uc.DBRepo,
		UtilsRepo: uc.UtilsRepo,
	})

	return &restRouters.RestInteractor{
		Usecase:    uc,
		Middleware: middleware,
	}
}

//GrpcGetInteractor gRPC get interactor and related usecase
func GrpcGetInteractor(cfg *config.ServiceConfig) *grpcRouters.GrpcInteractor {
	uc := grpcUC.New(
		&grpcUC.UCInteractor{
			Cfg:             cfg,
			HelpersRepo:     NewBaseModule().HelpersInteractor(cfg),
			DBRepo:          NewBaseModule().DBInteractor(cfg),
			CallWrapperRepo: NewBaseModule().CallWrapperInteractor(cfg),
			CacheRepo:       NewBaseModule().CacheInteractor(cfg),
			UtilsRepo:       NewBaseModule().UtilsInteractor(),
			OMDBRepo:        NewExternalModule().OMDBInteractor(cfg),
			SpotifyRepo:     NewExternalModule().SpotifyInteractor(cfg),
		},
	)

	return &grpcRouters.GrpcInteractor{
		GrpcInterface: uc,
	}
}

//CronGetInteractor job scheduler interactor and related usecase
func CronGetInteractor(cfg *config.ServiceConfig) *cronRouters.CronInteractor {
	ucScheduler := cronUC.New(
		&cronUC.UCInteractor{
			Cfg:             cfg,
			HelpersRepo:     NewBaseModule().HelpersInteractor(cfg),
			DBRepo:          NewBaseModule().DBInteractor(cfg),
			Firestore:       NewBaseModule().FirestoreInteractor(cfg),
			CacheRepo:       NewBaseModule().CacheInteractor(cfg),
			CallWrapperRepo: NewBaseModule().CallWrapperInteractor(cfg),
			UtilsRepo:       NewBaseModule().UtilsInteractor(),
			OMDBRepo:        NewExternalModule().OMDBInteractor(cfg),
			SpotifyRepo:     NewExternalModule().SpotifyInteractor(cfg),
		},
	)

	return &cronRouters.CronInteractor{
		CronInterface: ucScheduler,
	}
}

//ConsumerGetInteractor consumer interactor and related usecase
func ConsumerGetInteractor(cfg *config.ServiceConfig) *consumerRouters.ConsumerInteractor {
	ucConsumer := consumerUC.New(
		&consumerUC.UCInteractor{
			Cfg:             cfg,
			HelpersRepo:     NewBaseModule().HelpersInteractor(cfg),
			DBRepo:          NewBaseModule().DBInteractor(cfg),
			Firestore:       NewBaseModule().FirestoreInteractor(cfg),
			CacheRepo:       NewBaseModule().CacheInteractor(cfg),
			CallWrapperRepo: NewBaseModule().CallWrapperInteractor(cfg),
			UtilsRepo:       NewBaseModule().UtilsInteractor(),
			OMDBRepo:        NewExternalModule().OMDBInteractor(cfg),
			SpotifyRepo:     NewExternalModule().SpotifyInteractor(cfg),
		},
	)

	return &consumerRouters.ConsumerInteractor{
		ConsumerInterface: ucConsumer,
	}
}
