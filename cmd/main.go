package main

import (
	"log"
	"github.com/tarpal7/all-commerce"
	"github.com.tarpal7/all-commerce/pkg/handler"
)
func main() {
	handlers := new(handler.Handler)
	srv := new(allcommerce.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}