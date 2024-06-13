package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "learning_app"
	password = "03212164"
)

func ConnectDB() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		host, port, user, dbname, password)
	
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}