package postgres

import (
	"database/sql"
	"my_module/models"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) CreateUser(user models.User) error {
	_, err := u.DB.Exec(`
		INSERT INTO users(first_name, last_name, email, gender, age)
		VALUES($1, $2, $3, $4, $5)
	`, user.FirstName, user.LastName, user.Email, user.Gender, user.Age)

	return err
}

func (u *UserRepo) GetUserByID(id string) (models.User, error) {
	var user models.User

	err := u.DB.QueryRow(`
		SELECT id, first_name, last_name, email, gender, age
		FROM users
		WHERE id=$1
	`, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Gender, &user.Age)

	return user, err
}

func (u *UserRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User

	rows, err := u.DB.Query(`
		SELECT id, first_name, last_name, email, gender, age FROM users
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close() 

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Gender, &user.Age)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}


func (u *UserRepo) UpdateUser(user models.User) error {
	_, err := u.DB.Exec(`
		UPDATE users SET first_name=$1, last_name=$2, email=$3, gender=$4, age=$5 WHERE id=$6
	`, user.FirstName, user.LastName, user.Email, user.Gender, user.Age, user.ID)

	return err
}

func (u *UserRepo) DeleteUser(id string) error {
	_, err := u.DB.Exec(`
		DELETE FROM users WHERE id=$1
	`, id)

	return err
}
