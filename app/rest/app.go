package main

import (
	"log"

	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
	"github.com/alhamsya/boilerplate-go/lib/helpers/initialize"
	"github.com/alhamsya/boilerplate-go/transport/rest/handler"
)

func main() {
	cfg := initialize.GetConfig()

	err := customLog.InitializeLogging(cfg.RestServer.AccessLogFile, cfg.RestServer.ErrorLogFile)
	if err != nil {
		log.Fatalln("[GRPC] failed initialize logging:", err)
	}

	//init rest option
	server := restHandler.New(&restHandler.Handler{
		Cfg:        &cfg,
		Interactor: initialize.RestGetInteractor(&cfg),
	})

	//running service rest API
	if err := server.Run(); err != nil {
		log.Fatalln("[REST] service not running", err)
	}
}
