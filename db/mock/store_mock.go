// Code generated by MockGen. DO NOT EDIT.
// Source: bill-splitting/db/sqlc (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -package mockdb -destination db/mock/store_mock.go bill-splitting/db/sqlc Store
//
// Package mockdb is a generated GoMock package.
package mockdb

import (
	db "bill-splitting/db/sqlc"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateExpense mocks base method.
func (m *MockStore) CreateExpense(arg0 context.Context, arg1 db.CreateExpenseParams) (db.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExpense", arg0, arg1)
	ret0, _ := ret[0].(db.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExpense indicates an expected call of CreateExpense.
func (mr *MockStoreMockRecorder) CreateExpense(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExpense", reflect.TypeOf((*MockStore)(nil).CreateExpense), arg0, arg1)
}

// CreateExpenseTx mocks base method.
func (m *MockStore) CreateExpenseTx(arg0 context.Context, arg1 db.CreateExpenseTxParams) (*db.CreateExpenseTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExpenseTx", arg0, arg1)
	ret0, _ := ret[0].(*db.CreateExpenseTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExpenseTx indicates an expected call of CreateExpenseTx.
func (mr *MockStoreMockRecorder) CreateExpenseTx(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExpenseTx", reflect.TypeOf((*MockStore)(nil).CreateExpenseTx), arg0, arg1)
}

// CreateGroup mocks base method.
func (m *MockStore) CreateGroup(arg0 context.Context, arg1 string) (db.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroup", arg0, arg1)
	ret0, _ := ret[0].(db.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroup indicates an expected call of CreateGroup.
func (mr *MockStoreMockRecorder) CreateGroup(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroup", reflect.TypeOf((*MockStore)(nil).CreateGroup), arg0, arg1)
}

// CreateGroupMember mocks base method.
func (m *MockStore) CreateGroupMember(arg0 context.Context, arg1 db.CreateGroupMemberParams) (db.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroupMember", arg0, arg1)
	ret0, _ := ret[0].(db.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroupMember indicates an expected call of CreateGroupMember.
func (mr *MockStoreMockRecorder) CreateGroupMember(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroupMember", reflect.TypeOf((*MockStore)(nil).CreateGroupMember), arg0, arg1)
}

// CreateSettlement mocks base method.
func (m *MockStore) CreateSettlement(arg0 context.Context, arg1 db.CreateSettlementParams) (db.Settlement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSettlement", arg0, arg1)
	ret0, _ := ret[0].(db.Settlement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSettlement indicates an expected call of CreateSettlement.
func (mr *MockStoreMockRecorder) CreateSettlement(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSettlement", reflect.TypeOf((*MockStore)(nil).CreateSettlement), arg0, arg1)
}

// CreateSettlementsTx mocks base method.
func (m *MockStore) CreateSettlementsTx(arg0 context.Context, arg1 int64) (*db.CreateSettlementTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSettlementsTx", arg0, arg1)
	ret0, _ := ret[0].(*db.CreateSettlementTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSettlementsTx indicates an expected call of CreateSettlementsTx.
func (mr *MockStoreMockRecorder) CreateSettlementsTx(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSettlementsTx", reflect.TypeOf((*MockStore)(nil).CreateSettlementsTx), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateUserExpense mocks base method.
func (m *MockStore) CreateUserExpense(arg0 context.Context, arg1 db.CreateUserExpenseParams) (db.UserExpense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserExpense", arg0, arg1)
	ret0, _ := ret[0].(db.UserExpense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserExpense indicates an expected call of CreateUserExpense.
func (mr *MockStoreMockRecorder) CreateUserExpense(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserExpense", reflect.TypeOf((*MockStore)(nil).CreateUserExpense), arg0, arg1)
}

// DeleteExpense mocks base method.
func (m *MockStore) DeleteExpense(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExpense", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteExpense indicates an expected call of DeleteExpense.
func (mr *MockStoreMockRecorder) DeleteExpense(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExpense", reflect.TypeOf((*MockStore)(nil).DeleteExpense), arg0, arg1)
}

// DeleteGroup mocks base method.
func (m *MockStore) DeleteGroup(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroup indicates an expected call of DeleteGroup.
func (mr *MockStoreMockRecorder) DeleteGroup(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroup", reflect.TypeOf((*MockStore)(nil).DeleteGroup), arg0, arg1)
}

// DeleteGroupMember mocks base method.
func (m *MockStore) DeleteGroupMember(arg0 context.Context, arg1 db.DeleteGroupMemberParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroupMember", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroupMember indicates an expected call of DeleteGroupMember.
func (mr *MockStoreMockRecorder) DeleteGroupMember(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupMember", reflect.TypeOf((*MockStore)(nil).DeleteGroupMember), arg0, arg1)
}

// DeleteSettlement mocks base method.
func (m *MockStore) DeleteSettlement(arg0 context.Context, arg1 db.DeleteSettlementParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSettlement", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSettlement indicates an expected call of DeleteSettlement.
func (mr *MockStoreMockRecorder) DeleteSettlement(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSettlement", reflect.TypeOf((*MockStore)(nil).DeleteSettlement), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// DeleteUserExpense mocks base method.
func (m *MockStore) DeleteUserExpense(arg0 context.Context, arg1 db.DeleteUserExpenseParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserExpense", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserExpense indicates an expected call of DeleteUserExpense.
func (mr *MockStoreMockRecorder) DeleteUserExpense(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserExpense", reflect.TypeOf((*MockStore)(nil).DeleteUserExpense), arg0, arg1)
}

// GetExpense mocks base method.
func (m *MockStore) GetExpense(arg0 context.Context, arg1 int64) (db.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpense", arg0, arg1)
	ret0, _ := ret[0].(db.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpense indicates an expected call of GetExpense.
func (mr *MockStoreMockRecorder) GetExpense(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpense", reflect.TypeOf((*MockStore)(nil).GetExpense), arg0, arg1)
}

// GetGroup mocks base method.
func (m *MockStore) GetGroup(arg0 context.Context, arg1 int64) (db.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroup", arg0, arg1)
	ret0, _ := ret[0].(db.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockStoreMockRecorder) GetGroup(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockStore)(nil).GetGroup), arg0, arg1)
}

// GetGroupMember mocks base method.
func (m *MockStore) GetGroupMember(arg0 context.Context, arg1 db.GetGroupMemberParams) (db.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMember", arg0, arg1)
	ret0, _ := ret[0].(db.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupMember indicates an expected call of GetGroupMember.
func (mr *MockStoreMockRecorder) GetGroupMember(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMember", reflect.TypeOf((*MockStore)(nil).GetGroupMember), arg0, arg1)
}

// GetSettlement mocks base method.
func (m *MockStore) GetSettlement(arg0 context.Context, arg1 db.GetSettlementParams) (db.Settlement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettlement", arg0, arg1)
	ret0, _ := ret[0].(db.Settlement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettlement indicates an expected call of GetSettlement.
func (mr *MockStoreMockRecorder) GetSettlement(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettlement", reflect.TypeOf((*MockStore)(nil).GetSettlement), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserExpense mocks base method.
func (m *MockStore) GetUserExpense(arg0 context.Context, arg1 db.GetUserExpenseParams) (db.UserExpense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserExpense", arg0, arg1)
	ret0, _ := ret[0].(db.UserExpense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserExpense indicates an expected call of GetUserExpense.
func (mr *MockStoreMockRecorder) GetUserExpense(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserExpense", reflect.TypeOf((*MockStore)(nil).GetUserExpense), arg0, arg1)
}

// ListGroupExpenses mocks base method.
func (m *MockStore) ListGroupExpenses(arg0 context.Context, arg1 int64) ([]db.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupExpenses", arg0, arg1)
	ret0, _ := ret[0].([]db.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupExpenses indicates an expected call of ListGroupExpenses.
func (mr *MockStoreMockRecorder) ListGroupExpenses(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupExpenses", reflect.TypeOf((*MockStore)(nil).ListGroupExpenses), arg0, arg1)
}

// ListGroupMembers mocks base method.
func (m *MockStore) ListGroupMembers(arg0 context.Context, arg1 int64) ([]db.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupMembers", arg0, arg1)
	ret0, _ := ret[0].([]db.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupMembers indicates an expected call of ListGroupMembers.
func (mr *MockStoreMockRecorder) ListGroupMembers(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupMembers", reflect.TypeOf((*MockStore)(nil).ListGroupMembers), arg0, arg1)
}

// ListGroupSettlements mocks base method.
func (m *MockStore) ListGroupSettlements(arg0 context.Context, arg1 int64) ([]db.Settlement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupSettlements", arg0, arg1)
	ret0, _ := ret[0].([]db.Settlement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupSettlements indicates an expected call of ListGroupSettlements.
func (mr *MockStoreMockRecorder) ListGroupSettlements(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupSettlements", reflect.TypeOf((*MockStore)(nil).ListGroupSettlements), arg0, arg1)
}

// ListNonSettledGroupExpenses mocks base method.
func (m *MockStore) ListNonSettledGroupExpenses(arg0 context.Context, arg1 int64) ([]db.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNonSettledGroupExpenses", arg0, arg1)
	ret0, _ := ret[0].([]db.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNonSettledGroupExpenses indicates an expected call of ListNonSettledGroupExpenses.
func (mr *MockStoreMockRecorder) ListNonSettledGroupExpenses(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNonSettledGroupExpenses", reflect.TypeOf((*MockStore)(nil).ListNonSettledGroupExpenses), arg0, arg1)
}

// ListUserExpenses mocks base method.
func (m *MockStore) ListUserExpenses(arg0 context.Context, arg1 int64) ([]db.UserExpense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserExpenses", arg0, arg1)
	ret0, _ := ret[0].([]db.UserExpense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserExpenses indicates an expected call of ListUserExpenses.
func (mr *MockStoreMockRecorder) ListUserExpenses(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserExpenses", reflect.TypeOf((*MockStore)(nil).ListUserExpenses), arg0, arg1)
}

// UpdateExpense mocks base method.
func (m *MockStore) UpdateExpense(arg0 context.Context, arg1 db.UpdateExpenseParams) (db.Expense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExpense", arg0, arg1)
	ret0, _ := ret[0].(db.Expense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateExpense indicates an expected call of UpdateExpense.
func (mr *MockStoreMockRecorder) UpdateExpense(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExpense", reflect.TypeOf((*MockStore)(nil).UpdateExpense), arg0, arg1)
}

// UpdateGroup mocks base method.
func (m *MockStore) UpdateGroup(arg0 context.Context, arg1 db.UpdateGroupParams) (db.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroup", arg0, arg1)
	ret0, _ := ret[0].(db.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGroup indicates an expected call of UpdateGroup.
func (mr *MockStoreMockRecorder) UpdateGroup(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroup", reflect.TypeOf((*MockStore)(nil).UpdateGroup), arg0, arg1)
}

// UpdateSettlement mocks base method.
func (m *MockStore) UpdateSettlement(arg0 context.Context, arg1 db.UpdateSettlementParams) (db.Settlement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSettlement", arg0, arg1)
	ret0, _ := ret[0].(db.Settlement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSettlement indicates an expected call of UpdateSettlement.
func (mr *MockStoreMockRecorder) UpdateSettlement(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSettlement", reflect.TypeOf((*MockStore)(nil).UpdateSettlement), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}

// UpdateUserExpense mocks base method.
func (m *MockStore) UpdateUserExpense(arg0 context.Context, arg1 db.UpdateUserExpenseParams) (db.UserExpense, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserExpense", arg0, arg1)
	ret0, _ := ret[0].(db.UserExpense)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserExpense indicates an expected call of UpdateUserExpense.
func (mr *MockStoreMockRecorder) UpdateUserExpense(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserExpense", reflect.TypeOf((*MockStore)(nil).UpdateUserExpense), arg0, arg1)
}
