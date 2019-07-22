package viewmodel

import (
	"github.com/rcdmk/shortest-flight-path/domain/entity"
)

// Route model for API binding
type Route struct {
	Flights []Flight `json:"flights,omitempty"`
}

// BuildRoute returns a Route object ready for response
func BuildRoute(routes []entity.Route) Route {
	route := Route{}

	for _, entry := range routes {
		route.Flights = append(route.Flights, BuildFlight(entry))
	}

	return route
}
