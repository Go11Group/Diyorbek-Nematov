package main

import (
	"client/api"
	"client/service"
	"log"
)

func main() {
	services := service.New()

	router := api.Router(services)

	log.Println("")

	log.Fatal(router.Run())
}
