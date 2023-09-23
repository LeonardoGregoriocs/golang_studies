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
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?utf8mb4&parseTime=True&loc=Local"
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

	// Select two rows
	var productLimit []Product
	db.Limit(2).Find(&productLimit)
	fmt.Printf("\n Select two rows\n")
	for _, product := range productLimit {
		fmt.Println(product)
	}

	// Select with off set
	var productOffSet []Product
	db.Limit(4).Offset(2).Find(&productOffSet)
	fmt.Printf("\n Select Off set\n")
	for _, product := range productOffSet {
		fmt.Println(product)
	}

	// Select with where
	var productWhere []Product
	db.Where("price > ?", 100).Find(&productWhere)
	fmt.Printf("\n Select with where\n")
	for _, product := range productWhere {
		fmt.Println(product)
	}

	// Select with Like
	var productLike []Product
	db.Where("name LIKE ?", "%P%").Find(&productLike)
	fmt.Printf("\n Select with Like\n")
	for _, product := range productLike {
		fmt.Println(product)
	}

	// Update product cod 1
	var p Product
	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)

	// Delete item
	db.Delete(&p2)
}
