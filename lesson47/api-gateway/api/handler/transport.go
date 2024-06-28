package handler

import (
	t "client/generated/transport"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBusScheduleHandler(ctx *gin.Context) {
	busNumber := ctx.Param("bus_number")

	req := t.GetBusScheduleRequest{BusNumber: busNumber}

	resp, err := h.Transport.GetBusSchedule(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) TrackBusLocationHandler(ctx *gin.Context) {
	busNumber := ctx.Param("bus_number")

	req := t.TrackBusLocationRequest{BusNumber: busNumber}

	resp, err := h.Transport.TrackBusLocation(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ReportTrafficJamHandler(ctx *gin.Context) {
	busNumber := ctx.Param("bus_number")
	location := ctx.Query("location")

	req := t.ReportTrafficJamRequest{Location: location, BusNumber: busNumber}

	resp, err := h.Transport.ReportTrafficJam(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
