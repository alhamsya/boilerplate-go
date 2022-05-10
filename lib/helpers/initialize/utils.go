package initialize

import (
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
)

//GetConfig get config by name
func GetConfig() (cfg config.ServiceConfig) {
	cfg.ReadConfig("main")
	cfg.ReadConfig("toggle")
	cfg.ReadConfig("scheduler")
	cfg.ReadConfig("pubsub")
	return cfg
}
