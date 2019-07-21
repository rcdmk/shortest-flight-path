package config

// Config holds configuration values for the entire application
type Config struct {
	DB DBConfig `mapstructure:"db"`
}

// DBConfig holds DB related configuration options
// In a production scenario the secrets would be loaded from AWS SSM
// and the config files would only have the secret keys instead.
type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Schema   string `mapstructure:"schema"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}
