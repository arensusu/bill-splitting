package db

import (
	"bill-splitting/helper"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomGroup(t *testing.T) Group {
	name := helper.RandomString(10)
	lineId := helper.RandomString(20)

	group, err := testStore.CreateGroup(context.Background(), CreateGroupParams{
		Name:   name,
		LineID: sql.NullString{String: lineId, Valid: lineId != ""},
	})

	require.NoError(t, err)
	require.NotEmpty(t, group)

	require.NotZero(t, group.ID)
	require.Equal(t, name, group.Name)

	return group
}

func TestCreateGroup(t *testing.T) {
	createRandomGroup(t)
}

func TestGetGroup(t *testing.T) {
	group1 := createRandomGroup(t)

	group2, err := testStore.GetGroup(context.Background(), group1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, group2)

	require.Equal(t, group1.ID, group2.ID)
	require.Equal(t, group1.Name, group2.Name)
}

func TestUpdateGroup(t *testing.T) {
	group1 := createRandomGroup(t)

	newName := helper.RandomString(10)
	param := UpdateGroupParams{
		ID:   group1.ID,
		Name: newName,
	}
	group2, err := testStore.UpdateGroup(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, group2)

	require.Equal(t, group1.ID, group2.ID)
	require.Equal(t, newName, group2.Name)
}

func TestDeleteGroup(t *testing.T) {
	group1 := createRandomGroup(t)

	err := testStore.DeleteGroup(context.Background(), group1.ID)

	require.NoError(t, err)

	group2, err := testStore.GetGroup(context.Background(), group1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, group2)
}

func TestListGroups(t *testing.T) {
	user := createRandomUser(t)
	groups := []Group{}
	for i := 0; i < 10; i++ {
		groups = append(groups, createRandomGroup(t))
	}

	for i := 0; i < 5; i += 1 {
		testStore.CreateMember(context.Background(), CreateMemberParams{
			GroupID: groups[i].ID,
			UserID:  user.ID,
		})
	}

	actualGroups, err := testStore.ListGroups(context.Background(), user.ID)
	require.NoError(t, err)
	require.Len(t, actualGroups, 5)

	for i := 0; i < 5; i += 1 {
		require.Equal(t, groups[i], actualGroups[i])
	}
}
