package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"transaction/models"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while decoding JSON, err: %s", err.Error())))
		return
	}

	err = h.User.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while creating user, err: %s", err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created successfully"))
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.User.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while Decode, err: %s", err.Error())))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/users/")
	id, err := strconv.Atoi(idStr)
	fmt.Println(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.User.GetUserById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while decoding JSON, err: %s", err.Error())))
		return
	}

	err = h.User.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while updating user, err: %s", err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user updated successfully"))
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.User.DeleteUser(id)
	if err != nil {
		http.Error(w, "Error occurred while deleting the user", http.StatusInternalServerError)
		return
	}

	http.Error(w, "User deleted successfully", http.StatusOK)
}
