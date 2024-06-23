package handler

import (
	"fmt"
	"net/http"
	"metro-service/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCardHandler(ctx *gin.Context) {
	var card models.Card
	if err := ctx.BindJSON(&card); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := h.Card.CreateCard(card); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create card",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Card created successfully",
	})
}

func (h *Handler) GetCardsHandler(ctx *gin.Context) {
	cards, err := h.Card.GetCards()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get cards",
		})
		return
	}

	ctx.JSON(http.StatusOK, cards)
}

func (h *Handler) GetCardByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	card, err := h.Card.GetCardByID(id)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Card not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, card)
}

func (h *Handler) UpdateCardHandler(ctx *gin.Context) {
	var card models.Card
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&card); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	card.ID = id

	if err := h.Card.UpdateCard(card); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update card",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Card updated successfully",
	})
}

func (h *Handler) DeleteCardHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	if err := h.Card.DeleteCard(id); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Card not found or already deleted",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Card deleted successfully",
	})
}
