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

func Get(fileName string) *Config {
	c := &Config{
		Viper: viper.New(),
	}
	workDir, _ := os.Getwd()
	c.Viper.SetConfigName(fileName)
	c.Viper.SetConfigType("yml")
	c.Viper.AddConfigPath(workDir + "/config")
	err := c.Viper.ReadInConfig()

	if err != nil {
		log.Fatalln(err)
	}
	if err := c.Viper.Unmarshal(c); err != nil {
		log.Fatalln(err)
	}
	return c
}
