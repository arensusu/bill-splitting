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
INSERT INTO settlements (id, payer_id, payee_id, amount, date)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, payer_id, payee_id, amount, date
`

type CreateSettlementParams struct {
	ID      int32     `json:"id"`
	PayerID int32     `json:"payer_id"`
	PayeeID int32     `json:"payee_id"`
	Amount  string    `json:"amount"`
	Date    time.Time `json:"date"`
}

func (q *Queries) CreateSettlement(ctx context.Context, arg CreateSettlementParams) (Settlement, error) {
	row := q.db.QueryRowContext(ctx, createSettlement,
		arg.ID,
		arg.PayerID,
		arg.PayeeID,
		arg.Amount,
		arg.Date,
	)
	var i Settlement
	err := row.Scan(
		&i.ID,
		&i.PayerID,
		&i.PayeeID,
		&i.Amount,
		&i.Date,
	)
	return i, err
}

const deleteSettlement = `-- name: DeleteSettlement :exec
DELETE FROM settlements
WHERE id = $1
`

func (q *Queries) DeleteSettlement(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteSettlement, id)
	return err
}

const getSettlement = `-- name: GetSettlement :one
SELECT id, payer_id, payee_id, amount, date
FROM settlements
WHERE id = $1
`

func (q *Queries) GetSettlement(ctx context.Context, id int32) (Settlement, error) {
	row := q.db.QueryRowContext(ctx, getSettlement, id)
	var i Settlement
	err := row.Scan(
		&i.ID,
		&i.PayerID,
		&i.PayeeID,
		&i.Amount,
		&i.Date,
	)
	return i, err
}

const updateSettlement = `-- name: UpdateSettlement :one
UPDATE settlements
SET payer_id = $2, payee_id = $3, amount = $4, date = $5
WHERE id = $1
RETURNING id, payer_id, payee_id, amount, date
`

type UpdateSettlementParams struct {
	ID      int32     `json:"id"`
	PayerID int32     `json:"payer_id"`
	PayeeID int32     `json:"payee_id"`
	Amount  string    `json:"amount"`
	Date    time.Time `json:"date"`
}

func (q *Queries) UpdateSettlement(ctx context.Context, arg UpdateSettlementParams) (Settlement, error) {
	row := q.db.QueryRowContext(ctx, updateSettlement,
		arg.ID,
		arg.PayerID,
		arg.PayeeID,
		arg.Amount,
		arg.Date,
	)
	var i Settlement
	err := row.Scan(
		&i.ID,
		&i.PayerID,
		&i.PayeeID,
		&i.Amount,
		&i.Date,
	)
	return i, err
}