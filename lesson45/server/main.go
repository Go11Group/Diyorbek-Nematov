package main

import (
	pb "library/generated/library_server"
	"library/server/handler"
	"library/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Error connection Database!")
	}
	lib := postgres.NewLibraryRepo(db)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}

	defer listener.Close()

	s := grpc.NewServer()

	pb.RegisterLibraryServiceServer(s,  &handler.Server{Liblary: lib})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
