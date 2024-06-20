package postgres

import (
	"database/sql"
	"server/models"
)

type EmployeeRepo struct {
	DB *sql.DB
}

func NewEmployeeRepo(db *sql.DB) *EmployeeRepo {
	return &EmployeeRepo{DB: db}
}

func (e *EmployeeRepo) CreateEmployee(emp models.Employee) error {
	_, err := e.DB.Exec(`
		INSERT INTO employees(id, name, position, salary)
		VALUES ($1, $2, $3, $4)
	`, emp.ID, emp.Name, emp.Position, emp.Salary)

	return err
}


func (e *EmployeeRepo) GetEmployee(id int) (models.Employee, error) {
	var emp models.Employee

	err := e.DB.QueryRow("SELECT id, name, position, salary FROM employees WHERE id=$1", id).
		Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary)
	
	
	return emp, err
}

func (e *EmployeeRepo) UpdateEmployee(emp models.Employee) error {
	_, err := e.DB.Exec(`
		UPDATE employees SET name=$1, position=$2, salary=$3 WHERE id=$4
	`, emp.Name, emp.Position, emp.Salary, emp.ID)

	return err
}

func (e *EmployeeRepo) DeleteEmploee(id int) error {
	_, err := e.DB.Exec(`
		DELETE FROM employees WHERE id=$1
	`, id)

	return err
}