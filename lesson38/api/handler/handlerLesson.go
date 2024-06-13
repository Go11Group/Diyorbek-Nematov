package handler

import (
	"learning_app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateLessonHandler(ctx *gin.Context) {
	var lesson models.Lesson
	if err := ctx.ShouldBindJSON(&lesson); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := h.Lesson.CreateLesson(lesson); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create lesson",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Lesson created successfully",
	})
}

func (h *Handler) GetLessonsHandler(ctx *gin.Context) {
	lessons, err := h.Lesson.GetLessons()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get lessons",
		})
		return
	}

	ctx.JSON(http.StatusOK, lessons)
}

func (h *Handler) GetLessonByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.Lesson.GetLessonByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Lesson not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateLessonHandler(ctx *gin.Context) {
	var lesson models.Lesson
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&lesson); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	lesson.ID = id

	if err := h.Lesson.UpdateLesson(lesson); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update lesson",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Lesson updated successfully",
	})
}

func (h *Handler) DeleteLessonHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.Lesson.DeleteLesson(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Lesson not found or already deleted",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Lesson deleted successfully",
	})
}

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

	lessons, err := h.Lesson.GetAllLessons(fLesson)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get lessons",
		})
		return
	}

	ctx.JSON(http.StatusOK, lessons)
}