package main

import (
	"context"
	"fmt"
	pb "library/generated/library_server"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewLibraryServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addBook := pb.AddBookRequest{
		Title: "Obid Ketmon",
		Author: "Alisher Navoiy",
		YearPublished: 1932,
	}

	respAdd, err := c.AddBook(ctx, &addBook)
	if err != nil {
		log.Fatal("Error could not add book: ", err)
	}

	fmt.Printf("Kitob muvaffaqiyatli qo'shildi: %s\n\n", respAdd.BookId)

	searchBook := pb.SearchBookRequest{
		Query: "Abdulla Qodiriy",
	}

	respSerach, err := c.SearchBook(ctx, &searchBook)
	if err != nil {
		log.Fatal("Error in search book: ", err)
	}

	fmt.Println("Qidiruv natijasi: ")
	for _, v := range respSerach.Books {
		fmt.Printf("BookId: %s\nTitle: %s\nAuthor: %s\nYearch published: %d\n\n", v.BookId, v.Title, v.Author, v.YearPublished)
	}

	borrowBook := pb.BorrowBookRequest{
		BookId: "a54d48f1-061a-4bb0-9b18-9441b9689653",
		UserId: "83b35b47-f1f3-418f-802c-339c7ae05bdb",
	}

	respBorrow, err := c.BorrowBook(ctx, &borrowBook)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	if respBorrow.Success {
		fmt.Printf("%t Kitob ijaraga berildi\n", respBorrow.Success)
	} else {
		fmt.Printf("%t Kitobni ijaraga berilmadi\n", respBorrow.Success)
	}
}