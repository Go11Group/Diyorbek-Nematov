package postgres

import (
	"database/sql"
	"my_module/models"
	"time"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (p *ProductRepo) Create(product models.Product) error {
	tx, err := p.DB.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = p.DB.Exec(`
		INSERT INTO product (id, Name, price, created_at)
		VALUES
			($1, $2, $3, $4)
	`, product.ID, product.Name, product.Price, time.Now())

	return err
}

func (p *ProductRepo) GetProductByName(name string) (int, error) {
	count := 0
	err := p.DB.QueryRow("SELECT count(1) FROM product WHERE name = $1", name).Scan(&count)

	return count, err
}

func (p *ProductRepo) GetProductByNamePrice (name string, price float64) (int, error) {
	var count int

	err := p.DB.QueryRow("SELECT count(1) FROM product WHERE price = $1 and name = $2", price, name).Scan(&count)

	return count, err
}

func (p *ProductRepo) GetProductByPriceName (price float64, name string) (int, error) {
	var count int

	err := p.DB.QueryRow("SELECT count(1) FROM product WHERE name=$1 and price = $2", name, price).Scan(&count)

	return count, err
}


