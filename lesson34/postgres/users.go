package postgres

import (
	"database/sql"
	"transaction/models"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo (db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) CreateUser(user models.User) error {
	_, err := u.DB.Exec(`
		INSERT INTO users(id, username, email, password) 
		VALUES
			($1, $2, $3, $4)
	`, user.ID, user.UserName, user.Email, user.Password)

	return err
}

func (u *UserRepo) GetUsers() ([]models.User, error) {
	var users []models.User

	rows, err := u.DB.Query(`
		SELECT u.id, username, email, password, name
		FROM users as u 
		INNER JOIN user_products as up ON u.id = up.user_id
		INNER JOIN products as p ON up.product_id = p.id
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.ProductName)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepo) GetUserById(id int) (models.User, error) {
	var user models.User

	err := u.DB.QueryRow(`
		SELECT u.id, username, email, password, name
		FROM users as u 
		INNER JOIN user_products as up ON u.id = up.user_id
		INNER JOIN products as p ON up.product_id = p.id
		WHERE id = $1
	`, id).
		Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.ProductName)

	return user, err
}

func (u *UserRepo) UpdateUser(user models.User) error {
	_, err := u.DB.Exec(`
		UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4
	`, user.UserName, user.Email, user.Password, user.ID)
	return err
}


func (u *UserRepo) DeleteUser(id int) error {
	_, err := u.DB.Exec(`
		DELETE FROM users WHERE id = $1
	`, id)

	return err
}
