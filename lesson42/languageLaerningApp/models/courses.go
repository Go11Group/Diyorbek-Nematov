package models

type Course struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type FilterCourse struct {
	Title       string `json:"title"`
	Description string `jons:"description"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
}

type PopularCourse struct {
	CourseID         string `json:"course_id"`
	CourseTitle      string `json:"course_title"`
	EnrollmentsCount int    `json:"enrollment_count"`
}
