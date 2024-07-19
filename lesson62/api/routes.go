package api

import (
	"students/api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(handle *handler.Handler) *gin.Engine {
	router := gin.Default()

	student := router.Group("api/student")
	{
		student.POST("/", handle.CreateStudentHandler)
		student.GET("/", handle.GetStudentsHandler)
		student.PUT("/:student-id", handle.UpdateStudentHandler)
		student.DELETE("/:student-id", handle.DeleteStudentHandler)
		student.GET("/subjects", handle.GetSubjectsForStudentHandler)
	}

	subject := router.Group("/api/subject")
	{
		subject.POST("/", handle.CreateSubjectHandler)
		subject.GET("/", handle.GetStudentsHandler)
		subject.PUT("/:subject-id", handle.UpdateSubjectHandler)
		subject.DELETE("/:subject-id", handle.DeleteSubjectHandler)
		subject.GET("/students", handle.GetStudentsForSubjectHandler)
	}

	router.POST("/api/associate", handle.AssociateStudentWithSubjectHandler)

	return router
}
