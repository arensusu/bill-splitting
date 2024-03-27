-- name: CreateMember :one
INSERT INTO members (group_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetMember :one
SELECT *
FROM members
WHERE id = $1;

-- name: ListMembersOfGroup :many
SELECT *
FROM members
WHERE group_id = $1;

-- name: DeleteMember :exec
DELETE FROM members
WHERE group_id = $1 AND user_id = $2;
