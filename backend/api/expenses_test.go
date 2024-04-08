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
		body          createExpenseRequest
		buildStub     func(t *testing.T, mockStore *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: createExpenseRequest{
				GroupID:     1,
				Amount:      100,
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
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name: "InternalErrorOfGetGroupMember",
			body: createExpenseRequest{
				GroupID:     group.ID,
				Amount:      100,
				Description: "test",
				Date:        "2022-01-01",
			},
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetMembership(gomock.Any(), gomock.Any()).Times(1).Return(db.Member{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "InternalErrorOfCreateExpenseTx",
			body: createExpenseRequest{
				GroupID:     group.ID,
				Amount:      100,
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
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name: "BadRequest",
			body: createExpenseRequest{
				GroupID:     0,
				Amount:      100,
				Description: "test",
				Date:        "2022-01-01",
			},
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().GetMember(gomock.Any(), gomock.Any()).Times(0)
				mockStore.EXPECT().CreateExpense(gomock.Any(), gomock.Any()).Times(0)
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

			server := newTestServer(mockStore)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/expenses"
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
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			groupID: group.ID,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().ListExpenses(gomock.Any(), gomock.Eq(group.ID)).Times(1)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
		{
			name:    "InternalError",
			groupID: group.ID,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().ListExpenses(gomock.Any(), gomock.Eq(group.ID)).Times(1).Return(nil, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoder.Code)
			},
		},
		{
			name:    "InvalidID",
			groupID: 0,
			buildStub: func(t *testing.T, mockStore *mockdb.MockStore) {
				mockStore.EXPECT().ListExpenses(gomock.Any(), gomock.Any()).Times(0)
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

			server := newTestServer(mockStore)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/expenses/%d", tc.groupID)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			addAuthentication(t, req, server.tokenMaker, user.ID, time.Minute)
			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)
		})
	}
}
