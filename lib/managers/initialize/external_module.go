package initialize

import (
	"github.com/alhamsya/boilerplate-go/domain/repositorys"
	"github.com/alhamsya/boilerplate-go/infrastructure/externals/omdb"
	"github.com/alhamsya/boilerplate-go/infrastructure/externals/spotify"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
)

func NewExternalModule() *ExternalRepo {
	return &ExternalRepo{}
}

func (e *ExternalRepo) OMDBInteractor(cfg *config.ServiceConfig) repositorys.OMDBRepo {
	return omdb.New(&omdb.OMDB{
		Cfg: cfg,
	})
}

func (e *ExternalRepo) SpotifyInteractor(cfg *config.ServiceConfig) repositorys.SpotifyRepo {
	return spotify.New(&spotify.Spotify{
		Cfg: cfg,
	})
}
