package postgres

import (
	"database/sql"
	"fmt"
	"learning_app/models"
	"learning_app/pkg"
)

type EnrollmentRepo struct {
	DB *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{DB: db}
}

func (e *EnrollmentRepo) CreateEnrollment(enrolment models.Enrollment) error {
	fmt.Println(enrolment)
	_, err := e.DB.Exec(`
		INSERT INTO enrollments (
			user_id,
			course_id,
			enrollment_date
		) VALUES($1, $2, $3)
	`, enrolment.UserID, enrolment.CourseID, enrolment.EnrollmentDate)

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
		WHERE enrollment_id = $1 AND deleted_at = 0
	`, id).Scan(&enrollment.ID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate)
	fmt.Println(err)
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
	query += " updated_at = CURRENT_TIMESTAMP WHERE enrollment_id = :id AND deleted_at = 0"
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
		UPDATE Enrollments 
			SET deleted_at = DATE_PART('epoch', CURRENT_TIMESTAMP)::INT
		 WHERE enrollment_id = $1 AND deleted_at = 0
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
		args   []interface{}
		filter string
	)

	query := "SELECT enrollment_id, user_id, course_id, enrollment_date FROM Enrollments WHERE deleted_at = 0 "
	fmt.Println(fEnrollment)
	if fEnrollment.UserID != "" {
		params["user_id"] = fEnrollment.UserID
		filter += "AND user_id = :user_id "
	}
	if fEnrollment.CourseID != "" {
		params["course_id"] = fEnrollment.CourseID
		filter += "AND course_id = :course_id "
	}
	if fEnrollment.EnrollmentDate != "" {
		params["enrollment_date"] = fEnrollment.EnrollmentDate
	}
	if fEnrollment.Offset > 0 {
		params["offset"] = fEnrollment.Offset
		filter += "OFFSET :offset "
	}
	if fEnrollment.Limit > 0 {
		params["limit"] = fEnrollment.Limit
		filter += "LIMIT :limit "
	}

	query += filter

	query, args = pkg.ReplaceQueryParams(query, params)
	fmt.Println(query)
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
