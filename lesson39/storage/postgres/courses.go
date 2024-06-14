package postgres

import (
	"database/sql"
	"fmt"
	"learning_app/models"
	"learning_app/pkg"
	"time"
)

// course uchun coursening kurudlarini birlashtirish uchun struct
type CourseRepo struct {
	DB *sql.DB
}

// yangi courseRepo yaratish
func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{DB: db}
}

// Course uchun CreateCourse methodi 
func (c *CourseRepo) CreateCourse(course models.Course) error {
	_, err := c.DB.Exec(`
		INSERT INTO courses (
			title,
			description,
			created_at,
			updated_at
		) VALUES($1, $2, $3, $4)
	`, course.Title, course.Description, time.Now(), time.Now())

	return err
}

// course table uchun malumotni id sini boyicha qidish
func (c *CourseRepo) GetCourseByID(id string) (models.Course, error) {
	var course models.Course
	err := c.DB.QueryRow(`
		SELECT 
			course_id,
			title,
			description
		FROM courses
		WHERE deleted_at = 0 AND course_id = $1
	`, id).Scan(&course.ID, &course.Title, &course.Description)

	return course, err
}

// barcha courselarni olish
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

	// databasedan kelgan malumotlarni scan qilish
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

// course tableni update qilish uchun mthod
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
    if len(params) > 0 {
        query = query[:len(query)-2] + ", "
    }
    query += "updated_at = CURRENT_TIMESTAMP WHERE course_id = :id AND deleted_at = 0"
    params["id"] = course.ID

    query, args := pkg.ReplaceQueryParams(query, params)

    _, err := c.DB.Exec(query, args...)
    if err != nil {
        return err
    }
    return nil
}

// course table uchun delete methodi id bo'yicha qidirib o'chiradi
func (c *CourseRepo) DeleteCourse(id string) error {
	res, err := c.DB.Exec(`
		UPDATE courses 
			SET deleted_at = DATE_PART('epoch', CURRENT_TIMESTAMP)::INT
		WHERE deleted_at = 0 AND course_id = $1
	`, id)

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

// course tableni turli columnlari qiymatlari bialn filterlash
func (c *CourseRepo) GetAllCourses(fCourse models.FilterCourse) ([]models.Course, error) {
	var (
		params = make(map[string]interface{})
		args    []interface{}
		filter string
	)

	query := "SELECT course_id, title, description  FROM courses WHERE deleted_at = 0 "
	// filterlash uchun kerakli ma'lumotlarni olish
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
		filter += "OFFSET :offset "
	}
	if fCourse.Limit > 0 {
		params["limit"] = fCourse.Limit
		filter += "LIMIT :limit "
	}

	query += filter

	query, args = pkg.ReplaceQueryParams(query, params)

	// databasega filterlsh uchun query jo'natish
	rows, err := c.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	// filter qilingan malumotlarni olish
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