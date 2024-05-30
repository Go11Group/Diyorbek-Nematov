package postgres

import (
	"database/sql"
	"github.com/Go11Group/at_lesson/lesson28/model"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}

func (c *CourseRepo) GetAllCars() ([]model.Course, error) {
	rows, err := c.DB.Query(`
		SELECT id, name
		FROM course
	`)
	if err != nil {
		return nil, err
	}

	var courses []model.Course
	var course model.Course
	for rows.Next() {
		err = rows.Scan(&course.Id, &course.Name)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (c *CourseRepo) GetByID(id string) (model.Course, error) {
	var course model.Course

	err := c.DB.QueryRow(`
		SELECT id, name
		FROM course
		WHERE id=$1
	`, id).
	Scan(&course.Id, &course.Name,)

	return course, err
}


func (c *CourseRepo) Create(course *model.Course) error {
	_, err := c.DB.Exec(`
		INSERT INTO course (id, name)
		VALUES
			($1, $2)
	`, course.Id, course.Name)

	if err != nil {
		return err
	}

	return nil
}

func (c *CourseRepo) Update(course *model.Course) error {
	_, err := c.DB.Exec(`
		UPDATE course SET name=$1 WHERE id=$2
	`, course.Name, course.Id)

	if err != nil {
		return err
	}

	return nil
}

func (c *CourseRepo) Delete(id string) error {
	_, err := c.DB.Exec(`
		DELETE FROM course WHERE id=$1
	`, id)

	if err != nil {
		return err
	}

	return nil
}