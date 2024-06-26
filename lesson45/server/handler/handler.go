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
	resp, err := s.Liblary.CreateBook(in)
	return &pb.AddBookResponse{BookId: resp}, err
}

func (s *Server) SearchBook(ctx context.Context, in *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	resp, err := s.Liblary.GetBookById(in.Query)

	return &pb.SearchBookResponse{Books: resp}, err
}

func (s *Server) BorrowBook(ctx context.Context, in *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	resp, err := s.Liblary.BorrowBook(in.UserId, in.BookId)
	
	return resp, err
}