package handler

import (
	"bytes"
	"client/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler
func (h *Handler) CreateUserHandler(ctx *gin.Context) {
	var user models.User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	byteUser, err := json.Marshal(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while marshaling JSON",
		})
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8081/users/", bytes.NewBuffer(byteUser))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating request",
		})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error sending request",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		ctx.JSON(resp.StatusCode, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

// GetUserByIDHandler
func (h *Handler) GetUserByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	url := fmt.Sprintf("http://localhost:8081/users/%s", id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating request",
		})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error sending request",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ctx.JSON(resp.StatusCode, gin.H{
			"message": "Failed to get user",
		})
		return
	}

	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding response",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// UpdateUserHandler
func (h *Handler) UpdateUserHandler(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	byteUser, err := json.Marshal(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while marshaling JSON",
		})
		return
	}

	url := fmt.Sprintf("http://localhost:8081/users/%s", id)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(byteUser))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating request",
		})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error sending request",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		ctx.JSON(resp.StatusCode, gin.H{
			"message": "Failed to update user",
			"details": string(body),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

// deleteUserHandler
func (h *Handler) DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)

	url := fmt.Sprintf("http://localhost:8081/users/%s", id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating request",
		})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error sending request",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		ctx.JSON(resp.StatusCode, gin.H{
			"message": "Failed to delete user",
			"details": string(body),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

// GET ALL Users FILTER
func (h *Handler) GetAllUsers(ctx *gin.Context) {
	var fUser models.FilterUser

	// Bind query parameters to fUser
	if err := ctx.ShouldBindQuery(&fUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameters",
		})
		return
	}

	// Set default limit and offset
	defaultLimit := 10
	defaultOffset := 0

	// Parse and validate limit parameter
	if limitStr := ctx.Query("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid limit parameter",
			})
			return
		}
		fUser.Limit = limit
	} else {
		fUser.Limit = defaultLimit
	}

	// Parse and validate offset parameter
	if offsetStr := ctx.Query("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid offset parameter",
			})
			return
		}
		fUser.Offset = offset
	} else {
		fUser.Offset = defaultOffset
	}

	// Create URL for request
	url := fmt.Sprintf("http://localhost:8081/users/?Name=%s&Email=%s&Birthday=%s&Limit=%d&Offset=%d",
		url.QueryEscape(fUser.Name), url.QueryEscape(fUser.Email), url.QueryEscape(fUser.Birthday), fUser.Limit, fUser.Offset)
	fmt.Println(fUser, url)

	// Create new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating request",
		})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Send HTTP request
	resp, err := h.Client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error sending request",
		})
		return
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		ctx.JSON(resp.StatusCode, gin.H{
			"message": "Failed to get users",
			"details": string(body),
		})
		return
	}

	// Decode response body
	var users []models.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding response",
		})
		return
	}

	// Return users in response
	ctx.JSON(http.StatusOK, users)
}
