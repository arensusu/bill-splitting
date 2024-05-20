package main

import (
	"bill-splitting-linebot/proto"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

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
		}
	}
}

func (s *LineBotServer) messageHandler(event webhook.MessageEvent) {
	var replyMessage linebot.SendingMessage

	var userId, lineGroupId string
	switch source := event.Source.(type) {
	case webhook.UserSource:
		userId = source.UserId
	case webhook.GroupSource:
		userId = source.UserId
		lineGroupId = source.GroupId
	default:
		log.Printf("Unknown source: %v", source)
	}

	switch message := event.Message.(type) {
	// Handle only on text message
	case webhook.TextMessageContent:

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

		var groupId int32
		if lineGroupId != "" {
			groupId, err = s.getGroup(token, lineGroupId)
			if err != nil {
				group, err := s.MsgApi.GetGroupSummary(lineGroupId)
				if err != nil {
					log.Println("GetGroupSummary err:", err)
					replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
					break
				}
				if groupId, err = s.createGroup(token, lineGroupId, group.GroupName); err != nil {
					log.Println("createGroup err:", err)
					replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
					break
				}
			}

			if err = s.checkMembership(token, groupId); err != nil {
				err = s.addMembership(token, groupId)
				if err != nil {
					log.Println("addMember err:", err)
					replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
					break
				}
			}
		} else {
			groupId, err = s.getGroup(token, userId)
			if err != nil {
				log.Println("getGroup err:", err)
				replyMessage = linebot.NewTextMessage("發生錯誤，請稍後再試")
				break
			}
		}

		msgList := strings.Split(message.Text, "\n")
		if len(msgList) == 3 {
			msg := createExpense(token, groupId, msgList[0], msgList[1], msgList[2])
			replyMessage = linebot.NewTextMessage(msg)
		} else if strings.Contains(msgList[0], "支出") {
			imgUrl, err := s.getExpenseImage(token, groupId, msgList[0])
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

	if _, err = s.createGroup(token, "", "個人"); err != nil {
		log.Println("createGroup err:", err)
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

func (s *LineBotServer) createGroup(token, lineGroupId, groupName string) (int32, error) {
	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	createGroupResp, err := s.GrpcClient.CreateLineGroup(ctx, &proto.CreateLineGroupRequest{
		Name:   groupName,
		LineId: lineGroupId,
	})
	if err != nil {
		return 0, fmt.Errorf("CreateLineGroup err: %v", err)
	}

	if createGroupResp.Name != groupName || createGroupResp.LineId != lineGroupId {
		return 0, errors.New("group name or line id is not match")
	}

	return createGroupResp.Id, nil
}

func (s *LineBotServer) getGroup(token, lineGroupId string) (int32, error) {
	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	group, err := s.GrpcClient.GetLineGroup(ctx, &proto.GetLineGroupRequest{
		LineId: lineGroupId,
	})
	if err != nil {
		return 0, fmt.Errorf("GetLineGroup err: %v", err)
	}
	return group.GetId(), nil
}

func (s *LineBotServer) getExpenseImage(token string, groupId int32, summaryType string) (string, error) {
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
	resp, err := s.GrpcClient.CreateExpenseSummaryChart(ctx, &proto.CreateExpenseSummaryChartRequest{
		GroupId:   int32(groupId),
		StartDate: startTime.Format("2006-01-02"),
		EndDate:   endTime.Format("2006-01-02"),
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://arensusu.ddns.net/api/v1/images/%s", resp.Url), nil
}

func (s *LineBotServer) checkMembership(token string, groupId int32) error {
	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err := s.GrpcClient.GetMembership(ctx, &proto.GetMembershipRequest{
		GroupId: groupId,
	})
	return err
}

func (s *LineBotServer) addMembership(token string, groupId int32) error {
	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err := s.GrpcClient.AddMembership(ctx, &proto.AddMembershipRequest{
		GroupId: groupId,
	})
	return err
}
