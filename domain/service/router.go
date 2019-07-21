package service

import (
	"github.com/rcdmk/shortest-flight-path/domain"
	"github.com/rcdmk/shortest-flight-path/domain/contract"
	"github.com/rcdmk/shortest-flight-path/domain/entity"
)

// router service is responsible for managing route data
type router struct {
	db contract.DataManager
}

// NewRouter returns a new Router service instance
func NewRouter(db contract.DataManager) contract.RouterService {
	return &router{
		db: db,
	}
}

// GetShortestRoute returns the shortest route between two airports
func (r *router) GetShortestRoute(sourceAirportIATA3 string, destAirportIATA3 string) (stops []entity.Route, err error) {
	if sourceAirportIATA3 == destAirportIATA3 {
		return nil, domain.ErrSameRouteSourceAndDestination
	}

	_, err = r.db.Airports().GetByCode(sourceAirportIATA3)
	if err != nil {
		if err == domain.ErrNotFound {
			err = domain.ErrInvalidRouteOrigin
		}

		return nil, err
	}

	_, err = r.db.Airports().GetByCode(destAirportIATA3)
	if err != nil {
		if err == domain.ErrNotFound {
			err = domain.ErrInvalidRouteDestination
		}

		return nil, err
	}

	stops = make([]entity.Route, 0)

	routes, err := r.db.Routes().GetAllDepartingFromAirport(sourceAirportIATA3)
	if err != nil {
		return nil, err
	}

	for _, route := range routes {
		if route.Destination == destAirportIATA3 {
			stops = append(stops, route)
		}
	}

	if len(stops) == 0 {
		return nil, domain.ErrNotFound
	}

	return stops, nil
}
