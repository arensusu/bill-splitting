package api

import (
	mockdb "bill-splitting/db/mock"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestReplaceSettlementAPI(t *testing.T) {
	user := randomUser()
	group := randomGroup()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		groupID       int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			groupID: group.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateSettlementsTx(gomock.Any(), gomock.Eq(group.ID)).Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:    "InternalError",
			groupID: group.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateSettlementsTx(gomock.Any(), gomock.Eq(group.ID)).Times(1).Return(nil, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:    "BadRequest",
			groupID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateSettlementsTx(gomock.Any(), gomock.Any()).Times(0)
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
			tc.buildStubs(mockStore)

			server := newTestServer(mockStore)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/api/v1/groups/%d/settlements", tc.groupID)
			req := httptest.NewRequest(http.MethodPut, url, nil)

			addAuthentication(t, req, server.tokenMaker, user.ID, time.Minute)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)
		})
	}
}
