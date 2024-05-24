package main

import (
	"log"

	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

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

	token, err := s.getAuthToken(source.UserId, profile.DisplayName)
	if err != nil {
		log.Println("getAuthToken err:", err)
		return
	}

	if _, err = s.createGroup(token, "", "個人"); err != nil {
		log.Println("createGroup err:", err)
		return
	}

	replyMessage := linebot.NewTextMessage("Welcome " + profile.DisplayName + "!")
	if _, err := s.Bot.ReplyMessage(event.ReplyToken, replyMessage).Do(); err != nil {
		log.Print(err)
	}
}
