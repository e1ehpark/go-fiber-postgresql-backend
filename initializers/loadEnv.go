package initializers

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"POSTGRES__HOST"`
	DBUserName string `mapstructure:"POSTGRES__USER"`
	DBUserPass string `mapstructure:"POSTGRES__PASSWORD"`
	DBName     string `mapstructure:"POSTGRES__DB"`
	DBPort     string `mapstructure:"POSTGRES__PORT"`

	ClientOrigin string `mapstructure:"CLIENT__ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("env")
	viper.SetConfigType("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}