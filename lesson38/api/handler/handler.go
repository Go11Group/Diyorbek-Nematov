package handler

import "learning_app/storage/postgres"



type Handler struct {
	User *postgres.UserRepo
	Course *postgres.CourseRepo
	Lesson *postgres.LessonRepo
	Enrollment *postgres.EnrollmentRepo
}
