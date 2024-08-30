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
	"os"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
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

	summary, err := s.store.SummarizeExpensesWithinDate(ctx, db.SummarizeExpensesWithinDateParams{
		GroupID:   req.GroupId,
		StartTime: startDate,
		EndTime:   endDate,
	})
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(summary)
	if err != nil {
		return nil, err
	}

	hasher := sha256.New()
	hasher.Write(data)
	hashBytes := hasher.Sum(nil)

	values := make([]chart.Value, len(summary))
	for i, v := range summary {
		values[i] = chart.Value{
			Value: float64(v.Total),
			Label: fmt.Sprintf("%s $%d", v.Category.String, v.Total),
		}
	}

	fontBytes, err := os.ReadFile("./msjh.ttc")
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	pie := chart.PieChart{
		Title:  fmt.Sprintf("%s ~ %s", req.StartDate, req.EndDate),
		Width:  500,
		Height: 600,
		Values: values,
		Font:   font,
	}

	f, err := os.Create(fmt.Sprintf("/var/images/%x.png", hashBytes))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = pie.Render(chart.PNG, f)
	if err != nil {
		return nil, err
	}

	return &proto.CreateExpenseSummaryChartResponse{
		Url: fmt.Sprintf("%x.png", hashBytes),
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
		MemberID:    member.ID,
		Amount:      fmt.Sprint(amount),
		Description: req.Description,
		Category:    sql.NullString{String: req.Category, Valid: req.Category != ""},
		Date:        time.Now(),
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
