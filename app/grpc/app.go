package main

import (
	"github.com/alhamsya/boilerplate-go/app"

	grpcHandler "github.com/alhamsya/boilerplate-go/infrastructure/services/inter/grpc/handler"
)

func main()  {
	cfg := app.GetConfig()

	//init databases
	dbService, err := app.GetDatabase(&cfg, "tes")
	if err != nil {
		panic("database")
	}

	// init grpc option
	server := grpcHandler.New(&grpcHandler.Handler{
		Cfg: &cfg,
		Interactor: app.GrpcGetInteractor(&cfg, dbService),
	})

	if err := server.Run(); err != nil {
		panic("server")
	}
}