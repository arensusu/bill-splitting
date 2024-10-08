package db

import (
	"bill-splitting/helper"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	param := CreateUserParams{
		ID:       helper.RandomString(32),
		Username: helper.RandomString(10),
	}

	user, err := testStore.CreateUser(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotZero(t, user.ID)
	require.Equal(t, param.Username, user.Username)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testStore.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)
	newUsername := helper.RandomString(10)
	param := UpdateUserParams{
		ID:       user1.ID,
		Username: newUsername,
	}

	user2, err := testStore.UpdateUser(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, newUsername, user2.Username)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	err := testStore.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testStore.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestGetUserByUsername(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testStore.GetUserByUsername(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
}
