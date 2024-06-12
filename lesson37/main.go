package main

import (
	"log"
	"my_module/gin/handler"
	postgres "my_module/gin/postgres/storage"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Databse connectionda xatolik bor, ", err.Error())
	}

	u := postgres.NewUserRepo(db)

	r := handler.NewHandler(handler.Handler{User: u})

	err = r.Run(":8080")

	if err != nil {
		log.Fatal("Ginni run qilishda xatolik bor, ", err.Error())
	}
	
}
