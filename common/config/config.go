package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Viper  *viper.Viper
	Server *Server
	DB     *DB
	Redis  *Redis
	Log    *Log
}

const (
	red    = "38;2;255;0;0"
	yellow = "38;2;255;165;0"
	gray   = "38;2;100;100;100"
	def    = "38;2;100;150;200"
	green  = "38;2;85;107;47"
)

func Get(fileName string) *Config {
	v := viper.New()

	workDir, _ := os.Getwd()
	v.SetConfigName(fileName)
	v.SetConfigType("yml")
	v.AddConfigPath(workDir + "/config")

	v.SetDefault("log.color.debug", green)
	v.SetDefault("log.color.trace", green)
	v.SetDefault("log.color.warn", yellow)
	v.SetDefault("log.color.error", red)
	v.SetDefault("log.color.fatal", red)
	v.SetDefault("log.color.panic", red)
	v.SetDefault("log.color.panic", red)
	v.SetDefault("log.color.info", def)

	c := &Config{
		Viper: v,
	}

	err := v.ReadInConfig()

	if err != nil {
		log.Fatalln(err)
	}
	if err := c.Viper.Unmarshal(c); err != nil {
		log.Fatalln(err)
	}
	return c
}
