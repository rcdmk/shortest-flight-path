// Package service holds implementations of domain service interfaces that deal with
// domain logic and business rules.
package service

import (
	"github.com/rcdmk/shortest-flight-path/domain/contract"
)

// NewRouter returns a new Router service instance
func NewRouter(db contract.DataManager) contract.RouterService {
	return &router{
		db: db,
	}
}
