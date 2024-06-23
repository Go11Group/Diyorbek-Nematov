package handler

import (
	"api-gateway/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStation(ctx *gin.Context) {
	url := "http://localhost:8082/stations"
	req, err := http.NewRequestWithContext(ctx.Request.Context(), http.MethodPost, url, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		ctx.JSON(resp.StatusCode, gin.H{"error": "Unexpected status code"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Station created successfully"})
}

func (h *Handler) GetStations(ctx *gin.Context) {
	url := "http://localhost:8082/stations/"

	req, err := http.NewRequestWithContext(ctx.Request.Context(), http.MethodGet, url, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ctx.JSON(resp.StatusCode, gin.H{"error": "Unexpected status code"})
		return
	}

	var station []models.Station
	if err := json.NewDecoder(resp.Body).Decode(&station); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response body"})
		return
	}

	ctx.JSON(http.StatusOK, station)
}

func (h *Handler) GetStation(ctx *gin.Context) {
	id := ctx.Param("id")
	url := fmt.Sprintf("http://localhost:8082/stations/%s", id)

	req, err := http.NewRequestWithContext(ctx.Request.Context(), http.MethodGet, url, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ctx.JSON(resp.StatusCode, gin.H{"error": "Unexpected status code"})
		return
	}

	var station models.Station
	if err := json.NewDecoder(resp.Body).Decode(&station); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response body"})
		return
	}

	ctx.JSON(http.StatusOK, station)
}

func (h *Handler) UpdateStation(ctx *gin.Context) {
	id := ctx.Param("id")
	url := fmt.Sprintf("http://localhost:8082/stations/%s", id)

	req, err := http.NewRequestWithContext(ctx.Request.Context(), http.MethodPut, url, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ctx.JSON(resp.StatusCode, gin.H{"error": "Unexpected status code"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Station updated successfully"})
}

func (h *Handler) DeleteStation(ctx *gin.Context) {
	id := ctx.Param("id")
	url := fmt.Sprintf("http://localhost:8082/stations/%s", id)

	req, err := http.NewRequestWithContext(ctx.Request.Context(), http.MethodDelete, url, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ctx.JSON(resp.StatusCode, gin.H{"error": "Unexpected status code"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Station deleted successfully"})
}
