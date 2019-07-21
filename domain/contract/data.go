package contract

import (
	"github.com/rcdmk/shortest-flight-path/domain/entity"
)

// DataManager is the core data handling interface that holds repository references
type DataManager interface {
	Airlines() AirlineRepo
	Airports() AirportRepo
}

// AirlineRepo holds methods to fetch airline data
type AirlineRepo interface {
	GetByID(airlineID string) entity.Airline
}

// AirportRepo holds methods to fetch airport data
type AirportRepo interface {
	GetByCode(iata3 string) entity.Airport
}

// RouteRepo holds methods to fetch route data
type RouteRepo interface {
	GetAllDepartingFromAirport(airportIATA3 string) []entity.Route
}
