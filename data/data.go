package data

import (
	"github.com/rcdmk/shortest-flight-path/data/datamysql"
	"github.com/rcdmk/shortest-flight-path/domain/contract"
	"github.com/rcdmk/shortest-flight-path/infra/config"
)

// New returns a new initialized data manager instance
func New(cfg config.DBConfig) (contract.DataManager, error) {
	dm, err := datamysql.New(cfg)
	if err != nil {
		return nil, err
	}

	return dm, nil
}
