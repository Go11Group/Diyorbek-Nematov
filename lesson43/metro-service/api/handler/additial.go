package handler

import (
	"database/sql"
	"metro-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBalance(ctx *gin.Context) {
	id := ctx.Query("user_id")
	balance, err := h.Additial.GetBalance(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create card",
		})
		return
	}

	ctx.JSON(http.StatusOK, balance)
}

func (h *Handler) CheckBalance(ctx *gin.Context) {
	userId := ctx.Query("user_id")

	result, err := h.Additial.CheckBalance(userId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusOK, models.BalanceResponse{Balance: 0, BalanceStatus: "Insufficient"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, result)
}
