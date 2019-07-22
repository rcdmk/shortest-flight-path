package controller

import (
	"net/http"
	"strings"

	"github.com/rcdmk/shortest-flight-path/server/viewmodel"

	"github.com/labstack/echo"
	"github.com/rcdmk/shortest-flight-path/domain"
	"github.com/rcdmk/shortest-flight-path/domain/contract"
	"github.com/rcdmk/shortest-flight-path/infra/config"
	"github.com/rcdmk/shortest-flight-path/infra/errors"
)

// Route handles flight route related API routes
type Route struct {
	cfg           *config.Config
	routerService contract.RouterService
}

// NewRoute returns a new RouteController instance
func NewRoute(cfg *config.Config, routerService contract.RouterService) *Route {
	return &Route{
		cfg:           cfg,
		routerService: routerService,
	}
}

// getOriginAndDestination gets and validates origin and destination codes
func getOriginAndDestination(c echo.Context) (origin string, destination string, err error) {
	origin = c.QueryParam("origin")
	destination = c.QueryParam("destination")

	if strings.TrimSpace(origin) == "" {
		return origin, destination, domain.ErrInvalidRouteOrigin
	}

	if strings.TrimSpace(destination) == "" {
		return origin, destination, domain.ErrInvalidRouteDestination
	}

	return origin, destination, nil
}

// HandleShortestRoute returns the shortest route betweem two airports
func (ctr *Route) HandleShortestRoute(c echo.Context) error {
	origin, destination, err := getOriginAndDestination(c)
	if err != nil {
		return errorResult(c, err)
	}

	shortesRoute, err := ctr.routerService.GetShortestRoute(origin, destination)
	if err != nil {
		if err == domain.ErrNotFound {
			err = errors.NewNotFoundError("no route exists between origin and destination")
		}
		return errorResult(c, err)
	}

	result := viewmodel.BuildRoute(shortesRoute)

	return c.JSON(http.StatusOK, result)
}
