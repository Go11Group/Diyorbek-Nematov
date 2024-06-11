package main

import (
	"log"
	"my_module/handler"
	postgres "my_module/postgres/storege"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Error open database", err.Error())
	}
	defer db.Close()

	u := postgres.NewUserRepo(db)

	r := handler.NewHandler(handler.Handler{User: u})

	r.Run(":8080")
}
