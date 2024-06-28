package main

import (
	"client/api"
	"client/service"
	"log"
)

func main() {
	services := service.New()

	router := api.Router(services)

	log.Println("Starting server on port 8080 ... ")
	log.Fatal(router.Run(":8080"))
}
