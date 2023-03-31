package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// product := NewProduct("Notebook", 1549.99)
	// err = insertProduct(db, product)
	// if err != nil {
	// 	panic(err)
	// }
	// product.Price = 1600.0
	// err = updateProduct(db, product)
	// if err != nil {
	// 	panic(err)
	// }
	// p, err := selectOneProduct(db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Product: %v, Price: %.2f", p.Name, p.Price)
	products, err := selectMultipleProduct(db)
	for _,p  := range products {
		fmt.Printf("Id: %v - Product: %v, Price: %.2f\n", p.ID, p.Name, p.Price)

	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare(`INSERT INTO products(id, name, price) VALUES(?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare(`UPDATE products set name = ?, price = ? where id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("SELECT id, name, price from products where id = ? ")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var p Product 
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func selectMultipleProduct(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var p []Product
	for rows.Next() {
		var p2 Product
		err = rows.Scan(&p2.ID, &p2.Name, &p2.Price)
		if err != nil {
			return nil, err
		}
		p = append(p, p2)
	}
	return p, nil

}