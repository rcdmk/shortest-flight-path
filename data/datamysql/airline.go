package datamysql

import (
	"github.com/rcdmk/shortest-flight-path/domain/entity"
)

// airlineRepo holds methods to fetch airline data
type airlineRepo struct {
	conn executor
}

// newAirlineRepo returns a new airline repo instance
func newAirlineRepo(conn executor) *airlineRepo {
	return &airlineRepo{
		conn: conn,
	}
}

// GetByCode returns an airline instance with the given two letter code
func (repo *airlineRepo) GetByCode(airlineCode string) (airline entity.Airline, err error) {
	const query = `
		SELECT 		a.name,
					a.code,
					a.threeDigitCode,
					c.name
		
		FROM 		tb_airline 		a

		INNER JOIN 	tb_country 		c
			ON 		a.country_id 	= 	c.country_id

		WHERE 		a.code 			= 	?
		
		LIMIT 1;
	`

	row := repo.conn.QueryRow(query, airlineCode)

	err = row.Scan(
		&airline.Name,
		&airline.Code,
		&airline.ThreeDigitCode,
		&airline.Country,
	)
	if err != nil {
		return airline, parseError(err)
	}

	return airline, nil
}
