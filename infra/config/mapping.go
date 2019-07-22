package config

import "time"

// Config holds configuration values for the entire application
type Config struct {
	DB     DBConfig     `mapstructure:"db"`
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

// ServerConfig holds API server configuration options
type ServerConfig struct {
	Port   int    `mapstructure:"port"`
	Prefix string `mapstructure:"prefix"`
}
