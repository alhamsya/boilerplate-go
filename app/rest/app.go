package main

import (
	"log"

	"github.com/alhamsya/boilerplate-go/app"
	"github.com/alhamsya/boilerplate-go/service/inter/rest/handler"
)

func main() {
	cfg := app.GetConfig()

	//init rest option
	server := restHandler.New(&restHandler.Handler{
		Cfg: &cfg,
		Interactor: app.RestGetInteractor(&cfg),
	})

	//running service rest API
	if err := server.Run(); err != nil {
		log.Fatalln("[REST] service not running", err)
	}
}
