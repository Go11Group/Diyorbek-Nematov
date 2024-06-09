package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solution/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) GetAllProblems(w http.ResponseWriter, r *http.Request) {
	problems, err := h.Problem.GetAllProblems()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving problems: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(problems)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding problems to JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	problemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid problem ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	problem, err := h.Problem.GetProblemByID(problemID)

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid problem ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(problem)
}

func (h *Handler) CreateProblem(w http.ResponseWriter, r *http.Request) {
	var problem models.Problem
	err := json.NewDecoder(r.Body).Decode(&problem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error decoding problem data: %s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	err = h.Problem.CreateProblem(problem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error creating problem: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	problemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid problem ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	var updateProblem models.Problem
	err = json.NewDecoder(r.Body).Decode(&updateProblem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error decoding problem data: %s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	updateProblem.ID = problemID

	err = h.Problem.UpdateProblem(updateProblem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error updating problem: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func (h *Handler) DeleteProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	problemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid problem ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	err = h.Problem.DeleteProblem(problemID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error deleting problem: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
