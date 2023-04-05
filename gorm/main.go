package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:primaryKey`
	Name  string
	Price float64
	gorm.Model
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
}

type Category struct {
	ID   int `gorm:primaryKey`
	Name string
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:primaryKey`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:1234@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// category := Category{Name: "Miscellaneous"}
	// db.Create(&category)

	// product := Product{Name: "Car Toy", Price: 5, CategoryID: category.ID}
	// db.Create(&product)

	// serialNumber := SerialNumber{Number: "a10", ProductID: 2}
	// db.Create(&serialNumber)

	var products []Product
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	fmt.Println(products[0].SerialNumber.Number)
}
