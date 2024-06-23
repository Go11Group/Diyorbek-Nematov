package handler

import (
	"api-gateway/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTerminal(ctx *gin.Context) {
	url := "http://localhost:8082/terminals"
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

	ctx.JSON(http.StatusCreated, gin.H{"message": "Terminal created successfully"})
}

func (h *Handler) GetTerminals(ctx *gin.Context) {
	url := "http://localhost:8082/terminals/"

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

	var terminal []models.Terminal
	if err := json.NewDecoder(resp.Body).Decode(&terminal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response body"})
		return
	}

	ctx.JSON(http.StatusOK, terminal)
}


func (h *Handler) GetTerminal(ctx *gin.Context) {
	id := ctx.Param("id")
	url := fmt.Sprintf("http://localhost:8082/terminals/%s", id)

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

	var terminal models.Terminal
	if err := json.NewDecoder(resp.Body).Decode(&terminal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response body"})
		return
	}

	ctx.JSON(http.StatusOK, terminal)
}

func (h *Handler) UpdateTerminal(ctx *gin.Context) {
	id := ctx.Param("id")
	url := fmt.Sprintf("http://localhost:8082/terminals/%s", id)

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

	ctx.JSON(http.StatusOK, gin.H{"message": "Terminal updated successfully"})
}

func (h *Handler) DeleteTerminal(ctx *gin.Context) {
	id := ctx.Param("id")
	url := fmt.Sprintf("http://localhost:8082/terminals/%s", id)

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

	ctx.JSON(http.StatusOK, gin.H{"message": "Terminal deleted successfully"})
}
