-- name: CreateGroupMember :one
INSERT INTO group_members (group_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetGroupMember :one
SELECT *
FROM group_members
WHERE group_id = $1 AND user_id = $2;

-- name: ListGroupMembers :many
SELECT users.id, users.username
FROM (SELECT * FROM group_members WHERE group_id = $1) AS group_members, users
WHERE group_members.user_id = users.id;

-- name: DeleteGroupMember :exec
DELETE FROM group_members
WHERE group_id = $1 AND user_id = $2;
