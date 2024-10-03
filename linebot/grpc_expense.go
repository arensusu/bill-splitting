package main

import (
	"bill-splitting-linebot/proto"
	"bill-splitting-linebot/utils"
	"context"
	"crypto/sha256"
	"encoding/json"
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

	_, err = s.GrpcClient.CreateExpense(ctx, &proto.CreateExpenseRequest{
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

	startDate := startTime.Format("2006-01-02")
	endDate := endTime.Format("2006-01-02")

	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	summariesResp, err := s.GrpcClient.ListExpenseSummary(ctx, &proto.ListExpenseSummaryRequest{
		GroupId:   groupId,
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		return "", err
	}
	summaries := summariesResp.Summaries

	expensesResp, err := s.GrpcClient.ListExpense(ctx, &proto.ListExpenseRequest{
		GroupId:   groupId,
		StartDate: startDate,
		EndDate:   endDate,
	})
	if err != nil {
		return "", err
	}
	expenses := expensesResp.Expenses

	dataStr, err := json.Marshal(summaries)
	if err != nil {
		return "", fmt.Errorf("failed to marshal summary: %w", err)
	}

	hasher := sha256.New()
	hasher.Write(dataStr)
	hashBytes := hasher.Sum(nil)

	values := make([]float64, len(summaries))
	legends := make([]string, len(summaries))
	total := 0.0
	for i, v := range summaries {
		total += v.Total
		values[i] = v.Total
		legends[i] = v.Category
	}

	data := make([][4]string, len(expenses))
	for i, expense := range expenses {
		data[i] = [4]string{
			expense.Date,
			expense.Category,
			expense.Description,
			fmt.Sprintf("%.0f", expense.Amount),
		}
	}

	title := fmt.Sprintf("%s ~ %s", startDate, endDate)
	subtitle := fmt.Sprintf("Total: %.0f", total)
	path := fmt.Sprintf("/var/images/%x.html", hashBytes)

	err = utils.CreatePieChart(values, legends, title, subtitle, data, path)
	if err != nil {
		return "", fmt.Errorf("failed to create pie chart: %w", err)
	}

	return fmt.Sprintf("https://arensusu.ddns.net/images/%x.html", hashBytes), nil
}

func (s *LineBotServer) getTrendingImage(token string, groupId uint32, summaryType string) (string, error) {

	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	now := time.Now()

	values := make([]map[string]float64, 10)
	legends := make([]string, 10)

	for i := 0; i < 10; i += 1 {
		var start, end time.Time
		switch summaryType {
		case "周趨勢", "週趨勢":
			start = now.AddDate(0, 0, int(time.Sunday)-int(now.Weekday())-7).AddDate(0, 0, -7*i)
			end = start.AddDate(0, 0, 6)
		case "月趨勢":
			start = time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, now.Location()).AddDate(0, -i, 0)
			end = start.AddDate(0, 1, -1)
		}

		resp, err := s.GrpcClient.ListExpenseSummary(ctx, &proto.ListExpenseSummaryRequest{
			GroupId:   groupId,
			StartDate: start.Format("2006-01-02"),
			EndDate:   end.Format("2006-01-02"),
		})
		if err != nil {
			return "", err
		}

		legends[i] = fmt.Sprintf("%s ~ %s", start.Format("2006/01/02"), end.Format("2006/01/02"))
		values[i] = make(map[string]float64)
		for _, v := range resp.Summaries {
			values[i][v.Category] = v.Total
		}
	}

	dataStr, err := json.Marshal(values)
	if err != nil {
		return "", err
	}

	hasher := sha256.New()
	hasher.Write(dataStr)
	hashBytes := hasher.Sum(nil)

	title := summaryType
	path := fmt.Sprintf("/var/images/%x.html", hashBytes)
	utils.CreateTrendingChart(values, legends, title, path)

	return fmt.Sprintf("https://arensusu.ddns.net/images/%x.html", hashBytes), nil
}
