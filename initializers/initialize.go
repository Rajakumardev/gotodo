package initializers

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PORT       string `mapstructure:"PORT"`
	DBHOST     string `mapstructure:"DBHOST"`
	DBUSER     string `mapstructure:"DBUSER"`
	DBPASSWORD string `mapstructure:"DBPASSWORD"`
	DBPORT     string `mapstructure:"DBPORT"`
}

func LoadConfig(path string) (config Config, err error) {

	viper := viper.New()

	viper.AddConfigPath(path)
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
