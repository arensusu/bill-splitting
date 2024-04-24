// Code generated by MockGen. DO NOT EDIT.
// Source: bill-splitting/db/sqlc (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -package mockdb -destination db/mock/store.go bill-splitting/db/sqlc Store
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

// CreateGroupInvitation mocks base method.
func (m *MockStore) CreateGroupInvitation(arg0 context.Context, arg1 db.CreateGroupInvitationParams) (db.GroupInvitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroupInvitation", arg0, arg1)
	ret0, _ := ret[0].(db.GroupInvitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroupInvitation indicates an expected call of CreateGroupInvitation.
func (mr *MockStoreMockRecorder) CreateGroupInvitation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroupInvitation", reflect.TypeOf((*MockStore)(nil).CreateGroupInvitation), arg0, arg1)
}

// CreateGroupTx mocks base method.
func (m *MockStore) CreateGroupTx(arg0 context.Context, arg1 db.CreateGroupTxParams) (db.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGroupTx", arg0, arg1)
	ret0, _ := ret[0].(db.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGroupTx indicates an expected call of CreateGroupTx.
func (mr *MockStoreMockRecorder) CreateGroupTx(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGroupTx", reflect.TypeOf((*MockStore)(nil).CreateGroupTx), arg0, arg1)
}

// CreateMember mocks base method.
func (m *MockStore) CreateMember(arg0 context.Context, arg1 db.CreateMemberParams) (db.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMember", arg0, arg1)
	ret0, _ := ret[0].(db.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMember indicates an expected call of CreateMember.
func (mr *MockStoreMockRecorder) CreateMember(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMember", reflect.TypeOf((*MockStore)(nil).CreateMember), arg0, arg1)
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
func (m *MockStore) CreateSettlementsTx(arg0 context.Context, arg1 int32) ([]db.Settlement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSettlementsTx", arg0, arg1)
	ret0, _ := ret[0].([]db.Settlement)
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

// DeleteExpense mocks base method.
func (m *MockStore) DeleteExpense(arg0 context.Context, arg1 int32) error {
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
func (m *MockStore) DeleteGroup(arg0 context.Context, arg1 int32) error {
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

// DeleteGroupInvitation mocks base method.
func (m *MockStore) DeleteGroupInvitation(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGroupInvitation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGroupInvitation indicates an expected call of DeleteGroupInvitation.
func (mr *MockStoreMockRecorder) DeleteGroupInvitation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGroupInvitation", reflect.TypeOf((*MockStore)(nil).DeleteGroupInvitation), arg0, arg1)
}

// DeleteMember mocks base method.
func (m *MockStore) DeleteMember(arg0 context.Context, arg1 db.DeleteMemberParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMember", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMember indicates an expected call of DeleteMember.
func (mr *MockStoreMockRecorder) DeleteMember(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMember", reflect.TypeOf((*MockStore)(nil).DeleteMember), arg0, arg1)
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
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 string) error {
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

// GetExpense mocks base method.
func (m *MockStore) GetExpense(arg0 context.Context, arg1 int32) (db.GetExpenseRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpense", arg0, arg1)
	ret0, _ := ret[0].(db.GetExpenseRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpense indicates an expected call of GetExpense.
func (mr *MockStoreMockRecorder) GetExpense(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpense", reflect.TypeOf((*MockStore)(nil).GetExpense), arg0, arg1)
}

// GetGroup mocks base method.
func (m *MockStore) GetGroup(arg0 context.Context, arg1 int32) (db.Group, error) {
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

// GetGroupInvitation mocks base method.
func (m *MockStore) GetGroupInvitation(arg0 context.Context, arg1 string) (db.GroupInvitation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupInvitation", arg0, arg1)
	ret0, _ := ret[0].(db.GroupInvitation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupInvitation indicates an expected call of GetGroupInvitation.
func (mr *MockStoreMockRecorder) GetGroupInvitation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupInvitation", reflect.TypeOf((*MockStore)(nil).GetGroupInvitation), arg0, arg1)
}

// GetMember mocks base method.
func (m *MockStore) GetMember(arg0 context.Context, arg1 int32) (db.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMember", arg0, arg1)
	ret0, _ := ret[0].(db.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMember indicates an expected call of GetMember.
func (mr *MockStoreMockRecorder) GetMember(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMember", reflect.TypeOf((*MockStore)(nil).GetMember), arg0, arg1)
}

// GetMembership mocks base method.
func (m *MockStore) GetMembership(arg0 context.Context, arg1 db.GetMembershipParams) (db.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMembership", arg0, arg1)
	ret0, _ := ret[0].(db.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembership indicates an expected call of GetMembership.
func (mr *MockStoreMockRecorder) GetMembership(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembership", reflect.TypeOf((*MockStore)(nil).GetMembership), arg0, arg1)
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
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
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

// GetUserByUsername mocks base method.
func (m *MockStore) GetUserByUsername(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStoreMockRecorder) GetUserByUsername(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUserByUsername), arg0, arg1)
}

// ListExpenses mocks base method.
func (m *MockStore) ListExpenses(arg0 context.Context, arg1 int32) ([]db.ListExpensesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListExpenses", arg0, arg1)
	ret0, _ := ret[0].([]db.ListExpensesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExpenses indicates an expected call of ListExpenses.
func (mr *MockStoreMockRecorder) ListExpenses(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExpenses", reflect.TypeOf((*MockStore)(nil).ListExpenses), arg0, arg1)
}

// ListGroups mocks base method.
func (m *MockStore) ListGroups(arg0 context.Context, arg1 string) ([]db.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroups", arg0, arg1)
	ret0, _ := ret[0].([]db.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroups indicates an expected call of ListGroups.
func (mr *MockStoreMockRecorder) ListGroups(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockStore)(nil).ListGroups), arg0, arg1)
}

// ListMembersOfGroup mocks base method.
func (m *MockStore) ListMembersOfGroup(arg0 context.Context, arg1 int32) ([]db.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMembersOfGroup", arg0, arg1)
	ret0, _ := ret[0].([]db.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMembersOfGroup indicates an expected call of ListMembersOfGroup.
func (mr *MockStoreMockRecorder) ListMembersOfGroup(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMembersOfGroup", reflect.TypeOf((*MockStore)(nil).ListMembersOfGroup), arg0, arg1)
}

// ListNonSettledExpenses mocks base method.
func (m *MockStore) ListNonSettledExpenses(arg0 context.Context, arg1 int32) ([]db.ListNonSettledExpensesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNonSettledExpenses", arg0, arg1)
	ret0, _ := ret[0].([]db.ListNonSettledExpensesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNonSettledExpenses indicates an expected call of ListNonSettledExpenses.
func (mr *MockStoreMockRecorder) ListNonSettledExpenses(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNonSettledExpenses", reflect.TypeOf((*MockStore)(nil).ListNonSettledExpenses), arg0, arg1)
}

// ListSettlements mocks base method.
func (m *MockStore) ListSettlements(arg0 context.Context, arg1 int32) ([]db.Settlement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSettlements", arg0, arg1)
	ret0, _ := ret[0].([]db.Settlement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSettlements indicates an expected call of ListSettlements.
func (mr *MockStoreMockRecorder) ListSettlements(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSettlements", reflect.TypeOf((*MockStore)(nil).ListSettlements), arg0, arg1)
}

// SummarizeExpensesWithinDate mocks base method.
func (m *MockStore) SummarizeExpensesWithinDate(arg0 context.Context, arg1 db.SummarizeExpensesWithinDateParams) ([]db.SummarizeExpensesWithinDateRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SummarizeExpensesWithinDate", arg0, arg1)
	ret0, _ := ret[0].([]db.SummarizeExpensesWithinDateRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SummarizeExpensesWithinDate indicates an expected call of SummarizeExpensesWithinDate.
func (mr *MockStoreMockRecorder) SummarizeExpensesWithinDate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SummarizeExpensesWithinDate", reflect.TypeOf((*MockStore)(nil).SummarizeExpensesWithinDate), arg0, arg1)
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
