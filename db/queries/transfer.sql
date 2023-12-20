-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateTransfer :one
INSERT INTO transfers (
    to_account_id,
    from_account_id,
    amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: ListAccountTransfers :many
SELECT * FROM transfers
WHERE to_account_id = $3 OR from_account_id = $3
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;


