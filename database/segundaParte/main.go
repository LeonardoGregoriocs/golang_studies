package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 1000.0,
	// })

	// Create batch
	// products := []Product{
	// 	{Name: "PS5", Price: 5000},
	// 	{Name: "PS4", Price: 2000},
	// 	{Name: "Camisete", Price: 50.00},
	// }
	// db.Create(&products)

	// Select one
	var product, productTwo Product
	fmt.Printf("Select One with filter id\n")
	db.First(&product, 1)
	fmt.Println(product)

	fmt.Printf("\nSelect One with filter name\n")
	db.First(&productTwo, "name = ?", "PS5")
	fmt.Println(productTwo)

	// Select All
	var products []Product
	db.Find(&products)
	fmt.Printf("\nSelect All\n")
	for _, product := range products {
		fmt.Println(product)

	}

}
