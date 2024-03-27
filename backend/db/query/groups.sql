-- name: CreateGroup :one
INSERT INTO groups (name)
VALUES ($1)
RETURNING *;

-- name: GetGroup :one
SELECT *
FROM groups
WHERE id = $1
LIMIT 1;

-- name: ListGroups :many
SELECT id, name
FROM groups, (SELECT group_id FROM members WHERE user_id = $1) AS group_members
WHERE groups.id = group_members.group_id;

-- name: UpdateGroup :one
UPDATE groups
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteGroup :exec
DELETE FROM groups
WHERE id = $1;