package postgres

import (
	"database/sql"
	"fmt"
	pb "library/generated/library_server"
)

type LibraryRepo struct {
	DB *sql.DB
}

func NewLibraryRepo(db *sql.DB) *LibraryRepo {
	return &LibraryRepo{DB: db}
}

func (l *LibraryRepo) CreateBook(book *pb.AddBookRequest) (string, error) {
	var book_id string
	err := l.DB.QueryRow(`
		INSERT INTO books (
			title, 
			author, 
			year_published
		)
			VALUES($1, $2, $3) 
			RETURNING book_id
	`, book.Title, book.Author, book.YearPublished).Scan(&book_id)

	return book_id, err
}

func (l *LibraryRepo) GetBookById(query string) ([]*pb.Book, error) {
	var books []*pb.Book
	rows, err := l.DB.Query(`
		SELECT 
			book_id, 
			title, 
			author, 
			year_published
		FROM
			books
		WHERE
			title = $1 OR author = $1
	`, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := &pb.Book{}

		err = rows.Scan(&book.BookId, &book.Title, &book.Author, &book.YearPublished)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (l *LibraryRepo) BorrowBook(userId, bookId string) (*pb.BorrowBookResponse, error) {
	result, err := l.DB.Exec(`
		UPDATE books SET user_id = $1, is_rental = true WHERE user_id IS NULL and book_id = $2
	`, userId, bookId)
	if err != nil {
		return nil, fmt.Errorf("failed to update book rental status: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("failed to get affected rows: %v", err)
	}

	response := &pb.BorrowBookResponse{
		Success: rowsAffected > 0,
	}

	return response, nil
}
