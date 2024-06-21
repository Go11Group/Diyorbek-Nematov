package handler

import (
	"client/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// --------------- CREATE COURSE BY USER ---------------------
func (h *Handler) GetCoursesbyUser(ctx *gin.Context) {
	id := ctx.Param("user_id")

	url := fmt.Sprintf("http://localhost:8081/additional/users/%s/courses", id)

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
			"message": "Failed to get courses for user",
			"details": string(body),
		})
		return
	}

	var courseByUser []models.CourseByUser
	if err := json.NewDecoder(resp.Body).Decode(&courseByUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding response",
		})
		return
	}

	ctx.JSON(http.StatusOK, courseByUser)
}

// --------------- Lessons By Course -----------------
func (h *Handler) GetLessonsByCourse(ctx *gin.Context) {
	id := ctx.Param("course_id")
	url := fmt.Sprintf("http://localhost:8081/additional/courses/%s/lessons", id)

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
			"message": "Failed to get lessons for course",
			"details": string(body),
		})
		return
	}

	var lessonsByCourse []models.LessonsByCourse
	if err := json.NewDecoder(resp.Body).Decode(&lessonsByCourse); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding response",
		})
		return
	}

	ctx.JSON(http.StatusOK, lessonsByCourse)
}

// ---------------- Enrolled Users By Course -------------------------
func (h *Handler) GetEnrolledUsersbyCourse(ctx *gin.Context) {
	id := ctx.Param("course_id")
	
	url := fmt.Sprintf("http://localhost:8081/additional/courses/%s/enrollments", id)

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
			"message": "Failed to get enrolled users for course",
			"details": string(body),
		})
		return
	}

	var enrolledUsersbyCourse []models.EnrolledUsersByCourse
	if err := json.NewDecoder(resp.Body).Decode(&enrolledUsersbyCourse); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding response",
		})
		return
	}

	ctx.JSON(http.StatusOK, enrolledUsersbyCourse)
}

// ------------- SEARCH USER ----------------------------
func (h *Handler) SearchUsers(ctx *gin.Context) {
	var searchUser models.SearchUser

	if err := ctx.ShouldBindQuery(&searchUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameters",
		})
		return
	}
	
	var url = fmt.Sprintf("http://localhost:8081/additional/users/search?Name=%s&Email=%s&AgeFrom=%d&AgeTo=%d",
		url.QueryEscape(searchUser.Name), url.QueryEscape(searchUser.Email), searchUser.AgeFrom, searchUser.AgeTo)
	
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
			"message": "Failed to get users",
			"details": string(body),
		})
		return
	}

	var results models.Result
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding response",
		})
		return
	}

	ctx.JSON(http.StatusOK, results)
}


// ---------------------- GET MOST POPULAR -----------------------
func (h *Handler) GetMostPopularCourses(ctx *gin.Context) {
	startDate := ctx.Query("StartDate")
	endDate := ctx.Query("EndDate")
	// Most popular courses olish uchun logika
	
	url := fmt.Sprintf("http://localhost:8081/additional/courses/popular?StartDate=%sEndDate=%s",
		url.QueryEscape(startDate), url.QueryEscape(endDate))

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
			"message": "Failed to get popular course",
			"details": string(body),
		})
		return
	}

	var popularCourses models.PopularCourse
	if err := json.NewDecoder(resp.Body).Decode(&popularCourses); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error decoding response",
		})
		return
	}

	// Natijani qaytarish
	ctx.JSON(http.StatusOK, popularCourses)
}
