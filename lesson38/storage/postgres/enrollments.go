package postgres

import (
	"database/sql"
	"fmt"
	"learning_app/models"
	"learning_app/pkg"
	"time"
)

type EnrollmentRepo struct {
	DB *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{DB: db}
}

func (e *EnrollmentRepo) CreateEnrollment(enrolment models.Enrollment) error {
	_, err := e.DB.Exec(`
		INSERT INOT enrollments (
			enrollment_id,
			user_id,
			course_id,
			enrollment_date,
			created_at,
			updated_at
		) VALUES($1, $2, $3, $4)
	`, enrolment.ID, enrolment.UserID, enrolment.CourseID, enrolment.EnrollmentDate, time.Now(), time.Now())

	return err
}

func (e *EnrollmentRepo) GetEnrollmentByID(id string) (models.Enrollment, error) {
	var enrollment models.Enrollment
	err := e.DB.QueryRow(`
		SELECT 
			enrollment_id,
			user_id,  
			course_id,
			enrollment_date
		FROM enrollments 
		WHERE id = $1 AND deleted_at = 0
	`).Scan(&enrollment.ID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate)

	return enrollment, err
}

func (e *EnrollmentRepo) GetEnrollments() ([]models.Enrollment, error) {
	var enrollments []models.Enrollment
	rows, err := e.DB.Query(`
		SELECT 
			enrollment_id,
			user_id, 
			course_id,
			enrollment_date
		FROM enrollments 
		WHERE deleted_at = 0
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var enrollment models.Enrollment

		err = rows.Scan(&enrollment.ID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate)
		if err != nil {
			return nil, err
		}

		enrollments = append(enrollments, enrollment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return enrollments, err
}

func (e *EnrollmentRepo) UpdateEnrollment(enrollment models.Enrollment) error {
	params := make(map[string]interface{})
	var query = "UPDATE enrollments SET "
	if enrollment.UserID != "" {
		query += "user_id = :user_id, "
		params["user_id"] = enrollment.UserID
	}
	if enrollment.CourseID != "" {
		query += "course_id = :course_id, "
		params["course_id"] = enrollment.CourseID
	}
	query += " updated_at = CURRENT_TIMESTAMP WHERE id = :id AND deleted_at = 0"
	params["id"] = enrollment.ID

	query, args := pkg.ReplaceQueryParams(query, params)

	_, err := e.DB.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (e *EnrollmentRepo) DeleteEnrollment(id string) error {
	res, err := e.DB.Exec(`
		UPDATE enrollmets 
			SET deleted_at = DATE_PART('epoch', CURRENT_TIMESTAMP)::INT
		 WHERE id = $1 AND deleted_at = 0
	`, id)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected, user with id %s not found or already deleted", id)
	}

	return nil
}

func (e *EnrollmentRepo) GetAllEnrollments(fEnrollment models.FilterEnrollment) ([]models.Enrollment, error) {
	var (
		params = make(map[string]interface{})
		args    []interface{}
		filter string
	)

	query := "SELECT enrollment_id, user_id, course_id, enrollment_date FROM users WHERE deleted_at = 0 "

	if fEnrollment.UserID != "" {
		params["user_id"] = fEnrollment.UserID
		filter += "AND user_id = :user_id "
	}
	if fEnrollment.CourseID != "" {
		params["course_id"] = fEnrollment.CourseID
		filter += "AND course_id = :course_id "
	}
	if fEnrollment.EnrollmentDate != nil {
		params["enrollment_date"] = fEnrollment.EnrollmentDate.Format("2006-01-02")
	}
	if fEnrollment.Offset > 0 {
		params["offset"] = fEnrollment.Offset
		filter += "AND OFFSET = :offset "
	}
	if fEnrollment.Limit > 0 {
		params["limit"] = fEnrollment.Limit
		filter += "AND LIMIT = :limit "
	}

	query += filter

	query, args = pkg.ReplaceQueryParams(query, params)

	rows, err := e.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var enrollments []models.Enrollment
	for rows.Next() {
		var enrollment models.Enrollment

		err = rows.Scan(&enrollment.ID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate)

		if err != nil {
			return nil, err
		}

		enrollments = append(enrollments, enrollment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return enrollments, err
}