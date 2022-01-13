package memorycache

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/vmihailenco/msgpack/v5"

	"github.com/rcdmk/shortest-flight-path/domain"
)

func New() *Cache {
	return &Cache{
		entries:               make(map[string]cacheEntry),
		DefaultExpirationTime: time.Hour,
	}
}

type cacheEntry struct {
	Value      string
	Expiration time.Time
}

type Cache struct {
	entries map[string]cacheEntry

	DefaultExpirationTime time.Duration
}

func (cm *Cache) getEntry(ctx context.Context, key string) (cacheEntry, error) {
	entry, exists := cm.entries[key]
	if !exists || entry.Expiration.Before(time.Now()) {
		if exists && entry.Expiration.Before(time.Now()) {
			delete(cm.entries, key)
		}

		return entry, domain.ErrCacheMiss
	}

	return entry, nil
}

// Get retrieves an entry from the cache
func (cm *Cache) Get(ctx context.Context, key string) (string, error) {
	entry, err := cm.getEntry(ctx, key)
	if err != nil {
		return "", err
	}

	return entry.Value, nil
}

// Set adds an entry to the cache
func (cm *Cache) Set(ctx context.Context, key string, value string) error {
	cm.entries[key] = cacheEntry{
		Value:      value,
		Expiration: time.Now().Add(cm.DefaultExpirationTime),
	}

	return nil
}

// GetStruct retrieves an entry from the cache and unmarshal it onto a struct pointer
func (cm *Cache) GetStruct(ctx context.Context, key string, destination interface{}) error {
	destinationType := reflect.TypeOf(destination)
	if destinationType.Kind() != reflect.Ptr {
		return fmt.Errorf("cache: destination must be a pointer")
	}

	value, err := cm.Get(ctx, key)
	if err != nil {
		return err
	}

	err = msgpack.Unmarshal([]byte(value), destination)
	if err != nil {
		return fmt.Errorf("cache: error unsmarshalling cache entry - %w", err)
	}

	return nil
}

// SetStruct marshals a struct and adds it to the cache
func (cm *Cache) SetStruct(ctx context.Context, key string, value interface{}) error {
	bytes, err := msgpack.Marshal(value)
	if err != nil {
		return fmt.Errorf("cache: error unsmarshalling cache entry - %w", err)
	}

	err = cm.Set(ctx, key, string(bytes))
	if err != nil {
		return err
	}

	return nil
}

// Expire sets the expiration time for an entry
func (cm *Cache) Expire(ctx context.Context, key string, duration time.Duration) error {
	entry, err := cm.getEntry(ctx, key)
	if err != nil {
		if errors.Is(err, domain.ErrCacheMiss) {
			return nil
		}

		return fmt.Errorf("cache: error setting key expiration - %w", err)
	}

	entry.Expiration = time.Now().Add(duration)

	err = cm.SetStruct(ctx, key, entry)
	if err != nil {
		return fmt.Errorf("cache: error setting key expiration - %w", err)
	}

	return nil
}

// Invalidate removes an item from the cache
func (cm *Cache) Invalidate(ctx context.Context, key string) error {
	delete(cm.entries, key)

	return nil
}
