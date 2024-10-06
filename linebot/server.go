package main

import (
	"bill-splitting-linebot/proto"
	"fmt"
	"log/slog"
	"net/http"
	"os"

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

func (s *LineBotServer) Start(addr string) {
	slog.Info(fmt.Sprintf("Listening on %s", addr))
	http.ListenAndServe(addr, nil)
}

func (s *LineBotServer) callbackHandler(w http.ResponseWriter, r *http.Request) {
	cb, err := webhook.ParseRequest(os.Getenv("LINEBOT_SECRET"), r)
	if err != nil {
		slog.Error("parse request err:", slog.Any("error", err))

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
		}
	}
}
