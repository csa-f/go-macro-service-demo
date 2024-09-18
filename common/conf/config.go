package conf

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	config *Config
)

type Config struct {
	Viper  *viper.Viper
	Server *Server
}

func GetConfig() *Config {
	return config
}

func LoadConfig(serverName string) {
	config := Config{
		Viper: viper.New(),
	}
	workDir, _ := os.Getwd()
	config.Viper.SetConfigName(serverName)
	config.Viper.SetConfigType("yml")
	config.Viper.AddConfigPath(workDir + "/conf")
	err := config.Viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}
