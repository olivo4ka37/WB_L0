package main

import (
	"github.com/olivo4ka37/WB_L0/internal/cache"
	"github.com/olivo4ka37/WB_L0/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	err := initConfig()
	if err != nil {
		logrus.Fatalf("error initializing configs:%s", err.Error())
	}

	database, err := repository.NewPostgresConnection(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to connect database:%s", err.Error())
	}

	cache := cache.NewCache()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
