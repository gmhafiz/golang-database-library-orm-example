-- name: CreateUser :one
INSERT INTO users (first_name, middle_name, last_name, email, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: ListUsers :many
SELECT *
FROM users
ORDER BY last_name;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;