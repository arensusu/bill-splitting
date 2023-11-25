// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"time"
)

type Expense struct {
	ID          int64     `json:"id"`
	GroupID     int64     `json:"group_id"`
	PayerID     int64     `json:"payer_id"`
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type Group struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type GroupMember struct {
	GroupID   int64     `json:"group_id"`
	UserID    int64     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Settlement struct {
	ID      int64     `json:"id"`
	PayerID int64     `json:"payer_id"`
	PayeeID int64     `json:"payee_id"`
	Amount  int64     `json:"amount"`
	Date    time.Time `json:"date"`
}

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserExpense struct {
	ExpenseID int64 `json:"expense_id"`
	UserID    int64 `json:"user_id"`
	Share     int64 `json:"share"`
}
