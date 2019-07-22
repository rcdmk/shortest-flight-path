package datamock

import (
	"github.com/rcdmk/shortest-flight-path/domain/contract"
	"github.com/stretchr/testify/mock"
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
	mock.Mock

	airlines *AirlineRepo
	airports *AirportRepo
	routes   *RouteRepo
}

// Close simlates closing the connection
func (dm *DataManager) Close() error {
	args := dm.Called()
	return args.Error(0)
}

// Airlines returns an airlines repo instance
func (dm *DataManager) Airlines() contract.AirlineRepo {
	dm.Called()
	return dm.airlines
}

// Airports returns an airport repo instance
func (dm *DataManager) Airports() contract.AirportRepo {
	dm.Called()
	return dm.airports
}

// Routes returns a route repo instance
func (dm *DataManager) Routes() contract.RouteRepo {
	dm.Called()
	return dm.routes
}
