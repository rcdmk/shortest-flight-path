// Package datamysql implements the DataManager interface for MySQL storage
package datamysql

import (
	"database/sql"

	"github.com/rcdmk/shortest-flight-path/infra/config"
)

// New returns a new MySQL DataManager instance
func New(cfg config.DBConfig) (*conn, error) {
	db, err := connect(cfg)
	if err != nil {
		return nil, err
	}

	dm := &conn{
		db:       db,
		airlines: newAirlineRepo(db),
		airports: newAirportRepo(db),
		routes:   newRouteRepo(db),
	}

	return dm, nil
}

// executor maps sql pacakge methods common to pools and transactions
type executor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}
