package service

import (
	"context"
	pb "weather/generated/weather"
	"weather/storage/postgres"
)

type WeatherService struct {
	pb.UnimplementedWeatherServiceServer

	Weather *postgres.WeatherRepo
}

func (w *WeatherService) GetCurrentWeather(ctx context.Context, in *pb.CurrentWeatherRequest) (*pb.CurrentWeatherResponse, error) {
	resp, err := w.Weather.CurrentWeather(in)

	return resp, err
}

func (w *WeatherService) GetWeatherForecast(ctx context.Context, in *pb.WeatherForecastRequest) (*pb.WeatherForecastResponse, error) {
	resp, err := w.Weather.WeatherForecast(in)

	return resp, err
}

func (w *WeatherService) ReportWeatherCondition(ctx context.Context, in *pb.ReportWeatherConditionRequest) (*pb.ReportWeatherConditionResponse, error) {
	resp, err := w.Weather.ReportWeatherCondition(in)

	return resp, err
}
