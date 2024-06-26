// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: expenses.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createExpense = `-- name: CreateExpense :one
INSERT INTO expenses (member_id, amount, description, date, category)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, member_id, amount, description, date, is_settled, category
`

type CreateExpenseParams struct {
	MemberID    int32          `json:"member_id"`
	Amount      string         `json:"amount"`
	Description string         `json:"description"`
	Date        time.Time      `json:"date"`
	Category    sql.NullString `json:"category"`
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, createExpense,
		arg.MemberID,
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
SELECT member_id, amount, description, date
FROM expenses
WHERE id = $1
`

type GetExpenseRow struct {
	MemberID    int32     `json:"member_id"`
	Amount      string    `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (q *Queries) GetExpense(ctx context.Context, id int32) (GetExpenseRow, error) {
	row := q.db.QueryRowContext(ctx, getExpense, id)
	var i GetExpenseRow
	err := row.Scan(
		&i.MemberID,
		&i.Amount,
		&i.Description,
		&i.Date,
	)
	return i, err
}

const listExpenses = `-- name: ListExpenses :many
SELECT expenses.id, member_id, amount, description, date, is_settled, category, members.id
FROM expenses, (SELECT id FROM members WHERE group_id = $1) AS members
WHERE expenses.member_id = members.id
`

type ListExpensesRow struct {
	ID          int32          `json:"id"`
	MemberID    int32          `json:"member_id"`
	Amount      string         `json:"amount"`
	Description string         `json:"description"`
	Date        time.Time      `json:"date"`
	IsSettled   bool           `json:"is_settled"`
	Category    sql.NullString `json:"category"`
	ID_2        int32          `json:"id_2"`
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
SELECT expenses.id, member_id, amount, description, date, is_settled, category, members.id
FROM expenses, (SELECT id FROM members WHERE group_id = $1) AS members
WHERE expenses.member_id = members.id AND is_settled = false
`

type ListNonSettledExpensesRow struct {
	ID          int32          `json:"id"`
	MemberID    int32          `json:"member_id"`
	Amount      string         `json:"amount"`
	Description string         `json:"description"`
	Date        time.Time      `json:"date"`
	IsSettled   bool           `json:"is_settled"`
	Category    sql.NullString `json:"category"`
	ID_2        int32          `json:"id_2"`
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
SELECT category, SUM(amount) as total
FROM (SELECT id, member_id, amount, description, date, is_settled, category FROM expenses WHERE date BETWEEN $2 AND $3) as expenses, (SELECT id FROM members WHERE group_id = $1) AS members
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
	Total    int64          `json:"total"`
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
RETURNING id, member_id, amount, description, date, is_settled, category
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
	)
	return i, err
}
