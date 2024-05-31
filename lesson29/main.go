package main

import (
	"log"
	"my_module/orm"
)

func main() {
	db, err := orm.ConnectDB()
	if err != nil {
		panic(err)
	}

	us := orm.NewUser(db)
	// err = us.Create(models.User{
	// 	Model: gorm.Model{ID: 1},
	// 	FirstName:  "Akobir",
	// 	LastName:   "Usmonov",
	// 	Email:      "usmonovakobir04@gmai.com",
	// 	Password:   "qwerty2004",
	// 	Age:        20,
	// 	Field:      "FrontEnd",
	// 	Gender:     "male",
	// 	IsEmployee: false,
	// })

	// if err != nil {
	// 	log.Fatal("Error insert databases", err.Error())
	// }

	// err = us.Create(models.User{
	// 	Model: gorm.Model{ID: 2},
	// 	FirstName: "Asadbek",
	// 	LastName: "Aymatov",
	// 	Email: "asadbekaymatov01@gmail.com",
	// 	Password: "0103asadbek",
	// 	Age: 23,
	// 	Field: "Kiberxavsizlik",
	// 	Gender: "male",
	// 	IsEmployee: true,
	// })

	// if err != nil {
	// 	log.Fatal("Error isert databases", err.Error())
	// }

	// users, err := us.GetAllUsers()
	// if err != nil {
	// 	log.Fatal("Error read in database", err.Error())
	// }

	// fmt.Println(users)

	// user, err := us.GetByIdUser(2)
	// if err != nil {
	// 	log.Fatal("Error in read Database", err.Error())
	// }

	// fmt.Println(user)

	// err = us.Update(models.User{
	// 	Model:     gorm.Model{ID: 2},
	// 	FirstName: "Anvar",
	// 	LastName:  "Narzullayev",
	// 	Email:     "narzullayevanvar@gmail.com",
	// 	Password:  "asdfghjkl",
	// 	Age:       45,
	// })

	// if err != nil {
	// 	log.Fatal("Error update Databases", err.Error())
	// }

	err = us.Delete(1)
	if err != nil {
		log.Fatal("Error delete in databases", err.Error())
	}
}
