package datamysql

import (
	"github.com/rcdmk/shortest-flight-path/domain/entity"
)

// airportRepo holds methods to fetch airport data
type airportRepo struct {
	conn executor
}

// newAirportRepo returns a new airport repo instance
func newAirportRepo(conn executor) *airportRepo {
	return &airportRepo{
		conn: conn,
	}
}

// GetByCode returns an airport instance with the given IATA 3 code
func (repo *airportRepo) GetByCode(iata3 string) (airport entity.Airport, err error) {
	const query = `
		SELECT 		a.name,
					a.city,
					a.country,
					a.iata3,
					a.latitude,
					a.longitude
		
		FROM 		tb_airport 	a

		WHERE 		a.iata3 	= 	?
		
		LIMIT 1;
	`

	row := repo.conn.QueryRow(query, iata3)

	err = row.Scan(
		&airport.Name,
		&airport.City,
		&airport.Country,
		&airport.IATA3Code,
		&airport.Latitude,
		&airport.Longitude,
	)
	if err != nil {
		return airport, parseError(err)
	}

	return airport, nil
}
