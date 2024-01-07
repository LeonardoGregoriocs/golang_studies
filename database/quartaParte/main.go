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
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

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

	// create serial number
	db.Create(&SerialNumber{
		Number:    "123456E",
		ProductID: product.ID,
	})

	// select product with categories
	var productsCategory []Product
	db.Preload("Category").Preload("SerialNumber").Find(&productsCategory)
	for _, product := range productsCategory {
		fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	}
}
