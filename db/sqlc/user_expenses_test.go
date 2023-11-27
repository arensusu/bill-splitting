package db_test

import (
	"bill-splitting/db/helper"
	db "bill-splitting/db/sqlc"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUserExpense(t *testing.T) db.UserExpense {
	user := createRandomUser(t)
	expense := createRandomExpense(t)

	param := db.CreateUserExpenseParams{
		ExpenseID: expense.ID,
		UserID:    user.ID,
		Share:     helper.RandomInt64(1, 100),
	}
	userExpense, err := testStore.Queries.CreateUserExpense(context.Background(), param)

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

	param := db.GetUserExpenseParams{
		ExpenseID: userExpense1.ExpenseID,
		UserID:    userExpense1.UserID,
	}
	userExpense2, err := testStore.Queries.GetUserExpense(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, userExpense2)

	require.Equal(t, userExpense1.ExpenseID, userExpense2.ExpenseID)
	require.Equal(t, userExpense1.UserID, userExpense2.UserID)
	require.Equal(t, userExpense1.Share, userExpense2.Share)
}

func TestUpdateUserExpense(t *testing.T) {
	userExpense1 := createRandomUserExpense(t)

	newShare := helper.RandomInt64(1, 100)
	param := db.UpdateUserExpenseParams{
		ExpenseID: userExpense1.ExpenseID,
		UserID:    userExpense1.UserID,
		Share:     newShare,
	}

	userExpense2, err := testStore.Queries.UpdateUserExpense(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, userExpense2)

	require.Equal(t, userExpense1.ExpenseID, userExpense2.ExpenseID)
	require.Equal(t, userExpense1.UserID, userExpense2.UserID)
	require.Equal(t, newShare, userExpense2.Share)
}

func TestDeleteUserExpense(t *testing.T) {
	userExpense1 := createRandomUserExpense(t)

	deleteParam := db.DeleteUserExpenseParams{
		ExpenseID: userExpense1.ExpenseID,
		UserID:    userExpense1.UserID,
	}
	err := testStore.Queries.DeleteUserExpense(context.Background(), deleteParam)

	require.NoError(t, err)

	getParam := db.GetUserExpenseParams{
		ExpenseID: userExpense1.ExpenseID,
		UserID:    userExpense1.UserID,
	}
	userExpense2, err := testStore.Queries.GetUserExpense(context.Background(), getParam)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, userExpense2)
}
