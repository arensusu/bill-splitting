// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import (
	"time"
)

type Expense struct {
	ID          int32     `json:"id"`
	MemberID    int32     `json:"member_id"`
	Amount      string    `json:"amount"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	IsSettled   bool      `json:"is_settled"`
}

type Group struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type GroupInvitation struct {
	Code    string `json:"code"`
	GroupID int32  `json:"group_id"`
}

type Member struct {
	ID      int32  `json:"id"`
	GroupID int32  `json:"group_id"`
	UserID  string `json:"user_id"`
}

type Settlement struct {
	PayerID int32  `json:"payer_id"`
	PayeeID int32  `json:"payee_id"`
	Amount  string `json:"amount"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
