package databases

import (
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
	"github.com/alhamsya/boilerplate-go/lib/managers/database"
)

type ServiceDB struct {
	Cfg    *config.ServiceConfig
	Driver string
	db     *database.Store
	Name   string
}
