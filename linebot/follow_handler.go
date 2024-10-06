package main

import (
	"log/slog"

	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func (s *LineBotServer) followHandler(event webhook.FollowEvent) {
	source, ok := event.Source.(webhook.UserSource)
	if !ok {
		slog.Error("unknown source:", "error", "source is not user source")
		return
	}

	profile, err := s.MsgApi.GetProfile(source.UserId)
	if err != nil {
		slog.Error("get profile err:", slog.Any("error", err))
		return
	}

	token, err := s.getAuthToken(source.UserId, profile.DisplayName)
	if err != nil {
		slog.Error("get auth token err:", slog.Any("error", err))
		return
	}

	if _, err = s.createGroup(token, "", "個人"); err != nil {
		slog.Error("create group err:", slog.Any("error", err))
		return
	}

	replyMessage := linebot.NewTextMessage("Welcome " + profile.DisplayName + "!")
	if _, err := s.Bot.ReplyMessage(event.ReplyToken, replyMessage).Do(); err != nil {
		slog.Error("reply message err:", slog.Any("error", err))
	}
}
