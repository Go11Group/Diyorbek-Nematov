package postgres

import (
	"auth-service/models"
	"database/sql"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) Register(req models.RegisterReq) error {
	query := "INSERT INTO users (id, username, email, password) VALUES ($1, $2, $3, $4) RETURNING id"
	var id string
	err := u.DB.QueryRow(query, req.ID, req.Username, req.Email, req.Password).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) GetByID(req models.GetProfileByIdReq) (*models.GetProfileByIdResp, error) {
	query := "SELECT id, username, email FROM users WHERE id = $1"
	user := &models.GetProfileByIdResp{}
	err := u.DB.QueryRow(query, req.ID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) Profile(req models.GetProfileReq) (*models.GetProfileResp, error) {
	query := "SELECT id, username, email, password FROM users WHERE email = $1"
	user := &models.GetProfileResp{}
	err := u.DB.QueryRow(query, req.Email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) EmailExists(email string) (bool, error) {
	var exists bool
	err := u.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email).Scan(&exists)
	return exists, err
}
