package api

import (
	"students/api/handler"
	"students/api/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(handle *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("api/register", handle.Register)
	router.POST("api/login", handle.Login)
	r := router.Group("/")
	r.Use(middleware.CasbinMiddleware(handle.Enforcer))

	student := r.Group("api/student")
	{
		student.POST("/", handle.CreateStudentHandler)
		student.GET("/", handle.GetStudentsHandler)
		student.PUT("/:student-id", handle.UpdateStudentHandler)
		student.DELETE("/:student-id", handle.DeleteStudentHandler)
		student.GET("/subjects", handle.GetSubjectsForStudentHandler)
	}

	subject := r.Group("/api/subject")
	{
		subject.POST("/", handle.CreateSubjectHandler)
		subject.GET("/", handle.GetStudentsHandler)
		subject.PUT("/:subject-id", handle.UpdateSubjectHandler)
		subject.DELETE("/:subject-id", handle.DeleteSubjectHandler)
		subject.GET("/students", handle.GetStudentsForSubjectHandler)
	}

	r.POST("/api/associate", handle.AssociateStudentWithSubjectHandler)

	return router
}
