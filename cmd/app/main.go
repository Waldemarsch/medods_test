package main

import (
	"context"
	"fmt"
	server "github.com/Waldemarsch/medods_test"
	handlers "github.com/Waldemarsch/medods_test/internal/handler"
	"github.com/Waldemarsch/medods_test/internal/infrastructure"
	"github.com/Waldemarsch/medods_test/internal/service"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatal("Error while initializing configs: ", err)
	}
	addressMongo := viper.GetStringMap("mongodb.address")
	var mongoURI string
	if viper.GetString("mongodb.credentials.exists") == "1" {
		creds := viper.GetStringMap("mongodb.credentials")
		mongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
			creds["login"],
			creds["password"],
			addressMongo["ip"],
			addressMongo["port"],
			addressMongo["db"])
	} else {
		mongoURI = fmt.Sprintf("mongodb://%s:%s/%s",
			addressMongo["ip"],
			addressMongo["port"],
			addressMongo["db"])
	}

	inf := infrastructure.NewInfrastructure(mongoURI)
	usecase := service.NewService(inf)
	handler := handlers.NewHandler(usecase)

	ctx := context.Background()

	srvr := new(server.Server)
	go func() {
		if err := srvr.Run(viper.GetString("server.port"), handler.InitRoutes()); err != nil {
			log.Fatal("Error while running server: ", err)
		}
	}()

	log.Println("App has successfully started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srvr.Shutdown(ctx); err != nil {
		log.Fatal("Error occurred on server shutdown: ", err)
	}

	if err := inf.Repository.CloseDB(ctx); err != nil {
		log.Fatal("Error occurred while disconnecting from DB: ", err)
	}

}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
