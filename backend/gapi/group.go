package gapi

import (
	db "bill-splitting/db/sqlc"
	"bill-splitting/proto"
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateLineGroup(ctx context.Context, req *proto.CreateLineGroupRequest) (*proto.CreateLineGroupResponse, error) {
	payload, err := s.authorize(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if req.LineId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid line id")
	}

	group, err := s.store.CreateGroupTx(ctx, db.CreateGroupTxParams{
		Name:   req.Name,
		UserID: payload.UserID,
		LineId: req.LineId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}

	return &proto.CreateLineGroupResponse{
		Id:     group.ID,
		LineId: group.LineID.String,
		Name:   group.Name,
	}, nil
}

func (s *Server) GetLineGroup(ctx context.Context, req *proto.GetLineGroupRequest) (*proto.GetLineGroupResponse, error) {
	_, err := s.authorize(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if req.LineId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid line id")
	}

	group, err := s.store.GetLineGroup(ctx, sql.NullString{
		String: req.LineId,
		Valid:  true,
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &proto.GetLineGroupResponse{
		Id:     group.ID,
		LineId: group.LineID.String,
		Name:   group.Name,
	}, nil
}

func (s *Server) AddMembership(ctx context.Context, req *proto.AddMembershipRequest) (*proto.AddMembershipResponse, error) {
	payload, err := s.authorize(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if req.GroupId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid group id or user id")
	}

	member, err := s.store.CreateMember(ctx, db.CreateMemberParams{
		GroupID: req.GroupId,
		UserID:  payload.UserID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, err.Error())
	}

	return &proto.AddMembershipResponse{
		Id:      member.ID,
		GroupId: member.GroupID,
		UserId:  member.UserID,
	}, nil
}

func (s *Server) GetMembership(ctx context.Context, req *proto.GetMembershipRequest) (*proto.GetMembershipResponse, error) {
	payload, err := s.authorize(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	if req.GroupId == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid group id or user id")
	}

	member, err := s.store.GetMembership(ctx, db.GetMembershipParams{
		GroupID: req.GroupId,
		UserID:  payload.UserID,
	})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &proto.GetMembershipResponse{
		Id:      member.ID,
		GroupId: member.GroupID,
		UserId:  member.UserID,
	}, nil
}
