package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "library"
	password = "03212164"
)

type Author struct {
	ID        int
	Name      string
	Birthdate string
}

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	/*
		Go dasturlash tilida db.Query funksiyasi ma'lumotlar bazasidan (DB) ma'lumot olish uchun ishlatiladi.
		Bu funksiya SQL so'rovini bajaradi va natijani *sql.Rows tipidagi struktura shaklida qaytaradi.
		Bu struktura ma'lumotlar qatorlarini o'qish uchun ishlatiladi.
	*/
	rows, err := db.Query(`SELECT * FROM author`)
	if err != nil {
		panic(err)
	}

	author := []Author{}
	for rows.Next() {
		var id int
		var name, birthdate string
		err = rows.Scan(&id, &name, &birthdate)
		if err != nil {
			panic(err)
		}
		author = append(author, Author{ID: id, Name: name, Birthdate: birthdate})
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, v := range author {
		fmt.Printf("ID :%d\nName :%s\nBirthDate :%s\n\n", v.ID, v.Name, v.Birthdate)
	}
}
