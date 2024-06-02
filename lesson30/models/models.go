package models

type User struct {
	ID int
	UserName string
	Email string
	Password string
	ProductName string
}

type Product struct {
	ID int
	Name string
	Description string
	Price float64
	StockQuantity int
	UserName string
}

type UserProduct struct {
	ID int
	UserID int
	ProductID int
}