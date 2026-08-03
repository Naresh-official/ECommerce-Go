-- name: GetSellerById :one
SELECT * FROM sellers WHERE id = $1;

-- name: GetSellerByStoreName :one
SELECT * FROM sellers WHERE store_name = $1;

-- name: GetSellerByOwnerId :one
SELECT * FROM sellers WHERE owner_id = $1;

-- name: CreateSeller :one
INSERT INTO sellers
(owner_id, store_name, description)
VALUES
($1, $2, $3)
RETURNING *;

-- name: UpdateSeller :one
UPDATE sellers
SET store_name = $2, description = $3
WHERE id = $1
RETURNING *;

-- name: MarkSellerAsVerified :one
UPDATE sellers
SET is_verified = TRUE
WHERE id = $1
RETURNING *;

-- name: MarkSellerAsUnverified :one
UPDATE sellers
SET is_verified = FALSE
WHERE id = $1
RETURNING *;

-- name: MarkSellerAsActive :one
UPDATE sellers
SET is_active = TRUE
WHERE id = $1
RETURNING *;

-- name: MarkSellerAsInactive :one
UPDATE sellers
SET is_active = FALSE
WHERE id = $1
RETURNING *;