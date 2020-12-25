package main

import (
	"log"
	"github.com/tarpal7/all-commerce"
)
func main() {
	srv := new(allcommerce.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}