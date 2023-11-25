-- name: CreateGroupMember :one
INSERT INTO group_members (group_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetGroupMember :one
SELECT *
FROM group_members
WHERE group_id = $1 AND user_id = $2;

-- name: DeleteGroupMember :exec
DELETE FROM group_members
WHERE group_id = $1 AND user_id = $2;
