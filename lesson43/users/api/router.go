package api

import (
	"user-service/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(handle *handler.Handler) *gin.Engine {
	r := gin.Default()

	user := r.Group("/users")
	{
		user.GET("/", handle.GetUsersHandler)
		user.POST("/", handle.CreateUserHandler)
		user.GET("/:id", handle.GetUserByIDHandler)
		user.PUT("/:id", handle.UpdateUserHandler)
		user.DELETE("/:id", handle.DeleteUserHandler)
	}

	return r
}