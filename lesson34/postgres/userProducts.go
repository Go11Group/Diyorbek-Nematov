package postgres

import (
	"database/sql"
	"transaction/models"
)

type UserProductRepo struct {
	DB *sql.DB
}

func NewUserProductRepo(db *sql.DB) *UserProductRepo {
	return &UserProductRepo{DB: db}
}

func (up *UserProductRepo) CreateUserProduct(userProduct models.UserProduct) error {
	tx , err := up.DB.Begin()
	if err != nil {
		return err
	}

	defer func ()  {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = up.DB.Exec(`
		INSERT INTO user_products (user_id, product_id) 
		VALUES
			($1, $2) 
	`, userProduct.UserID, userProduct.ProductID)

	return err
}

func (up *UserProductRepo) GetUserProductByID(id int) (models.UserProduct, error) {
	var userProduct models.UserProduct
	err := up.DB.QueryRow(`
		SELECT id, user_id, product_id FROM user_products WHERE id = $1
	`, id).Scan(&userProduct.ID, &userProduct.UserID, &userProduct.ProductID)
	return userProduct, err
}

func (up *UserProductRepo) GetUserProducts() ([]models.UserProduct, error) {
	rows, err := up.DB.Query(`
		SELECT id, user_id, product_id FROM user_products
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userProducts []models.UserProduct
	for rows.Next() {
		var userProduct models.UserProduct
		err := rows.Scan(&userProduct.ID, &userProduct.UserID, &userProduct.ProductID)
		if err != nil {
			return nil, err
		}
		userProducts = append(userProducts, userProduct)
	}
	return userProducts, nil
}

func (up *UserProductRepo) UpdateUserProduct(userProduct models.UserProduct) error {
	tx, err := up.DB.Begin()
	if err != nil {
		return err
	}

	defer func ()  {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = up.DB.Exec(`
		UPDATE user_products SET user_id = $1, product_id = $2 WHERE id = $3
	`, userProduct.UserID, userProduct.ProductID, userProduct.ID)

	return err
}

func (up *UserProductRepo) DeleteUserProduct(id int) error {
	tx, err := up.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = tx.Exec(`
		DELETE FROM user_products WHERE id = $1
	`, id)
	return err
}