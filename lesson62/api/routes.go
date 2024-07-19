package api

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()

	subject := router.Group("/api/subject")
	{
		subject.POST()
	}

	return router
}
