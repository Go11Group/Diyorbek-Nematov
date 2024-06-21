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

// CREATE Lesson Handler
func (h *Handler) CreateLessonHandler(ctx *gin.Context) {
	var lesson models.Lesson
	if err := ctx.BindJSON(&lesson); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	byteLesson, err := json.Marshal(lesson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while marshaling JSON",
		})
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8081/lessons/", bytes.NewBuffer(byteLesson))
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
			"message": "Failed to create lesson",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Lesson created successfully",
	})
}

// GetLesson By ID Handler
func (h *Handler) GetLessonByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	url := fmt.Sprintf("http://localhost:8081/lessons/%s", id)

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
			"message": "Failed to get lesson",
		})
		return
	}

	var lesson models.Lesson
	if err := json.NewDecoder(resp.Body).Decode(&lesson); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding response",
		})
		return
	}

	ctx.JSON(http.StatusOK, lesson)
}

// Update Lesson Handler
func (h *Handler) UpdateLessonHandler(ctx *gin.Context) {
	var lesson models.Lesson
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&lesson); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	byteLesson, err := json.Marshal(lesson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while marshaling JSON",
		})
		return
	}

	url := fmt.Sprintf("http://localhost:8081/lessons/%s", id)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(byteLesson))
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
			"message": "Failed to update lesson",
			"details": string(body),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Lesson updated successfully",
	})
}

// delete Lesson Handler
func (h *Handler) DeleteLessonHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	url := fmt.Sprintf("http://localhost:8081/lessons/%s", id)

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
			"message": "Failed to delete lesson",
			"details": string(body),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

// Get All Lessons
func (h *Handler) GetAllLessons(ctx *gin.Context) {
	var fLesson models.FilterLesson

	if err := ctx.ShouldBindQuery(&fLesson); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameters",
		})
		return
	}

	defaultLimit := 10
	defaultOffset := 0

	if limitStr := ctx.Query("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid limit parameter",
			})
			return
		}
		fLesson.Limit = limit
	} else {
		fLesson.Limit = defaultLimit
	}

	if offsetStr := ctx.Query("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid offset parameter",
			})
			return
		}
		fLesson.Offset = offset
	} else {
		fLesson.Offset = defaultOffset
	}

	url := fmt.Sprintf("http://localhost:8081/lessons/?CourseID=%s&Title=%s&Content=%s&limit=%d&offset=%d",
		fLesson.CourseID, url.QueryEscape(fLesson.Title), url.QueryEscape(fLesson.Content), fLesson.Limit, fLesson.Offset)

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
			"message": "Failed to get lessons",
			"details": string(body),
		})
		return
	}

	var lessons []models.Lesson
	if err := json.NewDecoder(resp.Body).Decode(&lessons); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding response",
		})
		return
	}

	ctx.JSON(http.StatusOK, lessons)
}
