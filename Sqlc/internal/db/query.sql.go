// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :exec
INSERT INTO products (id, name, description) 
VALUES (?, ?, ?)
`

type CreateProductParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) error {
	_, err := q.db.ExecContext(ctx, createProduct, arg.ID, arg.Name, arg.Description)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products WHERE id = ?
`

func (q *Queries) DeleteProduct(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getColors = `-- name: GetColors :many
SELECT id, product_id, name, description, price, pricefinal FROM colors
WHERE id = ?
`

func (q *Queries) GetColors(ctx context.Context, id string) ([]Color, error) {
	rows, err := q.db.QueryContext(ctx, getColors, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Color
	for rows.Next() {
		var i Color
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Pricefinal,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProducts = `-- name: ListProducts :many
SELECT id, name, description FROM products
`

func (q *Queries) ListProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products SET name = ?, description = ?
WHERE id = ?
`

type UpdateProductParams struct {
	Name        string
	Description sql.NullString
	ID          string
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.ExecContext(ctx, updateProduct, arg.Name, arg.Description, arg.ID)
	return err
}