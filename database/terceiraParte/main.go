package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// create category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Notebook",
		Price:      1000.0,
		CategoryID: category.ID,
	})

	// select product
	var products []Product
	db.Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}

	// select product with categories
	var productsCategory []Product
	db.Preload("Category").Find(&productsCategory)
	for _, product := range productsCategory {
		fmt.Println(product)
	}
}
