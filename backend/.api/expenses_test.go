package api

import (
	mockdb "bill-splitting/db/mock"
	db "bill-splitting/db/sqlc"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateExpenseAPI(t *testing.T) {
	user := randomUser()
	group := randomGroup()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCases := []struct {
		name          string
		groupID       int32
		body          createExpenseJSONRequest
		buildStub     func(t *testing.T, mockStore *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			groupID: group.ID,
			body: createExpenseJSONRequest{
				Amount:      "100",
				Description: "test",
				Date:        "2022-01-01",
			},
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetMembership(gomock.Any(), gomock.Any()).Times(1).Return(db.Member{
					GroupID: group.ID,
					UserID:  user.ID,
				}, nil)
				mockStore.EXPECT().CreateExpense(gomock.Any(), gomock.Any()).Times(1).Return(db.Expense{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code, recorder.Body)
			},
		},
		{
			name:    "InternalErrorOfGetGroupMember",
			groupID: group.ID,
			body: createExpenseJSONRequest{
				Amount:      "100",
				Description: "test",
				Date:        "2022-01-01",
			},
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetMembership(gomock.Any(), gomock.Any()).Times(1).Return(db.Member{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:    "InternalErrorOfCreateExpenseTx",
			groupID: group.ID,
			body: createExpenseJSONRequest{
				Amount:      "100",
				Description: "test",
				Date:        "2022-01-01",
			},
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetMembership(gomock.Any(), gomock.Any()).Times(1).Return(db.Member{
					GroupID: group.ID,
					UserID:  user.ID,
				}, nil)
				mockStore.EXPECT().CreateExpense(gomock.Any(), gomock.Any()).Times(1).Return(db.Expense{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:    "BadRequest",
			groupID: 0,
			body: createExpenseJSONRequest{
				Amount:      "100",
				Description: "test",
				Date:        "2022-01-01",
			},
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetMember(gomock.Any(), gomock.Any()).Times(0)
				mockStore.EXPECT().CreateExpense(gomock.Any(), gomock.Any()).Times(0)
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

			url := fmt.Sprintf("/api/v1/groups/%d/expenses", tc.groupID)
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))

			addAuthentication(t, req, server.tokenMaker, user.ID, time.Minute)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListExpensesAPI(t *testing.T) {
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
				mockStore.EXPECT().ListExpenses(gomock.Any(), gomock.Eq(group.ID)).Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:    "InternalError",
			groupID: group.ID,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().ListExpenses(gomock.Any(), gomock.Eq(group.ID)).Times(1).Return(nil, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:    "InvalidID",
			groupID: 0,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().ListExpenses(gomock.Any(), gomock.Any()).Times(0)
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

			url := fmt.Sprintf("/api/v1/groups/%d/expenses", tc.groupID)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			addAuthentication(t, req, server.tokenMaker, user.ID, time.Minute)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)
		})
	}
}
