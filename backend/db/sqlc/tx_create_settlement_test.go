package db

import (
	"bill-splitting/helper"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateSettlementTx(t *testing.T) {
	group := createRandomGroup(t)
	users := []User{}
	for i := 0; i < 5; i += 1 {
		users = append(users, createRandomUser(t))
	}
	for _, user := range users {
		testStore.CreateGroupMember(context.Background(), CreateGroupMemberParams{
			GroupID: group.ID,
			UserID:  user.ID,
		})
	}

	for i := range users {
		testStore.CreateExpenseTx(context.Background(), CreateExpenseTxParams{
			GroupID: group.ID,
			PayerID: users[i].ID,
			Amount:  helper.RandomInt64(1, 1000),
			Date:    helper.RandomDate(),
		})
	}

	settlements, err := testStore.CreateSettlementsTx(context.Background(), group.ID)

	require.NoError(t, err)
	require.NotEmpty(t, settlements)

	settlements, err = testStore.CreateSettlementsTx(context.Background(), group.ID)

	require.NoError(t, err)
	require.NotEmpty(t, settlements)

}
