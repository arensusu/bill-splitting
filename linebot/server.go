package main

import (
	"bill-splitting-linebot/proto"
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type LineBotServer struct {
	Bot        *linebot.Client
	MsgApi     *messaging_api.MessagingApiAPI
	GrpcClient proto.BillSplittingClient
}

func NewLineBotServer(bot *linebot.Client, msgApi *messaging_api.MessagingApiAPI, grpcClient proto.BillSplittingClient) *LineBotServer {
	return &LineBotServer{Bot: bot, MsgApi: msgApi, GrpcClient: grpcClient}
}

func (s *LineBotServer) callbackHandler(w http.ResponseWriter, r *http.Request) {
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

	for _, event := range cb.Events {
		switch e := event.(type) {
		case webhook.MessageEvent:
			s.messageHandler(e)
		case webhook.FollowEvent:
			s.followHandler(e)
		case webhook.PostbackEvent:
			data := e.Postback.Data
			log.Printf("Unknown message: Got postback: " + data)
		case webhook.BeaconEvent:
			log.Printf("Got beacon: " + e.Beacon.Hwid)
		}
	}
}

func (s *LineBotServer) messageHandler(event webhook.MessageEvent) {
	var replyMessage linebot.SendingMessage

	var userId, lineGroupId string
	switch source := event.Source.(type) {
	case webhook.UserSource:
		userId = source.UserId
		lineGroupId = userId
	case webhook.GroupSource:
		userId = source.UserId
		lineGroupId = source.GroupId
	default:
		log.Printf("Unknown source: %v", source)
	}

	switch message := event.Message.(type) {
	// Handle only on text message
	case webhook.TextMessageContent:
		_ = lineGroupId

		profile, err := s.MsgApi.GetProfile(userId)
		if err != nil {
			log.Println("GetProfile err:", err)
			replyMessage = linebot.NewTextMessage("找不到使用者，請確認使用者是否將官方帳號加入好友")
			break
		}

		token, err := s.getAuthToken(userId, profile.DisplayName)
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
			imgUrl, err := getExpenseImage(s.GrpcClient, token, msgList[0])
			if err != nil {
				log.Println("getExpenseImage err:", err)
				replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
			} else {
				replyMessage = linebot.NewImageMessage(imgUrl, imgUrl)
			}
		}

	default:
		log.Printf("Unknown message: %v", message)
	}

	if _, err := s.Bot.ReplyMessage(event.ReplyToken, replyMessage).Do(); err != nil {
		log.Print(err)
	}
}

func (s *LineBotServer) followHandler(event webhook.FollowEvent) {
	source, ok := event.Source.(webhook.UserSource)
	if !ok {
		log.Printf("Unknown source: %v", source)
		return
	}

	profile, err := s.MsgApi.GetProfile(source.UserId)
	if err != nil {
		log.Println("GetProfile err:", err)
		return
	}

	_, err = s.getAuthToken(source.UserId, profile.DisplayName)
	if err != nil {
		log.Println("getAuthToken err:", err)
		return
	}

	replyMessage := linebot.NewTextMessage("Welcome " + profile.DisplayName + "!")
	if _, err := s.Bot.ReplyMessage(event.ReplyToken, replyMessage).Do(); err != nil {
		log.Print(err)
	}
}

func (s *LineBotServer) getAuthToken(userId string, displayName string) (string, error) {
	resp, err := s.GrpcClient.GetAuthToken(context.Background(), &proto.GetAuthTokenRequest{
		Id:       userId,
		Username: displayName,
	})
	if err != nil {
		return "", err
	}

	return resp.Token, nil
}
