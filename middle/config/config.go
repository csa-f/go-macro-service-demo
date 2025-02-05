package config

import commonConfig "github.com/csa-f/macro-service/common/config"

type Config struct {
	*commonConfig.Config
}

func Get() *Config {
	return &Config{commonConfig.Get("app")}
}
