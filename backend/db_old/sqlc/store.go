package db

import (
	"context"
	"database/sql"
)

type Store interface {
	Querier
	CreateGroupTx(ctx context.Context, arg CreateGroupTxParams) (Group, error)
	CreateSettlementsTx(ctx context.Context, groupID int32) ([]Settlement, error)
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
