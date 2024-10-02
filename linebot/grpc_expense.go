package main

import (
	"bill-splitting-linebot/proto"
	"context"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/grpc/metadata"
)

func (s *LineBotServer) createExpense(token string, groupId uint32, category, description, currency, amount string) string {
	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	date := time.Now().Format("2006-01-02")

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "新增失敗"
	}

	expense, err := s.GrpcClient.CreateExpense(ctx, &proto.CreateExpenseRequest{
		GroupId:        groupId,
		Category:       category,
		Description:    description,
		OriginCurrency: currency,
		OriginAmount:   amountFloat,
		Date:           date,
	})
	if err != nil {
		return "新增失敗"
	}

	return "新增成功"
}

func (s *LineBotServer) getExpenseImage(token string, groupId uint32, summaryType string) (string, error) {
	var startTime, endTime time.Time
	now := time.Now()

	switch summaryType {
	case "本月支出":
		startTime = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		endTime = startTime.AddDate(0, 1, -1)
	case "今年支出":
		startTime = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
		endTime = startTime.AddDate(1, 0, -1)
	case "本周支出", "本週支出":
		startTime = now.AddDate(0, 0, int(time.Sunday)-int(now.Weekday()))
		endTime = startTime.AddDate(0, 0, 6)
	case "上週支出", "上周支出":
		startTime = now.AddDate(0, 0, int(time.Sunday)-int(now.Weekday())-7)
		endTime = startTime.AddDate(0, 0, 6)
	}

	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err := s.GrpcClient.CreateExpenseSummaryChart(ctx, &proto.CreateExpenseSummaryChartRequest{
		GroupId:   groupId,
		StartDate: startTime.Format("2006-01-02"),
		EndDate:   endTime.Format("2006-01-02"),
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://arensusu.ddns.net/api/v1/images/%s", resp.Url), nil
}

func (s *LineBotServer) getTrendingImage(token string, groupId uint32, summaryType string) (string, error) {
	trendingType := ""
	switch summaryType {
	case "周趨勢", "週趨勢":
		trendingType = "week"
	case "月趨勢":
		trendingType = "month"
	}

	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err := s.GrpcClient.CreateTrendingImage(ctx, &proto.CreateTrendingImageRequest{
		GroupId: groupId,
		Type:    trendingType,
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://arensusu.ddns.net/api/v1/images/%s", resp.Url), nil
}
