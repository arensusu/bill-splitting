package main

import (
	"bill-splitting-linebot/proto"
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func (s *LineBotServer) getGroup(token, lineGroupId string) (uint32, error) {
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

func (s *LineBotServer) createGroup(token, lineGroupId, groupName string) (uint32, error) {
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

func (s *LineBotServer) checkMembership(token string, groupId uint32) error {
	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err := s.GrpcClient.GetMembership(ctx, &proto.GetMembershipRequest{
		GroupId: groupId,
	})
	return err
}

func (s *LineBotServer) addMembership(token string, groupId uint32) error {
	md := metadata.New(map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	_, err := s.GrpcClient.AddMembership(ctx, &proto.AddMembershipRequest{
		GroupId: groupId,
	})
	return err
}
