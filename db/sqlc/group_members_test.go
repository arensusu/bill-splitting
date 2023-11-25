package db_test

import (
	db "bill-splitting/db/sqlc"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomGroupMember(t *testing.T, groupID int64, userID int64) db.GroupMember {
	param := db.CreateGroupMemberParams{
		GroupID: groupID,
		UserID:  userID,
	}

	groupMember, err := testQueries.CreateGroupMember(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, groupMember)

	require.Equal(t, groupID, groupMember.GroupID)
	require.Equal(t, userID, groupMember.UserID)
	require.NotZero(t, groupMember.CreatedAt)

	return groupMember
}

func TestCreateGroupMember(t *testing.T) {
	group1 := createRandomGroup(t)
	user1 := createRandomUser(t)

	createRandomGroupMember(t, group1.ID, user1.ID)
}

func TestGetGroupMember(t *testing.T) {
	group1 := createRandomGroup(t)
	user1 := createRandomUser(t)
	groupMember1 := createRandomGroupMember(t, group1.ID, user1.ID)

	param := db.GetGroupMemberParams{
		GroupID: group1.ID,
		UserID:  user1.ID,
	}
	groupMember2, err := testQueries.GetGroupMember(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, groupMember2)

	require.Equal(t, groupMember1.GroupID, groupMember2.GroupID)
	require.Equal(t, groupMember1.UserID, groupMember2.UserID)
	require.WithinDuration(t, groupMember1.CreatedAt, groupMember2.CreatedAt, time.Second)
}

func TestDeleteGroupMember(t *testing.T) {
	group1 := createRandomGroup(t)
	user1 := createRandomUser(t)
	createRandomGroupMember(t, group1.ID, user1.ID)

	deleteParam := db.DeleteGroupMemberParams{
		GroupID: group1.ID,
		UserID:  user1.ID,
	}
	err := testQueries.DeleteGroupMember(context.Background(), deleteParam)

	require.NoError(t, err)

	getParam := db.GetGroupMemberParams{
		GroupID: group1.ID,
		UserID:  user1.ID,
	}
	groupMember2, err := testQueries.GetGroupMember(context.Background(), getParam)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, groupMember2)
}
