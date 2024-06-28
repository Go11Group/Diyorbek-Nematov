package handler

import (
	w "client/generated/weather"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCurrentWeatherHandler(ctx *gin.Context) {
	location := ctx.Query("location")

	req := w.CurrentWeatherRequest{Location: location}

	resp, err := h.Weather.GetCurrentWeather(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetWeatherForecastHandler(ctx *gin.Context) {
	location := ctx.Query("location")
	daysStr := ctx.Query("days")

	days, err := strconv.Atoi(daysStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	req := w.WeatherForecastRequest{Location: location, Days: int32(days)}

	resp, err := h.Weather.GetWeatherForecast(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ReportWeatherConditionHandler(ctx *gin.Context) {
	location := ctx.Query("location")
	date := ctx.Query("date")

	req := w.ReportWeatherConditionRequest{Location: location, Date: date}

	resp, err := h.Weather.ReportWeatherCondition(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, &resp)
}