package main

import (
	"log"
	pb "transport/generated/transport"
	"transport/service"
	"transport/storage/postgres"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	transportService := service.TransportService{Transport: postgres.NewTransportRepo(db)}

	pb.RegisterTransportServiceServer(s, &transportService)

	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
