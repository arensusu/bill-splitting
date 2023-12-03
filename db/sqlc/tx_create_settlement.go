package db

import (
	"context"
	"database/sql"
	"fmt"
)

type CreateSettlementTxResult struct {
	Settlements []Settlement `json:"settlements"`
}

func (s *SQLStore) CreateSettlementsTx(ctx context.Context, groupID int64) (*CreateSettlementTxResult, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	q := New(tx)
	_, err = q.GetGroup(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("create settlements tx: %w", err)
	}

	expenses, err := q.ListNonSettledExpenses(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("create settlements tx: %w", err)
	}

	settleAmounts := map[int64]map[int64]int64{}
	for _, expense := range expenses {
		userExpenses, err := q.ListUserExpenses(ctx, expense.ID)
		if err != nil {
			return nil, fmt.Errorf("create settlements tx: %w", err)
		}

		for _, userExpense := range userExpenses {
			if expense.PayerID == userExpense.UserID {
				continue
			}

			if userExpense.UserID < expense.PayerID {
				if _, ok := settleAmounts[userExpense.UserID]; !ok {
					settleAmounts[userExpense.UserID] = map[int64]int64{expense.PayerID: 0}
				}
				settleAmounts[userExpense.UserID][expense.PayerID] += userExpense.Share
			} else {
				if _, ok := settleAmounts[expense.PayerID]; !ok {
					settleAmounts[expense.PayerID] = map[int64]int64{userExpense.UserID: 0}
				}
				settleAmounts[expense.PayerID][userExpense.UserID] -= userExpense.Share
			}
		}

		_, err = q.UpdateExpense(ctx, UpdateExpenseParams{
			ID:          expense.ID,
			GroupID:     expense.GroupID,
			PayerID:     expense.PayerID,
			Amount:      expense.Amount,
			Description: expense.Description,
			Date:        expense.Date,
			IsSettled:   true,
		})
		if err != nil {
			return nil, fmt.Errorf("create settlements tx: %w", err)
		}
	}

	for payerID, payees := range settleAmounts {
		for payeeID, amount := range payees {
			if amount == 0 {
				continue
			}

			if amount < 0 {
				payerID, payeeID = payeeID, payerID
				amount *= -1
			}

			prevSettlement, err := q.GetSettlement(ctx, GetSettlementParams{
				GroupID: groupID,
				PayerID: payerID,
				PayeeID: payeeID,
			})

			if err != nil {
				if err != sql.ErrNoRows {
					return nil, fmt.Errorf("create settlements tx: %w", err)
				} else {
					_, err = q.CreateSettlement(ctx, CreateSettlementParams{
						GroupID: groupID,
						PayerID: payerID,
						PayeeID: payeeID,
						Amount:  amount,
					})
					if err != nil {
						return nil, fmt.Errorf("create settlements tx: %w", err)
					}
				}
			} else {
				if !(prevSettlement.IsConfirmed) {
					amount += prevSettlement.Amount
				}
				_, err = q.UpdateSettlement(ctx, UpdateSettlementParams{
					GroupID: groupID,
					PayerID: payerID,
					PayeeID: payeeID,
					Amount:  amount,
				})
				if err != nil {
					return nil, fmt.Errorf("create settlements tx: %w", err)
				}
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("create settlements tx: %w", err)
	}

	settlements, err := s.ListSettlements(ctx, groupID)
	if err != nil {
		return nil, fmt.Errorf("create settlements tx: %w", err)
	}

	return &CreateSettlementTxResult{
		Settlements: settlements,
	}, nil
}
