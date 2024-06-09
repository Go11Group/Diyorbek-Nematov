package main

import (
	"log"
	"solution/handler"
	postgres "solution/postgres/storoge"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	u := postgres.NewUserRepo(db)
	p := postgres.NewProblemRepo(db)
	sp := postgres.NewSolvedProblemRepo(db)

	server := handler.NewHandler(handler.Handler{User: u, Problem: p, SolvedProblem: sp})

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
