package api

import (
	"client/api/handler"
	"client/service"

	"github.com/gin-gonic/gin"
)

func Router(services *service.ServiceManager) *gin.Engine {
	router := gin.Default()

	handler := handler.NewHandler(services.Weather, services.Transport)

	router.GET("/transport/schedule/:bus_number", handler.GetBusScheduleHandler)
	router.GET("/transport/report_traffic_jam/:bus_number", handler.ReportTrafficJamHandler)
	router.GET("/transport/track_bus_location/:bus_number", handler.TrackBusLocationHandler)

	router.GET("/weather/current_weather", handler.GetCurrentWeatherHandler)
	router.GET("/weather/weather_forecast", handler.GetWeatherForecastHandler)
	router.GET("/weather/report_weather_condition", handler.ReportWeatherConditionHandler)


	return router
}