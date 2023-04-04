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
}

func main() {
	dsn := "root:1234@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})
	// products := []Product{
	// 	{Name: "Notebook", Price: 1000},
	// 	{Name: "Smartphone", Price: 2000},
	// 	{Name: "Iphone", Price: 3000},
	// }
	// db.Create(&products)
	var products []Product
	// db.First(&products, 1)
	// fmt.Println(products)
	// db.First(&products, "name = ?", "Iphone")
	// fmt.Println(products)
	db.Find(&products)
	fmt.Println(products)
}
