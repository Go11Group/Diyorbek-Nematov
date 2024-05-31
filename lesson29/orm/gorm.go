package orm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("postgres://postgres:03212164@localhost:5432//gorm?sslmode=disable"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
