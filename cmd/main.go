package main

import (
	"log"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/tarpal7/all-commerce"
	"github.com/tarpal7/all-commerce/pkg/handler"
	"github.com/tarpal7/all-commerce/pkg/repository"
	"github.com/tarpal7/all-commerce/pkg/service"
)
func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: 	  viper.GetString("db.host"),
		Port: 	  viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
			
	srv := new(allcommerce.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}