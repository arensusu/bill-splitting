// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: expenses.sql

package db

import (
	"context"
	"time"
)

const createExpense = `-- name: CreateExpense :one
INSERT INTO expenses (group_id, payer_id, amount, description, date)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, group_id, payer_id, amount, description, date
`

type CreateExpenseParams struct {
	GroupID     int64     `json:"group_id"`
	PayerID     int64     `json:"payer_id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, createExpense,
		arg.GroupID,
		arg.PayerID,
		arg.Amount,
		arg.Description,
		arg.Date,
	)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.PayerID,
		&i.Amount,
		&i.Description,
		&i.Date,
	)
	return i, err
}

const deleteExpense = `-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1
`

func (q *Queries) DeleteExpense(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteExpense, id)
	return err
}

const getExpense = `-- name: GetExpense :one
SELECT id, group_id, payer_id, amount, description, date
FROM expenses
WHERE id = $1
`

func (q *Queries) GetExpense(ctx context.Context, id int64) (Expense, error) {
	row := q.db.QueryRowContext(ctx, getExpense, id)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.PayerID,
		&i.Amount,
		&i.Description,
		&i.Date,
	)
	return i, err
}

const listGroupExpenses = `-- name: ListGroupExpenses :many
SELECT id, group_id, payer_id, amount, description, date
FROM expenses
WHERE group_id = $1
`

func (q *Queries) ListGroupExpenses(ctx context.Context, groupID int64) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, listGroupExpenses, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.GroupID,
			&i.PayerID,
			&i.Amount,
			&i.Description,
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

const updateExpense = `-- name: UpdateExpense :one
UPDATE expenses
SET group_id = $2, payer_id = $3, amount = $4, description = $5, date = $6
WHERE id = $1
RETURNING id, group_id, payer_id, amount, description, date
`

type UpdateExpenseParams struct {
	ID          int64     `json:"id"`
	GroupID     int64     `json:"group_id"`
	PayerID     int64     `json:"payer_id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (q *Queries) UpdateExpense(ctx context.Context, arg UpdateExpenseParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, updateExpense,
		arg.ID,
		arg.GroupID,
		arg.PayerID,
		arg.Amount,
		arg.Description,
		arg.Date,
	)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.GroupID,
		&i.PayerID,
		&i.Amount,
		&i.Description,
		&i.Date,
	)
	return i, err
}
