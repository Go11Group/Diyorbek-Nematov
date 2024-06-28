package postgres

import (
	"database/sql"
	pb "weather/generated/weather"
)

type WeatherRepo struct {
	DB *sql.DB
}

func NewWeatherRepo(db *sql.DB) *WeatherRepo {
	return &WeatherRepo{DB: db}
}

func (w *WeatherRepo) CurrentWeather(currWeather *pb.CurrentWeatherRequest) (*pb.CurrentWeatherResponse, error) {
	resp := pb.CurrentWeatherResponse{}

	err := w.DB.QueryRow(`
		SELECT
			temperature,
			humidity,
			wind_speed
		FROM
			weather_info
		WHERE 
			location = $1;
	`, currWeather.Location).Scan(&resp.Temperature, &resp.Humidity, &resp.WindSpeed)

	return &resp, err
}

func (w *WeatherRepo) WeatherForecast(wForecast *pb.WeatherForecastRequest) (*pb.WeatherForecastResponse, error) {
	weatherForcasts := []*pb.WeatherForecast{}

	rows, err := w.DB.Query(`
		SELECT 
			day, 
			temperature,
			humidity,
			wind_speed
		FROM
			weather_info
		WHERE 
			location = $1
		LIMIT $2
	`, wForecast.Location, wForecast.Days)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		weatherForcast := &pb.WeatherForecast{}
		
		if err = rows.Scan(&weatherForcast.Date, &weatherForcast.Temperature, &weatherForcast.Humidity, &weatherForcast.WindSpeed); err != nil {
			return nil, err
		}

		weatherForcasts = append(weatherForcasts, weatherForcast)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.WeatherForecastResponse{Forecasts: weatherForcasts}, nil
}

func (w *WeatherRepo) ReportWeatherCondition(condition *pb.ReportWeatherConditionRequest) (*pb.ReportWeatherConditionResponse, error) {
	resp := pb.ReportWeatherConditionResponse{}

	err := w.DB.QueryRow(`
		SELECT
			temperature,
			humidity,
			wind_speed,
			condition
		FROM
			weather_info
		WHERE
			location = $1 and day = $2
	`, condition.Location, condition.Date).
		Scan(&resp.Temperature, &resp.Humidity, &resp.WindSpeed, &resp.Condition)

	return &resp, err
}