package handler

import (
	"net/http"
	"transaction/postgres"
)

type Handler struct {
	User        *postgres.UserRepo
	Product     *postgres.ProductRepo
	UserProduct *postgres.UserProductRepo
}

func NewHandler(handler Handler) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/users", handler.createUser)
	mux.HandleFunc("GET /api/users", handler.getUsers)
	mux.HandleFunc("GET /api/users/", handler.getUser)
	mux.HandleFunc("DELETE /api/users/{id}", handler.deleteUser)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
