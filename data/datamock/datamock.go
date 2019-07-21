package datamock

import (
	"github.com/rcdmk/shortest-flight-path/domain/contract"
)

// New returns a mocked data manager implementation
func New() *DataManager {
	return &DataManager{
		airlines: new(AirlineRepo),
		airports: new(AirportRepo),
		routes:   new(RouteRepo),
	}
}

// DataManager is the mocked data manager implementation
type DataManager struct {
	airlines *AirlineRepo
	airports *AirportRepo
	routes   *RouteRepo
}

// Airlines returns an airlines repo instance
func (dm *DataManager) Airlines() contract.AirlineRepo {
	return dm.airlines
}

// Airports returns an airport repo instance
func (dm *DataManager) Airports() contract.AirportRepo {
	return dm.airports
}

// Routes returns a route repo instance
func (dm *DataManager) Routes() contract.RouteRepo {
	return dm.routes
}
