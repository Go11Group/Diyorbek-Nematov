package handler

import (
	"fmt"
	"learning_app/models"
	"net/http"
	"strconv"

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
	fmt.Println(err)
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

// GET All Users
func (h *Handler) GetAllUsers(ctx *gin.Context) {
	var fUser models.FilterUser

	if err := ctx.ShouldBindQuery(&fUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameters",
		})
		return
	}

	defaultLimit := 10
	defaultOffset := 0

	if limitStr := ctx.Query("Limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid limit parameter",
			})
			return
		}
		fUser.Limit = limit
	} else {
		fUser.Limit = defaultLimit
	}
	if offsetStr := ctx.Query("Offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid offset parameter",
			})
			return
		}
		fUser.Offset = offset
	} else {
		fUser.Offset = defaultOffset
	}

	users, err := h.User.GetAllUsers(fUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get users",
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
