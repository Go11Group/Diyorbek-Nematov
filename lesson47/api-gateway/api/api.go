package api

import (
	"client/api/handler"
	"client/service"

	"github.com/gin-gonic/gin"
)

func Router(services *service.ServiceMeneger) *gin.Engine {
	router := gin.Default()

	handler := handler.NewHandler(services.Weather, services.Transport)

	router.GET("/transport/schedule/:bus_number", handler.GetBusScheduleHandler)
	router.GET("/transport/")

	return router
}