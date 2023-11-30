package api

import (
	mockdb "bill-splitting/db/mock"
	db "bill-splitting/db/sqlc"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func requireBodyMatchGroupMember(t *testing.T, member db.GroupMember, body *bytes.Buffer) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotMember db.GroupMember
	err = json.Unmarshal(data, &gotMember)

	require.NoError(t, err)
	require.Equal(t, member, gotMember)
}

func requireBodyMatchGroupMembers(t *testing.T, members []db.GroupMember, body *bytes.Buffer) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotMembers []db.GroupMember
	err = json.Unmarshal(data, &gotMembers)

	require.NoError(t, err)
	require.Equal(t, members, gotMembers)
}

func TestCreateGroupMember(t *testing.T) {
	user := randomUser()
	group := randomGroup()

	groupMember := db.GroupMember{
		GroupID: group.ID,
		UserID:  user.ID,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		body          createGroupMemberRequest
		buildStub     func(t *testing.T, mockStore *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: createGroupMemberRequest{GroupID: group.ID, UserID: user.ID},
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().CreateGroupMember(gomock.Any(), gomock.Eq(db.CreateGroupMemberParams{
					GroupID: groupMember.GroupID,
					UserID:  groupMember.UserID,
				})).Times(1).Return(groupMember, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
				requireBodyMatchGroupMember(t, groupMember, recoder.Body)
			},
		},
		{
			name: "InternalError",
			body: createGroupMemberRequest{GroupID: group.ID, UserID: user.ID},
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().CreateGroupMember(gomock.Any(), gomock.Eq(db.CreateGroupMemberParams{
					GroupID: groupMember.GroupID,
					UserID:  groupMember.UserID,
				})).Times(1).Return(db.GroupMember{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "BadRequest",
			body: createGroupMemberRequest{},
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().CreateGroupMember(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			mockStore := mockdb.NewMockStore(ctrl)
			tc.buildStub(t, mockStore)

			server := NewServer(mockStore)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/group-members"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListGroupMembers(t *testing.T) {
	user := randomUser()
	group := randomGroup()
	groupMember := db.GroupMember{
		GroupID: group.ID,
		UserID:  user.ID,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		groupID       int64
		buildStub     func(t *testing.T, mockStore *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			groupID: group.ID,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().ListGroupMembers(gomock.Any(), gomock.Eq(group.ID)).Times(1).Return([]db.GroupMember{groupMember}, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
				requireBodyMatchGroupMembers(t, []db.GroupMember{groupMember}, recoder.Body)
			},
		},
		{
			name:    "InternalError",
			groupID: group.ID,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().ListGroupMembers(gomock.Any(), gomock.Eq(group.ID)).Times(1).Return([]db.GroupMember{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name:    "BadRequest",
			groupID: 0,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().ListGroupMembers(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			mockStore := mockdb.NewMockStore(ctrl)
			tc.buildStub(t, mockStore)

			server := NewServer(mockStore)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/group-members/%d", tc.groupID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
