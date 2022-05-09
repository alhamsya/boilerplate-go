package main

import (
	"log"

	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
	"github.com/alhamsya/boilerplate-go/lib/helpers/initialize"
	"github.com/alhamsya/boilerplate-go/transport/grpc/handler"
)

func main() {
	cfg := initialize.GetConfig()

	err := customLog.InitializeLogging(cfg.GrpcServer.AccessLogFile, cfg.GrpcServer.ErrorLogFile)
	if err != nil {
		log.Fatalln("[GRPC] failed initialize logging:", err)
	}

	//init grpc option
	server := grpcHandler.New(&grpcHandler.Handler{
		Cfg:        &cfg,
		Interactor: initialize.GrpcGetInteractor(&cfg),
	})

	//running service gRPC
	if err := server.Run(); err != nil {
		log.Fatalln("[GRPC] service not running", err)
	}
}
