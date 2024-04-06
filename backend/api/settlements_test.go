package api

import (
	mockdb "bill-splitting/db/mock"
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestReplaceSettlementAPI(t *testing.T) {
	group := randomGroup()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		body          replaceSettlementRequest
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: replaceSettlementRequest{GroupID: group.ID},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateSettlementsTx(gomock.Any(), gomock.Eq(group.ID)).Times(1)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "InternalError",
			body: replaceSettlementRequest{GroupID: group.ID},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateSettlementsTx(gomock.Any(), gomock.Eq(group.ID)).Times(1).Return(nil, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "BadRequest",
			body: replaceSettlementRequest{GroupID: 0},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateSettlementsTx(gomock.Any(), gomock.Any()).Times(0)
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
			tc.buildStubs(mockStore)

			server := newTestServer(mockStore)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/settlements"
			req := httptest.NewRequest(http.MethodPut, url, bytes.NewReader(data))

			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)
		})
	}
}
