package models

type Subject struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Credits int    `json:"credits"`
}

type Associate struct {
	StudentID string `json:"student_id"`
	SubjectID string `json:"subject_id"`
}
