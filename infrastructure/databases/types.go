package databases

import (
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
	"github.com/alhamsya/boilerplate-go/lib/helpers/database"
)

type ServiceDB struct {
	Cfg    *config.ServiceConfig
	Driver string
	DB     *database.Store
	Name   string
}
