package service

import (
	"math"

	"github.com/rcdmk/shortest-flight-path/domain"
	"github.com/rcdmk/shortest-flight-path/domain/contract"
	"github.com/rcdmk/shortest-flight-path/domain/entity"
)

// router service is responsible for managing route data
type router struct {
	db contract.DataManager
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

	var routeQueue = []entity.Route{
		entity.Route{
			Destination: sourceAirportIATA3,
		},
	}

	var shortestPaths = map[string][]entity.Route{}
	var visitedAirports = map[string]struct{}{
		sourceAirportIATA3: struct{}{},
	}

	for len(routeQueue) > 0 {
		current := routeQueue[0]
		routeQueue = routeQueue[1:]

		connections, err := r.db.Routes().GetAllDepartingFromAirport(current.Destination)
		if err != nil {
			return nil, err
		}

		for _, connection := range connections {
			// found complete route
			if connection.Destination == destAirportIATA3 {
				stops = shortestPaths[current.Destination]
				stops = append(stops, connection)
				return stops, nil
			}

			_, visited := visitedAirports[connection.Destination]
			if !visited {
				visitedAirports[connection.Destination] = struct{}{}

				connectionCount := math.MaxInt32

				connectionPath, exists := shortestPaths[connection.Destination]
				if exists {
					connectionCount = len(connectionPath)
				}

				currentPath := shortestPaths[current.Destination]
				currentCount := len(currentPath)

				if currentCount < connectionCount {
					shortestPaths[connection.Destination] = append(currentPath, connection)
				}

				routeQueue = append(routeQueue, connection)
			}
		}
	}

	if len(stops) == 0 {
		return nil, domain.ErrNotFound
	}

	return stops, nil
}
