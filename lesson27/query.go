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
	dbname   = "companydb"
	password = "03212164"
)

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

	// 1 - query

	fmt.Println("1 - Query...")
	rows, err := db.Query(`
		SELECT c.name, COUNT(emp.name) AS "employee_coun"
		FROM company AS c 
		LEFT JOIN employee AS emp ON c.id = emp.company_id
		GROUP BY c.id;
	`)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var name string
		var count int

		err = rows.Scan(&name, &count)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Company Name: %s\nEmployee Count: %d\n\n", name, count)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	// 2 - query
	fmt.Println("2 - Query...")
	rows2, err := db.Query(`
		SELECT e.name, position, c.name as company_name
		FROM employee e
		JOIN company c ON e.company_id = c.id;
	`)

	for rows2.Next() {
		var name, position, company_name string

		err = rows2.Scan(&name, &position, &company_name)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Name :%s\nPosition :%s\nCompany Name :%s\n\n", name, position, company_name)
	}

	if err = rows2.Err(); err != nil {
		panic(err)
	}

	// 3 - query
	fmt.Println("3 - Query...")
	var employeeCount int
	err = db.QueryRow("SELECT COUNT(*) FROM employee WHERE company_id = $1", "77cfff33-e4f4-4565-b420-54177e26458c").Scan(&employeeCount)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Company %s has %d employees.\n", "Amazon", employeeCount)


	// 4 - query
	fmt.Println("4 - Query...")
	var employeeName, position string
	err = db.QueryRow("SELECT name, position FROM employee WHERE id = $1", "e9342587-ee1f-4a23-a651-089c7692405e").Scan(&employeeName, &position)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Employee: %s, Position: %s\n", employeeName, position)

}
