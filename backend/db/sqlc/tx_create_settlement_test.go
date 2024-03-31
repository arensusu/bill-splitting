package db

import (
	"bill-splitting/helper"
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateSettlementTx(t *testing.T) {
	group := createRandomGroup(t)
	users := []User{}
	for i := 0; i < 5; i += 1 {
		users = append(users, createRandomUser(t))
	}

	members := []Member{}
	for _, user := range users {
		member, _ := testStore.CreateMember(context.Background(), CreateMemberParams{
			GroupID: group.ID,
			UserID:  user.ID,
		})
		members = append(members, member)
	}

	for _, member := range members {
		testStore.CreateExpense(context.Background(), CreateExpenseParams{
			MemberID: member.ID,
			Amount:   strconv.FormatInt(helper.RandomInt64(1, 1000), 10),
			Date:     helper.RandomDate(),
		})
	}

	settlements, err := testStore.CreateSettlementsTx(context.Background(), group.ID)
	require.NoError(t, err)
	require.NotEmpty(t, settlements)

	settlements, err = testStore.CreateSettlementsTx(context.Background(), group.ID)
	require.NoError(t, err)
	require.NotEmpty(t, settlements)
}
