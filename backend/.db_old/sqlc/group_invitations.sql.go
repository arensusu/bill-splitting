// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: group_invitations.sql

package db

import (
	"context"
)

const createGroupInvitation = `-- name: CreateGroupInvitation :one
INSERT INTO group_invitations (code, group_id)
VALUES ($1, $2)
RETURNING code, group_id
`

type CreateGroupInvitationParams struct {
	Code    string `json:"code"`
	GroupID int32  `json:"group_id"`
}

func (q *Queries) CreateGroupInvitation(ctx context.Context, arg CreateGroupInvitationParams) (GroupInvitation, error) {
	row := q.db.QueryRowContext(ctx, createGroupInvitation, arg.Code, arg.GroupID)
	var i GroupInvitation
	err := row.Scan(&i.Code, &i.GroupID)
	return i, err
}

const deleteGroupInvitation = `-- name: DeleteGroupInvitation :exec
DELETE FROM group_invitations
WHERE code = $1
`

func (q *Queries) DeleteGroupInvitation(ctx context.Context, code string) error {
	_, err := q.db.ExecContext(ctx, deleteGroupInvitation, code)
	return err
}

const getGroupInvitation = `-- name: GetGroupInvitation :one
SELECT code, group_id
FROM group_invitations
WHERE code = $1
`

func (q *Queries) GetGroupInvitation(ctx context.Context, code string) (GroupInvitation, error) {
	row := q.db.QueryRowContext(ctx, getGroupInvitation, code)
	var i GroupInvitation
	err := row.Scan(&i.Code, &i.GroupID)
	return i, err
}
