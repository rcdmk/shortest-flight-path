package datamysql

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/rcdmk/shortest-flight-path/domain"
	"github.com/rcdmk/shortest-flight-path/infra/config"
)

// buildConfig returns a MySQL driver config object instance
func buildConfig(cfg config.DBConfig) *mysql.Config {
	config := mysql.NewConfig()

	config.Net = "tcp"
	config.Addr = fmt.Sprintf("%s:%v", cfg.Host, cfg.Port)

	config.DBName = cfg.Schema
	config.User = cfg.User
	config.Passwd = cfg.Password

	config.ParseTime = true
	config.Params = map[string]string{
		// "transaction_isolation": "'READ-COMMITTED'",
	}

	return config
}

// connect creates a new MySQL connection pool, connects it to the database and tests the connection
func connect(cfg config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", buildConfig(cfg).FormatDSN())
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(cfg.MaxConnLife)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}

// parseError maps common DB errors to application and domain errors
// decoupling application code from storage implementation
func parseError(err error) error {
	if err == sql.ErrNoRows {
		return domain.ErrNotFound
	}

	return err
}
