package app

import (
	"fmt"

	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/alhamsya/boilerplate-go/infrastructure/databases"
	"github.com/alhamsya/boilerplate-go/infrastructure/wrapper"
	"github.com/alhamsya/boilerplate-go/lib/helpers/database"
	"github.com/alhamsya/boilerplate-go/service/exter/omdb"
	"github.com/alhamsya/boilerplate-go/service/inter/grpc/routers"
	"github.com/alhamsya/boilerplate-go/service/inter/rest/routers"
	"github.com/alhamsya/boilerplate-go/usecase/grpc"
	"github.com/alhamsya/boilerplate-go/usecase/rest"
)

type ModuleRepo struct {
	service *databases.DBService
	omdb    *omdb.OMDB
	wrapper *wrapper.Wrapper
}

//GetConfig get config by name
func GetConfig() (cfg config.MainConfig) {
	cfg.ReadConfig("main")
	return cfg
}

//GetDatabase get specific database
func GetDatabase(cfg *config.MainConfig, nameConfig string) (*database.Store, error) {
	dbCfg, ok := cfg.Databases[nameConfig]
	if !ok {
		return nil, fmt.Errorf("config database does not exist")
	}

	db, err := database.New(dbCfg, database.DriverMySQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//RestGetInteractor rest get interactor and related usecase
func RestGetInteractor(cfg *config.MainConfig, db *database.Store) *restRouters.RestInteractor {
	generalInteractor := GeneralInteractor(cfg, db)

	uc := restUC.New(&restUC.UCInteractor{
		Cfg:         cfg,
		ServiceRepo: generalInteractor.service,
		OMDBRepo:    generalInteractor.omdb,
		CallWrapper: generalInteractor.wrapper,
	})

	return &restRouters.RestInteractor{
		RestInterface: uc,
	}
}

//GrpcGetInteractor gRPC get interactor and related usecase
func GrpcGetInteractor(cfg *config.MainConfig, db *database.Store) *grpcRouters.GrpcInteractor {
	generalInteractor := GeneralInteractor(cfg, db)

	uc := grpcUC.New(
		&grpcUC.UcInteractor{
			Cfg:         cfg,
			ServiceRepo: generalInteractor.service,
			OMDBRepo:    generalInteractor.omdb,
			CallWrapper: generalInteractor.wrapper,
		},
	)

	return &grpcRouters.GrpcInteractor{
		GrpcInterface: uc,
	}
}

//GeneralInteractor general interactor for rest and gRPC
func GeneralInteractor(cfg *config.MainConfig, db *database.Store) *ModuleRepo {
	service := databases.New(
		&databases.DBService{
			DB: db,
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
		service: service,
		omdb:    omdbRepo,
		wrapper: cw,
	}
}
