package db

import (
	"context"
	"fmt"
	"time"
)

type CreateExpenseTxParams struct {
	PayerID     int64     `json:"payerId"`
	GroupID     int64     `json:"groupId"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type CreateExpenseTxResult struct {
	Expense      Expense       `json:"expense"`
	UserExpenses []UserExpense `json:"userExpenses"`
}

func (s *SQLStore) CreateExpenseTx(ctx context.Context, arg CreateExpenseTxParams) (CreateExpenseTxResult, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return CreateExpenseTxResult{}, err
	}
	defer tx.Rollback()

	q := New(tx)
	group_members, err := q.ListGroupMembers(ctx, arg.GroupID)
	if err != nil {
		return CreateExpenseTxResult{}, fmt.Errorf("create expense tx: %w", err)
	}

	expense, err := q.CreateExpense(ctx, CreateExpenseParams{
		PayerID:     arg.PayerID,
		GroupID:     arg.GroupID,
		Amount:      arg.Amount,
		Description: arg.Description,
		Date:        arg.Date,
	})
	if err != nil {
		return CreateExpenseTxResult{}, fmt.Errorf("create expense tx: %w", err)
	}

	share := arg.Amount / int64(len(group_members))
	userExpenses := []UserExpense{}
	for _, member := range group_members {
		userExpense, err := q.CreateUserExpense(ctx, CreateUserExpenseParams{
			ExpenseID: expense.ID,
			UserID:    member.UserID,
			Share:     share,
		})
		if err != nil {
			return CreateExpenseTxResult{}, fmt.Errorf("create expense tx: %w", err)
		}
		userExpenses = append(userExpenses, userExpense)
	}

	if err = tx.Commit(); err != nil {
		return CreateExpenseTxResult{}, fmt.Errorf("create expense tx: %w", err)
	}

	return CreateExpenseTxResult{
		Expense:      expense,
		UserExpenses: userExpenses,
	}, nil
}
