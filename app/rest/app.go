package main

import (
	"github.com/alhamsya/boilerplate-go/app"
	"github.com/alhamsya/boilerplate-go/service/inter/rest/handler"
)

func main() {
	cfg := app.GetConfig()

	//init databases
	dbService, err := app.GetDatabase(&cfg, "tes")
	if err != nil {
		panic("database")
	}

	// init rest option
	server := restHandler.New(&restHandler.Handler{
		Cfg: &cfg,
		Interactor: app.RestGetInteractor(&cfg, dbService),
	})

	if err := server.Run(); err != nil {
		panic("")
	}
}
