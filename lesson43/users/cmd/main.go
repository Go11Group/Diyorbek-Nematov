package main

import (
	"log"
	"user-service/api"
	"user-service/api/handler"
	"user-service/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := api.Router(handler.NewHandler(db))

	log.Fatal(router.Run(":8081"))
}