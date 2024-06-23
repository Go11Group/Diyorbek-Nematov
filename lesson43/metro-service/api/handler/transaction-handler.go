package handler

import (
	"fmt"
	"metro-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTransactionHandler(ctx *gin.Context) {
	var transaction models.Transaction
	if err := ctx.BindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := h.Transaction.CreateTransaction(transaction); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create transaction",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Transaction created successfully",
	})
}

func (h *Handler) GetTransactionsHandler(ctx *gin.Context) {
	transactions, err := h.Transaction.GetTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get transactions",
		})
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

func (h *Handler) GetTransactionByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	transaction, err := h.Transaction.GetTransactionByID(id)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Transaction not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

func (h *Handler) UpdateTransactionHandler(ctx *gin.Context) {
	var transaction models.Transaction
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	transaction.ID = id

	if err := h.Transaction.UpdateTransaction(transaction); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update transaction",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Transaction updated successfully",
	})
}

func (h *Handler) DeleteTransactionHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.Transaction.DeleteTransaction(id); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Transaction not found or already deleted",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Transaction deleted successfully",
	})
}
