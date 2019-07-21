package datamysql

import (
	"database/sql"

	"github.com/rcdmk/shortest-flight-path/domain/contract"
)

// conn interacts with low level MySQL connection pool
type conn struct {
	db *sql.DB

	airlines *airlineRepo
	airports *airportRepo
	routes   *routeRepo
}

// Airlines returns an airlines repo instance
func (conn *conn) Airlines() contract.AirlineRepo {
	return conn.airlines
}

// Airports returns an airport repo instance
func (conn *conn) Airports() contract.AirportRepo {
	return conn.airports
}

// Routes returns a route repo instance
func (conn *conn) Routes() contract.RouteRepo {
	return conn.routes
}
