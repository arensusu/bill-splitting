package db

import (
	"bill-splitting/helper"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateGroupTx(t *testing.T) {
	user := createRandomUser(t)
	groupName := helper.RandomString(10)

	group, err := testStore.CreateGroupTx(context.Background(), CreateGroupTxParams{
		Name:   groupName,
		UserID: user.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, group)
	require.Equal(t, groupName, group.Name)
	require.NotZero(t, group.ID)

	members, err := testStore.ListMembersOfGroup(context.Background(), group.ID)
	require.NoError(t, err)
	require.Len(t, members, 1)
	require.Equal(t, user.Username, members[0].Username)
}
