-- name: CreateGroup :one
INSERT INTO groups (name)
VALUES ($1)
RETURNING *;

-- name: GetGroup :one
SELECT *
FROM groups
WHERE name = $1
LIMIT 1;