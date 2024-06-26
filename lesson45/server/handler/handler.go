package handler

import (
	"context"
	pb "library/generated/library_server"
	"library/storage/postgres"
)

type Server struct {
	pb.UnimplementedLibraryServiceServer

	Liblary *postgres.LibraryRepo
}

func (s *Server) AddBook(ctx context.Context, in *pb.AddBookRequest) (*pb.AddBookResponse, error) {

	return &pb.AddBookResponse{}, nil
}