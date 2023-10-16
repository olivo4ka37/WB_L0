package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/olivo4ka37/WB_L0/internal/cache"
	"github.com/olivo4ka37/WB_L0/internal/models"
	"github.com/olivo4ka37/WB_L0/pkg/handler"
	"github.com/olivo4ka37/WB_L0/pkg/repository"
	"github.com/olivo4ka37/WB_L0/pkg/service"
	"github.com/olivo4ka37/WB_L0/server"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

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

	defer database.Close()

	cache := cache.NewCache()

	repos := repository.NewRepository(database)
	services := service.NewService(repos, cache)
	handlers := handler.NewHandler(services)
	server := new(WB_L0.Server)

	logrus.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running server: %s", err.Error())
		}
	}()

	nc, err := stan.Connect("test-cluster", "subscriber", stan.NatsURL("0.0.0.0:4222"))
	if err != nil {
		logrus.Fatal(err)
	}

	defer nc.Close()

	_, err = nc.Subscribe("order", func(m *stan.Msg) {
		fmt.Print(string(m.Data))
		var order models.Order
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Println("not valid json, cant unmarshal it")
			return
		}
		if err = (*service.OrderService).Create(services.ServiceOrder, ctx, order.OrderUID, order); err != nil {
			log.Println(err)
			return
		}
	})
	if err != nil {
		log.Println(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := server.Shutdown(ctx); err != nil {
		logrus.Print("failed to stop server:")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
