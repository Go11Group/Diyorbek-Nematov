package main

import (
	"log"
	"transaction/handler"
	"transaction/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	u := postgres.NewUserRepo(db)
	p := postgres.NewProductRepo(db)

	server := handler.NewHandler(handler.Handler{User: u, Product: p})

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
