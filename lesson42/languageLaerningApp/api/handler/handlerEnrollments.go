package handler

import (
	"fmt"
	"learning_app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateEnrollmentHandler(ctx *gin.Context) {
	var enrollment models.Enrollment
	if err := ctx.ShouldBindJSON(&enrollment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := h.Enrollment.CreateEnrollment(enrollment); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create enrollment",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Enrollment created successfully",
	})
}

func (h *Handler) GetEnrollmentsHandler(ctx *gin.Context) {
	enrollments, err := h.Enrollment.GetEnrollments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get Enrollments",
		})
		return
	}

	ctx.JSON(http.StatusOK, enrollments)
}

func (h *Handler) GetEnrollmentByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	fmt.Println(id)
	enrollment, err := h.Enrollment.GetEnrollmentByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Enrollment not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, enrollment)
}

func (h *Handler) UpdateEnrollmentHandler(ctx *gin.Context) {
	var enrollment models.Enrollment
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&enrollment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	enrollment.ID = id

	if err := h.Enrollment.UpdateEnrollment(enrollment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update enrollment",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Enrollment updated successfully",
	})
}

func (h *Handler) DeleteEnrollmentHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.Enrollment.DeleteEnrollment(id); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Enrollment not found or already deleted",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Enrollment deleted successfully",
	})
}

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

	enollments, err := h.Enrollment.GetAllEnrollments(fEnrollment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get enollments",
		})
		return
	}

	ctx.JSON(http.StatusOK, enollments)
}
