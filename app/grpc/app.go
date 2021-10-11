package main

import (
	"github.com/alhamsya/boilerplate-go/app"
	"github.com/alhamsya/boilerplate-go/service/inter/grpc/handler"
	"log"
)

func main()  {
	cfg := app.GetConfig()

	//init databases
	dbService, err := app.GetDatabase(&cfg, "tes")
	if err != nil {
		log.Fatalln("[GRPC] fail connect to database", err)
	}

	// init grpc option
	server := grpcHandler.New(&grpcHandler.Handler{
		Cfg: &cfg,
		Interactor: app.GrpcGetInteractor(&cfg, dbService),
	})

	if err := server.Run(); err != nil {
		log.Fatalln("[GRPC] service not running", err)
	}
}