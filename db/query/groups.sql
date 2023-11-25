-- name: CreateGroup :one
INSERT INTO groups (name)
VALUES ($1)
RETURNING *;

-- name: GetGroup :one
SELECT *
FROM groups
WHERE id = $1
LIMIT 1;

-- name: UpdateGroup :one
UPDATE groups
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteGroup :exec
DELETE FROM groups
WHERE id = $1;