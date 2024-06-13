package models

import "time"

type User struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Birthday *time.Time `json:"birthday"`
	Password string     `json:"password"`
}

type Course struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Lesson struct {
	ID       string `json:"id"`
	CourseID string `json:"course_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type Enrollment struct {
	ID             string     `json:"id"`
	UserID         string     `json:"user_id"`
	CourseID       string     `json:"course_d"`
	EnrollmentDate *time.Time `json:"enrollment_id"`
}
