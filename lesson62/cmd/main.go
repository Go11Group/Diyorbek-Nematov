package main

import (
	"log"
	"students/api"
	"students/api/handler"
	"students/storage/redis"
)

func main() {
	db := redis.ConnectRDB()

	router := api.NewRouter(handler.NewHandler(db))

	log.Fatal(router.Run())
}
