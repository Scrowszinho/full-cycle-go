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

-- name: CreateColor :exec
INSERT INTO colors (id, name, product_id, description, price, priceFinal)
VALUES (?, ?, ?, ?, ?, ?);

-- name: ListColors :many
SELECT c.*, p.name as category_name 
FROM colors c JOIN products p ON c.product_id = p.id;
