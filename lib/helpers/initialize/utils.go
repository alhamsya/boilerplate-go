package initialize

import (
	"github.com/alhamsya/boilerplate-go/infrastructure/caches"
	"github.com/alhamsya/boilerplate-go/infrastructure/databases"
	"github.com/alhamsya/boilerplate-go/infrastructure/external/omdb"
	"github.com/alhamsya/boilerplate-go/infrastructure/external/spotify"
	"github.com/alhamsya/boilerplate-go/infrastructure/firestores"
	"github.com/alhamsya/boilerplate-go/infrastructure/wrappers"
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
	"github.com/alhamsya/boilerplate-go/lib/helpers/database"
	"github.com/alhamsya/boilerplate-go/lib/utils"
	"github.com/alhamsya/boilerplate-go/usecase/helpers"
)

//GetConfig get config by name
func GetConfig() (cfg config.ServiceConfig) {
	cfg.ReadConfig("main")
	cfg.ReadConfig("toggle")
	cfg.ReadConfig("scheduler")
	cfg.ReadConfig("pubsub")
	return cfg
}

//GeneralInteractor general interactor
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
		wrapper:          cw,
		utils:            utilsService,
		helpers:          helpersService,
	}
}

//ApiInteractor external API
func ApiInteractor(cfg *config.ServiceConfig) *ExternalRepo {
	omdbRepo := omdb.New(&omdb.OMDB{
		Cfg: cfg,
	})

	spotifyRepo := spotify.New(&spotify.Spotify{
		Cfg: cfg,
	})

	return &ExternalRepo{
		omdb:    omdbRepo,
		spotify: spotifyRepo,
	}
}
