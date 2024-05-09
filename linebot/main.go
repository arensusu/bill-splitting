// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bill-splitting-linebot/proto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	bot, err := linebot.New(os.Getenv("LINEBOT_SECRET"), os.Getenv("LINEBOT_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	client, err := messaging_api.NewMessagingApiAPI(os.Getenv("LINEBOT_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal("NewMessagingApiAPI err:", err)
	}

	conn, err := grpc.Dial("api:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("conn err:", err)
		return
	}
	defer conn.Close()
	grpcClient := proto.NewBillSplittingClient(conn)

	server := NewLineBotServer(bot, client, grpcClient)

	http.HandleFunc("/callback", server.callbackHandler)
	port := os.Getenv("LINEBOT_PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func createExpense(token, category, description, amount string) string {
	groupId := 1
	uri := fmt.Sprintf("http://api:8080/api/v1/groups/%d/expenses", groupId)

	date := time.Now().Format("2006-01-02")
	body, err := json.Marshal(map[string]string{
		"category":    category,
		"description": description,
		"amount":      amount,
		"date":        date,
	})
	if err != nil {
		return "發生錯誤，請稍後再試"
	}

	req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		return "發生錯誤，請稍後再試"
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "發生錯誤，請稍後再試"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "發生錯誤，請稍後再試"
	}
	return "新增成功"
}

func getExpenseImage(client proto.BillSplittingClient, token, summaryType string) (string, error) {
	groupId := 1

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
		endTime = startTime.AddDate(0, 0, 7)
	}

	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err := client.CreateExpenseSummaryChart(ctx, &proto.CreateExpenseSummaryChartRequest{
		GroupId:   int32(groupId),
		StartDate: startTime.Format("2006-01-02"),
		EndDate:   endTime.Format("2006-01-02"),
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://arensusu.ddns.net/api/v1/images/%s", resp.Url), nil
}
