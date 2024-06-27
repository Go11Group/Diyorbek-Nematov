package main

import (
	"context"
	"fmt"
	"log"
	"transport/generated/transport"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

		TransportInterface(conn, ctx)
}

func TransportInterface(conn *grpc.ClientConn, ctx context.Context) {
	c := transport.NewTransportServiceClient(conn)

	fmt.Println("TransportSerivice ...")
	fmt.Print("1. GetBusSchedule\n2.TrackBusLocation\nReportTrafficJam\n")

	var choice int
	fmt.Print("Choice... > ")
	fmt.Scan(choice)

	switch choice {
	case 1:
		var busNumber string
		fmt.Print("Avtobus raqamini kiriting: ")
		fmt.Scan(&busNumber)
		resp, err := c.GetBusSchedule(ctx, &transport.GetBusScheduleRequest{BusNumber: busNumber})

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
		resp, err := c.TrackBusLocation(ctx, &transport.TrackBusLocationRequest{BusNumber: busNumber})

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
		fmt.Println(&location)

		resp, err := c.ReportTrafficJam(ctx, &transport.ReportTrafficJamRequest{BusNumber: busNumber, Location: location})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp)
	}
}
