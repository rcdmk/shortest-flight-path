package rediscache

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/msgpack/v5"

	"github.com/rcdmk/shortest-flight-path/domain"
	"github.com/rcdmk/shortest-flight-path/infra/config"
)

func New(cfg config.CacheConfig) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Host + ":" + strconv.Itoa(cfg.Port),
		Password:     cfg.Password,
		DB:           cfg.DBNumber,
		MaxConnAge:   cfg.MaxConnLife,
		MinIdleConns: cfg.MinIdleConns,
		PoolSize:     cfg.MaxOpenConns,
	})

	err := client.Ping(context.Background()).Err()

	cm := &Cache{
		redis: client,

		defaultExpirationTime: cfg.DefaultExpirationTime,
		prefix:                cfg.Prefix,
	}

	return cm, err
}

type Cache struct {
	redis *redis.Client

	defaultExpirationTime time.Duration
	prefix                string
}

func (cm *Cache) buildCacheKey(key string) string {
	return cm.prefix + key
}

// Get retrieves an entry from the cache
func (cm *Cache) Get(ctx context.Context, key string) (string, error) {
	key = cm.buildCacheKey(key)

	value, err := cm.redis.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", domain.ErrCacheMiss
		}

		return "", fmt.Errorf("cache: error fetching from cache - %w", err)
	}

	return value, nil
}

// Set adds an entry to the cache
func (cm *Cache) Set(ctx context.Context, key string, value string) error {
	key = cm.buildCacheKey(key)

	err := cm.redis.Set(ctx, key, value, cm.defaultExpirationTime).Err()
	if err != nil {
		return fmt.Errorf("cache: error adding entry to cache - %w", err)
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
		return fmt.Errorf("cache: error marshalling cache entry - %w", err)
	}

	err = cm.Set(ctx, key, string(bytes))
	if err != nil {
		return err
	}

	return nil
}

// Expire sets the expiration time for an entry
func (cm *Cache) Expire(ctx context.Context, key string, duration time.Duration) error {
	key = cm.buildCacheKey(key)

	err := cm.redis.Expire(ctx, key, duration).Err()
	if err != nil {
		return fmt.Errorf("cache: error setting key expiration - %w", err)
	}

	return nil
}

// Invalidate removes an item from the cache
func (cm *Cache) Invalidate(ctx context.Context, key string) error {
	key = cm.buildCacheKey(key)

	err := cm.redis.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("cache: error invalidating key - %w", err)
	}

	return nil
}
