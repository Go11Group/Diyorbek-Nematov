package handler

import (
	"fmt"
	"net/http"
	"user-service/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUserHandler(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := h.User.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func (h *Handler) GetUsersHandler(ctx *gin.Context) {
	users, err := h.User.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get users",
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (h *Handler) GetUserByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := h.User.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUserHandler(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	user.ID = id

	if err := h.User.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

func (h *Handler) DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
	if err := h.User.DeleteUser(id); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found or already deleted",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
