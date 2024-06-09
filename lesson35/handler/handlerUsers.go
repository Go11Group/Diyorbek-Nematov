package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solution/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.User.GetAllUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving users: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding users to JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid user ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	user, err := h.User.GetUserByID(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid user ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error decoding user data: %s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	err = h.User.CreateUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error creating user: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid user ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	var updateUser models.User
	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error decoding user data: %s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	updateUser.ID = userID

	err = h.User.UpdateUser(updateUser)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error updating user: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid user ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	err = h.User.DeleteUser(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Error deleting user: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
