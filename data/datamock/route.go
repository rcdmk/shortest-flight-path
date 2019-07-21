package datamock

import (
	"github.com/rcdmk/shortest-flight-path/domain/entity"
	"github.com/stretchr/testify/mock"
)

// RouteRepo holds methods to fetch route data
type RouteRepo struct {
	mock.Mock
}

// GetAllDepartingFromAirport returns a list of destination routes from a given source airport
func (repo *RouteRepo) GetAllDepartingFromAirport(airportIATA3 string) (routes []entity.Route, err error) {
	args := repo.Called(airportIATA3)
	return args.Get(0).([]entity.Route), args.Error(1)
}
