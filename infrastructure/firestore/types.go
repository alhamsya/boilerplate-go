package firestore

import (
	"cloud.google.com/go/firestore"
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
)

type ServiceFirestore struct {
	Cfg       *config.ServiceConfig
	UtilsRepo repository.UtilsRepo
	Clients   map[string]*firestore.Client
}
