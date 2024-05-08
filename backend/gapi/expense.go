package gapi

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/proto"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/v2"
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
		Title:  fmt.Sprintf("%s ~ %s", startDate, endDate),
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
