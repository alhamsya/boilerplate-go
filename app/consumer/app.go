package main

import (
	"context"
	"log"

	"github.com/alhamsya/boilerplate-go/lib/managers/custom_log"
	"github.com/alhamsya/boilerplate-go/lib/managers/initialize"
	"github.com/alhamsya/boilerplate-go/transport/consumer/handler"
)

func main() {
	cfg := initialize.GetConfig()

	err := customLog.InitializeLogging(cfg.ConsumerServer.AccessLogFile, cfg.ConsumerServer.ErrorLogFile)
	if err != nil {
		log.Fatalln("[CRON] failed initialize logging:", err)
	}

	//init consumer option
	server := consumerHandler.New(&consumerHandler.Handler{
		Cfg:        &cfg,
		Interactor: initialize.ConsumerGetInteractor(&cfg),
	})

	//running service consumer
	if err := server.Run(context.Background()); err != nil {
		log.Fatalln("[CONSUMER] service not running:", err)
	}
}
