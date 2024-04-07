package api

import (
	mockdb "bill-splitting/db/mock"
	db "bill-splitting/db/sqlc"
	"bill-splitting/helper"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func requireBodyMatchMember(t *testing.T, member db.Member, body *bytes.Buffer) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotMember db.Member
	err = json.Unmarshal(data, &gotMember)

	require.NoError(t, err)
	require.Equal(t, member, gotMember)
}

func TestAcceptGroupInvitation(t *testing.T) {
	user := randomUser()
	group := randomGroup()
	invitationCode := helper.RandomString(6)
	member := db.Member{
		ID:      int32(helper.RandomInt64(1, 100000)),
		GroupID: group.ID,
		UserID:  user.ID,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name           string
		invitationCode string
		buildStub      func(t *testing.T, mockStore *mockdb.MockStore)
		checkResponse  func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:           "OK",
			invitationCode: invitationCode,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().
					GetGroupInvitation(gomock.Any(), gomock.Eq(invitationCode)).
					Times(1).
					Return(db.GroupInvitation{
						Code:    invitationCode,
						GroupID: group.ID,
					}, nil)

				mockStore.EXPECT().
					DeleteGroupInvitation(gomock.Any(), gomock.Eq(invitationCode)).
					Times(1).
					Return(nil)

				mockStore.EXPECT().
					CreateMember(gomock.Any(), gomock.Eq(db.CreateMemberParams{
						GroupID: group.ID,
						UserID:  user.ID,
					})).
					Times(1).
					Return(member, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
				requireBodyMatchMember(t, member, recoder.Body)
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

			url := "/groups/invite/" + tc.invitationCode
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			addAuthentication(t, request, server.tokenMaker, user.ID, time.Minute)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchInvitationCode(t *testing.T, expectedCode db.GroupInvitation, body *bytes.Buffer) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var actualCode db.GroupInvitation
	err = json.Unmarshal(data, &actualCode)
	require.NoError(t, err)
	require.Equal(t, expectedCode, actualCode)
}

func TestCreateGroupInvitation(t *testing.T) {
	user := randomUser()
	group := randomGroup()
	invitationCode := db.GroupInvitation{
		Code:    helper.RandomString(6),
		GroupID: group.ID,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		body          createGroupInvitationParams
		buildStub     func(t *testing.T, mockStore *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: createGroupInvitationParams{GroupID: group.ID},
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().CreateGroupInvitation(gomock.Any(), gomock.Any()).Times(1).Return(invitationCode, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchInvitationCode(t, invitationCode, recorder.Body)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockStore := mockdb.NewMockStore(ctrl)
			tc.buildStub(t, mockStore)

			server := newTestServer(mockStore)
			recorder := httptest.NewRecorder()
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/group/invite"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			addAuthentication(t, request, server.tokenMaker, user.ID, time.Minute)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
