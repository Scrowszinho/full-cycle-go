// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"database/sql"
)

type Color struct {
	ID          string
	ProductID   string
	Name        string
	Description sql.NullString
	Price       float64
	Pricefinal  float64
}

type Product struct {
	ID          string
	Name        string
	Description sql.NullString
}
