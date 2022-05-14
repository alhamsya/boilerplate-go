package main

import (
	"context"
	"log"

	"github.com/alhamsya/boilerplate-go/lib/managers/custom_log"
	"github.com/alhamsya/boilerplate-go/lib/managers/initialize"
	"github.com/alhamsya/boilerplate-go/transport/cron/handler"
)

func main() {
	cfg := initialize.GetConfig()

	err := customLog.InitializeLogging(cfg.CronServer.AccessLogFile, cfg.CronServer.ErrorLogFile)
	if err != nil {
		log.Fatalln("[CRON] failed initialize logging:", err)
	}

	//init cron option
	server := cronHandler.New(&cronHandler.Handler{
		Cfg:        &cfg,
		Interactor: initialize.CronGetInteractor(&cfg),
	})

	//running service cron
	if err := server.Run(context.Background()); err != nil {
		log.Fatalln("[CRON] service not running:", err)
	}
}
