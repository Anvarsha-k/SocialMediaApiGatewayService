package config_apigw

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	AuthSvcUrl string `mapstructure:"AUTH_SVC_URL"`
}

func LoadConfig() (*Config, error) {
	var c Config

	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("------", err)
		return nil, err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("-----", err)
		return nil,err
	}
	fmt.Println(c)
	return &c,nil
}
