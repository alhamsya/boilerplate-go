package app

import (
	"fmt"
	"github.com/alhamsya/boilerplate-go/infrastructure/services/exter/omdb"

	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/alhamsya/boilerplate-go/infrastructure/databases"
	"github.com/alhamsya/boilerplate-go/lib/helpers/database"

	grpcRouters "github.com/alhamsya/boilerplate-go/infrastructure/services/inter/grpc/routers"
	httpRouters "github.com/alhamsya/boilerplate-go/infrastructure/services/inter/rest/routers"
	grpcUc "github.com/alhamsya/boilerplate-go/usecase/grpc"
	httpUc "github.com/alhamsya/boilerplate-go/usecase/rest"
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

func HttpGetInteractor(cfg *config.MainConfig, db *database.Store) *httpRouters.RestInteractor {
	generalInteractor := GeneralInteractor(cfg, db)

	uc := httpUc.NewInteractor(&httpUc.UcInteractor{
		Cfg:         cfg,
		ServiceRepo: generalInteractor.service,
		OMDBRepo:    generalInteractor.omdb,
	})

	return &httpRouters.RestInteractor{
		RestInterface: uc,
	}
}

func GrpcGetInteractor(cfg *config.MainConfig, db *database.Store) *grpcRouters.GrpcInteractor {
	generalInteractor := GeneralInteractor(cfg, db)

	uc := grpcUc.NewInteractor(&grpcUc.UcInteractor{
		Cfg:         cfg,
		ServiceRepo: generalInteractor.service,
		OMDBRepo:    generalInteractor.omdb,
	})

	return &grpcRouters.GrpcInteractor{
		GrpcInterface: uc,
	}
}

func GeneralInteractor(cfg *config.MainConfig, db *database.Store) *ModuleRepo {
	service := databases.NewDBService(
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
