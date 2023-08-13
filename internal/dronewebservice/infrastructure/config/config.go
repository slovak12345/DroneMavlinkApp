package config

import (
	"github.com/spf13/viper"
)

const configFile = "dronewebservice.conf"

type Config struct {
	HTTPHost string
	HTTPPort string
	Dialect  string
}

func NewConfig() (*Config, error) {
	viper.SetConfigName(configFile)
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, err
	}

	c := &Config{
		HTTPHost: viper.GetString("http.host"),
		HTTPPort: viper.GetString("http.port"),
		Dialect:  viper.GetString("http.implementation"),
	}

	return c, nil
}
