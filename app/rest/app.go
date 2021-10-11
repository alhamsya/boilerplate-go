package main

import (
	"github.com/alhamsya/boilerplate-go/app"
	"github.com/alhamsya/boilerplate-go/service/inter/rest/handler"
	"log"
)

func main() {
	cfg := app.GetConfig()

	//init databases
	dbService, err := app.GetDatabase(&cfg, "tes")
	if err != nil {
		log.Fatalln("[REST] fail connect to database", err)
	}

	// init rest option
	server := restHandler.New(&restHandler.Handler{
		Cfg: &cfg,
		Interactor: app.RestGetInteractor(&cfg, dbService),
	})

	if err := server.Run(); err != nil {
		log.Fatalln("[REST] service not running", err)
	}
}
