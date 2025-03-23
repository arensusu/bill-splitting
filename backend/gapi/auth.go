package gapi

import (
	"bill-splitting/model"
	"bill-splitting/proto"
	"bill-splitting/token"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"
	"gorm.io/gorm"
)

func (s *Server) authorize(ctx context.Context) (*token.JWTPayload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata is not provided")
	}

	values := md.Get("authorization")
	if len(values) == 0 {
		return nil, fmt.Errorf("authorization header is not provided")
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != "bearer" {
		return nil, fmt.Errorf("unsupported authorization type")
	}

	accessToken := fields[1]
	payload, err := s.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	log.Printf("payload: %+v", payload)
	return payload, nil
}

func (s *Server) GetAuthToken(ctx context.Context, req *proto.GetAuthTokenRequest) (*proto.GetAuthTokenResponse, error) {
	var user *model.User
	var err error
	if req.GetLineId() != "" {
		user, err = s.store.GetUserByLineID(req.GetLineId())
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("failed to get user: %w", err)
			}

			err = s.store.CreateUser(&model.User{
				LineID: req.GetLineId(),
			})
			if err != nil {
				return nil, fmt.Errorf("failed to create user: %w", err)
			}

			user, err = s.store.GetUserByLineID(req.GetLineId())
			if err != nil {
				return nil, fmt.Errorf("failed to get user: %w", err)
			}
		}
	} else if req.GetDiscordId() != "" {
		user, err = s.store.GetUserByDiscordID(req.GetDiscordId())
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("failed to get user: %w", err)
			}

			err = s.store.CreateUser(&model.User{
				DiscordId: req.GetDiscordId(),
			})
			if err != nil {
				return nil, fmt.Errorf("failed to create user: %w", err)
			}

			user, err = s.store.GetUserByDiscordID(req.GetDiscordId())
			if err != nil {
				return nil, fmt.Errorf("failed to get user: %w", err)
			}
		}
	} else {
		return nil, fmt.Errorf("line id or discord id is not provided")
	}

	token, _, err := s.tokenMaker.CreateToken(fmt.Sprintf("%d", user.ID), time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to create token: %w", err)
	}
	return &proto.GetAuthTokenResponse{Token: token}, nil
}
