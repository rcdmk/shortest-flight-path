package cache

import (
	"github.com/rcdmk/shortest-flight-path/domain/contract"
	"github.com/rcdmk/shortest-flight-path/infra/cache/memorycache"
)

// New returns a CacheManager instance
func New() (contract.CacheManager, error) {
	cm := memorycache.New()

	return cm, nil
}
