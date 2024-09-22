package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomMember(t *testing.T, groupId int32) Member {
	user := createRandomUser(t)
	param := CreateMemberParams{
		GroupID: groupId,
		UserID:  user.ID,
	}

	member, err := testStore.CreateMember(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, member)

	require.Equal(t, groupId, member.GroupID)
	require.Equal(t, user.ID, member.UserID)

	return member
}

func TestCreateMember(t *testing.T) {
	createRandomMember(t, createRandomGroup(t).ID)
}

func TestGetMember(t *testing.T) {
	member1 := createRandomMember(t, createRandomGroup(t).ID)

	member2, err := testStore.GetMember(context.Background(), member1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, member2)

	require.Equal(t, member1.GroupID, member2.GroupID)
	require.Equal(t, member1.UserID, member2.UserID)
}

func TestDeleteMember(t *testing.T) {
	member1 := createRandomMember(t, createRandomGroup(t).ID)

	deleteParam := DeleteMemberParams{
		GroupID: member1.GroupID,
		UserID:  member1.UserID,
	}
	err := testStore.DeleteMember(context.Background(), deleteParam)

	require.NoError(t, err)

	member2, err := testStore.GetMember(context.Background(), member1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, member2)
}
