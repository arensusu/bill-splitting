package db

import (
	"bill-splitting/helper"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUserExpense(t *testing.T) UserExpense {
	user := createRandomUser(t)
	expense := createRandomExpense(t)

	param := CreateUserExpenseParams{
		ExpenseID: expense.ID,
		UserID:    user.ID,
		Share:     helper.RandomInt64(1, 100),
	}
	userExpense, err := testStore.CreateUserExpense(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, userExpense)

	require.Equal(t, expense.ID, userExpense.ExpenseID)
	require.Equal(t, user.ID, userExpense.UserID)
	require.Equal(t, param.Share, userExpense.Share)

	return userExpense
}

func TestCreateUserExpense(t *testing.T) {
	createRandomUserExpense(t)
}

func TestGetUserExpense(t *testing.T) {
	userExpense1 := createRandomUserExpense(t)

	param := GetUserExpenseParams{
		ExpenseID: userExpense1.ExpenseID,
		UserID:    userExpense1.UserID,
	}
	userExpense2, err := testStore.GetUserExpense(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, userExpense2)

	require.Equal(t, userExpense1.ExpenseID, userExpense2.ExpenseID)
	require.Equal(t, userExpense1.UserID, userExpense2.UserID)
	require.Equal(t, userExpense1.Share, userExpense2.Share)
}

func TestUpdateUserExpense(t *testing.T) {
	userExpense1 := createRandomUserExpense(t)

	newShare := helper.RandomInt64(1, 100)
	param := UpdateUserExpenseParams{
		ExpenseID: userExpense1.ExpenseID,
		UserID:    userExpense1.UserID,
		Share:     newShare,
	}

	userExpense2, err := testStore.UpdateUserExpense(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, userExpense2)

	require.Equal(t, userExpense1.ExpenseID, userExpense2.ExpenseID)
	require.Equal(t, userExpense1.UserID, userExpense2.UserID)
	require.Equal(t, newShare, userExpense2.Share)
}

func TestDeleteUserExpense(t *testing.T) {
	userExpense1 := createRandomUserExpense(t)

	deleteParam := DeleteUserExpenseParams{
		ExpenseID: userExpense1.ExpenseID,
		UserID:    userExpense1.UserID,
	}
	err := testStore.DeleteUserExpense(context.Background(), deleteParam)

	require.NoError(t, err)

	getParam := GetUserExpenseParams{
		ExpenseID: userExpense1.ExpenseID,
		UserID:    userExpense1.UserID,
	}
	userExpense2, err := testStore.GetUserExpense(context.Background(), getParam)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, userExpense2)
}
