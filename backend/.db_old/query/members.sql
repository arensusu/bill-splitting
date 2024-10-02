-- name: CreateMember :one
INSERT INTO members (group_id, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetMembership :one
SELECT *
FROM members
WHERE group_id = $1 AND user_id = $2;

-- name: GetMember :one
SELECT *
FROM members
WHERE id = $1;

-- name: ListMembersOfGroup :many
SELECT members.id as id, username
FROM (SELECT * FROM members WHERE group_id = $1) AS members, users
WHERE members.user_id = users.id;

-- name: DeleteMember :exec
DELETE FROM members
WHERE group_id = $1 AND user_id = $2;
