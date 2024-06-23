package handler

import (
	"fmt"
	"net/http"
	"metro-service/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTerminalHandler(ctx *gin.Context) {
	var terminal models.Terminal
	if err := ctx.BindJSON(&terminal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := h.Terminal.CreateTerminal(terminal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create terminal",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Terminal created successfully",
	})
}

func (h *Handler) GetTerminalsHandler(ctx *gin.Context) {
	terminals, err := h.Terminal.GetTerminals()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get terminals",
		})
		return
	}

	ctx.JSON(http.StatusOK, terminals)
}

func (h *Handler) GetTerminalByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	terminal, err := h.Terminal.GetTerminalByID(id)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Terminal not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, terminal)
}

func (h *Handler) UpdateTerminalHandler(ctx *gin.Context) {
	var terminal models.Terminal
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&terminal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	terminal.ID = id

	if err := h.Terminal.UpdateTerminal(terminal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update terminal",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Terminal updated successfully",
	})
}

func (h *Handler) DeleteTerminalHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.Terminal.DeleteTerminal(id); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Terminal not found or already deleted",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Terminal deleted successfully",
	})
}
