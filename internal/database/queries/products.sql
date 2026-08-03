-- name: CreateProduct :one
INSERT INTO products
(seller_id, category_id, name, description, price, stock_quantity, images)
VALUES
($1, $2, $3, $4, $5, $6, $7)
RETURNING *; 

-- name: GetProductById :one
SELECT * FROM products WHERE id = $1;

-- name: GetProductsBySellerId :many
SELECT * FROM products WHERE seller_id = $1;

-- name: GetProductsByCategoryId :many
SELECT * FROM products WHERE category_id = $1;

-- name: UpdateProduct :one
UPDATE products
SET category_id = $2, name = $3, description = $4, images = $5
WHERE id = $1
RETURNING *;

-- name: UpdateProductPrice :one
UPDATE products
SET price = $2
WHERE id = $1
RETURNING *;

-- name: UpdateProductStockQuantity :one
UPDATE products
SET stock_quantity = $2
WHERE id = $1
RETURNING *;