package main

import (
	"client/generated/transport"
	"client/generated/weather"
	"client/service"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	services := service.New()

	fmt.Println("Choose Service:")
	fmt.Print("1. Transport Service\n2. Weather Service\n")

	var choice int
	fmt.Print("Choice... > ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		TransportInterface(ctx, services)
	case 2:
		WeatherInterface(ctx, services)
	default:
		fmt.Println("Invalid choice")
	}
}

func TransportInterface(ctx context.Context, serv *service.ServiceMeneger) {
	fmt.Println("TransportService ...")
	fmt.Print("1. GetBusSchedule\n2. TrackBusLocation\n3. ReportTrafficJam\n")

	var choice int
	fmt.Print("Choice... > ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var busNumber string
		fmt.Print("Avtobus raqamini kiriting: ")
		fmt.Scan(&busNumber)
		resp, err := serv.Transport.GetBusSchedule(ctx, &transport.GetBusScheduleRequest{BusNumber: busNumber})

		if err != nil {
			log.Fatal(err)
		}

		for _, v := range resp.Schedules {
			fmt.Println(v)
		}

	case 2:
		var busNumber string
		fmt.Print("Avtobus raqamini kiriting: ")
		fmt.Scan(&busNumber)
		resp, err := serv.Transport.TrackBusLocation(ctx, &transport.TrackBusLocationRequest{BusNumber: busNumber})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp)

	case 3:
		var busNumber string
		var location string
		fmt.Print("Avtobus raqamini kiriting: ")
		fmt.Scan(&busNumber)
		fmt.Print("Avtobus lokatsiyasini kiriting: ")
		fmt.Scan(&location)

		resp, err := serv.Transport.ReportTrafficJam(ctx, &transport.ReportTrafficJamRequest{BusNumber: busNumber, Location: location})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp)
	}
}

func WeatherInterface(ctx context.Context, serv *service.ServiceMeneger) {
	fmt.Println("WeatherService ...")
	fmt.Print("1. GetCurrentWeather\n2. GetWeatherForecast\n3. ReportWeatherCondition\n")

	var choice int
	fmt.Print("Choice... > ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var location string
		fmt.Print("Locationni kiriting: ")
		fmt.Scan(&location)

		resp, err := serv.Weather.GetCurrentWeather(ctx, &weather.CurrentWeatherRequest{Location: location})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp)

	case 2:
		var days int32
		var location string
		fmt.Print("Qaysi lokatsiyani obi-havosini bilishni ishyatsiz: ")
		fmt.Scan(&location)
		fmt.Print("Necha kunning obi-havosi kerak: ")
		fmt.Scan(&days)

		resp, err := serv.Weather.GetWeatherForecast(ctx, &weather.WeatherForecastRequest{Location: location, Days: days})

		if err != nil {
			log.Fatal(err)
		}

		for _, v := range resp.Forecasts {
			fmt.Println(v)
		}

	case 3:
		var date string
		var location string
		fmt.Print("Qaysi lokatsiyani obi-havosini bilishni ishyatsiz: ")
		fmt.Scan(&location)
		fmt.Print("Qaysi sannaning obi-havosi kerak: ")
		fmt.Scan(&date)

		resp, err := serv.Weather.ReportWeatherCondition(ctx, &weather.ReportWeatherConditionRequest{Location: location, Date: date})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp)
	}
}
