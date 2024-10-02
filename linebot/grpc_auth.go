package main

import (
	"bill-splitting-linebot/proto"
	"context"
)

func (s *LineBotServer) getAuthToken(userId string, displayName string) (string, error) {
	resp, err := s.GrpcClient.GetAuthToken(context.Background(), &proto.GetAuthTokenRequest{
		LineId:   userId,
		Username: displayName,
	})
	if err != nil {
		return "", err
	}

	return resp.Token, nil
}
