package main

import (
	"log"

	"github.com/alhamsya/boilerplate-go/app"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
	"github.com/alhamsya/boilerplate-go/transport/inter/cron/handler"
)

func main() {
	cfg := app.GetConfig()

	err := customLog.InitializeLogging(cfg.CronServer.AccessLogFile, cfg.CronServer.ErrorLogFile)
	if err != nil {
		log.Fatalln("[CRON] failed initialize logging:", err)
	}

	//init job option
	server := cronHandler.New(&cronHandler.Handler{
		Cfg:        &cfg,
		Interactor: app.CronGetInteractor(&cfg),
	})

	//running transport job
	if err := server.Run(); err != nil {
		log.Fatalln("[CRON] transport not running:", err)
	}
}
