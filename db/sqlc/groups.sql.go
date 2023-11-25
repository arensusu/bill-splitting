// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: groups.sql

package db

import (
	"context"
)

const createGroup = `-- name: CreateGroup :one
INSERT INTO groups (name)
VALUES ($1)
RETURNING id, name, created_at
`

func (q *Queries) CreateGroup(ctx context.Context, name string) (Group, error) {
	row := q.db.QueryRowContext(ctx, createGroup, name)
	var i Group
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const deleteGroup = `-- name: DeleteGroup :exec
DELETE FROM groups
WHERE id = $1
`

func (q *Queries) DeleteGroup(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteGroup, id)
	return err
}

const getGroup = `-- name: GetGroup :one
SELECT id, name, created_at
FROM groups
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetGroup(ctx context.Context, id int64) (Group, error) {
	row := q.db.QueryRowContext(ctx, getGroup, id)
	var i Group
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const updateGroup = `-- name: UpdateGroup :one
UPDATE groups
SET name = $2
WHERE id = $1
RETURNING id, name, created_at
`

type UpdateGroupParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateGroup(ctx context.Context, arg UpdateGroupParams) (Group, error) {
	row := q.db.QueryRowContext(ctx, updateGroup, arg.ID, arg.Name)
	var i Group
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}
