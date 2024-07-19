package handler

import (
	"net/http"
	"students/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateSubjectHandler(ctx *gin.Context) {
	var subject models.Subject

	if err := ctx.ShouldBindJSON(&subject); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Error: "Error bind json",
		})
		return
	}

	resp, err := h.SubjectRepo.CreateSubject(&subject)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "Error in subject created",
		})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetSubjectsHandler(ctx *gin.Context) {
	subjects, err := h.SubjectRepo.GetSubjects()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "Error in get all subjects",
		})
		return
	}

	ctx.JSON(http.StatusOK, subjects)
}

func (h *Handler) GetSubjectHandler(ctx *gin.Context) {
	id := ctx.Param("subject-id")

	subject, err := h.SubjectRepo.GetSubject(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "Error in get subject by id",
		})
		return
	}

	ctx.JSON(http.StatusOK, subject)
}

func (h *Handler) UpdateSubjectHandler(ctx *gin.Context) {
	id := ctx.Param("subject-id")
	var subject models.Subject

	if err := ctx.ShouldBindJSON(&subject); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Error: "Error bind json",
		})
		return
	}
	subject.ID = id

	resp, err := h.SubjectRepo.UpdateSubject(subject)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "Error in updated subject",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteSubjectHandler(ctx *gin.Context) {
	id := ctx.Param("subject-id")

	resp, err := h.StudentRepo.DeleteStudent(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "error in deleted subject",
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) AssociateStudentWithSubjectHandler(ctx *gin.Context) {
	var associate models.Associate
	if err := ctx.ShouldBindJSON(&associate); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Errors{
			Error: "Error bind json",
		})
		return
	}

	resp, err := h.SubjectRepo.AssociateStudentWithSubject(associate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "Error in associate student with subject",
		})
		return
	}

	ctx.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetStudentsForSubjectHandler(ctx *gin.Context) {
	id := ctx.Param("subject-id")

	resp, err := h.SubjectRepo.GetStudentsForSubject(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Errors{
			Error: "Error in get students for subject",
		})
		return 
	}

	ctx.JSON(http.StatusOK, resp)
}