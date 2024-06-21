package models

type Lesson struct {
	ID       string `json:"id"`
	CourseID string `json:"course_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type FilterLesson struct {
	CourseID string `json:"course_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type AdditialLesson struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}