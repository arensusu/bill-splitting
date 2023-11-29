package db

import (
	"bill-splitting/helper"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	param := CreateUserParams{
		Username: helper.RandomString(10),
		Password: helper.RandomString(10),
	}

	user, err := testStore.CreateUser(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotZero(t, user.ID)
	require.Equal(t, param.Username, user.Username)
	require.Equal(t, param.Password, user.Password)
	require.NotZero(t, user.CreatedAt)

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
	require.Equal(t, user1.Password, user2.Password)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	newUsername := helper.RandomString(10)
	newPassword := helper.RandomString(10)
	param := UpdateUserParams{
		ID:       user1.ID,
		Username: newUsername,
		Password: newPassword,
	}

	user2, err := testStore.UpdateUser(context.Background(), param)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, newUsername, user2.Username)
	require.Equal(t, newPassword, user2.Password)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
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
