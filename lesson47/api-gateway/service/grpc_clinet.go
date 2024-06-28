package service

import (
	"log"

	t "client/generated/transport"
	w "client/generated/weather"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceMeneger struct {
	Weather   w.WeatherServiceClient
	Transport t.TransportServiceClient
}

func New() *ServiceMeneger {

	// weather service connection
	weatherConn, err := grpc.NewClient(":50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("error while connecting to post service : ", err)
	}
	weatherService := w.NewWeatherServiceClient(weatherConn)

	// transport service connection
	transportConn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("error while connecting to post service : ", err)
	}
	transportService := t.NewTransportServiceClient(transportConn)

	return &ServiceMeneger{
		Weather:   weatherService,
		Transport: transportService,
	}

}

