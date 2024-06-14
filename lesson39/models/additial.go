package models

type CourseByUser struct {
	UserID  string   `json:"user_id"`
	Courses []Course `json:"course_id"`
}


type LessonsByCourse struct {
	CourseID string           `json:"course_id"`
	Lessons  []AdditialLesson `json:"lessons"`
}

type EnrolledUsersByCourse struct {
	CourseID      string `json:"course_id"`
	EnrolledUsers User   `json:"enrolled_users"`
}


type Result struct {
	Results []AdditialUser `json:"results"`
}

type TimeRange struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type Goal struct {
	TimePeriod    TimeRange       `json:"time_perion"`
	PopularCourse []PopularCourse `json:"popular_courses"`
}