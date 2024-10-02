package model

import (
	"bill-splitting/helper"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user := User{
		LineID:   helper.RandomString(32),
		Username: helper.RandomString(10),
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+) RETURNING "id"`).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.LineID, user.Username).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	err := testStore.CreateUser(&user)
	assert.NoError(t, err)
}

func TestGetUserByLineID(t *testing.T) {
	expectedUser := &User{
		LineID:   helper.RandomString(32),
		Username: helper.RandomString(10),
	}

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "line_id", "username"}).
		AddRow(1, nil, nil, nil, expectedUser.LineID, expectedUser.Username)

	mock.ExpectQuery(`^SELECT \* FROM "users" WHERE line_id = \$1 (.+)`).
		WithArgs(expectedUser.LineID, sqlmock.AnyArg()).
		WillReturnRows(rows)

	user, err := testStore.GetUserByLineID(expectedUser.LineID)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.LineID, user.LineID)
	assert.Equal(t, expectedUser.Username, user.Username)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestGetUserByUsername(t *testing.T) {
	expectedUser := &User{
		LineID:   helper.RandomString(32),
		Username: helper.RandomString(10),
	}

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "line_id", "username"}).
		AddRow(1, nil, nil, nil, expectedUser.LineID, expectedUser.Username)

	mock.ExpectQuery(`^SELECT \* FROM "users" WHERE username = \$1 (.+)`).
		WithArgs(expectedUser.Username, sqlmock.AnyArg()).
		WillReturnRows(rows)

	user, err := testStore.GetUserByUsername(expectedUser.Username)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.LineID, user.LineID)
	assert.Equal(t, expectedUser.Username, user.Username)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	userID := uint(1)

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET "deleted_at"=\$1 WHERE "users"."id" = \$2`).
		WithArgs(sqlmock.AnyArg(), userID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := testStore.DeleteUser(userID)
	assert.NoError(t, err)
}
