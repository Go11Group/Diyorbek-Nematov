package main

import (
	"fmt"
	"log"
	"transaction/models"
	"transaction/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Users
	users := []models.User{
		{ID: 1, UserName: "testuser", Email: "testuser@example.com", Password: "password123"},
		{ID: 2, UserName: "Alice", Email: "alice@example.com", Password: "password1", ProductName: "Laptop"},
		{ID: 3, UserName: "Bob", Email: "bob@example.com", Password: "password2", ProductName: "Headphones"},
		{ID: 4, UserName: "Charlie", Email: "charlie@example.com", Password: "password3", ProductName: "Smartwatch"},
		{ID: 5, UserName: "David", Email: "david@example.com", Password: "password4", ProductName: "Camera"},
		{ID: 6, UserName: "Eve", Email: "eve@example.com", Password: "password5", ProductName: "Tablet"},
	}

	// Products
	products := []models.Product{
		{ID: 1, Name: "testproduct", Description: "This is a test product", Price: 99.99, StockQuantity: 100},
		{ID: 2, Name: "Laptop", Description: "High-performance laptop for professional use", Price: 1299.99, StockQuantity: 20, UserName: "Alice"},
		{ID: 3, Name: "Headphones", Description: "Noise-canceling headphones for immersive sound experience", Price: 199.99, StockQuantity: 30, UserName: "Bob"},
		{ID: 4, Name: "Smartwatch", Description: "Fitness tracker with smartwatch features", Price: 149.99, StockQuantity: 40, UserName: "Charlie"},
		{ID: 5, Name: "Camera", Description: "Digital camera with advanced photography capabilities", Price: 899.99, StockQuantity: 15, UserName: "David"},
		{ID: 6, Name: "Tablet", Description: "Tablet device for entertainment and productivity", Price: 499.99, StockQuantity: 25, UserName: "Eve"},
	}

	// UserProducts
	userProducts := []models.UserProduct{
		{ID: 1, UserID: 1, ProductID: 1},
		{ID: 2, UserID: 2, ProductID: 2}, // Alice dan Laptop
		{ID: 3, UserID: 3, ProductID: 3}, // Bob dan Headphones
		{ID: 4, UserID: 4, ProductID: 4}, // Charlie dan Smartwatch
		{ID: 5, UserID: 5, ProductID: 5}, // David dan Camera
		{ID: 6, UserID: 6, ProductID: 6}, // Eve dan Tablet
	}

	userRepo := postgres.NewUserRepo(db)
	productRepo := postgres.NewProductRepo(db)
	userProductRepo := postgres.NewUserProductRepo(db)

	// ------------------- users tablega qiyamt kiritish

	for _, v := range users {
		err = userRepo.CreateUser(v)
		if err != nil {
			log.Fatalf("Failed to create user: %v", err)
		}
	}
	fmt.Println("User created successfully.")
	
	// --------------------- product tablega qiymat kiritish
	
	for _, v := range products {
		err = productRepo.CreateProduct(v)
		if err != nil {
			log.Fatalf("Failed to create product: %v", err)
		}
	}
	fmt.Println("Product created successfully.")
	
	// --------------------- user_product tablega qiymat berish
	for _, v := range userProducts {
		err = userProductRepo.CreateUserProduct(v)
		if err != nil {
			log.Fatalf("Failed to create user product: %v", err)
		}
	}
	fmt.Println("UserProduct created successfully.")
	


	user := users[0]
	product := products[0]
	userProduct := userProducts[0]

	

	// users tableni o'qish

	users1, err := userRepo.GetUsers()
	if err != nil {
		log.Fatalf("Failed to get users: %v", err)
	}
	fmt.Printf("Users: %+v\n", users)

	user.ID = users1[0].ID
	user.UserName = "updateduser"
	err = userRepo.UpdateUser(user)
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}
	fmt.Println("User updated successfully.")

	// products tableni o'qish
	products1, err := productRepo.GetProducts()
	if err != nil {
		log.Fatalf("Failed to get products: %v", err)
	}
	fmt.Printf("Products: %+v\n", products1)

	product.ID = products[0].ID
	product.Name = "updatedproduct"
	err = productRepo.UpdateProduct(product)
	if err != nil {
		log.Fatalf("Failed to update product: %v", err)
	}
	fmt.Println("Product updated successfully.")

	// user_products tableni o'qish
	userProducts1, err := userProductRepo.GetUserProducts()
	if err != nil {
		log.Fatalf("Failed to get user products: %v", err)
	}
	fmt.Printf("UserProducts: %+v\n", userProducts1)

	userProduct.ID = userProducts[0].ID
	userProduct.ProductID = product.ID
	err = userProductRepo.UpdateUserProduct(userProduct)
	if err != nil {
		log.Fatalf("Failed to update user product: %v", err)
	}
	fmt.Println("UserProduct updated successfully.")

	// delete

	err = userProductRepo.DeleteUserProduct(userProduct.ID)
	if err != nil {
		log.Fatalf("Failed to delete user product: %v", err)
	}
	fmt.Println("UserProduct deleted successfully.")

	err = productRepo.DeleteProduct(product.ID)
	if err != nil {
		log.Fatalf("Failed to delete product: %v", err)
	}
	fmt.Println("Product deleted successfully.")

	err = userRepo.DeleteUser(user.ID)
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}
	fmt.Println("User deleted successfully.")
}
