package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/rcdmk/shortest-flight-path/data"
	"github.com/rcdmk/shortest-flight-path/domain/service"
	"github.com/rcdmk/shortest-flight-path/infra/cache"
	"github.com/rcdmk/shortest-flight-path/infra/config"
	"github.com/rcdmk/shortest-flight-path/server"
	"github.com/rcdmk/shortest-flight-path/server/controller"
)

func main() {
	log.Println("Loading configuration...")

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connecting to DB server %s:%v", cfg.DB.Host, cfg.DB.Port)

	db, err := data.New(cfg.DB)
	if err != nil {
		if db != nil {
			db.Close()
		}

		log.Fatal("error initializing DB: ", err)
	}
	defer db.Close()

	log.Printf("Connecting to cache server %s:%v", cfg.Cache.Host, cfg.Cache.Port)

	cache, err := cache.New(cfg.Cache)
	if err != nil {
		log.Println("error initializing cache: ", err)
	}

	routerService := service.NewRouter(db, cache)

	routeController := controller.NewRoute(cfg, routerService)

	server := server.New(cfg.Server)
	server.RegisterRoutes(routeController)

	err = server.Start()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("error starting server: ", err)
		}

		log.Printf("server closed")
	}
}
