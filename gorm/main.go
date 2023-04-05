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
	Categories   []Category `gorm:"many2many:products_categories";`
	SerialNumber SerialNumber
}

type Category struct {
	ID       int `gorm:primaryKey`
	Name     string
	Products []Product `gorm:"many2many:products_categories";`
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
	// db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// category := Category{Name: "Clothing"}
	// db.Create(&category)

	// category2 := Category{Name: "Eletronics"}
	// db.Create(&category2)

	// product := Product{Name: "Smart-Watch", Price: 15, Categories: []Category{category, category2}}
	// db.Create(&product)

	// serialNumber := SerialNumber{Number: "123", ProductID: 1}
	// db.Create(&serialNumber)

	var categorys []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Preload("Products").Find(&categorys).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categorys {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			fmt.Println(product.Name)
		}
	}

}
