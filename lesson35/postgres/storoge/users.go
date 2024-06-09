package postgres

import (
	"database/sql"
	"solution/models"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) CreateUser(user models.User) (error) {
	_, err := u.DB.Exec(`
		INSERT INTO users (id, username, email, password)
		VALUES
			($1, $2, $3, $4)
	`, user.ID, user.Username, user.Email, user.Password)

	return err
}

func (u *UserRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User

	rows, err := u.DB.Query(`
		SELECT id, username, email, password FROM users
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepo) GetUserByID(id int) (models.User, error) {
	var user models.User

	err := u.DB.QueryRow(`
		SELECT id, username, email, password FROM users WHERE id = $1
	`, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	return user, err
}

func (u *UserRepo) UpdateUser(user models.User) error {
	_, err := u.DB.Exec(`
		UPDATE users SET username=$1, email=$2, password=$3 WHERE id=$4	
	`, user.Username, user.Email, user.Password, user.ID)

	return err
}

func (u *UserRepo) DeleteUser(id int) error {
	_, err := u.DB.Exec(`
		DELETE FROM users WHERE id=$1
	`, id)

	return err
}