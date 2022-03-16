package main

import (
	"log"

	"github.com/alhamsya/boilerplate-go/app"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
	"github.com/alhamsya/boilerplate-go/transport/inter/job/handler"
)

func main() {
	cfg := app.GetConfig()

	err := customLog.InitializeLogging(cfg.GrpcServer.AccessLogFile, cfg.GrpcServer.ErrorLogFile)
	if err != nil {
		log.Fatalln("[JOB] failed initialize logging:", err)
	}

	//init job option
	server := jobHandler.New(&jobHandler.Handler{
		Cfg:        &cfg,
		Interactor: app.JobGetInteractor(&cfg),
	})

	//running transport job
	if err := server.Run(); err != nil {
		log.Fatalln("[JOB] transport not running:", err)
	}
}
