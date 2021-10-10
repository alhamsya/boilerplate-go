package app

import (
	"fmt"

	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/alhamsya/boilerplate-go/infrastructure/databases"
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
}

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

func RestGetInteractor(cfg *config.MainConfig, db *database.Store) *restRouters.RestInteractor {
	generalInteractor := GeneralInteractor(cfg, db)

	uc := restUc.New(&restUc.UcInteractor{
		Cfg:         cfg,
		ServiceRepo: generalInteractor.service,
		OMDBRepo:    generalInteractor.omdb,
	})

	return &restRouters.RestInteractor{
		RestInterface: uc,
	}
}

func GrpcGetInteractor(cfg *config.MainConfig, db *database.Store) *grpcRouters.GrpcInteractor {
	generalInteractor := GeneralInteractor(cfg, db)

	uc := grpcUc.New(&grpcUc.UcInteractor{
		Cfg:         cfg,
		ServiceRepo: generalInteractor.service,
		OMDBRepo:    generalInteractor.omdb,
	})

	return &grpcRouters.GrpcInteractor{
		GrpcInterface: uc,
	}
}

func GeneralInteractor(cfg *config.MainConfig, db *database.Store) *ModuleRepo {
	service := databases.New(
		&databases.DBService{
			DB: db,
		},
	)

	omdbRepo := omdb.New(&omdb.OMDB{
		Cfg: cfg,
	})

	return &ModuleRepo{
		service: service,
		omdb:    omdbRepo,
	}
}
