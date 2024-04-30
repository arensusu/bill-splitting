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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("LINEBOT_SECRET"), os.Getenv("LINEBOT_ACCESS_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("LINEBOT_PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	cb, err := webhook.ParseRequest(os.Getenv("LINEBOT_SECRET"), r)
	log.Printf("%v", err)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	client, err := messaging_api.NewMessagingApiAPI(os.Getenv("LINEBOT_ACCESS_TOKEN"))
	if err != nil {
		log.Println("NewMessagingApiAPI err:", err)
		return
	}

	for _, event := range cb.Events {
		switch e := event.(type) {
		case webhook.MessageEvent:
			var replyMessage linebot.SendingMessage

			switch message := e.Message.(type) {
			// Handle only on text message
			case webhook.TextMessageContent:
				var userId, lineGroupId string
				switch source := e.Source.(type) {
				case webhook.UserSource:
					userId = source.UserId
					lineGroupId = userId
				case webhook.GroupSource:
					userId = source.UserId
					lineGroupId = source.GroupId
				default:
					log.Printf("Unknown source: %v", source)
				}
				_ = lineGroupId

				profile, err := client.GetProfile(userId)
				if err != nil {
					log.Println("GetProfile err:", err)
					replyMessage = linebot.NewTextMessage("找不到使用者，請確認使用者是否將官方帳號加入好友")
					break
				}

				token, err := getAuthToken(userId, profile.DisplayName)
				if err != nil {
					log.Println("getAuthToken err:", err)
					replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
					break
				}

				msgList := strings.Split(message.Text, "\n")
				if len(msgList) == 3 {
					msg := createExpense(token, msgList[0], msgList[1], msgList[2])
					replyMessage = linebot.NewTextMessage(msg)
				} else if strings.Contains(msgList[0], "支出") {
					imgUrl, err := getExpenseImage(token, msgList[0])
					if err != nil {
						replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
					} else {
						replyMessage = linebot.NewImageMessage(imgUrl, imgUrl)
					}
				}

			default:
				log.Printf("Unknown message: %v", message)
			}

			if _, err = bot.ReplyMessage(e.ReplyToken, replyMessage).Do(); err != nil {
				log.Print(err)
			}
		case webhook.FollowEvent:
			log.Printf("message: Got followed event")
		case webhook.PostbackEvent:
			data := e.Postback.Data
			log.Printf("Unknown message: Got postback: " + data)
		case webhook.BeaconEvent:
			log.Printf("Got beacon: " + e.Beacon.Hwid)
		}
	}
}

func getAuthToken(userId string, displayName string) (string, error) {
	uri := "http://api:8080/api/v1/auth/linebot"
	body, err := json.Marshal(map[string]string{"id": userId, "username": displayName})
	if err != nil {
		return "", fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequest("POST", uri, bytes.NewReader(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println(string(data))

	var res struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(data, &res); err != nil {
		return "", fmt.Errorf("json.Unmarshal: %w", err)
	}
	return res.Token, nil
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

func getExpenseImage(token, summaryType string) (string, error) {
	groupId := 1
	uri := fmt.Sprintf("http://api:8080/api/v1/groups/%d/expenses/summary", groupId)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", err
	}

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

	q := req.URL.Query()
	q.Add("startTime", startTime.Format("2006-01-02"))
	q.Add("endTime", endTime.Format("2006-01-02"))
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var res struct {
		Image string `json:"image"`
	}
	if err := json.Unmarshal(data, &res); err != nil {
		return "", fmt.Errorf("json.Unmarshal: %w", err)
	}
	return fmt.Sprintf("https://arensusu.ddns.net/api/v1/images/%s", res.Image), nil
}
