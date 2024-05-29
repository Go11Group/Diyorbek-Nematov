package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
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

	companeID := map[string]string{
		"Amazon":    uuid.NewString(),
		"Apple":     uuid.NewString(),
		"Microsoft": uuid.NewString(),
		"Alphabet":  uuid.NewString(),
		"Meta":      uuid.NewString(),
	}

	companyInsertQuery := fmt.Sprintf(`
		INSERT INTO company (id, name, address) 
		VALUES
			('%s', 'Amazon', '410 Terry Ave N, Seattle, WA 98109, USA'),
			('%s', 'Apple', '1 Apple Park Way, Cupertino, CA 95014, USA'),
			('%s', 'Microsoft', 'One Microsoft Way, Redmond, WA 98052, USA'),
			('%s', 'Alphabet', '1600 Amphitheatre Parkway, Mountain View, CA 94043, USA'),
			('%s', 'Meta', '1 Hacker Way, Menlo Park, CA 94025, USA');
	`, companeID["Amazon"], companeID["Apple"], companeID["Microsoft"], companeID["Alphabet"], companeID["Meta"])

	_, err = db.Exec(companyInsertQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Company added successfully.")

	employeeINsertQuery := fmt.Sprintf(`
		INSERT INTO employee (name, position, company_id) 
		VALUES
			('Alice Johnson', 'Software Engineer', '%s'),
			('Bob Smith', 'Data Scientist', '%s'),
			('Carol Williams', 'Product Manager', '%s'),
			('David Brown', 'Designer', '%s'),
			('Eve Davis', 'HR Manager', '%s'),
			('Frank Moore', 'DevOps Engineer', '%s'),
			('Grace Wilson', 'Marketing Specialist', '%s'),
			('Hank Taylor', 'Sales Manager', '%s'),
			('Ivy Anderson', 'Business Analyst', '%s'),
			('Jack White', 'Support Engineer', '%s');
	`, companeID["Amazon"], companeID["Amazon"], companeID["Apple"], companeID["Apple"], companeID["Microsoft"], companeID["Microsoft"],
		companeID["Alphabet"], companeID["Alphabet"], companeID["Meta"], companeID["Meta"])

	_, err = db.Exec(employeeINsertQuery)
	if err != nil {
		panic(err)
	}

	fmt.Println("Employees added successfully.")
}
