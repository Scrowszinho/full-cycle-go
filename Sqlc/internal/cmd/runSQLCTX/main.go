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

func main() {
	ctx := context.Background()
	dbNewConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/teste")
	if err != nil {
		panic(err)
	}
	defer dbNewConn.Close()

	queries := db.New(dbNewConn)

}
