package postgres

import (
	"database/sql"
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
			VALUE($1, $2, $3)
	`, book.Title, book.Author, book.YearPublished).Scan(&book_id)

	return book_id, err
}

func (l *LibraryRepo) GetBookById(b *pb.SearhcBookRequest) ([]pb.Book, error) {
	var books []pb.Book
	rows, err := l.DB.Query(`
		SELECT 
			book_id, 
			title, 
			author, 
			year_published
		FROM
			books
		WHERE
			title = $1 OR author = $2

	`)	
	return books, nil
}
