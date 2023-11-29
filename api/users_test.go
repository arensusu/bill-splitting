package api

import (
	mockdb "bill-splitting/db/mock"
	db "bill-splitting/db/sqlc"
	"bill-splitting/helper"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func randomUser() db.User {
	return db.User{
		ID:       helper.RandomInt64(1, 1000),
		Username: helper.RandomString(10),
		Password: helper.RandomString(10),
	}
}

func TestGetUserApi(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := randomUser()

	mockStore := mockdb.NewMockStore(ctrl)
	mockStore.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(user, nil)

	server := NewServer(mockStore)
	recoder := httptest.NewRecorder()

	url := fmt.Sprintf("/users/%d", user.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recoder, request)
	require.Equal(t, http.StatusOK, recoder.Code)
	requireBodyMatchUser(t, user, recoder.Body)
}

func requireBodyMatchUser(t *testing.T, user db.User, body *bytes.Buffer) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, user, gotUser)
}
