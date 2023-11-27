package db_test

import (
	"bill-splitting/db/helper"
	db "bill-splitting/db/sqlc"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateExpenseTx(t *testing.T) {
	group := createRandomGroup(t)
	users := []db.User{}
	for i := 0; i < 5; i += 1 {
		users = append(users, createRandomUser(t))
	}
	for _, user := range users {
		testStore.Queries.CreateGroupMember(context.Background(), db.CreateGroupMemberParams{
			GroupID: group.ID,
			UserID:  user.ID,
		})
	}

	param := db.CreateExpenseTxParams{
		GroupID: group.ID,
		PayerID: users[0].ID,
		Amount:  helper.RandomInt64(1, 1000),
		Date:    helper.RandomDate(),
	}
	result, err := testStore.CreateExpenseTx(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, param.GroupID, result.Expense.GroupID)
	require.Equal(t, param.PayerID, result.Expense.PayerID)
	require.Equal(t, param.Amount, result.Expense.Amount)
	require.WithinDuration(t, param.Date, result.Expense.Date, time.Second)

	require.Len(t, result.UserExpenses, len(users))

	total := int64(0)
	for i, ue := range result.UserExpenses {
		require.Equal(t, result.Expense.ID, ue.ExpenseID)
		require.Equal(t, users[i].ID, ue.UserID)
		require.Equal(t, param.Amount/int64(len(users)), ue.Share)
		total += ue.Share
	}
	require.LessOrEqual(t, total, param.Amount)
}
