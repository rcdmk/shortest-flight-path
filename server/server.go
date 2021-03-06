package server

import (
	"strconv"

	"github.com/rcdmk/shortest-flight-path/server/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rcdmk/shortest-flight-path/infra/config"
)

// API handles the API server
type API struct {
	echo *echo.Echo
	cfg  config.ServerConfig
}

// New returns a new API instance
func New(cfg config.ServerConfig) *API {
	e := echo.New()

	api := &API{
		echo: e,
		cfg:  cfg,
	}

	api.registerMiddleware()

	return api
}

// registerMiddleware registers global server middleware
func (srv *API) registerMiddleware() {
	srv.echo.Use(middleware.Logger())
	srv.echo.Use(middleware.Recover())

	return
}

// RegisterRoutes registers root routes
func (srv *API) RegisterRoutes(routeController *controller.Route) {
	root := srv.echo.Group(srv.cfg.Prefix)

	root.GET("/routes", routeController.HandleShortestRoute)
}

// Start starts the HTTP server listening on configured port
func (srv *API) Start() error {
	srv.echo.Logger.Info("Starting server on port ", srv.cfg.Port)
	return srv.echo.Start(":" + strconv.Itoa(srv.cfg.Port))
}
