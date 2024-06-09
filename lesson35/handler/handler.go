package handler

import (
	"net/http"
	postgres "solution/postgres/storoge"

	"github.com/gorilla/mux"
)

type Handler struct {
	User *postgres.UserRepo
	Problem *postgres.ProblemRepo
	SolvedProblem *postgres.SolvedProblemRepo
}

func NewHandler(handler Handler) *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/api/users", handler.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", handler.GetUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", handler.DeleteUser).Methods("DELETE")


    r.HandleFunc("/api/problems", handler.CreateProblem).Methods("POST")
    r.HandleFunc("/api/problems", handler.GetAllProblems).Methods("GET")
    r.HandleFunc("/api/problems/{id}", handler.GetProblem).Methods("GET")
    r.HandleFunc("/api/problems/{id}", handler.UpdateProblem).Methods("PUT")
    r.HandleFunc("/api/problems/{id}", handler.DeleteProblem).Methods("DELETE")

	r.HandleFunc("/solved_problems", handler.GetSolvedProblems).Methods("GET")
	r.HandleFunc("/users/{userId}/problems/{problemId}/solved_problem", handler.GetSolvedProblem).Methods("GET")
	r.HandleFunc("/solved_problems", handler.CreateSolvedProblem).Methods("POST")
	r.HandleFunc("/solved_problems/{id}", handler.UpdateSolvedProblem).Methods("PUT")
	r.HandleFunc("/solved_problems/{id}", handler.DeleteSolvedProblem).Methods("DELETE")
	

	return &http.Server{Addr: ":8080", Handler: r}
}