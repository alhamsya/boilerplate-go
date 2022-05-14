package caches

import (
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
	"github.com/go-redis/redis/v8"
)

type ServiceCache struct {
	Cfg   *config.ServiceConfig
	Redis *redis.Client
}
