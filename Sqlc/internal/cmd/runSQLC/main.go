package main

import (
	"context"
	"database/sql"
	"teste/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbNewConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/teste")
	if err != nil {
		panic(err)
	}
	defer dbNewConn.Close()

	queries := db.New(dbNewConn)

	// err = queries.CreateProduct(ctx, db.CreateProductParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "teste",
	// 	Description: sql.NullString{String: "Teste desc", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// err = queries.UpdateProduct(ctx, db.UpdateProductParams{
	// 	Name:        "Teste corrigido",
	// 	Description: sql.NullString{String: "Teste desc corrigido"},
	// 	ID:          "cfd7c491-e39f-4b31-9fd9-39c7d42ecccc",
	// })
	// if err != nil {
	// 	panic(err)
	// }

	err = queries.DeleteProduct(ctx, "2c3addb1-fc98-4339-895e-c3a62f4113d8")

	categories, err := queries.ListProducts(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

}
