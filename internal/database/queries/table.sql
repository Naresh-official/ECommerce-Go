-- name: GetAllAddressesOfUser :many
SELECT * FROM addresses WHERE user_id = $1;

-- name: GetAddressById :one
SELECT * FROM addresses WHERE id = $1;

-- name: CreateAddress :one
INSERT INTO addresses (
    user_id,
    address_line1,
    address_line2,
    city,
    state,
    postal_code,
    country,
    phone,
    is_default
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    NOT EXISTS (
        SELECT 1
        FROM addresses
        WHERE user_id = $1
    )
)
RETURNING *;

-- name: UpdateAddress :one
UPDATE addresses
SET address_line1 = $2, address_line2 = $3, city = $4, state = $5, postal_code = $6, country = $7, phone = $8, is_default = $9
WHERE id = $1
RETURNING *;

-- name: DeleteAddress :exec
DELETE FROM addresses WHERE id = $1;

-- name: SetDefaultAddress :exec
UPDATE addresses
SET is_default = CASE WHEN addresses.id = $1 THEN true ELSE false END
WHERE user_id = (SELECT user_id FROM addresses WHERE addresses.id = $1);