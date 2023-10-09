package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	err := initConfig()
	if err != nil {
		logrus.Fatalf("error initializing configs:%s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
