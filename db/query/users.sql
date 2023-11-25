-- name: CreateUser :one
INSERT INTO users (username, password)
VALUES ($1, $2)
RETURNING (id, username, created_at);

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET username = $2
WHERE id = $1
RETURNING (id, username, created_at);

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;