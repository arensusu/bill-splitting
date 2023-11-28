// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: settlements.sql

package db

import (
	"context"
	"time"
)

const createSettlement = `-- name: CreateSettlement :one
INSERT INTO settlements (group_id, payer_id, payee_id, amount, date)
VALUES ($1, $2, $3, $4, $5)
RETURNING group_id, payer_id, payee_id, amount, date
`

type CreateSettlementParams struct {
	GroupID int64     `json:"group_id"`
	PayerID int64     `json:"payer_id"`
	PayeeID int64     `json:"payee_id"`
	Amount  int64     `json:"amount"`
	Date    time.Time `json:"date"`
}

func (q *Queries) CreateSettlement(ctx context.Context, arg CreateSettlementParams) (Settlement, error) {
	row := q.db.QueryRowContext(ctx, createSettlement,
		arg.GroupID,
		arg.PayerID,
		arg.PayeeID,
		arg.Amount,
		arg.Date,
	)
	var i Settlement
	err := row.Scan(
		&i.GroupID,
		&i.PayerID,
		&i.PayeeID,
		&i.Amount,
		&i.Date,
	)
	return i, err
}

const deleteSettlement = `-- name: DeleteSettlement :exec
DELETE FROM settlements
WHERE group_id = $1 AND payer_id = $2 AND payee_id = $3
`

type DeleteSettlementParams struct {
	GroupID int64 `json:"group_id"`
	PayerID int64 `json:"payer_id"`
	PayeeID int64 `json:"payee_id"`
}

func (q *Queries) DeleteSettlement(ctx context.Context, arg DeleteSettlementParams) error {
	_, err := q.db.ExecContext(ctx, deleteSettlement, arg.GroupID, arg.PayerID, arg.PayeeID)
	return err
}

const getSettlement = `-- name: GetSettlement :one
SELECT group_id, payer_id, payee_id, amount, date
FROM settlements
WHERE group_id = $1 AND payer_id = $2 AND payee_id = $3
`

type GetSettlementParams struct {
	GroupID int64 `json:"group_id"`
	PayerID int64 `json:"payer_id"`
	PayeeID int64 `json:"payee_id"`
}

func (q *Queries) GetSettlement(ctx context.Context, arg GetSettlementParams) (Settlement, error) {
	row := q.db.QueryRowContext(ctx, getSettlement, arg.GroupID, arg.PayerID, arg.PayeeID)
	var i Settlement
	err := row.Scan(
		&i.GroupID,
		&i.PayerID,
		&i.PayeeID,
		&i.Amount,
		&i.Date,
	)
	return i, err
}

const listGroupSettlements = `-- name: ListGroupSettlements :many
SELECT group_id, payer_id, payee_id, amount, date
FROM settlements
WHERE group_id = $1
`

func (q *Queries) ListGroupSettlements(ctx context.Context, groupID int64) ([]Settlement, error) {
	rows, err := q.db.QueryContext(ctx, listGroupSettlements, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Settlement
	for rows.Next() {
		var i Settlement
		if err := rows.Scan(
			&i.GroupID,
			&i.PayerID,
			&i.PayeeID,
			&i.Amount,
			&i.Date,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSettlement = `-- name: UpdateSettlement :one
UPDATE settlements
SET amount = $4, date = $5
WHERE group_id = $1 AND payer_id = $2 AND payee_id = $3
RETURNING group_id, payer_id, payee_id, amount, date
`

type UpdateSettlementParams struct {
	GroupID int64     `json:"group_id"`
	PayerID int64     `json:"payer_id"`
	PayeeID int64     `json:"payee_id"`
	Amount  int64     `json:"amount"`
	Date    time.Time `json:"date"`
}

func (q *Queries) UpdateSettlement(ctx context.Context, arg UpdateSettlementParams) (Settlement, error) {
	row := q.db.QueryRowContext(ctx, updateSettlement,
		arg.GroupID,
		arg.PayerID,
		arg.PayeeID,
		arg.Amount,
		arg.Date,
	)
	var i Settlement
	err := row.Scan(
		&i.GroupID,
		&i.PayerID,
		&i.PayeeID,
		&i.Amount,
		&i.Date,
	)
	return i, err
}
