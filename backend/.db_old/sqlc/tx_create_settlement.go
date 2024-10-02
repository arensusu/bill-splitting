package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
)

func (s *SQLStore) CreateSettlementsTx(ctx context.Context, groupID int32) ([]Settlement, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return []Settlement{}, err
	}
	defer tx.Rollback()

	q := New(tx)
	_, err = q.GetGroup(ctx, groupID)
	if err != nil {
		return []Settlement{}, fmt.Errorf("create settlements tx: %w", err)
	}

	expenses, err := q.ListNonSettledExpenses(ctx, groupID)
	if err != nil {
		return []Settlement{}, fmt.Errorf("create settlements tx: %w", err)
	}

	members, err := q.ListMembersOfGroup(ctx, groupID)
	if err != nil {
		return []Settlement{}, fmt.Errorf("create settlements tx: %w", err)
	}

	settleAmounts := map[int32]map[int32]int64{}
	for _, expense := range expenses {
		expenseAmount, err := strconv.ParseInt(expense.Amount, 10, 64)
		if err != nil {
			return []Settlement{}, fmt.Errorf("create settlements tx: %w", err)
		}
		share := expenseAmount / int64(len(members))

		for _, member := range members {
			if expense.MemberID == member.ID {
				continue
			}

			if member.ID < expense.MemberID {
				if _, ok := settleAmounts[member.ID]; !ok {
					settleAmounts[member.ID] = map[int32]int64{expense.MemberID: 0}
				}
				settleAmounts[member.ID][expense.MemberID] += share
			} else {
				if _, ok := settleAmounts[expense.MemberID]; !ok {
					settleAmounts[expense.MemberID] = map[int32]int64{member.ID: 0}
				}
				settleAmounts[expense.MemberID][member.ID] -= share
			}
		}

		_, err = q.UpdateExpense(ctx, UpdateExpenseParams{
			ID:          expense.ID,
			MemberID:    expense.MemberID,
			Amount:      expense.Amount,
			Description: expense.Description,
			Date:        expense.Date,
			IsSettled:   true,
		})
		if err != nil {
			return []Settlement{}, fmt.Errorf("create settlements tx: %w", err)
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
				PayerID: payerID,
				PayeeID: payeeID,
			})

			if err != nil {
				if err != sql.ErrNoRows {
					return []Settlement{}, fmt.Errorf("create settlements tx: %w", err)
				} else {
					_, err = q.CreateSettlement(ctx, CreateSettlementParams{
						PayerID: payerID,
						PayeeID: payeeID,
						Amount:  strconv.FormatInt(amount, 10),
					})
					if err != nil {
						return []Settlement{}, fmt.Errorf("create settlements tx: %w", err)
					}
				}
			} else {
				prevAmount, err := strconv.ParseInt(prevSettlement.Amount, 10, 64)
				if err != nil {
					return []Settlement{}, fmt.Errorf("create settlements tx: %w", err)
				}
				amount += prevAmount
				_, err = q.UpdateSettlement(ctx, UpdateSettlementParams{
					PayerID: payerID,
					PayeeID: payeeID,
					Amount:  strconv.FormatInt(amount, 10),
				})
				if err != nil {
					return []Settlement{}, fmt.Errorf("create settlements tx: %w", err)
				}
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return []Settlement{}, fmt.Errorf("create settlements tx: %w", err)
	}

	settlements, err := s.ListSettlements(ctx, groupID)
	if err != nil {
		return []Settlement{}, fmt.Errorf("create settlements tx: %w", err)
	}

	return settlements, nil
}
