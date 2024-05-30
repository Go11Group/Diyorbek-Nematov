package main

import (
	"fmt"
	"log"

	"github.com/Go11Group/at_lesson/lesson28/storage/postgres"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	st := postgres.NewStudentRepo(db)

	users, err := st.GetAllStudents()
	if err != nil {
		log.Fatal("Error read in database", err.Error())
	}

	user, _ := st.GetByID(users[2].ID)

	fmt.Println("----------  Users  ----------")
	for i, v := range users {
		fmt.Printf("ID :%d\nName :%s\nAge :%d\nCourse :%s\n\n", i, v.Name, v.Age, v.Course)
	}
	fmt.Println(user)

	cr := postgres.NewCourseRepo(db)
	courses, err := cr.GetAllCars()
	if err != nil {
		log.Fatal("Error read in database", err.Error())
	}

	fmt.Println("----------  Course  ---------")
	for _, v := range courses {
		fmt.Printf("ID: %s\nName: %s\n\n", v.Id, v.Name)
	}
}
