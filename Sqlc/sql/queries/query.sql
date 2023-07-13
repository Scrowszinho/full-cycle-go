-- name: ListProducts :many
SELECT * FROM products;

-- name: GetColors :many
SELECT * FROM colors
WHERE id = ?;

-- name: CreateProduct :exec
INSERT INTO products (id, name, description) 
VALUES (?, ?, ?);

-- name: UpdateProduct :exec
UPDATE products SET name = ?, description = ?
WHERE id = ?;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = ?;
