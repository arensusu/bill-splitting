package gapi

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/proto"
	"bill-splitting/utils"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateExpenseSummaryChart(ctx context.Context, req *proto.CreateExpenseSummaryChartRequest) (*proto.CreateExpenseSummaryChartResponse, error) {
	payload, err := s.authorize(ctx)
	if err != nil {
		return nil, err
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, err
	}

	_, err = s.store.GetMembership(ctx, db.GetMembershipParams{
		GroupID: req.GroupId,
		UserID:  payload.UserID,
	})
	if err != nil {
		return nil, err
	}

	expenses, err := s.store.ListExpensesWithinDate(ctx, db.ListExpensesWithinDateParams{
		GroupID:   req.GroupId,
		StartTime: startDate,
		EndTime:   endDate,
	})
	if err != nil {
		return nil, err
	}

	summary, err := s.store.SummarizeExpensesWithinDate(ctx, db.SummarizeExpensesWithinDateParams{
		GroupID:   req.GroupId,
		StartTime: startDate,
		EndTime:   endDate,
	})
	if err != nil {
		return nil, err
	}

	dataStr, err := json.Marshal(summary)
	if err != nil {
		return nil, err
	}

	hasher := sha256.New()
	hasher.Write(dataStr)
	hashBytes := hasher.Sum(nil)

	values := make([]float64, len(summary))
	legends := make([]string, len(summary))
	total := 0.0
	for i, v := range summary {
		value, _ := strconv.ParseFloat(v.Total, 64)

		total += value
		values[i] = value
		legends[i] = v.Category.String
	}

	data := make([][4]string, len(expenses))
	for i, expense := range expenses {
		data[i] = [4]string{
			expense.Date.Format("2006/01/02"),
			expense.Category.String,
			expense.Description,
			strings.Split(expense.Amount, ".")[0],
		}
	}

	title := fmt.Sprintf("%s ~ %s", req.StartDate, req.EndDate)
	subtitle := fmt.Sprintf("Total: %.0f", total)
	path := fmt.Sprintf("/var/images/%x.html", hashBytes)

	err = utils.CreatePieChart(values, legends, title, subtitle, data, path)
	if err != nil {
		return nil, fmt.Errorf("failed to create pie chart: %w", err)
	}

	return &proto.CreateExpenseSummaryChartResponse{
		Url: fmt.Sprintf("%x.html", hashBytes),
	}, nil
}

func (s *Server) CreateExpense(ctx context.Context, req *proto.CreateExpenseRequest) (*proto.CreateExpenseResponse, error) {
	payload, err := s.authorize(ctx)
	if err != nil {
		return nil, err
	}

	if req.GroupId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid group id")
	}

	group, err := s.store.GetGroup(ctx, req.GroupId)
	if err != nil {
		return nil, fmt.Errorf("group not found: %w", err)
	}

	member, err := s.store.GetMembership(ctx, db.GetMembershipParams{
		GroupID: req.GroupId,
		UserID:  payload.UserID,
	})
	if err != nil {
		return nil, err
	}

	amount, err := utils.GetExchangeAmount(req.OriginCurrency, group.Currency.String, req.OriginAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to get exchanged amount: %w", err)
	}

	expense, err := s.store.CreateExpense(ctx, db.CreateExpenseParams{
		MemberID:       member.ID,
		Amount:         fmt.Sprint(amount),
		OriginCurrency: sql.NullString{String: req.OriginCurrency, Valid: true},
		OriginAmount:   sql.NullString{String: fmt.Sprint(req.OriginAmount), Valid: true},
		Description:    req.Description,
		Category:       sql.NullString{String: req.Category, Valid: req.Category != ""},
		Date:           time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return &proto.CreateExpenseResponse{
		Id:          expense.ID,
		Category:    expense.Category.String,
		Date:        expense.Date.Format("2006-01-02"),
		Amount:      amount,
		Description: expense.Description,
	}, nil
}
