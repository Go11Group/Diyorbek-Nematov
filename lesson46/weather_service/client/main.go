package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"weather/generated/weather"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	WeatherInterface(conn, ctx)

}

func WeatherInterface(conn *grpc.ClientConn, ctx context.Context) {
	c := weather.NewWeatherServiceClient(conn)

	fmt.Println("TransportSerivice ...")
	fmt.Print("1. GetCurrentWeather\n2.GetWeatherForecast\n3.ReportWeatherCondition\n")

	var choice int
	fmt.Print("Choice... > ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var location string
		fmt.Print("Locationni kiriting: ")
		fmt.Scan(&location)

		resp, err := c.GetCurrentWeather(ctx, &weather.CurrentWeatherRequest{Location: location})
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
		fmt.Println(&days)

		resp, err := c.GetWeatherForecast(ctx, &weather.WeatherForecastRequest{Location: location, Days: days})

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
		fmt.Println(&date)

		resp, err := c.ReportWeatherCondition(ctx, &weather.ReportWeatherConditionRequest{Location: location, Date: date})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp)
	}
}
