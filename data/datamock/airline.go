package datamock

import (
	"github.com/rcdmk/shortest-flight-path/domain/entity"
	"github.com/stretchr/testify/mock"
)

// AirlineRepo holds methods to fetch airline data
type AirlineRepo struct {
	mock.Mock
}

// GetByCode returns an airline instance with the given two letter code
func (repo *AirlineRepo) GetByCode(airlineCode string) (airline entity.Airline, err error) {
	args := repo.Called(airlineCode)
	return args.Get(0).(entity.Airline), args.Error(1)
}
