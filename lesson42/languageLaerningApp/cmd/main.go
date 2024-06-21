package main

import (
	"learning_app/api"
	"learning_app/api/handler"
	"learning_app/storage/postgres"
	"log"
)

func main() {
	// databases connection
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// crudlar uchun structlarni yaratish
	u := postgres.NewUserRepo(db)
	c := postgres.NewCourseRepo(db)
	l := postgres.NewLessonRepo(db)
	e := postgres.NewEnrollmentRepo(db)
	a := postgres.NewAdditionalRepo(db)

	// handler stuructini yaratish
	handler := handler.Handler{
		User:       u,
		Course:     c,
		Lesson:     l,
		Enrollment: e,
		Additional: a,
	}

	// gin freamworkini run qilish
	router := api.Router(handler)

	err = router.Run(":8081")
	log.Fatal(err)
}
