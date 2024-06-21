package api

import (
	"learning_app/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(handler handler.Handler) *gin.Engine {

	router := gin.Default()

	user := router.Group("/users")
	{
		user.POST("/", handler.CreateUserHandler)
		user.GET("/", handler.GetAllUsers)
		user.GET("/:id", handler.GetUserByIDHandler)
		user.PUT("/:id", handler.UpdateUserHandler)
		user.DELETE("/:id", handler.DeleteUserHandler)
	}

	course := router.Group("/courses")
	{
		course.POST("/", handler.CreateCourseHandler)
		course.GET("/", handler.GetAllCourses)
		course.GET("/:id", handler.GetCourseByIDHandler)
		course.PUT("/:id", handler.UpdateCourseHandler)
		course.DELETE("/:id", handler.DeleteCourseHandler)
	}

	lesson := router.Group("/lessons")
	{
		lesson.POST("/", handler.CreateLessonHandler)
		lesson.GET("/", handler.GetAllLessons)
		lesson.GET("/:id", handler.GetLessonByIDHandler)
		lesson.PUT("/:id", handler.UpdateLessonHandler)
		lesson.DELETE("/:id", handler.DeleteLessonHandler)
	}

	enrollment := router.Group("/enrollments")
	{
		enrollment.POST("/", handler.CreateEnrollmentHandler)
		enrollment.GET("/", handler.GetAllEnrollments)
		enrollment.GET("/:id", handler.GetEnrollmentByIDHandler)
		enrollment.PUT("/:id", handler.UpdateEnrollmentHandler)
		enrollment.DELETE("/:id", handler.DeleteEnrollmentHandler)
	}
	
	additional := router.Group("/additional") 
	{
		additional.GET("/users/:user_id/courses", handler.GetCoursesbyUser)
		additional.GET("/courses/:course_id/lessons", handler.GetLessonsByCourse)
		additional.GET("/courses/:course_id/enrollments", handler.GetEnrolledUsersbyCourse)
		additional.GET("/users/search", handler.SearchUsers)
		additional.GET("/courses/popular", handler.GetMostPopularCourses)
	}
	return router
}
