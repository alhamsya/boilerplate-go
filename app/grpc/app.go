package main

import (
	"log"

	"github.com/alhamsya/boilerplate-go/app"
	"github.com/alhamsya/boilerplate-go/service/inter/grpc/handler"
)

func main()  {
	cfg := app.GetConfig()

	//init grpc option
	server := grpcHandler.New(&grpcHandler.Handler{
		Cfg: &cfg,
		Interactor: app.GrpcGetInteractor(&cfg),
	})

	//running service gRPC
	if err := server.Run(); err != nil {
		log.Fatalln("[GRPC] service not running", err)
	}
}