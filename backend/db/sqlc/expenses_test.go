package db

import (
	"bill-splitting/helper"
	"context"
	"database/sql"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomExpense(t *testing.T, memberID int32, category string) Expense {
	param := CreateExpenseParams{
		MemberID: memberID,
		Amount:   strconv.FormatInt(helper.RandomInt64(1, 1000), 10),
		Date:     helper.RandomDate(),
		Category: sql.NullString{
			String: category,
			Valid:  true,
		},
	}
	expense, err := testStore.CreateExpense(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, expense)

	require.NotZero(t, expense.ID)
	require.Equal(t, memberID, expense.MemberID)
	require.Equal(t, param.Amount, expense.Amount)
	require.Equal(t, param.Category.String, expense.Category.String)
	require.Equal(t, param.Date.Year(), expense.Date.Year())
	require.Equal(t, param.Date.Month(), expense.Date.Month())
	require.Equal(t, param.Date.Day(), expense.Date.Day())

	return expense
}

func TestCreateExpense(t *testing.T) {
	member := createRandomMember(t, createRandomGroup(t).ID)
	category := helper.RandomString(5)
	createRandomExpense(t, member.ID, category)
}

func TestGetExpense(t *testing.T) {
	member := createRandomMember(t, createRandomGroup(t).ID)
	category := helper.RandomString(5)
	expense1 := createRandomExpense(t, member.ID, category)

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
	member := createRandomMember(t, createRandomGroup(t).ID)
	category := helper.RandomString(5)
	expense := createRandomExpense(t, member.ID, category)

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
	member := createRandomMember(t, createRandomGroup(t).ID)
	category := helper.RandomString(5)
	expense1 := createRandomExpense(t, member.ID, category)

	err := testStore.DeleteExpense(context.Background(), expense1.ID)

	require.NoError(t, err)

	expense2, err := testStore.GetExpense(context.Background(), expense1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, expense2)
}

func TestListExpenses(t *testing.T) {
	member := createRandomMember(t, createRandomGroup(t).ID)
	category := helper.RandomString(5)
	for i := 0; i < 10; i++ {
		createRandomExpense(t, member.ID, category)
	}

	expenses, err := testStore.ListExpenses(context.Background(), member.GroupID)
	require.NoError(t, err)
	require.NotEmpty(t, expenses)

	for _, expense := range expenses {
		require.NotEmpty(t, expense)

		member, err := testStore.GetMember(context.Background(), expense.MemberID)
		require.NoError(t, err)
		require.Equal(t, member.GroupID, member.GroupID)
	}
}

func TestListNonSettledExpenses(t *testing.T) {
	member := createRandomMember(t, createRandomGroup(t).ID)
	category := helper.RandomString(5)

	var lastExpense Expense
	for i := 0; i < 10; i++ {
		lastExpense = createRandomExpense(t, member.ID, category)
	}

	expenses1, err := testStore.ListExpenses(context.Background(), member.GroupID)

	require.NoError(t, err)
	require.NotEmpty(t, expenses1)

	for _, expense := range expenses1 {
		require.NotEmpty(t, expense)

		member, err := testStore.GetMember(context.Background(), expense.MemberID)
		require.NoError(t, err)
		require.Equal(t, member.GroupID, member.GroupID)
	}

	_, err = testStore.UpdateExpense(context.Background(), UpdateExpenseParams{
		ID:        lastExpense.ID,
		MemberID:  lastExpense.MemberID,
		Amount:    lastExpense.Amount,
		Date:      lastExpense.Date,
		IsSettled: true,
	})
	require.NoError(t, err)

	expenses2, err := testStore.ListNonSettledExpenses(context.Background(), member.GroupID)

	require.NoError(t, err)
	require.Equal(t, len(expenses1)-1, len(expenses2))
}

func TestListSumOfExpensesWithCategory(t *testing.T) {
	member := createRandomMember(t, createRandomGroup(t).ID)
	category := helper.RandomString(5)
	for i := 0; i < 10; i++ {
		createRandomExpense(t, member.ID, category)
	}

	expenses, err := testStore.ListExpenses(context.Background(), member.GroupID)
	require.NoError(t, err)
	require.NotEmpty(t, expenses)

	expectedSum := int64(0)
	for _, expense := range expenses {
		if expense.Category.String == category {
			amount, err := strconv.ParseInt(expense.Amount, 10, 64)
			require.NoError(t, err)
			expectedSum += amount
		}
	}

	categorySums, err := testStore.SummarizeExpensesWithinDate(context.Background(), SummarizeExpensesWithinDateParams{
		GroupID:   member.GroupID,
		StartTime: time.Now().AddDate(0, 0, -365),
		EndTime:   time.Now().AddDate(0, 0, 365),
	})
	require.NoError(t, err)

	for _, sum := range categorySums {
		if sum.Category.String == category {
			require.Equal(t, expectedSum, sum.Total)
		}
	}
}
