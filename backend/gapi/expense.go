package gapi

import (
	"bill-splitting/model"
	"bill-splitting/proto"
	"bill-splitting/utils"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateExpense(ctx context.Context, req *proto.CreateExpenseRequest) (*proto.CreateExpenseResponse, error) {
	payload, err := s.authorize(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to authorize: %w", err)
	}

	if req.GroupId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid group id")
	}

	group, err := s.store.GetGroup(uint(req.GroupId))
	if err != nil {
		return nil, fmt.Errorf("group not found: %w", err)
	}

	user, err := s.store.GetUserByLineID(payload.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	member, err := s.store.GetMembership(uint(req.GroupId), user.ID)
	if err != nil {
		return nil, err
	}

	amount, err := utils.GetExchangeAmount(req.OriginCurrency, group.Currency, req.OriginAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to get exchanged amount: %w", err)
	}

	err = s.store.CreateExpense(&model.Expense{
		Member:           *member,
		Category:         req.Category,
		ConvertedAmount:  amount,
		OriginalAmount:   req.OriginAmount,
		OriginalCurrency: req.OriginCurrency,
		Date:             time.Now(),
		Description:      req.Description,
	})
	if err != nil {
		return nil, err
	}
	return &proto.CreateExpenseResponse{}, nil
}

func (s *Server) ListExpense(ctx context.Context, req *proto.ListExpenseRequest) (*proto.ListExpenseResponse, error) {
	payload, err := s.authorize(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.store.GetUserByLineID(payload.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	_, err = s.store.GetMembership(uint(req.GroupId), user.ID)
	if err != nil {
		return nil, err
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %w", err)
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end date: %w", err)
	}

	expenses, err := s.store.ListExpensesWithinDate(uint(req.GroupId), startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to list expenses: %w", err)
	}

	protoExpenses := make([]*proto.Expense, len(expenses))
	for i, expense := range expenses {
		protoExpenses[i] = &proto.Expense{
			Id:          uint32(expense.ID),
			Category:    expense.Category,
			Description: expense.Description,
			Amount:      expense.ConvertedAmount,
			Date:        expense.Date.Format("2006-01-02"),
		}
	}

	return &proto.ListExpenseResponse{
		Expenses: protoExpenses,
	}, nil
}

func (s *Server) ListExpenseSummary(ctx context.Context, req *proto.ListExpenseSummaryRequest) (*proto.ListExpenseSummaryResponse, error) {
	payload, err := s.authorize(ctx)
	if err != nil {
		return nil, err
	}

	user, err := s.store.GetUserByLineID(payload.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	_, err = s.store.GetMembership(uint(req.GroupId), user.ID)
	if err != nil {
		return nil, err
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %w", err)
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end date: %w", err)
	}

	summary, err := s.store.SummarizeExpensesWithinDate(uint(req.GroupId), startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to summarize expenses: %w", err)
	}

	protoSummary := make([]*proto.ExpenseSummary, len(summary))
	for i, s := range summary {
		protoSummary[i] = &proto.ExpenseSummary{
			Category: s.Category,
			Total:    s.Total,
		}
	}

	return &proto.ListExpenseSummaryResponse{
		Summaries: protoSummary,
	}, nil
}
