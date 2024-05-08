package gapi

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/proto"
	"bill-splitting/token"
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	return payload, nil
}

func (s *Server) GetAuthToken(ctx context.Context, req *proto.GetAuthTokenRequest) (*proto.GetAuthTokenResponse, error) {
	if req.GetId() == "" || req.GetUsername() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id or username")
	}
	user, err := s.store.GetUser(ctx, req.GetId())
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}

		user, err = s.store.CreateUser(ctx, db.CreateUserParams{
			ID:       req.Id,
			Username: req.Username,
		})
		if err != nil {
			return nil, err
		}
	}

	if user.Username != req.Username {
		return nil, fmt.Errorf("invalid username")
	}

	token, _, err := s.tokenMaker.CreateToken(user.ID, time.Hour)
	if err != nil {
		return nil, err
	}
	return &proto.GetAuthTokenResponse{Token: token}, nil
}
