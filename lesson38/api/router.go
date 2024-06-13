package api

import (
	"database/sql"
	"learning_app/api/handler"
	"learning_app/storage/postgres"

	"github.com/gin-gonic/gin"
)

func Router(db *sql.DB) *gin.Engine {
	u := postgres.NewUserRepo(db)
	c := postgres.NewCourseRepo(db)
	l := postgres.NewLessonRepo(db)
	e := postgres.NewEnrollmentRepo(db)

	handler := handler.Handler{
		User: u,
		Course: c,
		Lesson: l,
		Enrollment: e,
	}

	router := gin.Default()

	user := router.Group("/users")
	{
		user.POST("/", handler.CreateUserHandler)
		user.GET("/", handler.GetUsersHandler)
		user.GET("/:id", handler.GetUserByIDHandler)
		user.PUT("/:id", handler.UpdateUserHandler)
		user.DELETE("/:id", handler.DeleteUserHandler)
	}

	course := router.Group("/courses")
	{
		course.POST("/", handler.CreateCourseHandler)
		course.GET("/", handler.GetCourseHandler)
		course.GET("/:id", handler.GetCourseByIDHandler)
		course.PUT("/:id", handler.UpdateCourseHandler)
		course.DELETE("/:id", handler.DeleteCourseHandler)
	}

	lesson := router.Group("/lessons")
	{
		lesson.POST("/", handler.CreateLessonHandler)
		lesson.GET("/", handler.GetLessonsHandler)
		lesson.GET("/:id", handler.GetEnrollmentByIDHandler)
		lesson.PUT("/:id", handler.UpdateLessonHandler)
		lesson.DELETE("/:id", handler.DeleteLessonHandler)
	}

	enrollment := router.Group("/enrollments")
	{
		enrollment.POST("/", handler.CreateEnrollmentHandler)
		enrollment.GET("/", handler.GetEnrollmentsHandler)
		enrollment.GET("/:id", handler.GetEnrollmentByIDHandler)
		enrollment.PUT("/:id", handler.UpdateEnrollmentHandler)
		enrollment.DELETE("/:id", handler.DeleteEnrollmentHandler)
	}

	return router
}



