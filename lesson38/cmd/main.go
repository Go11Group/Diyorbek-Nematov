package main

import (
	"learning_app/api"
	"learning_app/storage/postgres"
	"log"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	
	router := api.Router(db)

	log.Fatal(router.Run())
}
