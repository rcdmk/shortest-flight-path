package mockcache

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

func New() *Cache {
	return &Cache{}
}

// Cache is a mocked cache manager instrumented for testing
type Cache struct {
	mock.Mock
}

// Get retrieves an entry from the cache
func (cm *Cache) Get(ctx context.Context, key string) (string, error) {
	args := cm.Called(ctx, key)

	return args.String(0), args.Error(1)
}

// Set adds an entry to the cache
func (cm *Cache) Set(ctx context.Context, key string, value string) error {
	args := cm.Called(ctx, key, value)

	return args.Error(0)
}

// GetStruct retrieves an entry from the cache and unmarshal it onto a struct pointer
func (cm *Cache) GetStruct(ctx context.Context, key string, destination interface{}) error {
	args := cm.Called(ctx, key, destination)

	return args.Error(0)
}

// SetStruct marshals a struct and adds it to the cache
func (cm *Cache) SetStruct(ctx context.Context, key string, value interface{}) error {
	args := cm.Called(ctx, key, value)

	return args.Error(0)
}

// Expire sets the expiration time for an entry
func (cm *Cache) Expire(ctx context.Context, key string, duration time.Duration) error {
	args := cm.Called(ctx, key, duration)

	return args.Error(0)
}

// Invalidate removes an item from the cache
func (cm *Cache) Invalidate(ctx context.Context, key string) error {
	args := cm.Called(ctx, key)

	return args.Error(0)
}
