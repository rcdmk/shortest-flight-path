package cache

import (
	"github.com/rcdmk/shortest-flight-path/domain/contract"
	"github.com/rcdmk/shortest-flight-path/infra/cache/rediscache"
	"github.com/rcdmk/shortest-flight-path/infra/config"
)

// New returns a CacheManager instance
func New(cfg config.CacheConfig) (contract.CacheManager, error) {
	cm, err := rediscache.New(cfg)
	if err != nil {
		return cm, err
	}

	return cm, nil
}
