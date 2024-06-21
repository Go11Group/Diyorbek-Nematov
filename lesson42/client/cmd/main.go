package main

import (
	"client/api"
	"client/api/handler"
	"log"
)

func main() {
	handler := handler.Handler{}

	router := api.Router(handler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
