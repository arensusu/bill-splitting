package db_test

import (
	db "bill-splitting/db/sqlc"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomGroupMember(t *testing.T) db.GroupMember {
	group := createRandomGroup(t)
	user := createRandomUser(t)
	param := db.CreateGroupMemberParams{
		GroupID: group.ID,
		UserID:  user.ID,
	}

	groupMember, err := testQueries.CreateGroupMember(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, groupMember)

	require.Equal(t, group.ID, groupMember.GroupID)
	require.Equal(t, user.ID, groupMember.UserID)
	require.NotZero(t, groupMember.CreatedAt)

	return groupMember
}

func TestCreateGroupMember(t *testing.T) {
	createRandomGroupMember(t)
}

func TestGetGroupMember(t *testing.T) {
	groupMember1 := createRandomGroupMember(t)

	param := db.GetGroupMemberParams{
		GroupID: groupMember1.GroupID,
		UserID:  groupMember1.UserID,
	}
	groupMember2, err := testQueries.GetGroupMember(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, groupMember2)

	require.Equal(t, groupMember1.GroupID, groupMember2.GroupID)
	require.Equal(t, groupMember1.UserID, groupMember2.UserID)
	require.WithinDuration(t, groupMember1.CreatedAt, groupMember2.CreatedAt, time.Second)
}

func TestDeleteGroupMember(t *testing.T) {
	groupMember1 := createRandomGroupMember(t)

	deleteParam := db.DeleteGroupMemberParams{
		GroupID: groupMember1.GroupID,
		UserID:  groupMember1.UserID,
	}
	err := testQueries.DeleteGroupMember(context.Background(), deleteParam)

	require.NoError(t, err)

	getParam := db.GetGroupMemberParams{
		GroupID: groupMember1.GroupID,
		UserID:  groupMember1.UserID,
	}
	groupMember2, err := testQueries.GetGroupMember(context.Background(), getParam)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, groupMember2)
}
