package firestore

import (
	"cloud.google.com/go/firestore"
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
)

type ServiceFirestore struct {
	Cfg       *config.ServiceConfig
	UtilsRepo repository.UtilsRepo
	clients   map[string]*firestore.Client
}
