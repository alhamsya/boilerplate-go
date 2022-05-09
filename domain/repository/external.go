package repository

import (
	"github.com/alhamsya/boilerplate-go/infrastructure/external/omdb"
	"github.com/alhamsya/boilerplate-go/infrastructure/external/spotify"
)

type OMDBRepo interface {
	GetListMovie(search string, page int64) (*omdb.OMDBList, error)
	GetDetailMovie(movieID string) (*omdb.OMDBDetail, error)
}

type SpotifyRepo interface {
	GetUser() (responseObject *spotify.Profile, err error)
}
