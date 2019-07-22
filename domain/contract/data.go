package contract

import (
	"github.com/rcdmk/shortest-flight-path/domain/entity"
)

// DataManager is the core data handling interface that holds repository references
type DataManager interface {
	Close() error

	Airlines() AirlineRepo
	Airports() AirportRepo
	Routes() RouteRepo
}

// AirlineRepo holds methods to fetch airline data
type AirlineRepo interface {
	GetByCode(code string) (entity.Airline, error)
}

// AirportRepo holds methods to fetch airport data
type AirportRepo interface {
	GetByCode(iata3 string) (entity.Airport, error)
}

// RouteRepo holds methods to fetch route data
type RouteRepo interface {
	GetAllDepartingFromAirport(airportIATA3 string) ([]entity.Route, error)
}
