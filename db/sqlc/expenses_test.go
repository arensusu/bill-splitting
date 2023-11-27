package db_test

import (
	"bill-splitting/db/helper"
	db "bill-splitting/db/sqlc"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomExpense(t *testing.T) db.Expense {
	group := createRandomGroup(t)
	user := createRandomUser(t)
	param := db.CreateExpenseParams{
		GroupID: group.ID,
		PayerID: user.ID,
		Amount:  helper.RandomInt64(1, 1000),
		Date:    helper.RandomDate(),
	}
	expense, err := testStore.Queries.CreateExpense(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, expense)

	require.NotZero(t, expense.ID)
	require.Equal(t, group.ID, expense.GroupID)
	require.Equal(t, user.ID, expense.PayerID)
	require.Equal(t, param.Amount, expense.Amount)
	require.WithinDuration(t, param.Date, expense.Date, time.Second)

	return expense
}

func TestCreateExpense(t *testing.T) {
	createRandomExpense(t)
}

func TestGetExpense(t *testing.T) {
	expense1 := createRandomExpense(t)

	expense2, err := testStore.Queries.GetExpense(context.Background(), expense1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, expense2)

	require.Equal(t, expense1.ID, expense2.ID)
	require.Equal(t, expense1.GroupID, expense2.GroupID)
	require.Equal(t, expense1.PayerID, expense2.PayerID)
	require.Equal(t, expense1.Amount, expense2.Amount)
	require.WithinDuration(t, expense1.Date, expense2.Date, time.Second)
}

func TestUpdateExpense(t *testing.T) {
	expense := createRandomExpense(t)

	newAmount := helper.RandomInt64(1, 1000)
	param := db.UpdateExpenseParams{
		ID:      expense.ID,
		GroupID: expense.GroupID,
		PayerID: expense.PayerID,
		Amount:  newAmount,
		Date:    expense.Date,
	}
	newExpense, err := testStore.Queries.UpdateExpense(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, newExpense)

	require.Equal(t, expense.ID, newExpense.ID)
	require.Equal(t, expense.GroupID, newExpense.GroupID)
	require.Equal(t, expense.PayerID, newExpense.PayerID)
	require.Equal(t, newAmount, newExpense.Amount)
	require.WithinDuration(t, expense.Date, newExpense.Date, time.Second)
}

func TestDeleteExpense(t *testing.T) {
	expense1 := createRandomExpense(t)

	err := testStore.Queries.DeleteExpense(context.Background(), expense1.ID)

	require.NoError(t, err)

	expense2, err := testStore.Queries.GetExpense(context.Background(), expense1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, expense2)
}
