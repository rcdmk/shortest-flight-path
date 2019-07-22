package config

import (
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("api")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath(path.Dir(os.Args[0]))
	viper.AddConfigPath(".")
}

// Load loads the configuration values from disk and environment variables and returns a config struct
func Load() (cfg *Config, err error) {
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg = new(Config)

	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
