package main

import (
	"log"

	"github.com/rcdmk/shortest-flight-path/data"
	"github.com/rcdmk/shortest-flight-path/domain/service"

	"github.com/rcdmk/shortest-flight-path/infra/config"
	"github.com/rcdmk/shortest-flight-path/server"
	"github.com/rcdmk/shortest-flight-path/server/controller"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	server := server.New(cfg.Server)

	db, err := data.New(cfg.DB)
	if err != nil {
		if db != nil {
			db.Close()
		}

		log.Fatal("error initializing DB: ", err)
	}

	routerService := service.NewRouter(db)

	routeController := controller.NewRoute(cfg, routerService)

	server.RegisterRoutes(routeController)

	err = server.Start()
	if err != nil {
		log.Fatal("error starting server: ", err)
	}

	return
}
