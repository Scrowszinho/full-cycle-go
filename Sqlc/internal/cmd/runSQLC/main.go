package main

import (
	"context"
	"database/sql"
	"teste/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	dbNewConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/teste")
	if err != nil {
		panic(err)
	}
	defer dbNewConn.Close()

	queries := db.New(dbNewConn)

	err = queries.CreateProduct(ctx, db.CreateProductParams{
		ID:          uuid.New().String(),
		Name:        "teste",
		Description: sql.NullString{String: "Teste desc"},
	})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListProducts(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

}
