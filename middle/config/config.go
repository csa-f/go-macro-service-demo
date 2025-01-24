package config

import commonConfig "github.com/csa-f/go-macro-service-demo/common/config"

type Config struct {
	*commonConfig.Config
}

func Get() *Config {
	return &Config{commonConfig.Get("app")}
}
