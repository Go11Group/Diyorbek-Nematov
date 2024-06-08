package handler

import (
	"net/http"
	"transaction/postgres"
)

type Handler struct {
	User        *postgres.UserRepo
	Product     *postgres.ProductRepo
}

func NewHandler(handler Handler) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/users", handler.createUser)
	mux.HandleFunc("GET /api/users", handler.getUsers)
	mux.HandleFunc("GET /api/users/", handler.getUser)
	mux.HandleFunc("PUT /api/users", handler.updateUser)
	mux.HandleFunc("DELETE /api/users/", handler.deleteUser)


	mux.HandleFunc("POST /api/products", handler.createProduct)
	mux.HandleFunc("GET /api/products", handler.getProducts)
	mux.HandleFunc("GET /api/products/", handler.getProduct)
	mux.HandleFunc("PUT /api/products", handler.updateProduct)
	mux.HandleFunc("DELETE /api/products/", handler.deleteProduct)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
