// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: user_expenses.sql

package db

import (
	"context"
)

const createUserExpense = `-- name: CreateUserExpense :one
INSERT INTO user_expenses (expense_id, user_id, share)
VALUES ($1, $2, $3)
RETURNING expense_id, user_id, share
`

type CreateUserExpenseParams struct {
	ExpenseID int64 `json:"expense_id"`
	UserID    int64 `json:"user_id"`
	Share     int64 `json:"share"`
}

func (q *Queries) CreateUserExpense(ctx context.Context, arg CreateUserExpenseParams) (UserExpense, error) {
	row := q.db.QueryRowContext(ctx, createUserExpense, arg.ExpenseID, arg.UserID, arg.Share)
	var i UserExpense
	err := row.Scan(&i.ExpenseID, &i.UserID, &i.Share)
	return i, err
}

const deleteUserExpense = `-- name: DeleteUserExpense :exec
DELETE FROM user_expenses
WHERE expense_id = $1 and user_id = $2
`

type DeleteUserExpenseParams struct {
	ExpenseID int64 `json:"expense_id"`
	UserID    int64 `json:"user_id"`
}

func (q *Queries) DeleteUserExpense(ctx context.Context, arg DeleteUserExpenseParams) error {
	_, err := q.db.ExecContext(ctx, deleteUserExpense, arg.ExpenseID, arg.UserID)
	return err
}

const getUserExpense = `-- name: GetUserExpense :one
SELECT expense_id, user_id, share
FROM user_expenses
WHERE expense_id = $1 and user_id = $2
`

type GetUserExpenseParams struct {
	ExpenseID int64 `json:"expense_id"`
	UserID    int64 `json:"user_id"`
}

func (q *Queries) GetUserExpense(ctx context.Context, arg GetUserExpenseParams) (UserExpense, error) {
	row := q.db.QueryRowContext(ctx, getUserExpense, arg.ExpenseID, arg.UserID)
	var i UserExpense
	err := row.Scan(&i.ExpenseID, &i.UserID, &i.Share)
	return i, err
}

const listUserExpenses = `-- name: ListUserExpenses :many
SELECT expense_id, user_id, share
FROM user_expenses
WHERE expense_id = $1
`

func (q *Queries) ListUserExpenses(ctx context.Context, expenseID int64) ([]UserExpense, error) {
	rows, err := q.db.QueryContext(ctx, listUserExpenses, expenseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserExpense{}
	for rows.Next() {
		var i UserExpense
		if err := rows.Scan(&i.ExpenseID, &i.UserID, &i.Share); err != nil {
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

const updateUserExpense = `-- name: UpdateUserExpense :one
UPDATE user_expenses
SET share = $3
WHERE expense_id = $1 and user_id = $2
RETURNING expense_id, user_id, share
`

type UpdateUserExpenseParams struct {
	ExpenseID int64 `json:"expense_id"`
	UserID    int64 `json:"user_id"`
	Share     int64 `json:"share"`
}

func (q *Queries) UpdateUserExpense(ctx context.Context, arg UpdateUserExpenseParams) (UserExpense, error) {
	row := q.db.QueryRowContext(ctx, updateUserExpense, arg.ExpenseID, arg.UserID, arg.Share)
	var i UserExpense
	err := row.Scan(&i.ExpenseID, &i.UserID, &i.Share)
	return i, err
}
