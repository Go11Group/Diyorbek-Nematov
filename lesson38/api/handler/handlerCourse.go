package handler

import (
	"learning_app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCourseHandler(ctx *gin.Context) {
	var course models.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := h.Course.CreateCourse(course); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create course",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Course created successfully",
	})
}

func (h *Handler) GetCourseHandler(ctx *gin.Context) {
	courses, err := h.Course.GetCourses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get courses",
		})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

func (h *Handler) GetCourseByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	course, err := h.Course.GetCourseByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Course not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, course)
}

func (h *Handler) UpdateCourseHandler(ctx *gin.Context) {
	var course models.Course
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	course.ID = id

	if err := h.Course.UpdateCourse(course); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update course",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Course updated successfully",
	})
}

func (h *Handler) DeleteCourseHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.Course.DeleteCourse(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Course not found or already deleted",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Course deleted successfully",
	})
}

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

	courses, err := h.Course.GetAllCourses(fCourse)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get courses",
		})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}
