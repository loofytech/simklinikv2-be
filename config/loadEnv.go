package config

import (
	"github.com/spf13/viper"
)

type ConfigDB struct {
	DBHost         string `mapstructure:"DB_HOST"`
	DBUserName     string `mapstructure:"DB_USER"`
	DBUserPassword string `mapstructure:"DB_PASS"`
	DBName         string `mapstructure:"DB_NAME"`
	DBPort         string `mapstructure:"DB_PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config ConfigDB, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
