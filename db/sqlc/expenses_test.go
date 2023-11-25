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

func createRandomExpense(t *testing.T, groupID int64, payerID int64) db.Expense {
	param := db.CreateExpenseParams{
		GroupID: groupID,
		PayerID: payerID,
		Amount:  helper.RandomInt64(1, 1000),
		Date:    helper.RandomDate(),
	}
	expense, err := testQueries.CreateExpense(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, expense)

	require.NotZero(t, expense.ID)
	require.Equal(t, groupID, expense.GroupID)
	require.Equal(t, payerID, expense.PayerID)
	require.Equal(t, param.Amount, expense.Amount)
	require.WithinDuration(t, param.Date, expense.Date, time.Second)

	return expense
}

func TestCreateExpense(t *testing.T) {
	group := createRandomGroup(t)
	user := createRandomUser(t)
	createRandomGroupMember(t, group.ID, user.ID)
}

func TestGetExpense(t *testing.T) {
	group := createRandomGroup(t)
	user := createRandomUser(t)
	expense1 := createRandomExpense(t, group.ID, user.ID)

	expense2, err := testQueries.GetExpense(context.Background(), expense1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, expense2)

	require.Equal(t, expense1.ID, expense2.ID)
	require.Equal(t, expense1.GroupID, expense2.GroupID)
	require.Equal(t, expense1.PayerID, expense2.PayerID)
	require.Equal(t, expense1.Amount, expense2.Amount)
	require.WithinDuration(t, expense1.Date, expense2.Date, time.Second)
}

func TestUpdateExpense(t *testing.T) {
	group := createRandomGroup(t)
	user := createRandomUser(t)
	expense := createRandomExpense(t, group.ID, user.ID)

	newAmount := helper.RandomInt64(1, 1000)
	param := db.UpdateExpenseParams{
		ID:      expense.ID,
		GroupID: group.ID,
		PayerID: user.ID,
		Amount:  newAmount,
		Date:    expense.Date,
	}
	newExpense, err := testQueries.UpdateExpense(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, newExpense)

	require.Equal(t, expense.ID, newExpense.ID)
	require.Equal(t, group.ID, newExpense.GroupID)
	require.Equal(t, user.ID, newExpense.PayerID)
	require.Equal(t, newAmount, newExpense.Amount)
	require.WithinDuration(t, expense.Date, newExpense.Date, time.Second)
}

func TestDeleteExpense(t *testing.T) {
	group := createRandomGroup(t)
	user := createRandomUser(t)
	expense1 := createRandomExpense(t, group.ID, user.ID)

	err := testQueries.DeleteExpense(context.Background(), expense1.ID)

	require.NoError(t, err)

	expense2, err := testQueries.GetExpense(context.Background(), expense1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, expense2)
}
