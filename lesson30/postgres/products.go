package postgres

import (
	"database/sql"
	"transaction/models"
)

type ProductRepo struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {

	return &ProductRepo{DB: db}
}

func (p *ProductRepo) CreateProduct(product models.Product) error {
	_, err := p.DB.Exec(`
		INSERT INTO products(id, name, description, price, stock_quantity)
		VALUES
			($1, $2, $3, $4, $5)
	`, product.ID, product.Name, product.Description, product.Price, product.StockQuantity)

	return err
}

func (p *ProductRepo) GetProducts() ([]models.Product, error) {
	var products []models.Product

	rows, err := p.DB.Query(`
		SELECT p.id, name, description, price, stock_quantity, username 
		FROM products as p
		INNER JOIN user_products as up ON p.id = up.product_id
		INNER JOIN users as u ON up.user_id = u.id;
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product models.Product

		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.StockQuantity, &product.UserName)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (u *ProductRepo) GetProductById(id int) (models.Product, error) {
	var product models.Product
	err := u.DB.QueryRow(`
		SELECT p.id, name, description, price, stock_quantity, username 
		FROM products as p
		INNER JOIN user_products as up ON p.id = up.product_id
		INNER JOIN users as u ON up.user_id = u.id
		WHERE p.id = $1
	`, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.StockQuantity, &product.UserName)

	return product, err
}

func (p *ProductRepo) UpdateProduct(product models.Product) error {
	_, err := p.DB.Exec(`
		UPDATE products SET name = $1, description = $2, price = $3, stock_quantity = $4 WHERE id = $5
	`, product.Name, product.Description, product.Price, product.StockQuantity, product.ID)

	return err
}

func (p *ProductRepo) DeleteProduct(id int) error {
	_, err := p.DB.Exec(`
		DELETE FROM products WHERE id = $1
	`, id)

	return err
}
