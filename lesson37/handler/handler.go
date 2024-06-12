package handler

import (
	postgres "my_module/gin/postgres/storage"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	User *postgres.UserRepo
}

func NewHandler(handler Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/users", handler.GetUsers)

	return r
}