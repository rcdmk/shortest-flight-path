package datamysql

import (
	"github.com/rcdmk/shortest-flight-path/domain/entity"
)

// routeRepo holds methods to fetch route data
type routeRepo struct {
	conn executor
}

// newRouteRepo returns a new route repo instance
func newRouteRepo(conn executor) *routeRepo {
	return &routeRepo{
		conn: conn,
	}
}

// GetAllDepartingFromAirport returns a list of destination routes from a given source airport
func (repo *routeRepo) GetAllDepartingFromAirport(airportIATA3 string) (routes []entity.Route, err error) {
	const query = `
		SELECT 		r.airline_code
					r.origin,
					r.destination
		
		FROM 		tb_route 	r

		WHERE 		r.origin 	= 	?
		
		LIMIT 1;
	`

	rows, err := repo.conn.Query(query, airportIATA3)
	if err != nil {
		return routes, parseError(err)
	}

	routes = make([]entity.Route, 0)
	var route entity.Route

	for rows.Next() {
		err = rows.Scan(
			&route.AirlineCode,
			&route.Origin,
			&route.Destination,
		)
		if err != nil {
			return routes, parseError(err)
		}

		routes = append(routes, route)
	}

	return routes, nil
}
