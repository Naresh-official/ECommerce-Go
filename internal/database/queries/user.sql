-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users
(name, email, password_hash,phone,role)
VALUES
($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetUserByEmailAndRole :one
SELECT * FROM users WHERE email = $1 AND role = $2;

-- name: UpdateUser :one
UPDATE users
SET name = $2, email = $3, phone = $4, role = $5
WHERE id = $1
RETURNING *;

-- name: UpdateUserPassword :one
UPDATE users
SET password_hash = $2
WHERE id = $1
RETURNING *;

-- name: UpdateRefreshToken :exec
UPDATE users
SET refresh_token = $2
WHERE id = $1;