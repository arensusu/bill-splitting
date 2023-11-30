// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"context"
)

type Querier interface {
	CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error)
	CreateGroup(ctx context.Context, name string) (Group, error)
	CreateGroupMember(ctx context.Context, arg CreateGroupMemberParams) (GroupMember, error)
	CreateSettlement(ctx context.Context, arg CreateSettlementParams) (Settlement, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateUserExpense(ctx context.Context, arg CreateUserExpenseParams) (UserExpense, error)
	DeleteExpense(ctx context.Context, id int64) error
	DeleteGroup(ctx context.Context, id int64) error
	DeleteGroupMember(ctx context.Context, arg DeleteGroupMemberParams) error
	DeleteSettlement(ctx context.Context, arg DeleteSettlementParams) error
	DeleteUser(ctx context.Context, id int64) error
	DeleteUserExpense(ctx context.Context, arg DeleteUserExpenseParams) error
	GetExpense(ctx context.Context, id int64) (Expense, error)
	GetGroup(ctx context.Context, id int64) (Group, error)
	GetGroupMember(ctx context.Context, arg GetGroupMemberParams) (GroupMember, error)
	GetSettlement(ctx context.Context, arg GetSettlementParams) (Settlement, error)
	GetUser(ctx context.Context, id int64) (User, error)
	GetUserExpense(ctx context.Context, arg GetUserExpenseParams) (UserExpense, error)
	ListExpenses(ctx context.Context, groupID int64) ([]Expense, error)
	ListGroupMembers(ctx context.Context, groupID int64) ([]GroupMember, error)
	ListGroupSettlements(ctx context.Context, groupID int64) ([]Settlement, error)
	ListNonSettledExpenses(ctx context.Context, groupID int64) ([]Expense, error)
	ListUserExpenses(ctx context.Context, expenseID int64) ([]UserExpense, error)
	UpdateExpense(ctx context.Context, arg UpdateExpenseParams) (Expense, error)
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) (Group, error)
	UpdateSettlement(ctx context.Context, arg UpdateSettlementParams) (Settlement, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserExpense(ctx context.Context, arg UpdateUserExpenseParams) (UserExpense, error)
}

var _ Querier = (*Queries)(nil)