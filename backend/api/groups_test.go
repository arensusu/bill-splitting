package api

import (
	mockdb "bill-splitting/db/mock"
	db "bill-splitting/db/sqlc"
	"bill-splitting/helper"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func randomGroup() db.Group {
	return db.Group{
		ID:   int32(helper.RandomInt64(1, 1000)),
		Name: helper.RandomString(10),
	}
}

func requireBodyMatchGroup(t *testing.T, group db.Group, body *bytes.Buffer) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotGroup db.Group
	err = json.Unmarshal(data, &gotGroup)

	require.NoError(t, err)
	require.Equal(t, group, gotGroup)
}

func TestGetGroupAPI(t *testing.T) {
	user := randomUser()
	group := randomGroup()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		groupID       int32
		buildStub     func(t *testing.T, mockStore *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			groupID: group.ID,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetGroup(gomock.Any(), gomock.Eq(group.ID)).Times(1).Return(group, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchGroup(t, group, recorder.Body)
			},
		},
		{
			name:    "NotFound",
			groupID: group.ID,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetGroup(gomock.Any(), gomock.Eq(group.ID)).Times(1).Return(db.Group{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "InternalError",
			groupID: group.ID,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetGroup(gomock.Any(), gomock.Eq(group.ID)).Times(1).Return(db.Group{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:    "InvalidID",
			groupID: 0,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetGroup(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			mockStore := mockdb.NewMockStore(ctrl)
			tc.buildStub(t, mockStore)

			server := newTestServer(mockStore)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/api/v1/groups/%d", tc.groupID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			addAuthentication(t, request, server.tokenMaker, user.ID, time.Minute)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestCreateGroupAPI(t *testing.T) {
	user := randomUser()
	group := randomGroup()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		body          createGroupRequest
		buildStub     func(t *testing.T, mockStore *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: createGroupRequest{Name: group.Name},
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().CreateGroupTx(gomock.Any(), gomock.Any()).Times(1).Return(group, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code, recorder.Body)
				requireBodyMatchGroup(t, group, recorder.Body)
			},
		},
		{
			name: "InternalError",
			body: createGroupRequest{Name: group.Name},
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().CreateGroupTx(gomock.Any(), gomock.Eq(db.CreateGroupTxParams{
					Name:   group.Name,
					UserID: user.ID,
				})).Times(1).Return(db.Group{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidName",
			body: createGroupRequest{Name: ""},
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().CreateGroupTx(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			mockStore := mockdb.NewMockStore(ctrl)
			tc.buildStub(t, mockStore)

			server := newTestServer(mockStore)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/api/v1/groups"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			addAuthentication(t, request, server.tokenMaker, user.ID, time.Minute)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchGroups(t *testing.T, groups []db.Group, body *bytes.Buffer) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotGroups []db.Group
	err = json.Unmarshal(data, &gotGroups)

	require.NoError(t, err)
	require.Equal(t, groups, gotGroups)
}

func TestListGroupsAPI(t *testing.T) {
	user := randomUser()
	groups := []db.Group{}
	for i := 0; i < 5; i += 1 {
		groups = append(groups, randomGroup())
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mockdb.NewMockStore(ctrl)
	mockStore.EXPECT().ListGroups(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(groups, nil)

	server := newTestServer(mockStore)
	recorder := httptest.NewRecorder()

	request, err := http.NewRequest(http.MethodGet, "/api/v1/groups", nil)
	require.NoError(t, err)

	addAuthentication(t, request, server.tokenMaker, user.ID, time.Minute)
	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchGroups(t, groups, recorder.Body)
}
