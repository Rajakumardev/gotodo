package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PORT string `mapstructure:"PORT"`
}

func load_config(path string) (config Config, err error) {

	viper := viper.New()

	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	config_read_err := viper.ReadInConfig()
	if config_read_err != nil {
		log.Fatal("[!]Unable to read environment configuration...")
		return
	}

	config_read_err = viper.Unmarshal(&config)
	if config_read_err != nil {
		log.Fatal("[!]Unable to copy environment configuratoin...")
		return
	}
	return
}
