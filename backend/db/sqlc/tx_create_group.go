package db

import (
	"context"
	"database/sql"
)

type CreateGroupTxParams struct {
	Name   string `json:"name"`
	LineId string
	UserID string `json:"userId"`
}

func (s *SQLStore) CreateGroupTx(ctx context.Context, arg CreateGroupTxParams) (Group, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return Group{}, err
	}
	defer tx.Rollback()

	q := New(tx)

	group, err := q.CreateGroup(ctx, CreateGroupParams{
		Name: arg.Name,
		LineID: sql.NullString{
			String: arg.LineId,
			Valid:  arg.LineId != "",
		},
	})
	if err != nil {
		return Group{}, err
	}

	_, err = q.CreateMember(ctx, CreateMemberParams{
		GroupID: group.ID,
		UserID:  arg.UserID,
	})
	if err != nil {
		return Group{}, err
	}

	err = tx.Commit()
	if err != nil {
		return Group{}, err
	}

	return group, nil
}
