-- name: CreateUser :one
INSERT INTO users (
    username,
    full_name,
    email,
    hashed_password
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE username = $1 LIMIT 1 
FOR NO KEY UPDATE;

-- name: UpdateUserEmail :one
UPDATE users
SET email = $2
WHERE username = $1
RETURNING *;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username
LIMIT $1
OFFSET $2;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;

