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
	ID       int `gorm:primaryKey`
	Name     string
	Products []Product
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

	// category := Category{Name: "Clothing"}
	// db.Create(&category)

	// product := Product{Name: "Shirt", Price: 15, CategoryID: 1}
	// db.Create(&product)

	var categorys []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categorys).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categorys {
		for _, product := range category.Products {
			fmt.Println(product)
		}
	}

}
