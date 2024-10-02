package gapi

import (
	"bill-splitting/model"
	"bill-splitting/proto"
	"bill-splitting/utils"
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateExpenseSummaryChart(ctx context.Context, req *proto.CreateExpenseSummaryChartRequest) (*proto.CreateExpenseSummaryChartResponse, error) {
	payload, err := s.authorize(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to authorize: %w", err)
	}

	user, err := s.store.GetUserByLineID(payload.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse start date: %w", err)
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse end date: %w", err)
	}

	endDate = endDate.AddDate(0, 0, 1)

	_, err = s.store.GetMembership(uint(req.GroupId), user.ID)
	if err != nil {
		return nil, errors.New("not a member of the group")
	}

	expenses, err := s.store.ListExpensesWithinDate(uint(req.GroupId), startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to list expenses: %w", err)
	}

	summary, err := s.store.SummarizeExpensesWithinDate(uint(req.GroupId), startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to summarize expenses: %w", err)
	}

	dataStr, err := json.Marshal(summary)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal summary: %w", err)
	}

	hasher := sha256.New()
	hasher.Write(dataStr)
	hashBytes := hasher.Sum(nil)

	values := make([]float64, len(summary))
	legends := make([]string, len(summary))
	total := 0.0
	for i, v := range summary {
		total += v.Total
		values[i] = v.Total
		legends[i] = v.Category
	}

	data := make([][4]string, len(expenses))
	for i, expense := range expenses {
		data[i] = [4]string{
			expense.Date.Format("2006/01/02"),
			expense.Category,
			expense.Description,
			fmt.Sprintf("%.0f", expense.OriginalAmount),
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

func (s *Server) CreateTrendingImage(ctx context.Context, req *proto.CreateTrendingImageRequest) (*proto.CreateTrendingImageResponse, error) {
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

	now := time.Now()

	values := make([]map[string]float64, 10)
	legends := make([]string, 10)

	for i := 0; i < 10; i += 1 {
		var start, end time.Time
		switch req.Type {
		case "week":
			start = now.AddDate(0, 0, int(time.Sunday)-int(now.Weekday())-7).AddDate(0, 0, -7*i)
			end = start.AddDate(0, 0, 6)
		case "month":
			start = time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, now.Location()).AddDate(0, -i, 0)
			end = start.AddDate(0, 1, -1)
		}

		summary, err := s.store.SummarizeExpensesWithinDate(uint(req.GroupId), start, end)
		if err != nil {
			return nil, err
		}

		legends[i] = fmt.Sprintf("%s ~ %s", start.Format("2006/01/02"), end.Format("2006/01/02"))
		values[i] = make(map[string]float64)
		for _, v := range summary {
			values[i][v.Category] = v.Total
		}
	}

	dataStr, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	hasher := sha256.New()
	hasher.Write(dataStr)
	hashBytes := hasher.Sum(nil)

	title := fmt.Sprintf("%s trend", req.Type)
	path := fmt.Sprintf("/var/images/%x.html", hashBytes)
	utils.CreateTrendingChart(values, legends, title, path)

	return &proto.CreateTrendingImageResponse{
		Url: fmt.Sprintf("%x.html", hashBytes),
	}, nil
}
