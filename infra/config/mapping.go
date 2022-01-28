package config

import "time"

// Config holds configuration values for the entire application
type Config struct {
	DB     DBConfig     `mapstructure:"db"`
	Cache  CacheConfig  `mapstructure:"cache"`
	Server ServerConfig `mapstructure:"server"`
}

// DBConfig holds DB related configuration options
// In a production scenario the secrets would be loaded from AWS SSM
// and the config files would only have the secret keys instead.
type DBConfig struct {
	Host         string        `mapstructure:"host"`
	Port         int           `mapstructure:"port"`
	Schema       string        `mapstructure:"schema"`
	User         string        `mapstructure:"user"`
	Password     string        `mapstructure:"password"`
	MaxConnLife  time.Duration `mapstructure:"max-conn-life"`
	MaxIdleConns int           `mapstructure:"max-idle-conns"`
	MaxOpenConns int           `mapstructure:"max-open-conns"`
}

// CacheConfig holds ache related configuration options
// In a production scenario the secrets would be loaded from AWS SSM
// and the config files would only have the secret keys instead.
type CacheConfig struct {
	Host                  string        `mapstructure:"host"`
	Port                  int           `mapstructure:"port"`
	DBNumber              int           `mapstructure:"db-number"`
	Password              string        `mapstructure:"password"`
	DefaultExpirationTime time.Duration `mapstructure:"default-expiration-time"`
	MaxConnLife           time.Duration `mapstructure:"max-conn-life"`
	MaxOpenConns          int           `mapstructure:"max-open-conns"`
	MinIdleConns          int           `mapstructure:"min-idle-conns"`
	Prefix                string        `mapstructure:"prefix"`
}

// ServerConfig holds API server configuration options
type ServerConfig struct {
	Port   int    `mapstructure:"port"`
	Prefix string `mapstructure:"prefix"`
}
