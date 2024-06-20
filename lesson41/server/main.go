package main

import (
	"log"
	"net/http"
	"server/api"
	"server/api/handler"
	"server/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	emp := postgres.NewEmployeeRepo(db)
	hand := handler.Handler{Emp: *emp}

	mux := api.Router(hand)

	log.Println("Server is running on :8080 ...")
	if err = http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}