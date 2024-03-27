package db

import (
	"bill-splitting/helper"
	"context"
	"database/sql"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomExpense(t *testing.T) Expense {
	group := createRandomGroup(t)
	member := createRandomMember(t, group.ID)
	param := CreateExpenseParams{
		MemberID: member.ID,
		Amount:   strconv.FormatInt(helper.RandomInt64(1, 1000), 10),
		Date:     helper.RandomDate(),
	}
	expense, err := testStore.CreateExpense(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, expense)

	require.NotZero(t, expense.ID)
	require.Equal(t, member.ID, expense.MemberID)
	require.Equal(t, param.Amount, expense.Amount)
	require.Equal(t, param.Date.Year(), expense.Date.Year())
	require.Equal(t, param.Date.Month(), expense.Date.Month())
	require.Equal(t, param.Date.Day(), expense.Date.Day())

	return expense
}

func TestCreateExpense(t *testing.T) {
	createRandomExpense(t)
}

func TestGetExpense(t *testing.T) {
	expense1 := createRandomExpense(t)

	expense2, err := testStore.GetExpense(context.Background(), expense1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, expense2)

	require.Equal(t, expense1.MemberID, expense2.MemberID)
	require.Equal(t, expense1.Amount, expense2.Amount)
	require.Equal(t, expense1.Date.Year(), expense2.Date.Year())
	require.Equal(t, expense1.Date.Month(), expense2.Date.Month())
	require.Equal(t, expense1.Date.Day(), expense2.Date.Day())
}

func TestUpdateExpense(t *testing.T) {
	expense := createRandomExpense(t)

	newAmount := strconv.FormatInt(helper.RandomInt64(1, 1000), 10)
	param := UpdateExpenseParams{
		ID:       expense.ID,
		MemberID: expense.MemberID,
		Amount:   newAmount,
		Date:     expense.Date,
	}
	newExpense, err := testStore.UpdateExpense(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, newExpense)

	require.Equal(t, expense.ID, newExpense.ID)
	require.Equal(t, expense.MemberID, newExpense.MemberID)
	require.Equal(t, newAmount, newExpense.Amount)
	require.Equal(t, expense.Date.Year(), newExpense.Date.Year())
	require.Equal(t, expense.Date.Month(), newExpense.Date.Month())
	require.Equal(t, expense.Date.Day(), newExpense.Date.Day())
}

func TestDeleteExpense(t *testing.T) {
	expense1 := createRandomExpense(t)

	err := testStore.DeleteExpense(context.Background(), expense1.ID)

	require.NoError(t, err)

	expense2, err := testStore.GetExpense(context.Background(), expense1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, expense2)
}

func TestListExpenses(t *testing.T) {
	var lastExpense Expense
	for i := 0; i < 10; i++ {
		lastExpense = createRandomExpense(t)
	}

	lastMember, err := testStore.GetMember(context.Background(), lastExpense.MemberID)
	require.NoError(t, err)
	expenses, err := testStore.ListExpenses(context.Background(), lastMember.GroupID)

	require.NoError(t, err)
	require.NotEmpty(t, expenses)

	for _, expense := range expenses {
		require.NotEmpty(t, expense)

		member, err := testStore.GetMember(context.Background(), expense.MemberID)
		require.NoError(t, err)
		require.Equal(t, member.GroupID, lastMember.GroupID)
	}
}

func TestListNonSettledExpenses(t *testing.T) {
	var lastExpense Expense
	for i := 0; i < 10; i++ {
		lastExpense = createRandomExpense(t)
	}

	lastMember, err := testStore.GetMember(context.Background(), lastExpense.MemberID)
	require.NoError(t, err)
	expenses1, err := testStore.ListExpenses(context.Background(), lastMember.GroupID)

	require.NoError(t, err)
	require.NotEmpty(t, expenses1)

	for _, expense := range expenses1 {
		require.NotEmpty(t, expense)

		member, err := testStore.GetMember(context.Background(), expense.MemberID)
		require.NoError(t, err)
		require.Equal(t, member.GroupID, lastMember.GroupID)
	}

	_, err = testStore.UpdateExpense(context.Background(), UpdateExpenseParams{
		ID:        lastExpense.ID,
		MemberID:  lastExpense.MemberID,
		Amount:    lastExpense.Amount,
		Date:      lastExpense.Date,
		IsSettled: true,
	})
	require.NoError(t, err)

	expenses2, err := testStore.ListNonSettledExpenses(context.Background(), lastMember.GroupID)

	require.NoError(t, err)
	require.Equal(t, len(expenses1)-1, len(expenses2))
}
