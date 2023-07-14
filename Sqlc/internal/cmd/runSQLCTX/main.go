package main

import (
	"context"
	"database/sql"
	"fmt"
	"teste/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

type ColorsDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewColorsDB(dbConn *sql.DB) *ColorsDB {
	return &ColorsDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *ColorsDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("Error on rollback: %v, original error: %v", rbErr, err)
		}
		return err
	}
	return tx.Commit()
}

type ProductParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

type ColorParams struct {
	ID          string
	Name        string
	Description sql.NullString
	ProductID   string
	Price       float64
	Pricefinal  float64
}

func (c *ColorsDB) CreateCourseAndCategory(ctx context.Context, argsProducts ProductParams, argsColor ColorParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateProduct(ctx, db.CreateProductParams{
			ID:          argsProducts.ID,
			Name:        argsProducts.Name,
			Description: argsProducts.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateColor(ctx, db.CreateColorParams{
			ID:          argsColor.ID,
			Name:        argsColor.Name,
			ProductID:   argsColor.ProductID,
			Description: argsColor.Description,
			Price:       argsColor.Price,
			Pricefinal:  argsColor.Pricefinal,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	dbNewConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/teste")
	if err != nil {
		panic(err)
	}
	defer dbNewConn.Close()

	queries := db.New(dbNewConn)

	// courseArgs := ProductParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Copo",
	// 	Description: sql.NullString{String: "Copo desc", Valid: true},
	// }

	// categoryArgs := ColorParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Azul",
	// 	Description: sql.NullString{String: "Azul desc", Valid: true},
	// 	ProductID:   courseArgs.ID,
	// 	Price:       10.00,
	// 	Pricefinal:  15.00,
	// }

	// colorsDB := NewColorsDB(dbNewConn)
	// err = colorsDB.CreateCourseAndCategory(ctx, courseArgs, categoryArgs)
	// if err != nil {
	// 	panic(err)
	// }

	colors, err := queries.ListColors(ctx)
	if err != nil {
		panic(err)
	}
	for _, color := range colors {
		println(color.ID, color.Name, color.Description.String, color.ProductID, color.Price, color.Pricefinal, color.CategoryName)
	}

}
