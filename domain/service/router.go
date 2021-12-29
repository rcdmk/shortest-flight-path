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

	stops, errPath := r.getShortestPath(sourceAirportIATA3, destAirportIATA3)
	if errPath != nil {
		return nil, errPath
	}

	if len(stops) == 0 {
		return nil, domain.ErrNotFound
	}

	return stops, nil
}

func (r *router) getShortestPath(sourceAirportIATA3 string, destAirportIATA3 string) ([]entity.Route, error) {
	stops := make([]entity.Route, 0)

	var routeQueue = []entity.Route{
		{
			Destination: sourceAirportIATA3,
		},
	}

	var shortestPaths = map[string][]entity.Route{}
	var visitedAirports = map[string]struct{}{
		sourceAirportIATA3: {},
	}

	for len(routeQueue) > 0 {
		current := routeQueue[0]
		routeQueue = routeQueue[1:]

		connections, err := r.db.Routes().GetAllDepartingFromAirport(current.Destination)
		if err != nil {
			return nil, err
		}

		for _, connection := range connections {
			if connection.Destination == destAirportIATA3 {
				// found complete route
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

	return stops, nil
}
