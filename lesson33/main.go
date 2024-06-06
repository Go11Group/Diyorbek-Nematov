package main

import (
	"net/http"
)

var Message string

func main() {
	http.HandleFunc("/http/methods", httpMethods)
	http.ListenAndServe(":8080", nil)
}

func httpMethods(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handelGet(w, r)
	case http.MethodPost:
		handelPost(w, r)
	case http.MethodPut:
		handelPut(w, r)
	case http.MethodDelete:
		handelDelete(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handelGet(w http.ResponseWriter, r *http.Request) {
	Message = "GET methodi orqali ishladi"
	w.Write([]byte(r.Method + ": " + Message))
}

func handelPost(w http.ResponseWriter, r *http.Request) {
	Message = "POST Methodi orqali ishadi"
	w.Write([]byte(r.Method + ": " + Message))
}

func handelPut(w http.ResponseWriter, r *http.Request) {
	Message = "PUT Methodi orqali ishladi"
	w.Write([]byte(r.Method + ": " + Message))
}

func handelDelete(w http.ResponseWriter, r *http.Request) {
	Message = "DELETE Methodi orqali ishladi"
	w.Write([]byte(r.Method + ": " + Message))
}
