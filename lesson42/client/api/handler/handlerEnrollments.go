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

// Create Enrollment Handler
func (h *Handler) CreateEnrollmentHandler(ctx *gin.Context) {
	var enrollment models.Enrollment
	if err := ctx.BindJSON(&enrollment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	byteEnroolment, err := json.Marshal(enrollment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while marshaling JSON",
		})
		return
	}

	url := "http://localhost:8081/enrollments/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteEnroolment))
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
			"message": "Failed to create enrollment",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Enrollment created successfully",
	})
}

// Get Enrollment By ID Handler
func (h *Handler) GetEnrollmentByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	url := fmt.Sprintf("http://localhost:8081/enrollments/%s", id)

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
			"message": "Failed to get enrollment",
		})
		return
	}

	var enrollment models.Enrollment
	if err := json.NewDecoder(resp.Body).Decode(&enrollment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding response",
		})
		return
	}

	ctx.JSON(http.StatusOK, enrollment)
}

// Update Enrollment Handler
func (h *Handler) UpdateEnrollmentHandler(ctx *gin.Context) {
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

	url := fmt.Sprintf("http://localhost:8081/enrollments/%s", id)

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
			"message": "Failed to update enrollment",
			"details": string(body),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Enrollment updated successfully",
	})
}

// Delete Enrollment Handler
func (h *Handler) DeleteEnrollmentHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)

	url := fmt.Sprintf("http://localhost:8081/enrollments/%s", id)

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
			"message": "Failed to delete enrollment",
			"details": string(body),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Enrollment deleted successfully",
	})
}

// Get All Enrollments
func (h *Handler) GetAllEnrollments(ctx *gin.Context) {
	var fEnrollment models.FilterEnrollment

	if err := ctx.ShouldBindQuery(&fEnrollment); err != nil {
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
		fEnrollment.Limit = limit
	} else {
		fEnrollment.Limit = defaultLimit
	}

	if offsetStr := ctx.Query("Offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid offset parameter",
			})
			return
		}
		fEnrollment.Offset = offset
	} else {
		fEnrollment.Offset = defaultOffset
	}

	url := fmt.Sprintf("http://localhost:8081/enrollments?UserID=%s&CourseID=%s&EnrollmentDate=%s&Limit=%d&Offset=%d",
		fEnrollment.UserID, fEnrollment.CourseID, url.QueryEscape(fEnrollment.EnrollmentDate), fEnrollment.Limit, fEnrollment.Offset)
	
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
		body, _ := io.ReadAll(resp.Body)
		ctx.JSON(resp.StatusCode, gin.H{
			"message": "Failed to get enrollments",
			"details": string(body),
		})
		return
	}

	var enrollments []models.Enrollment
	if err := json.NewDecoder(resp.Body).Decode(&enrollments); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding response",
		})
		return
	}

	ctx.JSON(http.StatusOK, enrollments)
}
