package main

import (
	"log"
	"net"
	pb "weather/generated/weather"
	"weather/service"
	"weather/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	weatherService := service.WeatherService{Weather: postgres.NewWeatherRepo(db)}

	pb.RegisterWeatherServiceServer(s, &weatherService)

	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
