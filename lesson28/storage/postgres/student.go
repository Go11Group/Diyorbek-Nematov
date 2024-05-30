package postgres

import (
	"database/sql"

	"github.com/Go11Group/at_lesson/lesson28/model"
)

type StudentRepo struct {
	Db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{Db: db}
}

func (u *StudentRepo) GetAllStudents() ([]model.User, error) {
	rows, err := u.Db.Query(`
		SELECT DISTINCT s.id, s.name, s.age, c.name 
		FROM student_course as sc 
		INNER JOIN student as s ON s.id = sc.student_id
		LEFT JOIN course as c ON c.id = sc.course_id;
	`)

	if err != nil {
		return nil, err
	}

	var users []model.User
	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Course)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *StudentRepo) GetByID(id string) (model.User, error) {
	var user model.User

	err := u.Db.QueryRow(`
		SELECT s.id, s.name, s.age, c.name 
		FROM student_course as sc 
		INNER JOIN student as s ON s.id = sc.student_id
		LEFT JOIN course as c ON c.id = sc.course_id
		WHERE s.id = $1;
	`, id).
		Scan(&user.ID, &user.Name, &user.Age, &user.Course)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *StudentRepo) Create(user model.User) error {
	_, err := u.Db.Exec(`
		INSER INTO student(id, name, age)
		VALUES
			($1, $2, $3)
	`, user.ID, user.Name, user.Age)

	if err != nil {
		return err
	}

	return nil
}

func (u *StudentRepo) Update(user model.User) error {
	_, err := u.Db.Exec(`
		UPDATE student name=$1, age=$2 WHERE id=$3 
	`, user.Name, user.Age, user.ID)

	if err != nil {
		return err
	}

	return nil
}

func (u *StudentRepo) Delete( id string) error {
	_, err := u.Db.Exec(`
		DELETE FROM student WHERE id=$1
	`, id)

	if err != nil {
		return nil
	}

	return nil
}
