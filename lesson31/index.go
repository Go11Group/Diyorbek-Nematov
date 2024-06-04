package main

import (
	"fmt"
	"log"
	"math/rand"
	"my_module/postgres"
	"time"

	"github.com/go-faker/faker/v4"
)

func main() {

	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("ERROR in database open", err.Error())
	}
	defer db.Close()

	p := postgres.NewProductRepo(db)

	// for i := 0; i < 1000000; i++ {
	// 	var product models.Product

	// 	product.ID = uuid.NewString()
	// 	product.Name = faker.FirstName()
	// 	product.Price = rand.Float64() * 1000

	// 	err := p.Create(product)
	// 	if err != nil {
	// 		log.Fatal("ERROR insert in database", err.Error())
	// 	}

	// 	if i%1000 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	p.DB.SetMaxOpenConns(10)
	p.DB.SetMaxIdleConns(10)

	for i := 0; i < 500; i++ {
		// go func() {
		// 	t := time.Now()
		// 	name := faker.FirstName()
		// 	count, err := p.GetProductByName(name)
		// 	if err != nil {
		// 		panic(err)
		// 	}

		// 	fmt.Println(i+1, count, time.Now().Sub(t))
		// }()

		// go func() {
		// 	t := time.Now()
		// 	name := faker.FirstName()
		// 	price := rand.Float64() * 1000

		// 	count, err := p.GetProductByPriceName(price, name)
		// 	if err != nil {
		// 		panic(err)
		// 	}

		// 	fmt.Println(i+1, count, time.Now().Sub(t))
		// }()

		go func() {
			t := time.Now()
			name := faker.FirstName()
			price := rand.Float64() * 1000

			count, err := p.GetProductByNamePrice(name, price)
			if err != nil {
				panic(err)
			}

			fmt.Println(i+1, count, time.Now().Sub(t))
		}()
	}

	time.Sleep(10 * time.Second)

}
