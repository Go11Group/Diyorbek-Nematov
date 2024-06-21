package handler

import (
	"fmt"
	"learning_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCoursesbyUser(ctx *gin.Context) {
	id := ctx.Param("user_id")
	fmt.Println(id)
	getCourseByUser, err := h.Additional.GetCoursesbyUser(id)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, getCourseByUser)
}

func (h *Handler) GetLessonsByCourse(ctx *gin.Context) {
	id := ctx.Param("course_id")

	lessonsByCourse, err := h.Additional.GetLessonsByCourse(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Course not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, lessonsByCourse)
}

func (h *Handler) GetEnrolledUsersbyCourse(ctx *gin.Context) {
	id := ctx.Param("course_id")
	enrolledUsersbyCourse, err := h.Additional.GetEnrolledUsersbyCourse(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Course not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, enrolledUsersbyCourse)
}

func (h *Handler) SearchUsers(ctx *gin.Context) {
	var searchUser models.SearchUser

	if err := ctx.ShouldBindQuery(&searchUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameters",
		})
		return	
	}
	fmt.Println(searchUser)
	results, err := h.Additional.SearchUsers(searchUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get users",
		})
		return
	}

	ctx.JSON(http.StatusOK, results)
}

func (h *Handler) GetMostPopularCourses(ctx *gin.Context) {
	startDate := ctx.Query("StartDate")
	endDate := ctx.Query("EndDate")
	// Most popular courses olish uchun logika
	popularCourses, err := h.Additional.GetMostPopularCourses(startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Natijani qaytarish
	ctx.JSON(http.StatusOK, popularCourses)
}
