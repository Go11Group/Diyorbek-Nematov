package handler

import (
	"fmt"
	"net/http"
	"metro-service/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStationHandler(ctx *gin.Context) {
	var stattion models.Station
	if err := ctx.ShouldBindJSON(&stattion); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if err := h.Station.CreateStation(stattion); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create station",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Station created successfully",
	})
}

func (h *Handler) GetStationsHandler(ctx *gin.Context) {
	stations, err := h.Station.GetStations()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get stations",
		})
		return
	}

	ctx.JSON(http.StatusOK, stations)
}

func (h *Handler) GetStationByIDHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	station, err := h.Station.GetStationByID(id)
	fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Station not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, station)
}

func (h *Handler) UpdateStationHandler(ctx *gin.Context) {
	var station models.Station
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(&station); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	station.ID = id

	if err := h.Station.UpdateStation(station); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update station",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Station updated successfully",
	})
}

func (h *Handler) DeleteStationHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.Station.DeleteStation(id); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Station not found or already deleted",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Station deleted successfully",
	})
}
