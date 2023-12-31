-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name, hashed_password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUsersByDetails :many
SELECT * FROM users WHERE name = $1 AND hashed_password = $2;

-- name: UpdateUser :one
UPDATE users
SET  updated_at = $2, name = $3, hashed_password = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;