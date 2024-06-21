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

// CREATE COURSE Handler
func (h *Handler) CreateCourseHandler(ctx *gin.Context) {
	var course models.Course
	if err := ctx.BindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	byteCourse, err := json.Marshal(course)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while marshaling JSON",
		})
		return
	}

	url := "http://localhost:8081/courses/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byteCourse))
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
			"message": "Failed to create course",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

// GET COURSE BY ID HANDLER
func (h *Handler) GetCourseByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	url := fmt.Sprintf("http://localhost:8081/courses/%s", id)

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
			"message": "Failed to get course",
		})
		return
	}

	var course models.Course
	if err := json.NewDecoder(resp.Body).Decode(&course); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding response",
		})
		return
	}

	ctx.JSON(http.StatusOK, course)
}

// Update Course Handler
func (h *Handler) UpdateCourseHandler(ctx *gin.Context) {
	var course models.Course
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	byteCourse, err := json.Marshal(course)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while marshaling JSON",
		})
		return
	}

	url := fmt.Sprintf("http://localhost:8081/courses/%s", id)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(byteCourse))
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
			"message": "Failed to update course",
			"details": string(body),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Course updated successfully",
	})
}

// delete Course Handler
func (h *Handler) DeleteCourseHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)

	url := fmt.Sprintf("http://localhost:8081/courses/%s", id)

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
			"message": "Failed to delete course",
			"details": string(body),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Course deleted successfully",
	})
}

// GET All Courses
func (h *Handler) GetAllCourses(ctx *gin.Context) {
	var fCourse models.FilterCourse

	if err := ctx.ShouldBindQuery(&fCourse); err != nil {
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
		fCourse.Limit = limit
	} else {
		fCourse.Limit = defaultLimit
	}

	if offsetStr := ctx.Query("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid offset parameter",
			})
			return
		}
		fCourse.Offset = offset
	} else {
		fCourse.Offset = defaultOffset
	} 

	url := fmt.Sprintf("http://localhost:8081/courses/?Title=%s&Description=%s&limit=%d&offset=%d",
		url.QueryEscape(fCourse.Title), url.QueryEscape(fCourse.Description),fCourse.Limit, fCourse.Offset)
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
			"message": "Failed to get courses",
			"details": string(body),
		})
		return
	}

	var courses []models.Course
	if err := json.NewDecoder(resp.Body).Decode(&courses); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding response",
		})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}
