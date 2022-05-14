package firestores

import (
	"cloud.google.com/go/firestore"
	"github.com/alhamsya/boilerplate-go/domain/repositorys"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
)

type ServiceFirestore struct {
	Cfg       *config.ServiceConfig
	UtilsRepo repositorys.UtilsRepo
	clients   map[string]*firestore.Client
}
