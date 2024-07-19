package handler

import (
	"net/http"
	"students/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStudentHandler(ctx *gin.Context) {
	var student models.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Error: "Error bind json",
		})
		return
	}

	resp, err := h.StudentRepo.CreateStudent(student)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "Error in user created",
		})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetStudentsHandler(ctx *gin.Context) {
	students, err := h.StudentRepo.GetStudents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "Error in get all students",
		})
		return
	}

	ctx.JSON(http.StatusOK, students)
}

func (h *Handler) GetStudentHandler(ctx *gin.Context) {
	id := ctx.Param("student-id")

	student, err := h.StudentRepo.GetStudent(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "Error in get student by id",
		})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (h *Handler) UpdateStudentHandler(ctx *gin.Context) {
	id := ctx.Param("student-id")
	var student models.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Error: "Error bind json",
		})
		return
	}
	student.ID = id

	resp, err := h.StudentRepo.UpdateStudent(student)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "Error in updated student",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteStudentHandler(ctx *gin.Context) {
	id := ctx.Param("student-id")

	resp, err := h.StudentRepo.DeleteStudent(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "error in deleted student",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetSubjectsForStudentHandler(ctx *gin.Context) {
	id := ctx.Param("student-id")

	resp, err := h.StudentRepo.GetSubjectsForStudent(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "Error in get subjets fo student",
		})
		return 
	}

	ctx.JSON(http.StatusOK, resp)
}