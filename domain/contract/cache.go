package contract

import (
	"context"
	"time"
)

// CacheManager is the core cache handling interface to store and retrieve cache entries
type CacheManager interface {
	// Get retrieves an entry from the cache
	Get(ctx context.Context, key string) (string, error)
	// Set adds an entry to the cache
	Set(ctx context.Context, key string, value string) error

	// GetStruct retrieves an entry from the cache and unmarshal it onto a struct pointer
	GetStruct(ctx context.Context, key string, destination interface{}) error
	// SetStruct marshals a struct and adds it to the cache
	SetStruct(ctx context.Context, key string, value interface{}) error

	// Expire sets the expiration time for an entry
	Expire(ctx context.Context, key string, duration time.Duration) error
	// Invalidate removes an item from the cache
	Invalidate(ctx context.Context, key string) error
}
