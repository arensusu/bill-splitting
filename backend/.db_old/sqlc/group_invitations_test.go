package db

import (
	"bill-splitting/helper"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomGroupInvitation(t *testing.T, groupID int32) GroupInvitation {
	code := helper.RandomString(8)
	param := CreateGroupInvitationParams{
		Code:    code,
		GroupID: groupID,
	}

	invite, err := testStore.CreateGroupInvitation(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, invite)
	require.Equal(t, code, invite.Code)
	require.Equal(t, groupID, invite.GroupID)

	return invite
}

func TestCreateGroupInvitation(t *testing.T) {
	group := createRandomGroup(t)
	createRandomGroupInvitation(t, group.ID)
}

func TestGetGroupInvitation(t *testing.T) {
	group := createRandomGroup(t)
	expected := createRandomGroupInvitation(t, group.ID)

	actual, err := testStore.GetGroupInvitation(context.Background(), expected.Code)
	require.NoError(t, err)
	require.NotEmpty(t, actual)
	require.Equal(t, expected.Code, actual.Code)
	require.Equal(t, expected.GroupID, actual.GroupID)
}

func TestDeleteGroupInvitation(t *testing.T) {
	group := createRandomGroup(t)
	invite := createRandomGroupInvitation(t, group.ID)

	err := testStore.DeleteGroupInvitation(context.Background(), invite.Code)
	require.NoError(t, err)

	actual, err := testStore.GetGroupInvitation(context.Background(), invite.Code)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, actual)
}
