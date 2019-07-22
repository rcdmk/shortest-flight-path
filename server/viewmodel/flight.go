package viewmodel

import (
	"github.com/rcdmk/shortest-flight-path/domain/entity"
)

// Flight model for API binding
type Flight struct {
	Origin      string `json:"origin,omitempty"`
	Destination string `json:"destination,omitempty"`
	Airline     string `json:"airline,omitempty"`
}

// BuildFlight returns a Flight object ready for response
func BuildFlight(route entity.Route) Flight {
	return Flight{
		Origin:      route.Origin,
		Destination: route.Destination,
		Airline:     route.AirlineCode,
	}
}
