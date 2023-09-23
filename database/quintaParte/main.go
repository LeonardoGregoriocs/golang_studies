package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
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
	category := Category{Name: "Celulares"}
	db.Create(&category)

	// create product
	db.Create(&Product{
		Name:       "Iphone 14",
		Price:      4000.0,
		CategoryID: category.ID,
	})

	// buscar product
	var product Product
	db.Where("Name = ?", "Iphone 14").Find(&product)

	// select product with categories
	var productsCategory []Product
	db.Preload("Category").Find(&productsCategory)
	for _, product := range productsCategory {
		fmt.Println(product.Name, product.Category.Name)
	}

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, category.Name)
		}
	}
}
