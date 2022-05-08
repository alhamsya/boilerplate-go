package main

import (
	"context"
	"log"

	"github.com/alhamsya/boilerplate-go/app"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
	"github.com/alhamsya/boilerplate-go/transport/consumer/handler"
)

func main() {
	cfg := app.GetConfig()

	err := customLog.InitializeLogging(cfg.ConsumerServer.AccessLogFile, cfg.ConsumerServer.ErrorLogFile)
	if err != nil {
		log.Fatalln("[CRON] failed initialize logging:", err)
	}

	//init consumer option
	server := consumerHandler.New(&consumerHandler.Handler{
		Cfg:        &cfg,
		Interactor: app.ConsumerGetInteractor(&cfg),
	})

	//running service consumer
	if err := server.Run(context.Background()); err != nil {
		log.Fatalln("[CONSUMER] service not running:", err)
	}
}
