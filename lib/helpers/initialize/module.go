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
	generalInteractor := GeneralInteractor(cfg)
	apiInteractor := ApiInteractor(cfg)

	uc := restUC.New(&restUC.UCInteractor{
		Cfg:             cfg,
		DBRepo:          generalInteractor.serviceDB,
		CallWrapperRepo: generalInteractor.wrapper,
		CacheRepo:       generalInteractor.serviceCache,
		UtilsRepo:       generalInteractor.utils,
		HelpersRepo:     generalInteractor.helpers,
		OMDBRepo:        apiInteractor.omdb,
		SpotifyRepo:     apiInteractor.spotify,
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
	apiInteractor := ApiInteractor(cfg)

	uc := grpcUC.New(
		&grpcUC.UCInteractor{
			Cfg:             cfg,
			DBRepo:          generalInteractor.serviceDB,
			CallWrapperRepo: generalInteractor.wrapper,
			CacheRepo:       generalInteractor.serviceCache,
			UtilsRepo:       generalInteractor.utils,
			HelpersRepo:     generalInteractor.helpers,
			OMDBRepo:        apiInteractor.omdb,
			SpotifyRepo:     apiInteractor.spotify,
		},
	)

	return &grpcRouters.GrpcInteractor{
		GrpcInterface: uc,
	}
}

//CronGetInteractor job scheduler interactor and related usecase
func CronGetInteractor(cfg *config.ServiceConfig) *cronRouters.CronInteractor {
	generalInteractor := GeneralInteractor(cfg)
	apiInteractor := ApiInteractor(cfg)

	ucScheduler := cronUC.New(
		&cronUC.UCInteractor{
			Cfg:             cfg,
			DBRepo:          generalInteractor.serviceDB,
			Firestore:       generalInteractor.serviceFirestore,
			CacheRepo:       generalInteractor.serviceCache,
			CallWrapperRepo: generalInteractor.wrapper,
			UtilsRepo:       generalInteractor.utils,
			HelpersRepo:     generalInteractor.helpers,
			OMDBRepo:        apiInteractor.omdb,
			SpotifyRepo:     apiInteractor.spotify,
		},
	)

	return &cronRouters.CronInteractor{
		CronInterface: ucScheduler,
	}
}

//ConsumerGetInteractor consumer interactor and related usecase
func ConsumerGetInteractor(cfg *config.ServiceConfig) *consumerRouters.ConsumerInteractor {
	generalInteractor := GeneralInteractor(cfg)
	apiInteractor := ApiInteractor(cfg)

	ucConsumer := consumerUC.New(
		&consumerUC.UCInteractor{
			Cfg:             cfg,
			DBRepo:          generalInteractor.serviceDB,
			Firestore:       generalInteractor.serviceFirestore,
			CacheRepo:       generalInteractor.serviceCache,
			CallWrapperRepo: generalInteractor.wrapper,
			UtilsRepo:       generalInteractor.utils,
			HelpersRepo:     generalInteractor.helpers,
			OMDBRepo:        apiInteractor.omdb,
			SpotifyRepo:     apiInteractor.spotify,
		},
	)

	return &consumerRouters.ConsumerInteractor{
		ConsumerInterface: ucConsumer,
	}
}
