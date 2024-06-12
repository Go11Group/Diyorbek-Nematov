package postgres

import (
	"database/sql"
	"my_module/gin/models"
	"my_module/gin/pkg"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) GetAll(f models.Filter) ([]models.User, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
	)

	query := `SELECT id, first_name, last_name, age, gender, nation, field, parent_name, city
	FROM users WHERE true `

	if len(f.Gender) > 0 {
		params["gender"] = f.Gender
		filter += " and gender = :gender "
	}

	if len(f.Nation) > 0 {
		params["nation"] = f.Nation
		filter += " and nation = :nation "
	}

	if len(f.Field) > 0 {
		params["field"] = f.Field
		filter += " and field = :field "
	}

	if f.Age > 0 {
		params["age"] = f.Age
		filter += " and age = :age "
	}

	if f.Limit > 0 {
		params["limit"] = f.Limit
		filter += " LIMIT :limit"
	}

	if f.Offset > 0 {
		params["offset"] = f.Offset
		filter += " OFFSET :offset"
	}

	query = query + filter

	query, arr = pkg.ReplaceQueryParams(query, params)
	rows, err := u.DB.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Age, &user.Gender, &user.Nation, &user.Field, &user.ParentName, &user.City)

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
