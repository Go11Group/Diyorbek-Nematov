package service

import (
    "log"

    t "client/generated/transport"
    w "client/generated/weather"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

type ServiceManager struct {
    Weather   w.WeatherServiceClient
    Transport t.TransportServiceClient
}

func New() *ServiceManager {
    // weather service connection
    weatherConn, err := grpc.Dial(":50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatal("error while connecting to weather service: ", err)
    }
    weatherService := w.NewWeatherServiceClient(weatherConn)

    // transport service connection
    transportConn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatal("error while connecting to transport service: ", err)
    }
    transportService := t.NewTransportServiceClient(transportConn)

    return &ServiceManager{
        Weather:   weatherService,
        Transport: transportService,
    }
}
