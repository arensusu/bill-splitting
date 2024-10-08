// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: expenses.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createExpense = `-- name: CreateExpense :one
INSERT INTO expenses (member_id, origin_currency, origin_amount, amount, description, date, category)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, member_id, amount, description, date, is_settled, category, origin_amount, origin_currency
`

type CreateExpenseParams struct {
	MemberID       int32          `json:"member_id"`
	OriginCurrency sql.NullString `json:"origin_currency"`
	OriginAmount   sql.NullString `json:"origin_amount"`
	Amount         string         `json:"amount"`
	Description    string         `json:"description"`
	Date           time.Time      `json:"date"`
	Category       sql.NullString `json:"category"`
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, createExpense,
		arg.MemberID,
		arg.OriginCurrency,
		arg.OriginAmount,
		arg.Amount,
		arg.Description,
		arg.Date,
		arg.Category,
	)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.MemberID,
		&i.Amount,
		&i.Description,
		&i.Date,
		&i.IsSettled,
		&i.Category,
		&i.OriginAmount,
		&i.OriginCurrency,
	)
	return i, err
}

const deleteExpense = `-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1
`

func (q *Queries) DeleteExpense(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteExpense, id)
	return err
}

const getExpense = `-- name: GetExpense :one
SELECT member_id, amount, description, date, origin_currency, origin_amount
FROM expenses
WHERE id = $1
`

type GetExpenseRow struct {
	MemberID       int32          `json:"member_id"`
	Amount         string         `json:"amount"`
	Description    string         `json:"description"`
	Date           time.Time      `json:"date"`
	OriginCurrency sql.NullString `json:"origin_currency"`
	OriginAmount   sql.NullString `json:"origin_amount"`
}

func (q *Queries) GetExpense(ctx context.Context, id int32) (GetExpenseRow, error) {
	row := q.db.QueryRowContext(ctx, getExpense, id)
	var i GetExpenseRow
	err := row.Scan(
		&i.MemberID,
		&i.Amount,
		&i.Description,
		&i.Date,
		&i.OriginCurrency,
		&i.OriginAmount,
	)
	return i, err
}

const listExpenses = `-- name: ListExpenses :many
SELECT expenses.id, member_id, amount, description, date, is_settled, category, origin_amount, origin_currency, members.id
FROM expenses, (SELECT id FROM members WHERE group_id = $1) AS members
WHERE expenses.member_id = members.id
`

type ListExpensesRow struct {
	ID             int32          `json:"id"`
	MemberID       int32          `json:"member_id"`
	Amount         string         `json:"amount"`
	Description    string         `json:"description"`
	Date           time.Time      `json:"date"`
	IsSettled      bool           `json:"is_settled"`
	Category       sql.NullString `json:"category"`
	OriginAmount   sql.NullString `json:"origin_amount"`
	OriginCurrency sql.NullString `json:"origin_currency"`
	ID_2           int32          `json:"id_2"`
}

func (q *Queries) ListExpenses(ctx context.Context, groupID int32) ([]ListExpensesRow, error) {
	rows, err := q.db.QueryContext(ctx, listExpenses, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListExpensesRow{}
	for rows.Next() {
		var i ListExpensesRow
		if err := rows.Scan(
			&i.ID,
			&i.MemberID,
			&i.Amount,
			&i.Description,
			&i.Date,
			&i.IsSettled,
			&i.Category,
			&i.OriginAmount,
			&i.OriginCurrency,
			&i.ID_2,
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

const listExpensesWithinDate = `-- name: ListExpensesWithinDate :many
SELECT expenses.id, member_id, amount, description, date, is_settled, category, origin_amount, origin_currency, members.id
FROM expenses, (SELECT id FROM members WHERE group_id = $1) AS members
WHERE expenses.member_id = members.id AND date BETWEEN $2 AND $3
`

type ListExpensesWithinDateParams struct {
	GroupID   int32     `json:"group_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type ListExpensesWithinDateRow struct {
	ID             int32          `json:"id"`
	MemberID       int32          `json:"member_id"`
	Amount         string         `json:"amount"`
	Description    string         `json:"description"`
	Date           time.Time      `json:"date"`
	IsSettled      bool           `json:"is_settled"`
	Category       sql.NullString `json:"category"`
	OriginAmount   sql.NullString `json:"origin_amount"`
	OriginCurrency sql.NullString `json:"origin_currency"`
	ID_2           int32          `json:"id_2"`
}

func (q *Queries) ListExpensesWithinDate(ctx context.Context, arg ListExpensesWithinDateParams) ([]ListExpensesWithinDateRow, error) {
	rows, err := q.db.QueryContext(ctx, listExpensesWithinDate, arg.GroupID, arg.StartTime, arg.EndTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListExpensesWithinDateRow{}
	for rows.Next() {
		var i ListExpensesWithinDateRow
		if err := rows.Scan(
			&i.ID,
			&i.MemberID,
			&i.Amount,
			&i.Description,
			&i.Date,
			&i.IsSettled,
			&i.Category,
			&i.OriginAmount,
			&i.OriginCurrency,
			&i.ID_2,
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

const listNonSettledExpenses = `-- name: ListNonSettledExpenses :many
SELECT expenses.id, member_id, amount, description, date, is_settled, category, origin_amount, origin_currency, members.id
FROM expenses, (SELECT id FROM members WHERE group_id = $1) AS members
WHERE expenses.member_id = members.id AND is_settled = false
`

type ListNonSettledExpensesRow struct {
	ID             int32          `json:"id"`
	MemberID       int32          `json:"member_id"`
	Amount         string         `json:"amount"`
	Description    string         `json:"description"`
	Date           time.Time      `json:"date"`
	IsSettled      bool           `json:"is_settled"`
	Category       sql.NullString `json:"category"`
	OriginAmount   sql.NullString `json:"origin_amount"`
	OriginCurrency sql.NullString `json:"origin_currency"`
	ID_2           int32          `json:"id_2"`
}

func (q *Queries) ListNonSettledExpenses(ctx context.Context, groupID int32) ([]ListNonSettledExpensesRow, error) {
	rows, err := q.db.QueryContext(ctx, listNonSettledExpenses, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListNonSettledExpensesRow{}
	for rows.Next() {
		var i ListNonSettledExpensesRow
		if err := rows.Scan(
			&i.ID,
			&i.MemberID,
			&i.Amount,
			&i.Description,
			&i.Date,
			&i.IsSettled,
			&i.Category,
			&i.OriginAmount,
			&i.OriginCurrency,
			&i.ID_2,
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

const summarizeExpensesWithinDate = `-- name: SummarizeExpensesWithinDate :many
SELECT category, SUM(amount::decimal)::decimal as total
FROM (SELECT id, member_id, amount, description, date, is_settled, category, origin_amount, origin_currency FROM expenses WHERE date BETWEEN $2 AND $3) as expenses, (SELECT id FROM members WHERE group_id = $1) AS members
WHERE expenses.member_id = members.id
GROUP BY category
`

type SummarizeExpensesWithinDateParams struct {
	GroupID   int32     `json:"group_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type SummarizeExpensesWithinDateRow struct {
	Category sql.NullString `json:"category"`
	Total    string         `json:"total"`
}

func (q *Queries) SummarizeExpensesWithinDate(ctx context.Context, arg SummarizeExpensesWithinDateParams) ([]SummarizeExpensesWithinDateRow, error) {
	rows, err := q.db.QueryContext(ctx, summarizeExpensesWithinDate, arg.GroupID, arg.StartTime, arg.EndTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SummarizeExpensesWithinDateRow{}
	for rows.Next() {
		var i SummarizeExpensesWithinDateRow
		if err := rows.Scan(&i.Category, &i.Total); err != nil {
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
SET member_id = $2, amount = $3, description = $4, date = $5, is_settled = $6
WHERE id = $1
RETURNING id, member_id, amount, description, date, is_settled, category, origin_amount, origin_currency
`

type UpdateExpenseParams struct {
	ID          int32     `json:"id"`
	MemberID    int32     `json:"member_id"`
	Amount      string    `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	IsSettled   bool      `json:"is_settled"`
}

func (q *Queries) UpdateExpense(ctx context.Context, arg UpdateExpenseParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, updateExpense,
		arg.ID,
		arg.MemberID,
		arg.Amount,
		arg.Description,
		arg.Date,
		arg.IsSettled,
	)
	var i Expense
	err := row.Scan(
		&i.ID,
		&i.MemberID,
		&i.Amount,
		&i.Description,
		&i.Date,
		&i.IsSettled,
		&i.Category,
		&i.OriginAmount,
		&i.OriginCurrency,
	)
	return i, err
}
