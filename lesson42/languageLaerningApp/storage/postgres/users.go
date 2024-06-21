package postgres

import (
	"database/sql"
	"fmt"
	"learning_app/models"
	"learning_app/pkg"
	"time"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) CreateUser(user models.User) error {
	_, err := u.DB.Exec(`
		INSERT INTO users (
			name, email, 
			birthday, 
			password, 
			created_at, 
			updated_at
		)
		VALUES
			($1, $2, $3, $4, $5, $6)
	`, user.Name, user.Email, user.Birthday, user.Password, time.Now(), time.Now())

	return err
}

func (u *UserRepo) GetUserByID(id string) (models.User, error) {
	var user models.User
	err := u.DB.QueryRow(`
		SELECT 
			user_id, 
			name, 
			email, 
			birthday, 
			password
		FROM users 
		WHERE user_id = $1 AND deleted_at = 0
	`, id).Scan(&user.ID, &user.Name, &user.Email, &user.Birthday, &user.Password)
	return user, err
}

func (u *UserRepo) GetUsers() ([]models.User, error) {
	var users []models.User
	rows, err := u.DB.Query(`
		SELECT 
			user_id, 
			name, 
			email, 
			birthday, 
			password
		FROM users 
		WHERE deleted_at = 0
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Birthday, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, err
}

func (u *UserRepo) UpdateUser(user models.User) error {
    params := make(map[string]interface{})
    var query = "UPDATE users SET "
    if user.Name != "" {
        query += "name = :name, "
        params["name"] = user.Name
    }
    if user.Email != "" {
        query += "email = :email, "
        params["email"] = user.Email
    }
    if user.Password != "" {
        query += "password = :password, "
        params["password"] = user.Password
    }
    if user.Birthday != "" {
        query += "birthday = :birthday, "
        params["birthday"] = user.Birthday
    }

    query += "updated_at = CURRENT_TIMESTAMP WHERE user_id = :id AND deleted_at = 0"
    params["id"] = user.ID
    query, args := pkg.ReplaceQueryParams(query, params)

    res, err := u.DB.Exec(query, args...)
    if err != nil {
        return err
    }

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected, user with id %s not found", user.ID)
	}
    return nil
}

func (u *UserRepo) DeleteUser(id string) error {
	res, err := u.DB.Exec(`
		UPDATE users 
			SET deleted_at = DATE_PART('epoch', CURRENT_TIMESTAMP)::INT
		WHERE user_id = $1 AND deleted_at = 0
	`, id)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected, user with id %s not found or already deleted", id)
	}

	return nil
}

func (u *UserRepo) GetAllUsers(fUser models.FilterUser) ([]models.User, error) {
	var (
		params = make(map[string]interface{})
		args   []interface{}
		filter string
	)
	fmt.Println(fUser)
	query := "SELECT user_id, name, email, birthday, password FROM users WHERE deleted_at = 0 "

	if fUser.Name != "" {
		params["name"] = fUser.Name
		filter += "AND name = :name "
	}
	if fUser.Email != "" {
		params["email"] = fUser.Email
		filter += "AND email = :email "
	}
	if fUser.Birthday != "" {
		params["birthday"] = fUser.Birthday
	}
	if fUser.Offset > 0 {
		params["offset"] = fUser.Offset
		filter += " OFFSET :offset "
	}
	if fUser.Limit > 0 {
		params["limit"] = fUser.Limit
		filter += "LIMIT :limit "
	}
	fmt.Println(filter)
	query += filter

	query, args = pkg.ReplaceQueryParams(query, params)
	rows, err := u.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Birthday, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, err
}
