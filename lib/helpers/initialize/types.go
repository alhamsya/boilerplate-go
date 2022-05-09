package initialize

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/infrastructure/caches"
	"github.com/alhamsya/boilerplate-go/infrastructure/databases"
	"github.com/alhamsya/boilerplate-go/infrastructure/external/omdb"
	"github.com/alhamsya/boilerplate-go/infrastructure/external/spotify"
	"github.com/alhamsya/boilerplate-go/infrastructure/firestores"
	"github.com/alhamsya/boilerplate-go/infrastructure/wrappers"
)

type ModuleRepo struct {
	serviceFirestore *firestores.ServiceFirestore
	serviceDB        *databases.ServiceDB
	serviceCache     *caches.ServiceCache
	wrapper          *wrappers.Wrapper
	utils            repository.UtilsRepo
	helpers          repository.HelpersRepo
}

type ExternalRepo struct {
	omdb    *omdb.OMDB
	spotify *spotify.Spotify
}
