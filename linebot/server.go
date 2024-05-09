package main

import (
	"bill-splitting-linebot/proto"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"google.golang.org/grpc/metadata"
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
		case webhook.JoinEvent:
			s.joinHandler(e)
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

	token, err := s.getAuthToken(source.UserId, profile.DisplayName)
	if err != nil {
		log.Println("getAuthToken err:", err)
		return
	}

	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err := s.GrpcClient.CreateLineGroup(ctx, &proto.CreateLineGroupRequest{
		Name:   "個人",
		LineId: source.UserId,
	})
	if err != nil {
		log.Println("CreateLineGroup err:", err)
		return
	}

	if resp.Name != "個人" || resp.LineId != source.UserId {
		log.Println("group name or line id is not match")
		return
	}

	replyMessage := linebot.NewTextMessage("Welcome " + profile.DisplayName + "!")
	if _, err := s.Bot.ReplyMessage(event.ReplyToken, replyMessage).Do(); err != nil {
		log.Print(err)
	}
}

func (s *LineBotServer) joinHandler(event webhook.JoinEvent) {
	source, ok := event.Source.(webhook.GroupSource)
	if !ok {
		log.Printf("Unknown source: %v", source)
		return
	}

	memberIds, err := s.MsgApi.GetGroupMembersIds(source.GroupId, "")
	if err != nil {
		log.Println("GetGroupMembersIds err:", err)
		return
	}

	profiles := []*messaging_api.GroupUserProfileResponse{}
	for _, memberId := range memberIds.MemberIds {
		profile, err := s.MsgApi.GetGroupMemberProfile(source.GroupId, memberId)
		if err != nil {
			log.Println("GetGroupMemberProfile err:", err)
			return
		}
		profiles = append(profiles, profile)
	}

	var token string
	for _, profile := range profiles {
		token, err = s.getAuthToken(source.UserId, profile.DisplayName)
		if err != nil {
			log.Println("getAuthToken err:", err)
			return
		}
	}

	group, err := s.MsgApi.GetGroupSummary(source.GroupId)
	if err != nil {
		log.Println("GetGroupSummary err:", err)
		return
	}

	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	createGroupResp, err := s.GrpcClient.CreateLineGroup(ctx, &proto.CreateLineGroupRequest{
		Name:   group.GroupName,
		LineId: source.GroupId,
	})
	if err != nil {
		log.Println("CreateLineGroup err:", err)
		return
	}

	if createGroupResp.Name != group.GroupName || createGroupResp.LineId != source.GroupId {
		log.Println("group name or line id is not match")
		return
	}

	for i := 0; i < len(memberIds.MemberIds)-1; i += 1 {
		resp, err := s.GrpcClient.AddGroupMember(ctx, &proto.AddGroupMemberRequest{
			GroupId: createGroupResp.Id,
			UserId:  memberIds.MemberIds[i],
		})
		if err != nil {
			log.Println("AddGroupMember err:", err)
			return
		}
		if resp.GroupId != createGroupResp.Id || resp.UserId != memberIds.MemberIds[i] {
			log.Println("group id or user id is not match")
			return
		}
	}

	replyMessage := linebot.NewTextMessage("Hello " + group.GroupName + "!")
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
