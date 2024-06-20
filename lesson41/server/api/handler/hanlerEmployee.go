package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"server/models"
	"strconv"
	"strings"
)

func (h *Handler) CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Malumotni JSONga decode qilishda xatolik yuz berdi"))
		return
	}

	err = h.Emp.CreateEmployee(emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Malumotni bazaga qo'shishda xatolik yuz berdi"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Xodim muvaffaqiyatli yaratildi"))
}

func (h *Handler) GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employees/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID ni olishda xatolik yuz berdi"))
		return
	}

	emp, err := h.Emp.GetEmployee(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Xodim topilmadi"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Ma'lumotni bazadan o'qishda xatolik yuz berdi"))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ma'lumotni JSONga encode qilishda xatolik yuz berdi"))
		return
	}
}

func (h *Handler) UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employees/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID ni olishda xatolik yuz berdi"))
		return
	}

	var emp models.Employee
	err = json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Malumotni JSONga decode qilishda xatolik yuz berdi"))
		return
	}
	emp.ID = id

	err = h.Emp.UpdateEmployee(emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ma'lumotni bazaga yangilashda xatolik yuz berdi"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Xodim muvaffaqiyatli yangilandi"))
}

func (h *Handler) DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employees/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID ni olishda xatolik yuz berdi"))
		return
	}

	err = h.Emp.DeleteEmploee(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Xodimni bazadan o'chirishda xatolik yuz berdi"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Xodim muvaffaqiyatli o'chirildi"))
}
