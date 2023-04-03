package main

import (
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
	db.Create(&Product{
		ID:    1,
		Name:  "Teste Product",
		Price: 10,
	})
}
