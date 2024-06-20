package api

import (
	"net/http"
	"server/api/handler"
)

func Router(handler handler.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /employees/", handler.GetEmployeeHandler)
	mux.HandleFunc("POST /employees", handler.CreateEmployeeHandler)
	mux.HandleFunc("PUT /employees/", handler.UpdateEmployeeHandler)
	mux.HandleFunc("DELETE /employees/", handler.DeleteEmployeeHandler)

	return mux
}