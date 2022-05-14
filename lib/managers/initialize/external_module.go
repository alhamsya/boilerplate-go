package initialize

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/infrastructure/external/omdb"
	"github.com/alhamsya/boilerplate-go/infrastructure/external/spotify"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
)

func NewExternalModule() *ExternalRepo {
	return &ExternalRepo{}
}

func (e *ExternalRepo) OMDBInteractor(cfg *config.ServiceConfig) repository.OMDBRepo {
	return omdb.New(&omdb.OMDB{
		Cfg: cfg,
	})
}

func (e *ExternalRepo) SpotifyInteractor(cfg *config.ServiceConfig) repository.SpotifyRepo {
	return spotify.New(&spotify.Spotify{
		Cfg: cfg,
	})
}
