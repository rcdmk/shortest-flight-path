package server

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rcdmk/shortest-flight-path/infra/config"
	"github.com/rcdmk/shortest-flight-path/server/controller"
)

// API handles the API server
type API struct {
	echo *echo.Echo
	cfg  config.ServerConfig
}

// New returns a new API instance
func New(cfg config.ServerConfig) *API {
	e := echo.New()

	e.HideBanner = true

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
}

// RegisterRoutes registers root routes
func (srv *API) RegisterRoutes(routeController *controller.Route) {
	root := srv.echo.Group(srv.cfg.Prefix)

	root.GET("/routes", routeController.HandleShortestRoute)
}

// Start starts the HTTP server listening on configured port
func (srv *API) Start() error {
	return srv.echo.Start(":" + strconv.Itoa(srv.cfg.Port))
}
