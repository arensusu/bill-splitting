package db

import (
	"context"
	"database/sql"
)

type Store interface {
	Querier
	CreateExpenseTx(ctx context.Context, arg CreateExpenseTxParams) (*CreateExpenseTxResult, error)
	CreateSettlementsTx(ctx context.Context, groupID int64) (*CreateSettlementTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
