package models

type Enrollment struct {
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	CourseID       string `json:"course_id"`
	EnrollmentDate string `json:"enrollment_date"`
}

type FilterEnrollment struct {
	UserID         string `json:"user_id"`
	CourseID       string `json:"course_id"`
	EnrollmentDate string `json:"enrollment_date"`
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
}



