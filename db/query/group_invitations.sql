-- name: CreateGroupInvitation :one
INSERT INTO group_invitations (code, group_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetGroupInvitation :one
SELECT *
FROM group_invitations
WHERE code = $1;

-- name: DeleteGroupInvitation :exec
DELETE FROM group_invitations
WHERE code = $1;
