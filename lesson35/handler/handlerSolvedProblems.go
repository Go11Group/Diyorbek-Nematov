package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solution/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) GetSolvedProblems(w http.ResponseWriter, r *http.Request) {
	solvedProblems, err := h.SolvedProblem.GetAllSolvedProblems()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving solved_problems: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(solvedProblems)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding solved problems to JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetSolvedProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userId"])

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid user ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	problemID, err := strconv.Atoi(params["problemId"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid problem ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	solvedProblem, err := h.SolvedProblem.GetSolvedProblemById(userID, problemID)

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid solved problem ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solvedProblem)
}

func (h *Handler) CreateSolvedProblem(w http.ResponseWriter, r *http.Request) {
	var solvedProblem models.SolvedProblem
	err := json.NewDecoder(r.Body).Decode(&solvedProblem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error decoding solved problem data: %s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	err = h.SolvedProblem.CreateSolvedProblem(solvedProblem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error creating solved problem: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateSolvedProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	solvedProblemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid solved problem ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	var updateSolvedProblem models.SolvedProblem
	err = json.NewDecoder(r.Body).Decode(&updateSolvedProblem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error decoding solved problem data: %s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	updateSolvedProblem.ID = solvedProblemID

	err = h.SolvedProblem.UpdateSolvedProblem(updateSolvedProblem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error updating solved problem: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func (h *Handler) DeleteSolvedProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	solvedProblemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid solved problem ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	err = h.SolvedProblem.DeleteSolvedProblem(solvedProblemID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error deleting solved problem: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
