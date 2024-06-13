package postgres

import (
	"database/sql"
	"fmt"
	"learning_app/models"
	"learning_app/pkg"
	"time"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{DB: db}
}

func (c *CourseRepo) CreateCourse(course models.Course) error {
	_, err := c.DB.Exec(`
		INSERT INTO courses (
			course_id,
			title,
			description,
			created_at,
			updated_at
		) VALUES($1, $2, $3, $4, $5)
	`, course.ID, course.Title, course.Description, time.Now(), time.Now())

	return err
}

func (c *CourseRepo) GetCourseByID(id string) (models.Course, error) {
	var course models.Course
	err := c.DB.QueryRow(`
		SELECT 
			course_id,
			title,
			description
		FROM courses
		WHERE deleted_at = 0
	`, id).Scan(&course.ID, &course.Title, &course.Description)

	return course, err
}

func (c *CourseRepo) GetCourses() ([]models.Course, error) {
	var courses []models.Course

	rows, err := c.DB.Query(`
		SELECT
			course_id,
			title,
			description
		FROM courses
		WHERE deleted_at = 0
	`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var course models.Course
		err = rows.Scan(&course.ID, &course.Title, &course.Description)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *CourseRepo) UpdateCourse(course models.Course) error {
	params := make(map[string]interface{})
	var query = "UPDATE courses SET "
	if course.Title != "" {
		query += "title = :title, "
		params["title"] = course.Title
	}
	if course.Description != "" {
		query += "description = :description, "
		params["description"] = course.Description
	}
	query = "updated_at = CURRENT_TIMESTAMP WHERE id = :id AND deleted_at = 0"
	params["id"] = course.ID

	query, args := pkg.ReplaceQueryParams(query, params)

	_, err := c.DB.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseRepo) DeleteCourse(id string) error {
	res, err := c.DB.Exec(`
		UPDATE courses 
			SET deleted_at = DATE_PART('epoch', CURRENT_TIMESTAMP)::INT
		WHERE deleted_at = 0 AND id = $1
	`, time.Now(), id)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected, course with id %s not found or already deleted", id)
	}

	return nil
}

func (c *CourseRepo) GetAllCourses(fCourse models.FilterCourse) ([]models.Course, error) {
	var (
		params = make(map[string]interface{})
		args    []interface{}
		filter string
	)

	query := "SELECT course_id, title, description  FROM courses WHERE deleted_at = 0 "

	if fCourse.Title != "" {
		params["title"] = fCourse.Title
		filter += "AND title = :title "
	}
	if fCourse.Description != "" {
		params["description"] = fCourse.Description
		filter += "AND description = :description "
	}
	if fCourse.Offset > 0 {
		params["offset"] = fCourse.Offset
		filter += "AND OFFSET = :offset "
	}
	if fCourse.Limit > 0 {
		params["limit"] = fCourse.Limit
		filter += "AND LIMIT = :limit "
	}

	query += filter

	query, args = pkg.ReplaceQueryParams(query, params)

	rows, err := c.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var courses []models.Course
	for rows.Next() {
		var course models.Course

		err = rows.Scan(&course.ID, &course.Title, &course.Description)

		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, err
}