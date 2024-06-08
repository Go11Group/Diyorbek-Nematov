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

func (h *Handler) createProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while decoding JSON, err: %s", err.Error())))
		return
	}

	err = h.Product.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while creating product, err: %s", err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("product created successfully"))
}

func (h *Handler) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Product.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while Decode, err: %s", err.Error())))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *Handler) getProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	fmt.Println(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	product, err := h.Product.GetProductById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *Handler) updateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while decoding JSON, err: %s", err.Error())))
		return
	}

	err = h.Product.UpdateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error while updating product, err: %s", err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("product updated successfully"))
}

func (h *Handler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.User.DeleteUser(id)
	if err != nil {
		http.Error(w, "Error occurred while deleting the product", http.StatusInternalServerError)
		return
	}

	http.Error(w, "Pruduct deleted successfully", http.StatusOK)
}
