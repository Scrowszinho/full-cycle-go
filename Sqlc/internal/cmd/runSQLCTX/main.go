package main

import (
	"context"
	"database/sql"
	"fmt"
	"teste/sqlc/internal/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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
			ID:          uuid.New().String(),
			Name:        argsProducts.Name,
			Description: argsProducts.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateColor(ctx, db.CreateColorParams{
			ID:          uuid.New().String(),
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

}
