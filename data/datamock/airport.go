package datamock

import (
	"github.com/rcdmk/shortest-flight-path/domain/entity"
	"github.com/stretchr/testify/mock"
)

// AirportRepo holds methods to fetch airport data
type AirportRepo struct {
	mock.Mock
}

// GetByCode returns an airport instance with the given two letter code
func (repo *AirportRepo) GetByCode(airportCode string) (airport entity.Airport, err error) {
	args := repo.Called(airportCode)
	return args.Get(0).(entity.Airport), args.Error(1)
}
