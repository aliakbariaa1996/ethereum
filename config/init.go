package config

import (
	"github.com/spf13/viper"
)

// Config Create private data struct to hold Config options.
type Config struct {
	Port string `yaml:"port"`
}

func InitConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./../config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		return conf, viper.ReadInConfig()
	}
	return conf, viper.ReadInConfig()
}
