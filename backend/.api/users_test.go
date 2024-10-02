package api

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/helper"
)

// import (
// 	mockdb "bill-splitting/db/mock"
// 	db "bill-splitting/db/sqlc"
// 	"bill-splitting/helper"
// 	"bytes"
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// 	"go.uber.org/mock/gomock"
// )

func randomUser() db.User {
	return db.User{
		ID:       helper.RandomString(32),
		Username: helper.RandomString(10),
	}
}

// func TestCreateUserAPI(t *testing.T) {
// 	user := randomUser()

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	testCases := []struct {
// 		name          string
// 		body          createUserRequest
// 		buildStub     func(t *testing.T, mockStore *mockdb.MockStore)
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: createUserRequest{Username: user.Username, Password: user.Password},
// 			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
// 				param := db.CreateUserParams{
// 					Username: user.Username,
// 					Password: user.Password,
// 				}
// 				mockStore.EXPECT().CreateUser(gomock.Any(), gomock.Eq(param)).Times(1).Return(user, nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 				requireBodyMatchUser(t, user, recorder.Body)
// 			},
// 		},
// 		{
// 			name: "InternalError",
// 			body: createUserRequest{Username: user.Username, Password: user.Password},
// 			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
// 				mockStore.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(1).Return(db.User{}, sql.ErrConnDone)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 		{
// 			name: "InvalidUsername",
// 			body: createUserRequest{Username: "", Password: user.Password},
// 			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
// 				mockStore.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(0)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			mockStore := mockdb.NewMockStore(ctrl)
// 			tc.buildStub(t, mockStore)

// 			server := newTestServer(t, mockStore)
// 			recorder := httptest.NewRecorder()

// 			data, err := json.Marshal(tc.body)
// 			require.NoError(t, err)

// 			req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewReader(data))
// 			require.NoError(t, err)

// 			server.router.ServeHTTP(recorder, req)
// 			tc.checkResponse(t, recorder)
// 		})
// 	}
// }

// func TestGetUserAPI(t *testing.T) {
// 	user := randomUser()

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	testCases := []struct {
// 		name          string
// 		userID        int64
// 		buildStub     func(t *testing.T, mockStore *mockdb.MockStore)
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name:   "OK",
// 			userID: user.ID,
// 			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
// 				mockStore.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(user, nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 				requireBodyMatchUser(t, user, recorder.Body)
// 			},
// 		},
// 		{
// 			name:   "NotFound",
// 			userID: user.ID,
// 			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
// 				mockStore.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(db.User{}, sql.ErrNoRows)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusNotFound, recorder.Code)
// 			},
// 		},
// 		{
// 			name:   "InternalError",
// 			userID: user.ID,
// 			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
// 				mockStore.EXPECT().GetUser(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(db.User{}, sql.ErrConnDone)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 		{
// 			name:   "InvalidID",
// 			userID: 0,
// 			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
// 				mockStore.EXPECT().GetUser(gomock.Any(), gomock.Any()).Times(0)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusBadRequest, recorder.Code)
// 			},
// 		},
// 	}

// 	mockStore := mockdb.NewMockStore(ctrl)
// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			tc.buildStub(t, mockStore)
// 			server := newTestServer(t, mockStore)
// 			recorder := httptest.NewRecorder()

// 			url := fmt.Sprintf("/users/%d", tc.userID)
// 			request, err := http.NewRequest(http.MethodGet, url, nil)
// 			require.NoError(t, err)

// 			server.router.ServeHTTP(recorder, request)
// 			tc.checkResponse(t, recorder)
// 		})
// 	}
// }

// func requireBodyMatchUser(t *testing.T, user db.User, body *bytes.Buffer) {
// 	data, err := io.ReadAll(body)
// 	require.NoError(t, err)

// 	var gotUser db.User
// 	err = json.Unmarshal(data, &gotUser)

// 	require.NoError(t, err)
// 	require.Equal(t, user, gotUser)
// }
