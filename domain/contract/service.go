package contract

import (
	"github.com/rcdmk/shortest-flight-path/domain/entity"
)

// RouterService holds methods for router domain service
type RouterService interface {
	GetShortestRoute(sourceAirportIATA3 string, destAirportIATA3 string) []entity.Route
}
