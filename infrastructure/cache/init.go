package cache

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/go-redis/redis/v8"
)

func New(this *ServiceCache) *ServiceCache {
	redisClient := redis.NewClient(&redis.Options{
		Network:            constCommon.NetworkTCP,
		Addr:               this.Cfg.Redis.Address,
		Dialer:             func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.Dial(constCommon.NetworkTCP, this.Cfg.Redis.Address)
		},
		Username:           this.Cfg.Redis.Username,
		Password:           this.Cfg.Redis.Password,
		DB:                 this.Cfg.Redis.DB,
		MinIdleConns:       this.Cfg.Redis.MinIdleConn,
		MaxConnAge:         time.Duration(this.Cfg.Redis.MaxConnAgeInSec) * time.Second,
		IdleTimeout:        time.Duration(this.Cfg.Redis.IdleTimeoutInSec) * time.Second,
	})

	if _, err := redisClient.Ping(redisClient.Context()).Result(); err != nil {
		log.Fatalf("[INIT] [REDIS] fail initialize redis: %v", err)
	}

	log.Printf("[IGNORE] [REDIS] initialize")

	return &ServiceCache{
		Cfg:   this.Cfg,
		Redis: redisClient,
	}
}