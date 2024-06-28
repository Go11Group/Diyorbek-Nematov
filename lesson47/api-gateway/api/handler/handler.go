package handler

import (
	t "client/generated/transport"
	w "client/generated/weather"
)

type Handler struct {
	Weather  w.WeatherServiceClient
	Transport t.TransportServiceClient
}

func NewHandler(weather w.WeatherServiceClient, transport t.TransportServiceClient) *Handler {
	return &Handler{Weather: weather, Transport: transport}
}