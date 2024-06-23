package main

import (
	"log"
	"metro-service/api"
	"metro-service/api/handler"
	"metro-service/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := api.Router(handler.NewHandler(db))

	log.Fatal(router.Run(":8082"))
}