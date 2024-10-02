package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"google.golang.org/api/option"
)

func (s *LineBotServer) messageHandler(event webhook.MessageEvent) {
	var replyMessage linebot.SendingMessage

	source, err := s.getSource(event.Source)
	if err != nil {
		log.Println("GetSource err:", err)
		replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
	}

	switch message := event.Message.(type) {
	// Handle only on text message
	case webhook.TextMessageContent:

		profile, err := s.MsgApi.GetProfile(source.UserId)
		if err != nil {
			log.Println("GetProfile err:", err)
			replyMessage = linebot.NewTextMessage("找不到使用者，請確認使用者是否將官方帳號加入好友")
			break
		}

		token, err := s.getAuthToken(source.UserId, profile.DisplayName)
		if err != nil {
			log.Println("getAuthToken err:", err)
			replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
			break
		}

		var groupId uint32
		if source.IsGroupChat {
			groupId, err = s.groupChatPreProcessing(token, source)
			if err != nil {
				log.Println("groupChatPreProcessing err:", err)
				replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
				break
			}
		} else {
			groupId, err = s.getGroup(token, source.UserId)
			if err != nil {
				log.Println("getGroup err:", err)
				replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
				break
			}
		}

		msgList := strings.Split(message.Text, "\n")
		if len(msgList) == 3 {
			price := strings.Split(msgList[2], " ")

			var msg string
			if len(price) == 1 {
				msg = s.createExpense(token, groupId, msgList[0], msgList[1], "TWD", msgList[2])
			} else {
				msg = s.createExpense(token, groupId, msgList[0], msgList[1], price[0], price[1])
			}
			replyMessage = linebot.NewTextMessage(msg)
		} else if strings.Contains(msgList[0], "支出") {
			imgUrl, err := s.getExpenseImage(token, groupId, msgList[0])
			if err != nil {
				log.Println("getExpenseImage err:", err)
				replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
			} else {
				replyMessage = linebot.NewTemplateMessage(imgUrl, linebot.NewButtonsTemplate("", "", "支出圖表", &linebot.URIAction{Label: "查看", URI: imgUrl}))
			}
		} else if strings.Contains(msgList[0], "趨勢") {
			imgUrl, err := s.getTrendingImage(token, groupId, msgList[0])
			if err != nil {
				log.Println("getExpenseImage err:", err)
				replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
			} else {
				replyMessage = linebot.NewTemplateMessage(imgUrl, linebot.NewButtonsTemplate("", "", "趨勢圖表", &linebot.URIAction{Label: "查看", URI: imgUrl}))
			}
		} else {
			aiResp, err := callGemini(context.Background(), message.Text)
			if err != nil {
				log.Println("callGemini err:", err)
				replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
			} else {
				replyMessage = linebot.NewTextMessage(aiResp)
			}
		}

	default:
		log.Printf("Unknown message: %v", message)
	}

	if _, err := s.Bot.ReplyMessage(event.ReplyToken, replyMessage).Do(); err != nil {
		log.Print(err)
	}
}

type Source struct {
	UserId      string
	GroupId     *string
	IsGroupChat bool
}

func (s *LineBotServer) getSource(source webhook.SourceInterface) (Source, error) {
	switch source := source.(type) {
	case webhook.UserSource:
		return Source{UserId: source.UserId, IsGroupChat: false}, nil
	case webhook.GroupSource:
		return Source{UserId: source.UserId, GroupId: &source.GroupId, IsGroupChat: true}, nil
	default:
		return Source{}, errors.New("unknown source type")
	}
}

func (s *LineBotServer) groupChatPreProcessing(token string, source Source) (uint32, error) {
	groupId, err := s.getGroup(token, *source.GroupId)
	if err != nil {
		group, err := s.MsgApi.GetGroupSummary(*source.GroupId)
		if err != nil {
			return 0, fmt.Errorf("GetGroupSummary err: %w", err)
		}
		if groupId, err = s.createGroup(token, *source.GroupId, group.GroupName); err != nil {
			return 0, fmt.Errorf("createGroup err: %w", err)
		}
	}

	if err = s.checkMembership(token, groupId); err == nil {
		return groupId, nil
	}

	if err = s.addMembership(token, groupId); err != nil {
		return 0, fmt.Errorf("addMembership err: %w", err)
	}
	return groupId, nil
}

func callGemini(ctx context.Context, message string) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Print(err)
	}
	defer client.Close()

	// Get the model
	model := client.GenerativeModel("gemini-1.5-flash")

	systemPrompt, err := os.ReadFile("prompt.txt")
	if err != nil {
		return "", fmt.Errorf("failed to read prompt file: %w", err)
	}

	// Generate content
	resp, err := model.GenerateContent(ctx, genai.Text(systemPrompt), genai.Text(message))
	if err != nil {
		log.Fatal(err)
	}

	// Print the response
	log.Printf("Response: %#v\n", resp)

	content, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return "", errors.New("failed to get text from response")
	}

	return string(content), nil
}
