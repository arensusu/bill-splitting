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
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func requireBodyMatchGroupMembers(t *testing.T, members []db.Member, body *bytes.Buffer) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotMembers []db.Member
	err = json.Unmarshal(data, &gotMembers)

	require.NoError(t, err)
	require.Equal(t, members, gotMembers)
}

func TestListGroupMembers(t *testing.T) {
	user := randomUser()
	group := randomGroup()
	groupMember := db.Member{
		GroupID: group.ID,
		UserID:  user.ID,
	}

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
				mockStore.EXPECT().ListMembersOfGroup(gomock.Any(), gomock.Eq(group.ID)).Times(1).Return([]db.Member{groupMember}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchGroupMembers(t, []db.Member{groupMember}, recorder.Body)
			},
		},
		{
			name:    "InternalError",
			groupID: group.ID,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().ListMembersOfGroup(gomock.Any(), gomock.Eq(group.ID)).Times(1).Return([]db.Member{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:    "BadRequest",
			groupID: 0,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().ListMembersOfGroup(gomock.Any(), gomock.Any()).Times(0)
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

			url := fmt.Sprintf("/api/v1/groups/%d/members", tc.groupID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			addAuthentication(t, request, server.tokenMaker, user.ID, time.Minute)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
